for port in $(seq 1 3)
do
	docker stop redis-${port}
	docker rm redis-${port}
	sudo rm -rf ./node_${port}
done
docker network rm redis_colony
