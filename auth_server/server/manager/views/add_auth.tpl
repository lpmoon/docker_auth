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
  <script type="text/javascript" src="/static/js/addauth.js"> </script>
</head>

<body>
<div class="container-fluid">
<div class="row">

{{template "/left_bar.tpl" .}}

<div class="col-sm-offset-2 panel panel-info">
  <!-- Default panel contents -->
  <div class="panel-heading ">添加权限</div>
  <div class="panel-body lead">
       该面板中用于为用户添加新的权限.
  </div>

<form class="form-horizontal" role="form" id="adduserform">
    <div class="form-group">
        <label for="firstname" class="col-sm-1 control-label">用户名</label>
        <div class="col-sm-5">
            <input type="text" class="form-control popover-user" title="警告" data-container="body" data-toggle="popover" data-placement="right" data-content="必须输入用户名" id="username"placeholder="请输入名字">
        </div>
    </div>

    <div class="form-group">
        <label for="lastname" class="col-sm-1 control-label">镜像名</label>
        <div class="col-sm-5">
            <input type="text" class="form-control popover-img" title="警告" data-container="body" data-toggle="popover" data-placement="right" data-content="必须输入镜像名" id="imagename" placeholder="请输入镜像名">
        </div>
    </div>

    <div class="form-group">
        <div class="col-sm-offset-1 col-sm-5">
            <label class="checkbox-inline">
                <input type="checkbox" id="pull_check" value="pull"> Pull
            </label>

            <label class="checkbox-inline">
                <input type="checkbox" id="push_check" value="push"> Push
            </label>
        </div>
    </div>

    <div class="form-group">
        <div class="col-sm-offset-1 col-sm-10">
             <button type="button" id="submit_btn" class="btn btn-default">添加</button>
        </div>
    </div>
</form>

</div>
</div>
</div>
</body>
</html>
