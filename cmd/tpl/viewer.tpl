<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <style>
        body { display: flex; height: 100vh; margin: 0; overflow: hidden; }
        textarea { width: 50%; height: 100%; font-family: monospace; padding: 10px; box-sizing: border-box; border: none; outline: none; resize: none; white-space: pre; }
        .divider { width: 5px; background: #ccc; }
        iframe { width: 50%; height: 100%; border: none; }
    </style>
</head>
<body>
<label for="code"></label>
<textarea id="code">{{ . }}</textarea>
<div class="divider"></div>
<iframe id="preview"></iframe>
<script>
const textarea = document.getElementById('code');
const iframe = document.getElementById('preview');
const updatePreview = function() {
    const code = textarea.value;
    iframe.src = 'data:text/html;charset=utf-8,' + encodeURIComponent(code);
}();
</script>
</body>
</html>
