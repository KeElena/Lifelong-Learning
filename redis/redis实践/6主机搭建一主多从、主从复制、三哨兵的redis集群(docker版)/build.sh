docker network create redis_colony --subnet 172.18.0.0/16
DIR=`pwd`
#初始化文件夹
for port in $(seq 1 6)
do
	mkdir -p ./redis_${port}/data
	mkdir ./redis_${port}/conf
done
#写入配置信息，并启动redis主从复制集群
for port in $(seq 1 3)
do
	cat << EOF >./redis_${port}/conf/redis.conf
bind 0.0.0.0
port 6379
dir /data
logfile "/data/redis_${port}.log"
daemonize no
replica-serve-stale-data yes
EOF
if test ${port} -eq 1
	then cat master.conf >>./redis_${port}/conf/redis.conf
	else cat slave.conf >>./redis_${port}/conf/redis.conf
fi

docker run -d --privileged=true -p 637${port}:6379 --name redis-${port} \
	\-v $DIR/redis_${port}/data:/data \
	\-v $DIR/redis_${port}/conf/redis.conf:/etc/redis/redis.conf \
	\--net redis_colony --ip 172.18.0.1${port} redis redis-server /etc/redis/redis.conf
done
#写入哨兵配置并启动
for port in $(seq 4 6)
do
	cat sentinel.conf>>./redis_${port}/conf/sentinel.conf
docker run -d --privileged=true -p 637${port}:26379 --name redis-${port}\
	\-v $DIR/redis_${port}/data:/data \
	\-v $DIR/redis_${port}/conf/sentinel.conf:/etc/redis/sentinel.conf \
	\--net redis_colony --ip 172.18.0.1${port} redis redis-sentinel /etc/redis/sentinel.conf
done