<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style>
  </style>
</head>

<body>
  <header>
    AuthManagerWeb
  </header>

{{range $index, $value := .names}}
    {{$index}}, {{$value}}
{{end}}

</body>
</html>
