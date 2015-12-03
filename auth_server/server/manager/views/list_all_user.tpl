<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/css/bootstrap.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/css/custom.css" type="text/css" />
  <script src="/static/js/jquery-2.1.4.min.js" type="text/javascript" ></script>
  <script src="/static/js/bootstrap.min.js" type="text/javascript" ></script>
  <script src="/static/js/list.js" type="text/javascript"></script>
</head>

<body>
<div class="container-fluid">
<div class="row">
{{template "/left_bar.tpl" .}}

<div class="col-sm-offset-2 panel panel-info">
  <!-- Default panel contents -->
  <div class="panel-heading ">所有用户</div>
  <div class="panel-body lead">
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
            <td id="value_{{$index}}">{{index $value 0}}</td>
            <td>
                <button type="button" class="btn btn-success query" id="query_btn_{{$index}}">&nbsp查看&nbsp</button>
            </td>
            <td>
                {{if eq "0" (index $value 1)}}
                <button type="button" disabled="disabled" class="btn btn-danger delete" id="delete_btn_{{$index}}">&nbsp删除&nbsp</button>
                {{else}}
                <button type="button" class="btn btn-danger delete" id="delete_btn_{{$index}}">&nbsp删除&nbsp</button>
                {{end}}
            </td>

        </tr>
    {{end}}
  </table>
</div>
</div>
</div>
</body>
</html>
