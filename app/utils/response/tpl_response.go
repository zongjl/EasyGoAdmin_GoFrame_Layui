package response

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gview"
)

// 通用tpl响应
type TplResp struct {
	r   *ghttp.Request
	tpl string
}

// 返回一个tpl响应
func BuildTpl(r *ghttp.Request, tpl string) *TplResp {
	var t = TplResp{
		r:   r,
		tpl: tpl,
	}
	return &t
}

// 返回一个错误的tpl响应
func ErrorTpl(r *ghttp.Request) *TplResp {
	var t = TplResp{
		r:   r,
		tpl: "error/error.html",
	}
	return &t
}

// 输出页面模板附加自定义函数
func (resp *TplResp) WriteTpl(params ...gview.Params) error {
	return resp.r.Response.WriteTpl(resp.tpl, params...)
}
