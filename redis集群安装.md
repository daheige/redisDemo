# redis cluster 集群

    redis cluster实战

# 安装 redis

    sudo mkdir -p  /data/redis/
    sudo mkdir -p  /data/web/
    sudo chown -R $USER /data
    cd /data/web/
    wget http://download.redis.io/releases/redis-4.0.2.tar.gz
    tar xzf redis-4.0.2.tar.gz
    cd redis-4.0.2
    make

# 配置 redis 文件

    mkdir /data/redis/{6380,6381,6382,6383,6384,6385}
    vim /data/redis/6380/redis.conf
    port 6380
    cluster-enabled yes
    cluster-config-file nodes-6380.conf
    cluster-node-timeout 5000
    appendonly yes

    依次添加6381-6385的redis.conf

# 启动 redis

    cd /data/web/redis-4.0.2
    依次在不同的终端中，启动如下命令
    $ ./src/redis-server /data/redis/6380/redis.conf

    $ ./src/redis-server /data/redis/6381/redis.conf

    $ ./src/redis-server /data/redis/6382/redis.conf

    $ ./src/redis-server /data/redis/6383/redis.conf

    $ ./src/redis-server /data/redis/6384/redis.conf

    $ ./src/redis-server /data/redis/6385/redis.conf

# 安装 ruby

    sudo apt-get ruby
    cd /data/web/redis-4.0.2
    cd src
    gem install redis

# 添加集群

    ./src/redis-trib.rb create --replicas 1 127.0.0.1:6380 127.0.0.1:6381 127.0.0.1:6382 127.0.0.1:6383 127.0.0.1:6384 127.0.0.1:6385

# 开始运行

    $ cd /data/web/redis-4.0.2
    $ ./src/redis-cli -h 127.0.0.9 -p 6380
    127.0.0.9:6380> set name daheige
    (error) MOVED 5798 127.0.0.1:6381
    127.0.0.9:6380> set name2 daheige
    OK
    127.0.0.9:6380> set name2 daheige
    OK
    127.0.0.9:6380> set name3 daheige
    OK
    127.0.0.9:6380> get name
    (error) MOVED 5798 127.0.0.1:6381
    127.0.0.9:6380> set name daheige2
    (error) MOVED 5798 127.0.0.1:6381
    127.0.0.9:6380> set name2 daheige
    OK
    127.0.0.9:6380> get name2
    "daheige"
    127.0.0.9:6380> get username
    (error) MOVED 14315 127.0.0.1:6382
    127.0.0.9:6380> get username
    (error) MOVED 14315 127.0.0.1:6382
    127.0.0.9:6380>
    $ ./src/redis-cli -h 127.0.0.1 -p 6382
    127.0.0.1:6382> get name
    (error) MOVED 5798 127.0.0.1:6381
    127.0.0.1:6382> get username
    (nil)
    127.0.0.1:6382> get username
    "daheige"
    127.0.0.1:6382> get username
    "daheige"
    127.0.0.1:6382>

# 采用脚步搭建集群

    cd /data/web/redis-4.0.2/utils/create-cluster
    sudo vim utils/create-cluster

    修改PORT=6379
    添加REDIS_ROOT,并切换到$REDIS_ROOT/utils/create-cluster
    修改内容如下：

    #!/bin/bash

    # Settings
    PORT=6390
    TIMEOUT=2000
    NODES=6
    REPLICAS=1
    REDIS_ROOT=/data/web/redis-4.0.2

    # You may want to put the above config parameters into config.sh in order to
    # override the defaults without modifying this script.

    if [ -a config.sh ]
    then
        source "config.sh"
    fi

    # Computed vars
    ENDPORT=$((PORT+NODES))

    cd $REDIS_ROOT/utils/create-cluster

    启动集群
    cd  /data/web/redis-4.0.2/utils/create-cluster
    ./create-cluster start

    停止集群
    ./create-cluster stop

    进入客户端
    cd /data/web/redis-4.0.2/src

# 参考文档
    
    https://github.com/go-redis/redis
    
    https://blog.csdn.net/cnzyyh/article/details/78543324

    https://my.oschina.net/lyyjason/blog/1842002?from=timeline/
