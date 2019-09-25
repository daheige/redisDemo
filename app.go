package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	/*client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}*/

	// https://blog.csdn.net/cnzyyh/article/details/78543324
	// https://my.oschina.net/lyyjason/blog/1842002?from=timeline
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6391",
			"127.0.0.1:6392",
			"127.0.0.1:6393",
			"127.0.0.1:6394",
			"127.0.0.1:6395",
			"127.0.0.1:6396",
		},
		PoolSize:     10, // PoolSize applies per cluster node and not for the whole cluster.
		MaxRetries:   2,  //重试次数
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second, //底层默认3s
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: 10,
		IdleTimeout:  100 * time.Second,
	})

	defer cluster.Close()

	str, err := cluster.Set("username", "daheige", 1000*time.Second).Result()
	log.Println(str, err)

	str, err = cluster.Set("myname", "daheige2", 1000*time.Second).Result()
	log.Println(str, err)

	log.Println(cluster.Get("myname").Result())

	//执行原生的命令
	cluster.Do("lpush", "mylist", "123")

	log.Println(cluster.Do("lpush", "mylist", "123").Result())

	cluster.Do("set", "mysex", 1)
	log.Println(cluster.Get("mysex").Int64()) //对结果进行处理返回int64,error

	i64, err := cluster.Get("myname2").Int64()
	if err != nil {
		if err == redis.Nil {
			log.Println("current key not exist")
		} else {
			log.Println(err)
		}

	} else {
		log.Println("i64: ", i64)
	}

}

/**
2019/09/25 22:06:28 OK <nil>
2019/09/25 22:06:28 OK <nil>
2019/09/25 22:06:28 daheige2 <nil>
2019/09/25 22:06:28 21 <nil>
2019/09/25 22:06:28 1 <nil>
2019/09/25 22:06:28 current key not exist

$ ./src/redis-cli -h 127.0.0.1 -p 6382
127.0.0.1:6382> get username
"daheige"
127.0.0.1:6382> get myname
"daheige2"
127.0.0.1:6382>
*/
