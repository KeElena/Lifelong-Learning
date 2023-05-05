#!/bin/bash
#reset first node
for host in $(seq 1 3)
do
	docker exec rabbitmq${host} rabbitmqctl stop_app;
	docker exec rabbitmq${host} rabbitmqctl reset;
	if test ${host} -gt 1
		then docker exec rabbitmq${host} rabbitmqctl join_cluster --ram rabbit@rabbitmq1;
	fi
	docker exec rabbitmq${host} rabbitmqctl start_app;
done
echo "Starting to create user."
docker exec rabbitmq1 rabbitmqctl add_user admin 123456;
echo "Set tags for new user."
docker exec rabbitmq1 rabbitmqctl set_user_tags admin administrator;
echo "Grant permissions to new user."
docker exec rabbitmq1 /bin/bash -c "rabbitmqctl set_permissions -p '/' admin '.*' '.*' '.*'";
