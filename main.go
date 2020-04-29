package main

import (
	"fmt"
	"github.com/testground/plan-templates/templates"
)

func main() {
	var ts templates.TemplateSet
	err := templates.Fill("/", ts)
	if err != nil {
		panic(err)
	}

	for _, t := range ts {
		fmt.Printf("\n\n---------------------\n")
		fmt.Println(t.Filename)
		fmt.Println(t.Template)
	}
}
