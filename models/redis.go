package models

import(

	"github.com/garyburd/redigo/redis"

	"github.com/astaxie/beego"

	"time"

	"fmt"

)

var (
	
	RedisPool  *redis.Pool
	
	REDIS_HOST string

	REDIS_DB   int

)

func init() {

	initRedisPool()

}

//初始化redis链接池
func initRedisPool() {

	REDIS_HOST = beego.AppConfig.String("Redis.host")

	REDIS_DB, _ = beego.AppConfig.Int("Redis.db")

	// 建立连接池
	RedisPool = &redis.Pool{

		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxIdle:     beego.AppConfig.DefaultInt("Redis.maxidle", 1),

		//最大的激活连接数，表示同时最多有N个连接
		MaxActive:   beego.AppConfig.DefaultInt("Redis.maxactive", 10),
		
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 180 * time.Second,

		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", REDIS_HOST)

			if err != nil {

				return nil, err

			}

			// 选择db
			//c.Do("SELECT", REDIS_DB)

			return c, nil

		},
	}

}

//获取缓存的token
func GetTokenFromRedis(tokenKey string) string {

	rds := RedisPool.Get()

	defer rds.Close()

	tokenValue,err := redis.String(rds.Do("GET",tokenKey)) 

	if err!=nil {

		fmt.Println(err)

		return ""

	}
	
	return tokenValue
}