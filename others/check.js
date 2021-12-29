while('block' == $(".messager-window").css("display")) {
    console.log("alert");
    $(".l-btn-small").click();
    $(".dialog-button").click();
    $(".vjs-play-control").click();
    break;
}

for (var i = 0; i < 60; i++) {
    setTimeout(() => {
        console.log("check");
        if('block' == $(".messager-window").css("display")) {
            console.log("alert");
            $(".l-btn-small").click();
            $(".dialog-button").click();
            $(".vjs-play-control").click();
        }
    }, 10000);
}

window.setInterval(function(){
    console.log("check");
        if('block' == $(".messager-window").css("display")) {
            console.log("alert");
            $(".l-btn-small").click();
            $(".dialog-button").click();
            $(".vjs-play-control").click();
        }
}, 10000);     

com.insigma.ajax({
    url:"/zxpx/tec/play/aj/loose",
    dataType:'json',
    data:{_method:'PUT'
        ,courseid:'CS2020030000008227'
        ,coursewareid:'CS2020030000008227CW2020030000021225'
        ,reqtdate:new Date().getTime()
        ,startWatchTime:p.currentTime()
        ,watchSession:'CS2020030000008227CW202003000002122596674B33CB9D46439A052F71346EC8E4'
        ,reqt:$('#e8888b74d1229efec6b4712e17cb6b7a_e').data('reqt')},
    type:'post',
    success:function(data){
        if(data.closeModal){
            p.pause();
            $.messager.alert({title:'错误',msg:data.resMsg,onClose:function(){
                window.opener=null;window.open('','_self');window.close();
            }});
            return;
        }
        //请求成功
        if(data.resData.reqt){
            $('#e8888b74d1229efec6b4712e17cb6b7a_e').data('reqt',data.resData.reqt);
            if(data.resData.progress){
                console.log(data.resData.progress);
                //clw add 20210809
                render2(data.resData.progress);
                if(data.resData.finishit && data.resData.finishit =='1'){
                    $('.learnpercent').html('<span >学习进度：<span>已完成</span></span>');
                } 
                /* render2(data.resData.progress); */
                //end
            }
        }else{
            p.pause();
            $.messager.alert({title:'错误',msg:'服务出错，请刷新重试！'});
            return;
        }
        
    }
    
});