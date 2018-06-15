package models

import (

	"strconv"

	"time"

	"fmt"

	"crypto/rand"

	"io"

	"golang.org/x/crypto/scrypt" 

	"github.com/astaxie/beego/orm"

	"loginServer/common"

)

const pwdHashBytes = 64

type User struct {

	//Id		  			string	 `bson:"accountId" json:"account_id"`		  	//账户id

	AccountId		  	string	 `bson:"accountId" json:"account_id"`	  	//账户名

	CreateAt          	string     `bson:"createAt" json:"create_at,omitempty"`	//创建时间

	UpdateAt          	string    `bson:"updateAt" json:"update_at,omitempty"`	//更新时间

	DeleteAt          	string    `bson:"deleteAt" json:"delete_at"`		 	//注销时间

	Nickname           	string    `bson:"nickname" json:"nickname"`			 	//用户昵称

	Gender            	string    `bson:"gender"   json:"gender"`			 	//性别

	Password          	string    `bson:"password" json:"password,omitempty"`  	//密码

	Email              	string    `bson:"email" json:"email"`				 	//邮箱

	Position           	string    `bson:"position" json:"position"`			 	//位置
	
	Locale             	string    `bson:"locale" json:"locale"`				 	//国籍

	HeadImgUrl         	string    `bson:"headImgUrl" json:"head_url"`    		//头像地址        

	Mobile             	string    `bson:"mobile" json:"mobile"`      		 	//电话

	Salt 			   	string    `bson:"salt" json:"salt"`					 	//加密 “盐”

	Secret 			   	string    `bson:"secret" json:"secret"`				 

	Role 			   	string    `bson:"role" json:"role"`
	
	Ip					string	  `bson:"ip" json:"ip"`

}

type RegistData struct {

	AcountId string

	Password string

	Nickname string
	
	Email 	string

}

//登陆数据
type LoginData struct {

	AuthToken	string 		`json:"auth_token"`

	AcountId 	string 		`json:"login_id"`
	
	Password  	string 		`json:"password"`

}

type LoginDataComfirm struct {

	Token 				string			`json:"token"`

	AcountId 			string			`json:"login_id"`

	UserInfo			*User			`json:"user_info"`

}

type MotifyPwd struct {

	AuthToken			string 			

	AcountId 			string			

	OldPwd				string			

	NewPwd				string			

}

func createSalt() (salt string,err error) {

	buf := make([]byte,pwdHashBytes)

	if _, err := io.ReadFull(rand.Reader,buf ); err !=nil {

		return "" , err

	}

	return fmt.Sprintf("%x", buf), nil

}

func createPassHashs(pwd string,salt string) (hash string,err error) {

	h, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, pwdHashBytes)

	if err != nil {
	
		return "", err

	}

	return fmt.Sprintf("%x", h), nil

}

//校验密码
func (u *User) CheckPasswd(pwd string ) (ok bool,err error) {

	hashpwd, err := createPassHashs(pwd,u.Salt)

	if err != nil {

		return false, err 	

	}

	return u.Password == hashpwd,nil

}

//查询用户
func (u *User) CheckIsExit(accountId string ) (bool,error) {

	println("accountId:"+accountId)

	o := orm.NewOrm()

	o.Using("default")

	err := o.Raw("select * from user_tb where account_id = ?", accountId).QueryRow(&u)

	if err != nil {

		fmt.Println(err)

		return false,nil

	} else {

		return true,nil

	}

}

//用户注册
func CreateUser(register *RegistData) (u *User, err error) {

	salt, err := createSalt()

	registDate := time.Now().Format("2006-01-02 15:04:05")

	if err != nil {

		return nil,err

	}

	hashP, err := createPassHashs(register.Password,salt)

	if err != nil {

		return nil,err

	}

	user := User{

		Nickname: 		register.Nickname,

		Password: 		hashP,

		Salt:			salt,

		CreateAt:		registDate,

		AccountId:  	register.AcountId,

		Email:			register.Email,

	}

	return &user,nil

}

func (u *User) InsertToDB() bool {

	o := orm.NewOrm()
	
	o.Using("default")
	
	_, err := o.Raw("insert into user_tb (nickname,password,email,create_at,role_id,account_id,secret,salt,gender,ip_addr) values (?,?,?,?,?,?,?,?,?,?)", u.Nickname, u.Password, u.Email, u.CreateAt, 1, u.AccountId, u.Secret, u.Salt,1,u.Ip).Exec()

	if err != nil {

		fmt.Println(err)

		return false

	} else {
		
		fmt.Println("insert ok")

		return true

	}

}

func (login *LoginData) GenerToken() (string ,error) {

	tokenData := TokenData{

		AcountId : login.AcountId,

	}

	secret := common.GetSecretStr()

	token, expressTime := common.CreateToken(tokenData.AcountId, secret)

	println("token: "+token )

	expresstimeStr := strconv.FormatInt(expressTime, 10)

	userToken, err := NewUserToken(&tokenData,token, expresstimeStr)

	if err != nil {

		return "",err

	}

	ok := userToken.InsertToDB()

	if !ok  {

		return "",nil

	}

	return userToken.Token,nil

}

func (u *User) ParesePwd(pwd string) bool {

	o := orm.NewOrm()

	o.Using("default")

	fmt.Println("ParesePwd")

	err := o.Raw("select password,salt from user_tb where account_id = ?", u.AccountId).QueryRow(&u)

	if err != nil {

		fmt.Println(err)

		return false

	}

	hashP, err := createPassHashs(pwd,u.Salt)

	if hashP == u.Password {

		return true

	}

	return false

}


func (u *User) UpdatePwd(newPwd string) bool {

	fmt.Println("UpdatePwd")

	o := orm.NewOrm()
	
	o.Using("default")
	
	newPassword,_ := createPassHashs(newPwd,u.Salt)

	_,err := o.Raw("update user_tb set password=? where account_id = ?",newPassword,u.AccountId).Exec()

	if err != nil {

		fmt.Println(err)

		return false

	} 

	return true

}