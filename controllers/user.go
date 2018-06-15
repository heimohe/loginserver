package controllers

import (

	_"encoding/json"

	"loginServer/models"

	"loginServer/common"

)

// Operations about Users
type UserController struct {

	BaseController

}

// @Title CreateUser
// @Description regist users
// @Param	body		body 	models.RegistData	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /regist [post]
func (u *UserController) Regist() {
	
	register := models.RegistData{}

	if err := u.ParseForm(&register); err != nil {
	
		u.Data["json"] =  common.ErrorInputData

		u.ServeJSON()

		return

	}

	user, err := models.CreateUser(&register)

	if err != nil {
	
		u.Data["json"] = common.ErrorSystem

		u.ServeJSON()

		return

	}

	ip := u.Ctx.Request.RemoteAddr

	user.Ip  = ip

	accountId := register.AcountId 

	if ok, _ := user.CheckIsExit(accountId) ; ok {

		u.Data["json"] = common.ErrorAccountExits

		u.ServeJSON()

		return

	}

	seccess := user.InsertToDB()

	if !seccess {

		u.Data["json"] = common.ErrorRegistFailed 

		u.ServeJSON()

		return

	}

	u.Data["json"] = common.SuccessRegist 

	u.ServeJSON()

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {

	u.Data["json"] = "logout success"

	u.ServeJSON()

}

// @Title login
// @Description Logs out current logged in user session
// @Param	AcountId		query 	string	true		"The login_id for login"
// @Param	Password		query 	string	true		"The password for login"
// @Success 200 {string} logout success
// @router /login [get]
func (u *UserController) Login() {

	login := models.LoginData{}

	if err := u.ParseForm(&login); err != nil {
		
		u.Data["json"] = common.ErrorInputData

		u.ServeJSON()

		return

	}

	println("Login AcountId:"+login.AcountId)

	user := models.User{}

	if _, err := user.CheckIsExit(login.AcountId); err != nil {
		
		u.Data["json"] = common.ErrorUserNotExits

		u.ServeJSON()

		return

	}
	
	if ok, err := user.CheckPasswd(login.Password); err != nil {

		u.Data["json"] = common.ErrorSystem

		u.ServeJSON()

		return

	} else if !ok {

		u.Data["json"] =  common.ErrorPwd

		u.ServeJSON()

		return

	}

	tokenStr, _ := login.GenerToken()

	if tokenStr == "" {

		u.Data["json"] =  common.ErrorDatabase

		u.ServeJSON()

		return

	}

	loginSuccess := common.SuccessLogin

	loginSuccess.Message = &models.LoginDataComfirm{
		
		Token: 				tokenStr,

		AcountId: 			login.AcountId,
			
		UserInfo:			&user,

	}

	u.Data["json"] = loginSuccess
	
	u.ServeJSON()

}

// @Title update user password
// @Description update user password
// @Param	body		body 	models.MotifyPwd	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /updatepwd [post]
func (u *UserController) UpdatePwd() {

	tokenOk := u.CheckToken()

	if !tokenOk {

		u.ServeJSON()

	 	return

	}

	motifyPwd := models.MotifyPwd{}

	if err := u.ParseForm(&motifyPwd); err != nil {
		
		u.Data["json"] = common.ErrorInputData

		u.ServeJSON()

		return

	}

	accountId := motifyPwd.AcountId 

	user := models.User{}

	ok, _ := user.CheckIsExit(accountId)

	if !ok {

		u.Data["json"] = common.ErrorUserNotExits

		u.ServeJSON()

		return

	}

	pwdOk := user.ParesePwd(motifyPwd.OldPwd) 

	if !pwdOk {

		u.Data["json"] = common.ErrorOldPWd

		u.ServeJSON()

		return

	}

	updateOk := user.UpdatePwd(motifyPwd.NewPwd) 

	if !updateOk {

		u.Data["json"] = common.ErrorDatabase

		u.ServeJSON()

		return

	}

	u.Data["json"] = common.SuccessUpdatePwd

	u.ServeJSON()

	return

}
