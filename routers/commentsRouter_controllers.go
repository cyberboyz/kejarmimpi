package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["relasi/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
