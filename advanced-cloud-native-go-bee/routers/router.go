package routers

import (
	"learning/advanced-cloud-native-go-bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
