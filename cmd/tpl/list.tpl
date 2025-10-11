{{- define "list" }}
    {{- nindent 0 "<ul>" }}
    {{- range .Cfg.TemplateData }}
        {{- printf "<li>%s</li>" . | nindent 4 }}
    {{- end }}
    {{- nindent 0 "</ul>" }}
{{- end }}
