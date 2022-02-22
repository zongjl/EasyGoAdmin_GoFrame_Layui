// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 布局管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'itemName', width: 150, title: '页面编号', align: 'center'}
            , {field: 'locDesc', width: 200, title: '页面位置描述', align: 'center'}
            , {field: 'type', width: 100, title: '类型', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 新闻资讯
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 其他
                    cls = "layui-btn-danger";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'typeDesc', width: 350, title: '类型描述', align: 'center'}
            , {field: 'image', width: 100, title: '图片路径', align: 'center', templet: function (d) {
                var imageStr = "";
                if (d.imageUrl) {
                    imageStr = '<a href="' + d.imageUrl + '" target="_blank"><img src="' + d.imageUrl + '" height="26" /></a>';
                }
                return imageStr;
              }
            }
            , {field: 'sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'createUserName', width: 100, title: '添加人', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateUserName', width: 100, title: '更新人', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("布局");

    }
});
