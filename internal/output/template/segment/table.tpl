{{- define "tableInner" }}
<table{{- if .Cfg.TableClass }} class="{{- .Cfg.TableClass }}"{{- end }}>
    <tbody>
    {{- range .Table.Row }}
        <tr>
            {{- range .Cell }}
                {{- $tag := "td" }}
                {{- if eq .Type "key" }}
                    {{- $tag = "th" }}
                {{- end }}
            {{ printf "<%s" $tag }}{{- if gt .RowSpan 1 }} rowspan="{{ .RowSpan }}" {{- end }}{{- if gt .ColSpan 1 }} colspan="{{ .ColSpan }}" {{- end }}>{{ .Value }}{{- printf "</%s>" $tag }}
            {{- end }}
        </tr>
    {{- end }}
    </tbody>
</table>
{{- end }}

{{- define "table" }}
    {{- $indent := default 0 .Indent }}
    {{- $data := . }}
    {{- range .Table }}
        {{- render "tableInner" (dict "Table" . "Cfg" $data.Cfg) | indent $indent}}
    {{- end }}
{{- end }}
