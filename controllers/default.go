package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.EnableXSRF = true
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) ListFood() {
	jsonInfo := c.GetString("jsoninfo");
	if jsonInfo == "" {
			c.Ctx.WriteString("json info empty")
		}
	c.Ctx.WriteString(jsonInfo)
}

func (c *MainController) TestPost() {
	jsonInfo := c.GetString("jsoninfo");
	c.Ctx.WriteString(jsonInfo)
}
