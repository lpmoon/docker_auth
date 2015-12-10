$(document).ready(function() {
    $("#submit_btn").bind({
        click: function() {
            var username = $("#username").val();    
            var imagename= $("#imagename").val();
            if (username == "") {
                $('.popover-user').popover('show');
                return;
            }    
            if (imagename == "") {
                $('.popover-img').popover('show');
                return;    
            }
                       
            var actions = [];
            $('input[type="checkbox"]:checked').each(function(){
                actions.push($(this).val());           
            }); 
            
            $.ajax({
                cache: true,
                type: "POST",
                dataType: "json",
                url: "/addauth",
                data: {"username": username, "imagename": imagename, "actions": actions}, // 你的formid
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
