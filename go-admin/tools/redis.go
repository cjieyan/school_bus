package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/tools/config"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// var Rdb redis.Conn

var once sync.Once
var redisPool *redis.Pool

/**
redis 初始化连接
*/
func GetRedisPool() *redis.Pool {
	//连接地址
	//db分区
	//密码
	once.Do(func() {
		//建立连接池
		redisPool = &redis.Pool{
			//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
			MaxIdle: config.RedisConfig.RedisMaxIdle,
			//最大的激活连接数，表示同时最多有N个连接
			MaxActive: config.RedisConfig.RedisMaxActive,
			//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
			IdleTimeout: 300 * time.Second,
			//建立连接
			Dial: func() (redis.Conn, error) {
				//logs.Info(RedisConn)
				c, err := redis.Dial("tcp", config.RedisConfig.RedisConn)
				if err != nil {
					return nil, fmt.Errorf("redis connection error: %s", err)
				}
				if config.RedisConfig.RedisPassword != "" {
					if _, authErr := c.Do("AUTH", config.RedisConfig.RedisPassword); authErr != nil {
						return nil, fmt.Errorf("redis auth password error: %s", authErr)
					}
				}
				//选择分区
				c.Do("SELECT", config.RedisConfig.RedisDbNum)
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
	})
	return redisPool

}

func Redisinit() redis.Conn {
	con := GetRedisPool().Get()
	if err := con.Err(); err != nil {
		panic("redis err")
	}
	return con
}

/**
redis  SET
*/
func RdbSet(key, v string) (bool, error) {
	rdb := Redisinit()
	_, err := rdb.Do("SET", key, v)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
redis  GET
*/
func RdbGet(key string) (string, error) {
	rdb := Redisinit()
	val, err := redis.String(rdb.Do("GET", key))
	fmt.Println("RdbGet.....", val, err)
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
	rdb := Redisinit()
	_, err := rdb.Do("EXPIRE", key, ex)
	if err != nil {
		//logs.Error("set error", err.Error())
		return err
	}
	return nil
}

func RdbSetExp(key, v string, ex int) error {
	rdb := Redisinit()
	_, err := rdb.Do("SET", key, v, "EX", ex)
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
	rdb := Redisinit()
	b, err := redis.Bool(rdb.Do("EXISTS", key))
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
	rdb := Redisinit()
	_, err := rdb.Do("DEL", key)
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
	rdb := Redisinit()
	value, _ := json.Marshal(data)
	n, _ := rdb.Do("SETNX", key, value)
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
	rdb := Redisinit()
	var jsonData map[string]string
	bv, err := redis.Bytes(rdb.Do("GET", key))
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
	rdb := Redisinit()
	_, err := rdb.Do("HSET", key, field, data)
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
	rdb := Redisinit()
	data, err := rdb.Do("HGET", key, field)
	if err != nil {
		//logs.Error("hGet error", err.Error())
		return nil, err
	}
	return data, nil
}

/**
redis HDEL
*/
func RdbHDel(key string, field string) error {
	rdb := Redisinit()
	_, err := rdb.Do("DEL", key, field)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return nil
}

func RdbHlen(key string) (int, error) {
	rdb := Redisinit()
	return redis.Int(rdb.Do("HLEN", key))
}

/**
redis hGetAll
return map
*/
func RdbHGetAll(key string) (map[string]string, error) {
	rdb := Redisinit()
	data, err2 := redis.StringMap(rdb.Do("HGETALL", key))
	_, err := data, err2
	if err != nil {
		//logs.Error("hGetAll error", err.Error())
		return nil, err
	}
	return data, nil
}

/**
redis HEXISTS
*/
func RdbHExists(key string, field string) bool {
	rdb := Redisinit()
	b, err := redis.Bool(rdb.Do("HEXISTS", key, field))
	if err != nil {
		//fmt.Println(err)
		return false
	}
	return b
}

/**
redis INCR 将 key 中储存的数字值增一
*/
func RdbIncr(key string) error {
	rdb := Redisinit()
	_, err := rdb.Do("INCR", key)
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
	rdb := Redisinit()
	_, err := rdb.Do("INCRBY", key, n)
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
	rdb := Redisinit()
	_, err := rdb.Do("DECR", key)
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
	rdb := Redisinit()
	_, err := rdb.Do("DECRBY", key, n)
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
	rdb := Redisinit()
	_, err := rdb.Do("SADD", key, v)
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
	rdb := Redisinit()
	data, err := redis.Strings(rdb.Do("SMEMBERS", key))
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
	rdb := Redisinit()
	b, err := redis.Bool(rdb.Do("SISMEMBER", key, v))
	if err != nil {
		//logs.Error("SISMEMBER error", err.Error())
		return false
	}
	return b
}

func RdbZAdd(key, v string) error {
	rdb := Redisinit()
	_, err := rdb.Do("ZADD", key, v)
	if err != nil {
		return err
	}
	return nil
}
