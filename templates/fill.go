package templates

import (
	"fmt"
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
func fillWalk(t TemplateSet) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		extension := ".template"
		fn := info.Name()
		fmt.Println(fn)
		if filepath.Ext(fn) != extension {
			return nil
		}
		buf, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		newFn := fn[:len(fn)-len(extension)]
		t = append(t, &Templated{Filename: newFn, Template: string(buf)})
		return nil
	}
}

// Walk the path of templates, and return the contents of the file.
func Fill(path string, t TemplateSet) error {
	info, err := pkger.Current()
	if err != nil {
		return err
	}
	fmt.Println(info)

	return pkger.Walk(path, fillWalk(t))
}
