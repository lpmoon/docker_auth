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
  <script type="text/javascript" src="/static/js/adduser.js" ></script>
</head>

<body>
<div class="container-fluid">
<div class="row">

{{template "/left_bar.tpl" .}}

<div class="col-sm-offset-2 panel panel-info">
  <!-- Default panel contents -->
  <div class="panel-heading ">添加用户</div>
  <div class="panel-body lead">
       该面板中用于创建新的用户,需要填入用户名和用户密码.用户密码需要按照如下步骤进行加密,
        <ol>
            <li>安装apache2-utils(比如在ubuntu上运行, <code>sudo apt-get install apache2-utils</code>)</li>
            <li>使用命令<code>htpasswd -bBc myfile myuser mypassword</code>即可生成密码.密码存放在当前目录的myfile文件中</li>
            <li>打开myfile文件取冒号后面的数据当做密码即可</li>
        </ol>
  </div>

<form class="form-horizontal" role="form" id="adduserform">
   <div class="form-group">
      <label for="firstname" class="col-sm-1 control-label">用户名</label>
      <div class="col-sm-5">
         <input type="text" class="form-control popover-user" title="警告" data-container="body" data-toggle="popover" data-placement="right" data-content="必须输入用户名" id="username"placeholder="请输入名字">
      </div>
   </div>
   <div class="form-group">
      <label for="lastname" class="col-sm-1 control-label">密码</label>
      <div class="col-sm-5">
         <input type="text" class="form-control popover-pwd" title="警告" data-container="body" data-toggle="popover" data-placement="right" data-content="必须输入密码" id="password" placeholder="请输入密码">
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
