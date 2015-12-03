<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="/static/css/bootstrap.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css" type="text/css" />
  <link rel="stylesheet" href="/static/css/custom.css" type="text/css" />
  <script type="text/javascript" src="/static/js/jquery-2.1.4.min.js" ></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js" ></script>
</head>

<body>
<div class="container-fluid">
<div class="row">

{{template "/left_bar.tpl" .}}

<div class="col-sm-offset-2 panel panel-info">
  <!-- Default panel contents -->
  <div class="panel-heading ">说明</div>
  <div class="panel-body lead">
       该项目基于开源项目<code>docker_auth</code>修改,添加了对于用户和权限的控制界面.避免了修改配置文件而带来的服务重启的问题.
       后端的数据存储在<code>MongoDB</code>中.
  </div>


</div>
</div>
</div>
</body>
</html>
