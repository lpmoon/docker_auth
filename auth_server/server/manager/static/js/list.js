$(document).ready(function(){
    $(".query").bind({
        click: function(){
                  var query_id = $(this).attr("id");
                  var lio = query_id.lastIndexOf("_");
                  var idx = query_id.substring(lio + 1);
                  var user = $("#value_" + idx).html();
                  location.href = "query?user=" + user;    
        }
    }); 
});

