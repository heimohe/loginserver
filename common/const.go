package common

type CommonError struct {

	Status   int    `json:"status"`

	Code     int    `json:"code"`
	
	Message  string `json:"message"`
	
	DevInfo  string `json:"dev_info"`
	
	MoreInfo string `json:"more_info"`

}

var (

	Error404          	= &CommonError{404, 404, "page not found", "page not found", ""}

	ErrorInputData    	= &CommonError{400, 10001, "数据输入错误", "客户端参数错误", ""}

	ErrorDatabase    	= &CommonError{500, 10002, "服务器错误", "数据库操作错误", ""}

	ErrorUserExits      = &CommonError{400, 10003, "用户信息已存在", "数据库记录重复", ""}

	ErrorAccountExits   = &CommonError{400, 10004, "用户名已存在", "数据库记录重复", ""}
	
	ErrorUserNotExits   = &CommonError{400, 10005, "用户信息不存在", "数据库记录不存在", ""}
	
	ErrorPwd        	= &CommonError{400, 10006, "用户信息不存在或密码不正确", "密码不正确", ""}
	
	ErrorUserOrPass   	= &CommonError{400, 10007, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确", ""}
	
	ErrorNoUserChange 	= &CommonError{400, 10008, "用户信息不存在或数据未改变", "数据库记录不存在或数据未改变", ""}
	
	ErrorInvalidUser  	= &CommonError{400, 10009, "用户信息不正确", "Session信息不正确", ""}
	
	ErrorOpenFile     	= &CommonError{500, 10010, "服务器错误", "打开文件出错", ""}
	
	ErrorWriteFile    	= &CommonError{500, 10011, "服务器错误", "写文件出错", ""}
	
	ErrorSystem       	= &CommonError{500, 10012, "服务器错误", "操作系统错误", ""}
	
	ErrorTokenInvalid   = &CommonError{400, 10013, "登录已过期", "验证token无效", ""}
	
	ErrorPermission   	= &CommonError{400, 10014, "没有权限", "没有操作权限", ""}

	ErrorRegistFailed   = &CommonError{400, 10015, "注册失败", "注册失败！", ""}

	ErrorOldPWd  	 	= &CommonError{400, 10015, "修改密码", "输入密的密码错误", ""}

)

type Success struct {

	Status   	string    		`json:"status"`

	Code     	int    			`json:"code"`
	
	Message  	interface{} 	`json:"message"`
	
	MoreInfo  	string 			`json:"more_info"`

}

var(

	SuccessRegist 		= &Success{"true",20001,nil,"注册成功"}

	SuccessLogin 		= &Success{"true",20002,nil,"登陆成功"}

	SuccessUpdatePwd 	= &Success{"true",20003,nil,"修改密码成功"}

)