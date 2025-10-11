package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"text/template"

	"github.com/bgrooot/json2table/pkg/json2table"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

type Case struct {
	data []byte
	cfg  config.Config
}

func main() {
	sample, _ := os.ReadFile("test/files/sample/sample.json")
	sampleArr, _ := os.ReadFile("test/files/sample/sample-array.json")
	typed, _ := os.ReadFile("test/files/sample/typed.json")
	err, _ := os.ReadFile("test/files/sample/error.json")
	base, _ := os.ReadFile("cmd/tpl/base.tpl")
	list, _ := os.ReadFile("cmd/tpl/list.tpl")

	caseMap := map[string]Case{
		"vertical":           {sample, config.Config{}},
		"horizontal":         {sample, config.Config{Orientation: "horizontal"}},
		"css-link":           {sample, config.Config{CSSWebLink: "https://cdn.jsdelivr.net/npm/bulma@1.0.4/css/bulma.min.css", TableClass: "table is-bordered is-striped is-hoverable is-fullwidth"}},
		"custom-css":         {sample, config.Config{CSS: "th { background-color: #ebffea; }"}},
		"json-array":         {sampleArr, config.Config{}},
		"json-array-divided": {sampleArr, config.Config{DivideArray: config.CreateNilBool(true)}},
		"typed":              {typed, config.Config{}},
		"no-wrap-with-html":  {sample, config.Config{WrapWithHTML: config.CreateNilBool(false)}},
		"custom-template": {sample, config.Config{
			CSS:          "h1 { color: red; } li { list-style-type: square; color: blue; }",
			TemplateData: []string{"item1", "item2", "item3"},
			Template:     []string{string(list), string(base)},
		}},
		"custom-template-file": {sample, config.Config{
			CSS:           "h1 { color: red; } li { list-style-type: square; color: blue; }",
			TemplateData:  []string{"item1", "item2", "item3"},
			TemplateFiles: []string{"cmd/tpl/list.tpl", "cmd/tpl/base.tpl"},
		}},
		"error-vertical":   {err, config.Config{}},
		"error-horizontal": {err, config.Config{Orientation: "horizontal"}},
	}

	for key, cs := range caseMap {
		run(key, string(cs.data), cs.cfg)
	}
}

func run(name string, jsonStr string, cfg config.Config) {
	table, err := json2table.Json2table(jsonStr, cfg)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("cmd/tpl/viewer.tpl"))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, table); err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Printf("[%s]\n", name)
	fmt.Printf("data:text/html;base64,%s\n", encoded)
}
