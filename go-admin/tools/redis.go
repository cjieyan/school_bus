package tools


import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-admin/tools/config"
	"time"
)

var Rdb redis.Conn




/**
redis 初始化连接
*/
func GetRedisPool() *redis.Pool {
	//连接地址
	RedisConn := config.RedisConfig.RedisConn
	//db分区
	RedisDbNum := config.RedisConfig.RedisDbNum
	//密码
	RedisPassword := config.RedisConfig.RedisPassword
	//建立连接池
	return &redis.Pool{
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: config.RedisConfig.RedisMaxIdle,
		//最大的激活连接数，表示同时最多有N个连接
		MaxActive: config.RedisConfig.RedisMaxActive,
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 300 * time.Second,
		//建立连接
		Dial: func() (redis.Conn, error) {
			//logs.Info(RedisConn)
			c, err := redis.Dial("tcp", RedisConn)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			if RedisPassword != "" {
				if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			//选择分区
			c.Do("SELECT", RedisDbNum)
			return c, nil
		},
		//ping
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}

}

func Redisinit() {
	Rdb = GetRedisPool().Get()
}

/**
redis  SET
*/
func RdbSet(key, v string) (bool, error) {
	b, err := redis.Bool(Rdb.Do("SET", key, v))
	if err != nil {
		//logs.Error("set error", err.Error())
		return false, err
	}
	return b, nil
}

/**
redis  GET
*/
func RdbGet(key string) (string, error) {
	val, err := redis.String(Rdb.Do("GET", key))
	if err != nil {
		//logs.Error("get error", err.Error())
		return "", err
	}

	return val, nil
}

/**
redis EXPIRE
*/
func RdbSetKeyExp(key string, ex int) error {
	_, err := Rdb.Do("EXPIRE", key, ex)
	if err != nil {
		//logs.Error("set error", err.Error())
		return err
	}
	return nil
}

func RdbSetExp(key, v string, ex int) error {
	_, err := Rdb.Do("SET", key, v, "EX", ex)
	if err != nil {
		//logs.Error("set error", err.Error())
		return err
	}
	return nil
}

/**
redis EXISTS
*/
func RdbCheck(key string) bool {
	b, err := redis.Bool(Rdb.Do("EXISTS", key))
	if err != nil {
		//fmt.Println(err)
		return false
	}
	return b
}

/**
redis DEL
*/
func RdbDel(key string) error {
	_, err := Rdb.Do("DEL", key)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return nil
}

/**
redis SETNX
*/
func RdbSetJson(key string, data interface{}) error {
	value, _ := json.Marshal(data)
	n, _ := Rdb.Do("SETNX", key, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

/**
redis GET
return map
*/
func RdbGetJson(key string) (map[string]string, error) {
	var jsonData map[string]string
	bv, err := redis.Bytes(Rdb.Do("GET", key))
	if err != nil {
		//logs.Error("get json error", err.Error())
		return nil, err
	}
	errJson := json.Unmarshal(bv, &jsonData)
	if errJson != nil {
		//logs.Error("json nil", err.Error())
		return nil, err
	}
	return jsonData, nil
}

/**
redis hSet 注意 设置什么类型 取的时候需要获取对应类型
*/
func RdbHSet(key string, field string, data interface{}) error {
	_, err := Rdb.Do("HSET", key, field, data)
	if err != nil {
		//logs.Error("hSet error", err.Error())
		return err
	}
	return nil
}

/**
redis hGet 注意 设置什么类型 取的时候需要获取对应类型
*/
func RdbHGet(key, field string) (interface{}, error) {
	data, err := Rdb.Do("HGET", key, field)
	if err != nil {
		//logs.Error("hGet error", err.Error())
		return nil, err
	}
	return data, nil
}

/**
redis hGetAll
return map
*/
func RdbHGetAll(key string) (map[string]string, error) {
	data, err2 := redis.StringMap(Rdb.Do("HGETALL", key))
	_, err := data, err2
	if err != nil {
		//logs.Error("hGetAll error", err.Error())
		return nil, err
	}
	return data, nil
}

/**
redis INCR 将 key 中储存的数字值增一
*/
func RdbIncr(key string) error {
	_, err := Rdb.Do("INCR", key)
	if err != nil {
		//logs.Error("INCR error", err.Error())
		return err
	}
	return nil

}

/**
redis INCRBY 将 key 所储存的值加上增量 n
*/
func RdbIncrBy(key string, n int) error {
	_, err := Rdb.Do("INCRBY", key, n)
	if err != nil {
		//logs.Error("INCRBY error", err.Error())
		return err
	}
	return nil
}

/**
redis DECR 将 key 中储存的数字值减一。
*/
func RdbDecr(key string) error {
	_, err := Rdb.Do("DECR", key)
	if err != nil {
		//logs.Error("DECR error", err.Error())
		return err
	}
	return nil
}

/**
redis DECRBY 将 key 所储存的值减去减量 n
*/
func RdbDecrBy(key string, n int) error {
	_, err := Rdb.Do("DECRBY", key, n)
	if err != nil {
		//logs.Error("DECRBY error", err.Error())
		return err
	}
	return nil
}

/**
redis SADD 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
*/
func RdbSAdd(key, v string) error {
	_, err := Rdb.Do("SADD", key, v)
	if err != nil {
		//logs.Error("SADD error", err.Error())
		return err
	}
	return nil
}

/**
redis SMEMBERS 返回集合 key 中的所有成员。
return map
*/
func RdbSMembers(key string) (interface{}, error) {
	data, err := redis.Strings(Rdb.Do("SMEMBERS", key))
	if err != nil {
		//logs.Error("json nil", err)
		return nil, err
	}
	return data, nil
}

/**
redis SISMEMBER 判断 member 元素是否集合 key 的成员。
return bool
*/
func RdbSISMembers(key, v string) bool {
	b, err := redis.Bool(Rdb.Do("SISMEMBER", key, v))
	if err != nil {
		//logs.Error("SISMEMBER error", err.Error())
		return false
	}
	return b
}

func RdbZAdd(key, v string) error{
	_, err := Rdb.Do("ZADD", key, v)
	if err != nil {
		return err
	}
	return nil
}