package routers

import (
	"pro2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/block", &controllers.BlockController{})
	beego.Router("/createBlock", &controllers.BlockController{}, "post:CreateBlock")
	beego.Router("/Peer", &controllers.PeerController{})
	beego.Router("/add", &controllers.PeerController{}, "post:AddBlock")

}
