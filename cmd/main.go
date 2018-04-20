package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	lingo "github.com/remast/gherkin-lingo"
	_ "github.com/remast/gherkin-lingo/rules"

	gherkin "github.com/cucumber/gherkin-go"
)

func main() {
	serverFiles, _ := filepath.Glob(filepath.Join(".", "*.feature"))
	for _, f := range serverFiles {
		gherkinString, _ := ioutil.ReadFile(f)

		r := strings.NewReader(string(gherkinString))

		gherkinDocument, err := gherkin.ParseGherkinDocument(r)

		if err != nil {
			fmt.Printf("Datei %v enthält kein gültiges Gherkin und wird ignoriert.%v\n", f, err)
			continue
		}

		issues := lingo.Apply(f, gherkinDocument)
		for _, issue := range issues {
			fmt.Println(issue.String())
		}

	}

}
