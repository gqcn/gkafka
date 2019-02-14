# clear
sudo docker rm $(sudo docker ps -a -q)

# zookeeper
sudo docker run -d --name zookeeper -p 2181 -t wurstmeister/zookeeper

# kafka
sudo docker run -d --name kafka --publish 9092:9092 --link zookeeper --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 --env KAFKA_ADVERTISED_HOST_NAME=127.0.0.1 --env KAFKA_ADVERTISED_PORT=9092 wurstmeister/kafka:latest
