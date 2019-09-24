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
			"127.0.0.1:6380",
			"127.0.0.1:6381",
			"127.0.0.1:6382",
			"127.0.0.1:6383",
			"127.0.0.1:6384",
			"127.0.0.1:6385",
		},
		PoolSize:     10, // PoolSize applies per cluster node and not for the whole cluster.
		MaxRetries:   2,
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
}

/**
2019/09/24 22:45:14 OK <nil>
2019/09/24 22:45:14 OK <nil>

$ ./src/redis-cli -h 127.0.0.1 -p 6382
127.0.0.1:6382> get username
"daheige"
127.0.0.1:6382> get myname
"daheige2"
127.0.0.1:6382>
*/
