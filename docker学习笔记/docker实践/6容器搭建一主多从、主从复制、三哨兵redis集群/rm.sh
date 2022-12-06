for port in $(seq 1 6)
do
	docker stop redis-${port}
	docker rm redis-${port}
	sudo rm -rf ./redis_${port}
done
docker network rm redis_colony

