String.prototype.startWith = function(str) {
    if (str == null || str== "" || this.length == 0 || str.length > this.length) {
        return false;
    }

    if (this.substr(0, str.length) == str) {
        return true;
    } else {
        return false;
    }
}

$(document).ready(function() {

    $(".deletebtn").bind({
        click: function() {
            var btn_id = $(this).attr("id");
            var lio = btn_id.lastIndexOf("_");
            var idx = btn_id.substring(lio + 1);
            
            var name = $("#img_" + idx).html();
            var user = ""
            if ($("#user").length == 0) {
                user = $("#user_" + idx).html();
            } else {
                user = $("#user").html();
            }

            $.ajax({
                method: "POST",
                url: "/deleteauth",
                data: {"user": user, "name": name}
            }).done(function(data){
                location.reload()
            }).fail(function(data){
                // 处理
            });
        }
    });

    // 修改按钮事件绑定
    $(".modifybtn").bind({
        click: function() {
            var btn_id = $(this).attr("id");
            var lio = btn_id.lastIndexOf("_");
            var idx = btn_id.substring(lio + 1);
            if ($(this).hasClass("modifybtn")) {
                // 隐藏的label和button显示
                
                // 1) pull text
                var pull_text_id = "pull_text_" + idx;
                if ($("#" + pull_text_id).hasClass("hidden")) {
                    $("#" + pull_text_id).removeClass("hidden");
                }
                // 2) pull btn
                var pull_btn_id = "pull_btn_" + idx;
                if ($("#" + pull_btn_id).hasClass("hidden")) {
                    $("#" + pull_btn_id).removeClass("hidden");
                }
                // 3) push text
                var push_text_id = "push_text_" + idx;
                if ($("#" + push_text_id).hasClass("hidden")) {
                    $("#" + push_text_id).removeClass("hidden");
                }
                // 4) push btn 
                var push_btn_id = "push_btn_" + idx;
                if ($("#" + push_btn_id).hasClass("hidden")) {
                    $("#" + push_btn_id).removeClass("hidden");
                }
        
                // 不该显示的label和button隐藏

                $(this).removeClass("modifybtn");
                $(this).addClass("cancelbtn");
                
                // 修改文本
                $(this).html("取消");
                
                return
            } else if ($(this).hasClass("cancelbtn")) {
                // 根据pull 和 push btn是btn-success还是btn-danger, 判断是要隐藏还是显示按钮
                var pull_btn_id = "pull_btn_" + idx;
                var push_btn_id = "push_btn_" + idx;
           
                var pull_btn = $("#" + pull_btn_id); 
                if (pull_btn.hasClass("btn-success")) {
                    // 隐藏btn
                    $("#" + "pull_text_" + idx).addClass("hidden");
                    pull_btn.addClass("hidden");
                } else if (pull_btn.hasClass("btn-danger")) {
                    // 隐藏btn和隐藏label
                    pull_btn.addClass("hidden");
                    // $("#" + "pull_text_" + idx).addClass("hidden");
                } else {
                
                }
           
                var push_btn = $("#" + push_btn_id);
                if (push_btn.hasClass("btn-success")) {
                    // 隐藏btn
                    push_btn.addClass("hidden");
                    $("#" + "push_text_" + idx).addClass("hidden");
                } else if (push_btn.hasClass("btn-danger")) {
                    // 隐藏btn和隐藏label
                    push_btn.addClass("hidden");
                    // $("#" + "push_text_" + idx).addClass("hidden");
                } else {
                
                }
                $(this).removeClass("cancelbtn");
                $(this).addClass("modifybtn");
                $(this).html("修改"); 
            }
        }
    });

    $(".control").bind({
        click: function() {
            var btn_id = $(this).attr("id");
            var lio = btn_id.lastIndexOf("_");
            var idx = btn_id.substring(lio + 1);
            // 获取镜像名称
            var name = $("#img_" + idx).html();
            if (name == "*") {
                name = "";    
            }
            // 获取用户名
            var user = ""
            if ($("#user").length == 0) {
                user = $("#user_" + idx).html();
            } else {
                user = $("#user").html();
            }
            // 获取删除还是添加
            
            var mtype;
            if ($(this).hasClass("btn-success")) {
                mtype = 1;
            } else if ($(this).hasClass("btn-danger")) {
                mtype = 2;
            }

            var ispull = btn_id.startWith("pull");
            $.ajax({
                method: "POST",
                url: "/modify",
                data: {"user": user, "name": name, "mtype": mtype, "ispull": ispull}
            }).done(function(data){
                // 处理
                if (mtype == 1) { // 添加
                    // 添加按钮变成删除按钮
                    $("#" + btn_id).removeClass("btn-success");
                    $("#" + btn_id).addClass("btn-danger");
                    $("#" + btn_id).html('<span class="glyphicon glyphicon-trash" aria-hidden="true"></span> Delete');
                } else { // 删除
                    // 删除按钮变成添加按钮
                    $("#" + btn_id).removeClass("btn-danger");
                    $("#" + btn_id).addClass("btn-success");
                    $("#" + btn_id).html('<span class="glyphicon glyphicon-plus" aria-hidden="true"></span> Add');
                }
            }).fail(function(data){
                // 处理
            });
            
        }
    });
    
});
