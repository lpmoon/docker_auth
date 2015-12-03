$(document).ready(function() {
    $("#submit_btn").bind({
        click: function() {
            var username = document.getElementById("username").value;    
            var password = document.getElementById("password").value;
            if (username == "") {
                $('.popover-user').popover('show');
                return;
            }    
            if (password == "") {
                $('.popover-pwd').popover('show');
                return;    
            }
            $.ajax({
                cache: true,
                type: "POST",
                url: "/adduser",
                data: {"username": username, "password": password}, // 你的formid
                error: function(request) {
                    // 添加
                },
                success: function(data) {
                    // 添加
                }
            });     
        }
    });
});
