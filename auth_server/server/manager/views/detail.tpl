<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/css/bootstrap.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/js/bootstrap.min.js" type="text/css" />
  <link rel="stylesheet" href="/static/js/jquery-2.1.4.min.js" type="text/css" />

</head>

<body>

<div class="panel panel-default">
  <!-- Default panel contents -->
  <div class="panel-heading">{{.user}}</div>
  <div class="panel-body">
    <pre>该面板中列出了当前用户对各个镜像的权限.
         1. 如果镜像名称为*, 表示该用户对所有镜像拥有某个权限
         2. 如果镜像名不为*, 表示该用户对特定的某个镜像拥有相应的权限
         3. 拉取权限为PULL, 推送权限为PUSH
         4. 用户可以点击操作按钮, 对该用户的权限做操作
    </pre>
  </div>

  <!-- Table -->
  <table class="table table-hover">
    <tr>
        <th>编号</th>
        <th>镜像名称</th>
        <th>权限</th>
    </tr>
    {{range $index, $value := .names}}
        <tr>
            <td>{{$index}}</td>
            <td>{{$value}}</td>
            <td><button type="button" class="btn btn-success">&nbsp查看&nbsp</button></td>
        </tr>
    {{end}}
  </table>
</div>

</body>
</html>
