# Set up

## Set up with docker for simple(single node) cluster

아래의 커맨드를 실행시키는 것을 통해 카프카 브로커를 실행 시킵니다.

```shell
docker-compose -f docker-compose.simple.yaml up -d
```

아래와 같은 커맨드 입력을 통해 간단한 테스트를 해볼 수 있습니다.

```shell
# 도커 컨테이너가 정상 실행중임을 확인

❯ docker ps
CONTAINER ID   IMAGE                             COMMAND                  CREATED         STATUS         PORTS                               NAMES
3222ed832215   confluentinc/cp-kafka:7.0.1       "/etc/confluent/dock…"   8 minutes ago   Up 8 minutes   0.0.0.0:9092->9092/tcp              broker
412f9ef6d715   confluentinc/cp-zookeeper:7.0.1   "/etc/confluent/dock…"   8 minutes ago   Up 8 minutes   2181/tcp, 2888/tcp, 3888/tcp        zookeeper

# 토픽 만들기

docker exec broker \
kafka-topics --bootstrap-server broker:9092 \
             --create \
             --topic quickstart

# 토필 리스팅

docker exec broker \
kafka-topics --bootstrap-server broker:9092 \
             --list


# 토픽에 메세지 보내기
#   > 인터렉티브 쉘이 열린다. 라인 단위로 메세지가 구분된다.
#   > 보내고자 하는 메세지를 다 보낸 뒤에는 Ctrl-D 를 눌러서 쉘을 종료하자.

docker exec --interactive --tty broker \
kafka-console-producer --bootstrap-server broker:9092 \
                       --topic quickstart

# 토픽에서 메세지 읽기

docker exec --interactive --tty broker \
kafka-console-consumer --bootstrap-server broker:9092 \
                       --topic quickstart \
                       --from-beginning

```

아래의 커맨드를 실행시키는 것을 통해 카프카 브로커를 종료합니다.

```shell
docker-compose -f docker-compose.simple.yaml down
```
