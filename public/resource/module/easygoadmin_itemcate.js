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
 * 栏目管理
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
              {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '栏目名称', align: 'left'}
            , {field: 'itemName', width: 200, title: '所属站点', align: 'center'}
            , {field: 'pinyin', width: 150, title: '拼音(全)', align: 'center'}
            , {field: 'code', width: 100, title: '拼音(简)', align: 'center'}
            , {field: 'is_cover', width: 100, title: '有无封面', align: 'center', templet(d) {
                    if (d.is_cover == 1) {
                        // 有封面
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">有封面</span>';
                    } else {
                        // 无封面
                        return '<span class="layui-btn layui-btn-danger layui-btn-xs">有封面</span>';
                    }
                }}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                if (d.status == 1) {
                    // 在用
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">在用</span>';
                } else {
                    // 停用
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">停用</span>';
                }
            }}
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'note', width: 200, title: '备注', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("栏目");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
