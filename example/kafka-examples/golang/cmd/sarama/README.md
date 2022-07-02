# Sarama Examples

## Simple Producer, Consumer

- purchases 라는 토픽에 간단히 메세지들을 읽고 쓰는 예제이다.
- 아래와 같은 명령을 실행해 먼저 토픽을 생성해 준다.

```sh
docker exec broker \
kafka-topics --bootstrap-server broker:9092 \
             --create \
             --topic purchases2 \
             --replication-factor 1 \
             --partitions 1
```

- 아래와 같은 명령을 실행해 프로듀서를 실행한다. 메세지가 발행되는 것을 확인한다.

```sh
go run ./simple_producer/producer.go -brokers localhost:9092
```

- 아래와 같은 명령을 실행해 컨슈머를 실행한다. 메세지가 컨슘되는 것을 확인한다.

```sh
go run ./simple_consumer/consumer.go -brokers localhost:9092 -group test_group -topics purchases
```
