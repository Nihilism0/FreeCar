// Code generated by hertz generator.

package User

import (
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/router/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	return middleware.CommentMW()
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __dminget_llusersMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __dminchangepasswordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __dmingetsomeusersMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __dmindeleteuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __dminupdateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	return middleware.CommentWithoutJWT()
}

func __dminloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _upload_vatarMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordMw() []app.HandlerFunc {
	return middleware.CommentWithoutJWT()
}

func _user0Mw() []app.HandlerFunc {
	return middleware.CommentMW()
}
