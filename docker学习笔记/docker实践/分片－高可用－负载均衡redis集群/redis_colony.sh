#/bin/bash
docker network create redis --subnet 172.18.0.0/16
for port in $(seq 1 6)
do
	mkdir -p /home/keqing/桌面/redis集群/node-${port}/conf
	touch /home/keqing/桌面/redis集群/node-${port}/conf/redis.conf
	cat << EOF >/home/keqing/桌面/redis集群/node-${port}/conf/redis.conf
port 6379
bind 0.0.0.0
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.18.0.1$port
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes
EOF
docker run -p 637$port:6379 -p 1637$port:16379 --name redis-${port} \
\-v /home/keqing/桌面/redis集群/node-${port}/data:/data \
\-v /home/keqing/桌面/redis集群/node-${port}/conf/redis.conf:/etc/redis/redis.conf \
\-d --net redis --ip 172.18.0.1$port redis /etc/redis/redis.conf
done