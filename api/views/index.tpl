<!doctype html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Eyas Blog Page</title>
</head>
<body>
<div id="app">we are mode: {{.Mode}}</div>
{{if eq .ClientMode "dev"}}
<script src="http://localhost:3000/static/app.js"></script>
{{else}}
<script src="/static/app.js"></script>
{{end}}
</body>
</html>