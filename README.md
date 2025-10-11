# 🧩 json2table

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/bgrooot/json2table.svg)](https://pkg.go.dev/github.com/bgrooot/json2table)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgrooot/json2table)](https://goreportcard.com/report/github.com/bgrooot/json2table)

A Go library that easily converts JSON strings into HTML table format.

## 🚀 Usage
```go
table, err := json2table.Json2table(jsonStr, config.Config{})
```

<img src="demo.gif" width="800" />

---

## ⚙️ Configuration
> Bold values indicate the defaults.

| Setting        | Type     | Values                      | Description                                                  | 
|----------------|----------|-----------------------------|--------------------------------------------------------------|
| `Orientation`  | string   | **vertical**<br/>horizontal | Determines the table layout direction.                       |
| `CSS`          | string   |                             | Inline CSS to be inserted inside a `<style>` tag.            |
| `CssWebLink`   | string   |                             | External CSS link (inserted as a `<link rel>` tag).          |
| `TableClass`   | string   |                             | Class name applied to the `<table>` element.                 |
| `TdClass`      | string   |                             | Class name added to `<td>` elements.                         |
| `TrClass`      | string   |                             | Class name added to `<tr>` elements.                         |
| `Template`     | []string |                             | List of template strings.                                    |
| `TemplateFile` | []string |                             | List of template file paths.                                 |
| `TemplateData` | any      |                             | Data to be passed into templates.                            |
| `DivideArray`  | bool     | true<br/>**false**          | Whether to separate JSON array items into individual tables. |
| `WrapWithHTML` | bool     | **true**<br/>false          | Whether to output the full HTML document.                    |

> ⚠️ When using bool fields, the default value (nil) can be explicitly differentiated by using the config.CreateNilBool helper function.

---

## 🧱 Examples
### Vertical Layout
```go
cfg := config.Config{Orientation: "vertical"}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/yyeKNWo)

### Horizontal Layout
```go
cfg := config.Config{Orientation: "horizontal"}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/ogbqXrV)

### Custom CSS Link
```go
cfg := config.Config{
  CSSWebLink: "https://cdn.jsdelivr.net/npm/bulma@1.0.4/css/bulma.min.css",
  TableClass: "table is-bordered is-striped is-hoverable is-fullwidth",
}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/bNEvdXw)

### Custom Inline CSS
```go
cfg := config.Config{CSS: "th { background-color: #ebffea; }"}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/JoGLdge)

### JSON Array
```go
cfg := config.Config{}
table, err := json2table.Json2table(string(jsonArrayStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/ByjroBK)

### JSON Array (Divided)
```go
cfg := config.Config{DivideArray: config.CreateNilBool(true)}
table, err := json2table.Json2table(string(jsonArrayStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/EaPEVYp)

### Custom Inline Template
```go
base, _ := os.ReadFile("base.tpl")
list, _ := os.ReadFile("list.tpl")
cfg := config.Config{
  CSS:          "h1 { color: red; } li { list-style-type: square; color: blue; }",
  TemplateData: []string{"item1", "item2", "item3"},
  Template:     []string{string(list), string(base)},
}}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/gbPeaOp)

### Custom Template File
```go
cfg := config.Config{
  CSS:          "h1 { color: red; } li { list-style-type: square; color: blue; }",
  TemplateData: []string{"item1", "item2", "item3"},
  TemplateFiles: []string{"cmd/tpl/list.tpl", "cmd/tpl/base.tpl"},
}}
table, err := json2table.Json2table(string(jsonStr), cfg)
```
🔗 [View Result](https://codepen.io/bgrooot/full/dPGmYyE)

---

### 🧾 Built-in Templates
- The following templates are pre-defined and can be used when writing custom templates.

| Template Name | Description                                              |
|---------------|----------------------------------------------------------|
| `base`        | The entry point for all templates.                       |
| `table`       | Converts JSON into `<table>` markup.                     |
| `css`         | Renders the configured CSS inside a `<style>` tag.       |
| `cssWebLink`  | Inserts the CssWebLink value as a `<link href>` element. | 

#### Example
```go
{{- define "base" }}
  {{- template "css" . }}
  {{- nindent 0 "<h1>CUSTOM TEMPLATE</h1>" }}
  {{- template "list" . }}
{{- end }}

{{- define "list" }}
  {{- nindent 0 "<ul>" }}
  {{- range .Cfg.TemplateData }}
    {{- printf "<li>%s</li>" . | nindent 4 }}
  {{- end }}
  {{- nindent 0 "</ul>" }}
{{- end }}
```
