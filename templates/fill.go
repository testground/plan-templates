package templates

import (
	"github.com/markbates/pkger"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Templated struct {
	Filename string
	Template string
}

type TemplateSet []*Templated

// fillFiles -- filepath.WalkFunc decorator
// Fills a slice of *Templated
func fillWalk(t *TemplateSet) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		extension := ".template"
		fn := info.Name()
		if filepath.Ext(fn) != extension {
			return nil
		}
		// the path comes in a format which includes go mod information.
		// don't expect to be able to open it with regular file-handling utilities.
		rdr, err := pkger.Open(path)
		if err != nil {
			return err
		}
		defer rdr.Close()

		buf, err := ioutil.ReadAll(rdr)
		if err != nil {
			return err
		}
		newFn := fn[:len(fn)-len(extension)]
		*t = append(*t, &Templated{Filename: newFn, Template: string(buf)})
		return nil
	}
}

// Walk the path of templates, and return the contents of the file.
func Fill(path string, t *TemplateSet) error {
	return pkger.Walk(path, fillWalk(t))
}
