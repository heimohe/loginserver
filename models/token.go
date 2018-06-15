package models

import(

	"time"

	"fmt"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

)

type UserToken struct {

	AcountId 	string 	`json:"acount_id"`	//用户名

	Token 		string 	`json:"token"`			//用户token
	
	CreateTime	string  `json:"create_time"`    //token创建时间

	ExpressTime string  `json:"express_time"`   //token有效期

}

type TokenData  struct {

	AcountId 	string 	`json:"acount_name"`

	Secret		string	`json:"secret"`

}

type AuthToken struct{
		
	Token 			string

	Authrization 	string

}

//创建用户token
func NewUserToken(r *TokenData, token string, expressTime string) (u *UserToken, err error) {

	createDate := time.Now().Format("2006-06-08 15:04:05")

	if err != nil {

		return nil, err

	}

	user := UserToken{

		AcountId:       r.AcountId,

		Token:       	token,

		ExpressTime:  	expressTime,

		CreateTime: 	createDate,

	}

	return &user, nil

}

func (u *UserToken) SaveToRedis() bool {

	rds := RedisPool.Get()

	defer rds.Close()

	tokenKey := u.AcountId

	tokenValue := u.Token

	ok,err := rds.Do("SET",tokenKey,tokenValue)

	if err!=nil {

		fmt.Println("statues err:",err)

		return false

	}

	fmt.Println("statues:",ok)
	
	return true

}

//插入或更新用户token
func (u *UserToken) InsertToDB() bool {

	o := orm.NewOrm()

	o.Using("default")

	_, err := o.Raw("replace into token_tb (acount_id,token,create_at,express_time) values (?,?,?,?)", u.AcountId, u.Token, u.CreateTime, u.ExpressTime).Exec()

	if err != nil {

		fmt.Println(err)

		return false

	} 

	fmt.Println("insert ok")

	return true

}


