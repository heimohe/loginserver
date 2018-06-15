package controllers

import (

	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	"loginServer/common"

	"loginServer/models"

)

func init(){

	orm.RegisterDriver("mysql", orm.DRMySQL)

	idleConns  := beego.AppConfig.DefaultInt("Mysql.idleConns",2)

	maxOpenConns  := beego.AppConfig.DefaultInt("Mysql.maxOpenConns",3)

	dataSource := "root:123456@tcp(192.168.0.110:3306)/mydb?charset=utf8"
	
	orm.RegisterDataBase("default", "mysql", dataSource, idleConns, maxOpenConns)

	orm.Debug = true

}

type BaseController struct {

	beego.Controller

}

//校验token
func (b *BaseController) CheckToken() bool {

	token := b.GetString("AuthToken") 

	fmt.Println("request token:"+token)

	secret := common.GetSecretStr()

	reslut := common.TokenAuth(token,secret )

	fmt.Println("request reslut:"+reslut)

	if reslut == "" {

		b.Data["json"] = common.ErrorTokenInvalid

		return false 

	}

	catchToken := models.GetTokenFromRedis(reslut)

	if catchToken != token {

		b.Data["json"] = common.ErrorTokenInvalid

		return false

	}

	return true 

}

