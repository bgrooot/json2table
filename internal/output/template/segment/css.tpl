{{- define "css" }}
    {{- if .Cfg.CSS }}
<style>{{ .Cfg.CSS }}</style>
    {{- end }}
{{- end }}
