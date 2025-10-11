{{- define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    {{- template "cssWebLink" (deepCopy . | merge (dict "Indent" 4)) }}
    <style>
        body { padding: 1.5em; }
        table { border-collapse: collapse; margin-bottom: 2em; }
        th, td { border: 1px solid black; padding: .5em; }
    </style>
    {{- template "css" . }}
</head>
<body>
    {{- template "table" (deepCopy . | merge (dict "Indent" 4)) }}
</body>
</html>
{{- end }}
