mkdir -p redis/conf
mkdir ./redis/data
cat << EOF >./redis/conf/redis.conf
port 6379
bind 0.0.0.0

save 900 1
save 300 10
save 60 10000

stop-writes-on-bgsave-error yes
rdbcompression yes
dir /data
daemonize no
EOF
docker run -d --privileged=true -p 6379:6379 --name Redis -v 路径/redis/data:/data -v 路径/redis/conf/redis.conf:/etc/redis/redis.conf redis redis-server /etc/redis/redis.conf