for port in $(seq 1 6)
do
	docker stop redis-$port
	docker rm redis-$port	
done
docker network rm redis
#cd /home/keqing/桌面/redis集群	
#sudo rm -rf *					#有风险操作