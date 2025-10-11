{{- define "base" }}
    {{- template "css" . }}
    {{- nindent 0 "<h1>CUSTOM TEMPLATE</h1>" }}
    {{- template "list" . }}
{{- end }}
