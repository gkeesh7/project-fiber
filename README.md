#Project Fiber

##We use this project to benchmark data offerings


###Architecture
![](resources/Project-Fiber-Architechture.png)


###Setup / Quickstart Guide

- Run Kafka Locally on your system first.
```bash
#Run Zookeeper
docker run -d --name zookeeper -p 2181:2181 -p 3181:3181 -v zookeeper-data:/data/zookeeper -e ZOOKEEPER_CLIENT_PORT=2181 --memory 512m confluentinc/cp-zookeeper:5.4.0

#Run Kafka
docker run -d --name kafka -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=<ip-address or hostname>:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://<ip-address or hostname>:9092 --memory 512m confluentinc/cp-kafka:5.5.0

#Create the kafka topic inside the running kakfa container
/usr/bin/kafka-topics --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic my-topic
```