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
  <div class="panel-heading">所有用户</div>
  <div class="panel-body">
    <p>该面板中列出了鉴权系统中的所有用户,可以进行权限查看,删除用户等基本操作.</p>
  </div>

  <!-- Table -->
  <table class="table table-hover">
    <tr>
        <th>编号</th>
        <th>用户名</th>
        <th>查看</th>
        <th>删除</th>
    </tr>
    {{range $index, $value := .names}}
        <tr>
            <td>{{$index}}</td>
            <td>{{$value}}</td>
            <td><button type="button" class="btn btn-success">&nbsp查看&nbsp</button></td>
            <td><button type="button" class="btn btn-danger">&nbsp删除&nbsp</button></td>
        </tr>
    {{end}}
  </table>
</div>

</body>
</html>
