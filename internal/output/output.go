package output

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

//go:embed template
var templateFS embed.FS

func ToHtml(table []model.Table, cfg config.Config) (string, error) {
	html := ""
	tmpl := newTemplate()
	tmpl, err := parseFiles(tmpl, cfg)
	if err != nil {
		return html, err
	}

	for _, tpl := range cfg.Template {
		tmpl, err = tmpl.Parse(tpl)
		if err != nil {
			return html, err
		}
	}

	var buf bytes.Buffer
	data := map[string]any{
		"Cfg":   cfg,
		"Table": table,
	}

	if err = tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		return html, err
	}

	html = buf.String()
	return html, nil
}

func render(tmpl *template.Template, name string, data any) (string, error) {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func defaultTemplateFiles(cfg config.Config) []string {
	var files []string
	var segment = []string{
		"template/segment/css-web-link.tpl",
		"template/segment/css.tpl",
		"template/segment/table.tpl",
	}

	if cfg.WrapWithHTML.IsTrue() {
		files = append(files, "template/base/html.tpl")
	} else {
		files = append(files, "template/base/simple.tpl")
	}

	return append(files, segment...)
}

func configTemplateFiles(cfg config.Config) []string {
	var files []string
	for _, templateFile := range cfg.TemplateFiles {
		matches, err := filepath.Glob(templateFile)
		if err == nil {
			files = append(files, matches...)
		}
	}

	return files
}

func newTemplate() *template.Template {
	tmpl := template.New("base")
	tmpl = tmpl.Funcs(sprig.FuncMap()).Funcs(template.FuncMap{
		"render": func(name string, data any) string {
			result, err := render(tmpl, name, data)
			if err != nil {
				return fmt.Sprintf("render error: %v", err)
			}

			return result
		},
	})

	return tmpl
}

func parseFiles(tmpl *template.Template, cfg config.Config) (*template.Template, error) {
	defFiles := defaultTemplateFiles(cfg)
	cfgFiles := configTemplateFiles(cfg)

	tmpl, err := tmpl.ParseFS(templateFS, defFiles...)
	if err != nil {
		return tmpl, err
	}

	if len(cfgFiles) > 0 {
		tmpl, err = tmpl.ParseFiles(cfgFiles...)
		if err != nil {
			return tmpl, err
		}
	}

	return tmpl, nil
}
