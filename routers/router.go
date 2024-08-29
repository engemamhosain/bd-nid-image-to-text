// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"strings"
	"tl_mlkit/auth"
	"tl_mlkit/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	apiVersion, _ := beego.AppConfig.String("apiVersion")
	versions := strings.Split(apiVersion, ",")
	for i, s := range versions {
		fmt.Println(i, s)
		ns := beego.NewNamespace("/mlkit/"+s,
			beego.NSNamespace("/object",
				beego.NSInclude(
					&controllers.ObjectController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSBefore(auth.Check),
			beego.NSNamespace("/reader",
				beego.NSInclude(
					&controllers.MlkitController{},
				),
			),
		)
		beego.AddNamespace(ns)
	}

}
