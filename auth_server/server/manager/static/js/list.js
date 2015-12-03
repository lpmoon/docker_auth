$(document).ready(function() {
    $(".query").bind({
        click: function() {
            var query_id = $(this).attr("id");
            var lio = query_id.lastIndexOf("_");
            var idx = query_id.substring(lio + 1);
            var user = $("#value_" + idx).html();
            location.href = "query?user=" + user;    
        }
    });

    $(".delete").bind({
        click: function() {
            var query_id = $(this).attr("id");
            var lio = query_id.lastIndexOf("_");
            var idx = query_id.substring(lio + 1);
            var user = $("#value_" + idx).html();
            $.ajax({
                 cache: true,
                 type: "POST",
                 url:  "/deleteuser",
                 data: {"username": user}, 
                 error: function(request) {},
                 success: function(data) {
                    location.reload();                          
                 }
            });
        }
    }); 
});

