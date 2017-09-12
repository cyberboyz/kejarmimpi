package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"] = append(beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"] = append(beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"] = append(beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"] = append(beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"] = append(beego.GlobalControllerRouter["interest/kejarmimpi/controllers:InterestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
