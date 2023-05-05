docker network create redis_colony --subnet 172.18.0.0/16
DIR=`pwd`
for port in $(seq 1 3)
do
	mkdir -p ./node_${port}/data
	mkdir ./node_${port}/conf
done

for port in $(seq 1 3)
do
	cat << EOF >./node_${port}/conf/redis.conf
bind 0.0.0.0
port 6379
dir /data
logfile "/data/node_${port}.log"
daemonize no
replica-serve-stale-data yes
EOF
if test ${port} -eq 1
	then cat master.conf >>./node_${port}/conf/redis.conf
	else cat slave.conf >>./node_${port}/conf/redis.conf
fi

docker run -d --privileged=true -p 637${port}:6379 --name redis-${port} \
	\-v $DIR/node_${port}/data:/data \
	\-v $DIR/node_${port}/conf/redis.conf:/etc/redis/redis.conf \
	\--net redis_colony --ip 172.18.0.1${port} redis redis-server /etc/redis/redis.conf
done