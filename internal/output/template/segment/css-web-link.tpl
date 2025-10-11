{{- define "cssWebLink" }}
    {{- if .Cfg.CSSWebLink }}
        {{- $indent := default 0 .Indent }}
        {{- printf `<link href="%s" rel="stylesheet" />` .Cfg.CSSWebLink | nindent $indent }}
    {{- end }}
{{- end }}
