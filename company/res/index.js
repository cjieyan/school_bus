layui.use('layer', function(){ //独立版的layer无需执行这一句
    var $ = layui.jquery, layer = layui.layer; //独立版的layer无需执行这一句

    //触发事件
    var active = {
        notice: function(){
            //示范一个公告层
            layer.open({
                type: 1
                ,title: false //不显示标题栏
                ,closeBtn: 2
                ,area: '740px;'
                ,shade: 0.8
                ,id: 'LAY_layuipro' //设定一个id，防止重复弹出
                ,moveType: 1 //拖拽模式，0或者1
                ,content: '<div class="pop-main"">' +
                    '<div class="pop-left"> <img src="./res/static/images/logo-pop.png" /></div>'+
                    '<div class="pop-right">' +
                    '<p class="pop-title">申请试用</p>' +
                    '<form class="layui-form" action="">' +
                    '  <div class="layui-form-item">' +
                    '    <div class="layui-input-block">' +
                    '      <input type="text" name="title" lay-verify="title" autocomplete="off" placeholder="姓名" class="layui-input">' +
                    '    </div>' +
                    '  </div>' +
                    '  <div class="layui-form-item">' +
                    '    <div class="layui-input-block">' +
                    '      <input type="text" name="username" lay-verify="required" lay-reqtext="联系电话是必填项，岂能为空？" placeholder="联系电话" autocomplete="off" class="layui-input">' +
                    '    </div>' +
                    '  </div>' +
                    '  <div class="layui-form-item layui-form-text">' +
                    '    <div class="layui-input-block">' +
                    '      <textarea placeholder="备注" class="layui-textarea"></textarea>' +
                    '    </div>' +
                    '  </div>' +
                    '  <div class="layui-form-item">' +
                    '    <div class="layui-input-block">' +
                    '      <button type="submit" class="layui-btn layui-btn-submit" lay-submit="" lay-filter="demo1">立即提交</button>' +
                    '    </div>' +
                    '  </div>' +
                    '</form>' +
                    '</div>' +
                    '</div>'
                ,success: function(layero){
                    var btn = layero.find('.layui-input-block');
                    btn.find('.layui-btn-submit').click(function(){
                        console.log(".....")
                        layer.closeAll();
                    });
                }
            });
        }
        ,offset: function(othis){
            var type = othis.data('type')
                ,text = othis.text();
            layer.open({
                type: 1
                ,offset: type //具体配置参考：http://www.layui.com/doc/modules/layer.html#offset
                ,id: 'layerDemo'+type //防止重复弹出
                ,content: '<div style="padding: 20px 100px;">'+ text +'</div>'
                ,btn: '关闭全部'
                ,btnAlign: 'c' //按钮居中
                ,shade: 0 //不显示遮罩
                ,yes: function(){
                    layer.closeAll();
                }
            });
        }
    };

    $('.layui-btn').on('click', function(){
        var othis = $(this), method = othis.data('method');
        active[method] ? active[method].call(this, othis) : '';
    });
});
