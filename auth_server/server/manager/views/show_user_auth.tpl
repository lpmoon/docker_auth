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
  <script type="text/javascript" src="/static/js/detail.js"></script>
  </script>
</head>

<body>
<div class="container-fluid">
<div class="row">

{{template "/left_bar.tpl" .}}

<div class="col-sm-offset-2 panel panel-info">
  <!-- Default panel contents -->
  <div class="panel-heading " id="user">{{.user}}</div>
  <div class="panel-body lead">
     该面板中列出了当前用户对各个镜像的权限.
<ol>
<li>如果镜像名称为*, 表示该用户对所有镜像拥有某个权限</li>
<li>如果镜像名不为*, 表示该用户对特定的某个镜像拥有相应的权限</li>
<li>拉取权限为PULL, 推送权限为PUSH</li>
<li>用户可以点击操作按钮, 对该用户的权限做操作</li>
</ol>
  </div>

  <!-- Table -->
  <table class="table table-hover">
    <tr>
        <th>编号</th>
        <th>镜像名称</th>
        <th>权限</th>
        <th>修改</th>
        <th>删除</th>
    </tr>
        {{range $idx, $value := $.detail}}
            <tr>
                <td>{{$idx}}</td>
                <td id="img_{{$idx}}">{{index $value 0}}</td>
                <td>
                    {{if eq "3" (index $value 1)}}
                        <!-- 显示 -->
                        <span class="label label-primary btn-xs" id="pull_text_{{$idx}}">&nbspPull&nbsp</span>
                        <button type="button" class="btn btn-danger btn-xs hidden control" id="pull_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-trash" aria-hidden="true"></span> Delete 
                        </button>
                        <!-- 显示 -->
                        <span class="label label-success btn-xs" id="push_text_{{$idx}}">&nbspPush&nbsp</span>
                        <button type="button" class="btn btn-danger btn-xs hidden control" id="push_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-trash" aria-hidden="true"></span> Delete
                        </button>
                    {{else if eq "2" (index $value 1)}}
                        <!-- 隐藏 -->
                        <span class="label label-primary btn-xs hidden" id="pull_text_{{$idx}}">&nbspPull&nbsp</span>
                        <button type="button" class="btn btn-success btn-xs hidden control" id="pull_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-plus" aria-hidden="true"></span> Add 
                        </button>
                        <!-- 显示-->
                        <span class="label label-success btn-xs" id="push_text_{{$idx}}">&nbspPush&nbsp</span>
                        <button type="button" class="btn btn-danger btn-xs hidden control" id="push_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-trash" aria-hidden="true"></span> Delete
                        </button>
                    {{else if eq "1" (index $value 1)}}
                        <!-- 显示-->
                        <span class="label label-primary btn-xs" id="pull_text_{{$idx}}">&nbspPull&nbsp</span>
                        <button type="button" class="btn btn-danger btn-xs hidden control" id="pull_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-trash" aria-hidden="true"></span> Delete
                        </button>
                        <!-- 隐藏-->
                        <span class="label label-success btn-xs hidden" id="push_text_{{$idx}}">&nbspPush&nbsp</span>
                        <button type="button" class="btn btn-success btn-xs hidden control" id="push_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-plus" aria-hidden="true"></span> Add
                        </button>
                    {{else}}
                        <!-- 隐藏-->
                        <span class="label label-primary btn-xs hidden" id="pull_text_{{$idx}}">&nbspPull&nbsp</span>
                        <button type="button" class="btn btn-success btn-xs hidden control" id="pull_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-plus" aria-hidden="true"></span> Add
                        </button>
                        <!-- 隐藏-->
                        <span class="label label-success btn-xs hidden" id="push_text_{{$idx}}">&nbspPush&nbsp</span>
                        <button type="button" class="btn btn-success btn-xs hidden control" id="push_btn_{{$idx}}">
                              <span class="glyphicon glyphicon-plus" aria-hidden="true"></span> Add 
                        </button>
                    {{end}}
                </td>
                <td>
                {{if eq "0" (index $value 2)}}
                    <button type="button" class="btn btn-primary btn-xs" disabled="disabled">
                        修改
                    </button>
                {{else}}
                    <button type="button" class="btn btn-primary btn-xs modifybtn" id="modify_btn_{{$idx}}">
                        修改
                    </button>
                {{end}}
                </td>
                <td>
                {{if eq "0" (index $value 2)}}
                    <button type="button" class="btn btn-primary btn-xs" disabled="disabled">
                        删除  
                    </button>
                {{else}}
                    <button type="button" class="btn btn-primary btn-xs deletebtn" id="delete_btn_{{$idx}}">
                        删除
                    </button>
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
