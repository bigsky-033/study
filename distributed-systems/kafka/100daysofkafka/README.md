# 100 Days of Code with Apache Kafka

<https://developer.confluent.io/100-days-of-code/> 를 따라 공부한 내용을 정리한 문서이다. 각 일자 별 공부한 자료들, 정리한 내용을 적고 필요하다면 간략한 메모를 남긴다.

## Day 1 (20220612)

- 공부한 자료
  - <https://kafka.apache.org/intro>
  - <https://youtu.be/FKgi3n-FyNU>
  - <https://developer.confluent.io/quickstart/kafka-on-confluent-cloud/>
- 정리한 내용
  - [Set up guide](setup.md)
- 메모
  - 카프카 관련 키워드라고 느낀 것들: Logs, Events, Event streaming, Micro services, Topic, Durable, Real time analytic, Distributed, Scalable, Ecosystem
  - Kafka 공식 홈페이지의 intro 문서의 내용이 생각보다 좋았다.
  - Kafka 공식 문서 중 Design 문서는 꼭 읽어 보아야겠다.

## Day 2 (20220613)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/events/>
  - <https://developer.confluent.io/learn-kafka/apache-kafka/get-started-hands-on/>
- 메모
  - 카프카에서 가장 중요한 것은 Event이다. Kafka에서 이벤트란, `an event is a thing that has happened, that's it` 이다. 어떤 것이든 될 수 있다. 예를 들어 IOT, Business process change, User interaction, Microservice output 등이 될 수도 있다.
  - Event를 다른 말로 표현하면 `Notification(when) + State`이다.
  - Kafka는 event를 key/value pair로 모델링한다. RDB와 같은 곳 에서와 달리 Key는 고유할 필요가 없다. RDB의 PK와 같은 것 이라기 보다는 시스템에서 엔티티의 식별자일 경우가 더 많다.
  - Hands on은 Confluent Cloud를 이용해 kafka를 사용하는 법을 다룬다. 이 과정을 진행하며 Confluent Cloud는 직접 가입해 사용하지는 않는다. 강의에서 보여주는 내용 정도만 보고 카프카 관련 실습이 필요한 경우 Day 1 에서 작성한 로컬 셋업 가이드를 참고해 실습한다.

## Day 3 (20220614)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/topics/>
  - <https://www.confluent.io/blog/okay-store-data-apache-kafka/>
- 메모
  - Topic은 비슷한 종류의 이벤트들을 모아두는 컨테이너이다.
  - Kafka를 queue라고 부르는 것은 엄밀히 말했을 때 정확하지 않을 수 있다. Topic은 `durable logs of events`이다.
    - Append only.
    - Offset을 통한 탐색만 가능하다. 기본적으로 index를 기반으로 한 탐색은 불가능하다.
    - event는 불변이다.
  - Log는 디스크에 쓰여진다.
  - confluent blog의 글은 kafka에 저장된 데이터를 source-of-truth 로 사용해도 괜찮을지에 대한 답변이다.
    - 일단 가능은 하다. 데이터의 리텐션을 forever로 설정하거나 topic의 log compaction을 활성화하면 데이터가 계속 보관되기 때문이다.
    - 실제로 카프카는 이런 용도에 적합하다: event sourcing의 데이터 저장소, streaming app의 데이터 저장소 + replay, CDC
    - 그럼 사람들은 왜 이걸 자연스럽지 못 하다고 여기는 걸까? 카프카가 메세지 큐로 인식되기 때문이다. 그러나 카프카는 전통적인 메세지 큐와는 다르다. 카프카는 전통적인 메세지 큐 보다는 분산 파일 시스템(또는 데이터베이스)에 더 가깝다.
      - 카프카는 log를 영속적으로 저장하기 때문에 이것을 제한 없이 다시 읽을 수 있다.
      - 카프카는 현대적인 분산 시스템이다. 클러스터로서 동작하고 확장가능하며 내부적으로 데이터를 복제해 fault-tolerance, high-availability 등을 달성한다.
      - 메세지의 단건 처리 외에도, real-time streaming 처리를 가능하게 한다.
  - 그러나 카프카를 메인 스토리지로 쓴다는 것은 운영에 어려움을 더할 수 있다.
  - 또한 다양한 쿼리에 있어서는 취약할 수 있다. 분석 DB, 검색 엔덱스, 캐시 등 각각의 스토리지 기술들은 고유한 장점이 있다.

## Day 4 (20220615)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/partitions/>
  - <https://developer.confluent.io/learn-kafka/apache-kafka/partitions-hands-on/>
- 메모
  - Kafka는 분산 시스템으로서 시스템의 부하를 분산하고 수평 확장이 가능하게 하기 위해 토픽의 로그를 여러 파티션으로 나누어 저장할 수 있게 한다.
  - 어떤 메세지가 토픽의 어떤 파티션으로 전달될지를 결정하는 것이 key의 역할이다. Key를 기반으로 해시를 해서 로그가 저장될 파티션을 결정한다. 같은 key를 가진 로그는 같은 파티션으로 전달됨이 보장된다.
  - 이로 인해 이론상 파티션 크기의 불균형이 생길 수 있고 이로 인해 문제가 생길 수 있긴 한데 현실에서 문제가 될 가능성은 높지 않다.
  - Hands on은 confluent kafka에 파티션 수가 서로 다른 토픽을 생성하고 확인하는 것 이었다. 별도로 해보지는 않았다.

## Day 5 (20220616)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/brokers/>
  - <https://docs.confluent.io/platform/current/control-center/brokers.html>
- 메모
  - 브로커의 단순한 정의는 카프카 프로세스를 실행하고 있는 무언가(컴퓨터, 인스턴스, 컨테이너 등)이다.
  - 카프카의 파티션을 관리하며 요청을 읽고 쓴다.

## Day 6 (20220618)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/producers/>
  - <https://docs.confluent.io/platform/current/clients/producer.html>
  - <https://github.com/confluentinc/confluent-kafka-go>
  - <https://github.com/Shopify/sarama>
- 메모
  - 카프카를 사용하는 어플리케이션 개발이란 프로듀서와 컨슈머를 개발하는 것을 의미한다. 카프카 프로듀서를 개발한다는 것은 토픽에서 메세지를 쓰기 위한 코드를 작성하는 것이다.
  - 카프카 프로듀서는 단순히 메세지를 토픽에 넣는 것 이외에 커넥션 풀의 관리, 네트워크 버퍼링, 파티셔닝 등의 역할을 한다.
  - confluent의 프로듀서 문서의 내용은 중요한 컨셉을 잘 설명하고 있다. 핵심 내용은 다음과 같다.
    - 카프카 프로듀서의 핵심 기능은 메세지의 파티셔닝, 메세지 전송에 대한 보장,배치/압축 처리를 통한 높은 처리량 달성과 같은 점이다.
    - 카프카 프로듀서의 설정을 분류한다면 크게 다음과 같은 카테고리 들로 분류할 수 있다. 이 분류를 통해 카프카 프로듀서의 핵심이 무엇인지를 알 수 있다. 각 분류의 자세한 내용은 문서를 참고하자.
      - Core Configuration
      - Message Durability
      - Message Orderding
      - Batching and Compression
      - Queuing Limits
  - golang 카프카 클라이언트로는 confluent에서 제공하는 confluent-kafka-go와 Shopify가 개발한 sarama가 유명한 것 같다. 간단히 검색을 통해 파악한 각각의 장단점은 다음과 같다.
    - confluent-kafka-go는 Librdkafka를 래핑한 라이브러리로 거기서 오는 장단점을 함께 갖고 있다. 장점은 안정적이라는 것, 단점은 cgo에서 오는 점들이다.
    - sarama의 경우 장점은 pure-go, 단점은 상대적으로 안정성이 떨어진다는 점이다.
  - 두 라이브러리 모두 공부를 해나가며 직접 사용해 볼 것이고 장단점을 더 자세히 파악해 보려고 한다.

## Day 7 (20220619)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/consumers/>
  - <https://docs.confluent.io/platform/current/clients/consumer.html>
  - <https://github.com/confluentinc/confluent-kafka-go>
- 메모
  - 카프카 컨슈머를 개발하는 것에 적용되는 일반적인 원칙들은 프로듀서에 적용되는 원칙들과 유사하다. 카프카 컨슈머를 개발한다는 것은 토픽에서 메세지를 읽기 위한 코드를 작성하는 것이다.
  - 카프카 컨슈머는 단순히 메세지를 토픽에서 가져오는 것 이외에 커넥션 풀의 관리, 네트워크 관련된 것들을 처리한다. 단 카프카 프로듀서에 비해 조금 더 많은 것들이 있다. 카프카는 컨슈머 그룹이라는 개념이 있고 이것을 통해 컨슈머가 스케일링을 할 수 있다.
  - 토픽의 컨슈머를 늘리는 것을 통해 스케일링을 하는 것이 카프카만의 새로운 아이디어는 아니다. 다만 이렇게 했을 때 전통적인 메세징 시스템들은 메세지의 순서를 온전히 보장하지 못 하는 경우가 많다. 카프카의 경우 메세지의 key가 null이 아니라면, 같은 key의 메세지들에 대해서는 순서가 보장된다.
  - 컨슈머 그룹, 오프셋, 리밸런싱과 관련된 것들이 카프카 컨슈머에서 중요한 내용이다. confluent 문서 중 consumer 문서의 concept 부분에서 다루는 두 가지 내용도 컨슈머 그룹과 오프셋 관리이다. Confluent concept 문서의 내용만 주의 깊게 봐도 상당히 많은 내용을 배울 수 있다.
    - 컨슈머 관리 관련해서는 주키퍼를 사용하는 예전의 방식과 주키퍼 없이 카프카 자체적으로 처리하는 방식이 있다. 문서에서는 후자의 방식을 설명한다. 이 방식은 각 컨슈머 그룹에 대해 클러스터 내 브로커 중 하나가 coordinator 역할을 하는 방식이다. Coordinator를 통해 컨슈머 그룹에 join 하고 coordinator는 리밸런싱 등의 작업을 한다. 또한 컨슈머 그룹의 각 컨슈머는 coordinator에게 주기적으로 heartbeat를 보내게 되는데 만약 일정 시간 동안(session timeout) heartbeat를 보내지 못 하게 되면 그룹에서 추방된다.
    - 오프셋 관리와 관련해서는 몇 가지 중요한 설정이 있다.
      - 컨슈머 그룹이 생성되고 최초로 컨슈밍을 시작할 위치는 `auto.offset.reset` 설정에 의해 결정된다.
      - 컨슈머는 메세지를 컨슈밍하고 난 뒤에 올바르게 커밋을 해야 하는데 만약 컨슈머가 오프셋을 커밋하기 전에 크래시가 되고 컨슈머 그룹의 다른 컨슈머가 파티션을 할당받게 된다면 어느 오프셋에서부터 컨슈밍을 할지는 reset policy에 따라 결정된다.
      - 오프셋 관리는 자동으로 할 수도 있고 수동으로 할 수도 있는데 올바르게 오프셋을 관리하는 것은 delivery semantics(at-least-once, exactly-once) 관점에서 매우 중요하다.
        - auto commit를 사용하는 at-least-once 이다.
        - `auto.commit.interval.ms` 설정을 통해 auto commit의 주기를 설정할 수 있다.
        - `enable.auto.commit` 설정에 값을 주는 것을 통해 auto commit를 키고 끌 수 있다.
    - 전반적인 처리량(throughput)과 관련된 설정들도 몇 가지 있다.
      - poll을 했을 때 가져올 데이터의 양과 관련된 설정으로 `fetch.min.bytes`라는 설정이 있다. 이 값이 설정되어 있으면 충분한 데이터가 있거나 또는 `fetch.max.wait.ms` 설정에 설정된 값이 지나기 전까지 브로커는 홀딩한다.
      - async commit을 하는 방법도 있다. 요청했던 데이터가 전부 처리되기 전에 컨슈머는 커밋을 하고 return 해버리는 방법이다. 그러나 이 경우 컨슈머는 커밋에 fail 한다고 하더라도 커밋을 재시도하지 않는다. 또한 커밋의 순서가 꼬일수도 있다. 또한 커밋의 실패와 리밸런싱이 함께 발생하게 되면 중복된 데이터를 컨슈밍할 수도 있다. 따라서 async 커밋은 sync 커밋에 비해 위험하다고 여겨져야 한다.
        - async commit을 통해 달성할 수 있는 것은 at-least-once이다. at-most-once를 달성하기 위해서는 다음 메세지를 컨슈밍 하기 이전에 커밋이 성공했는 지를 알 수 있어야 한다.
    - 각 리밸런싱은 두 가지 phase가 있다: partition revocation, partition assignment. Partition revocation은 항상 리밸런싱 전에 호출되기 때문에 파티션이 재할당 되기 전에 오프셋을 커밋할 마지막 기회이다. Partition assiginment는 항상 리밸런싱이 일어난 이후에 호출되기 때문에 할당된 파티션의 초기 오프셋 위치를 설정할 수 있다.
    - 카프카 컨슈머의 설정을 분류한다면 크게 다음과 같은 카테고리 들로 분류할 수 있다. 이 분류를 통해 카프카 컨슈머의 핵심이 무엇인지를 알 수 있다. 각 분류의 자세한 내용은 문서를 참고하자.
      - Core Configuration
      - Group Configuration
      - Offest Management
  - Java 구현체와 librdkafka-based clients(C/C++, Python, Go 그리고 C#)의 내부 동작 방식이 좀 다르다.

## Day 8 (20220620)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/kafka-connect/>
  - <https://developer.confluent.io/learn-kafka/apache-kafka/kafka-connect-hands-on/>
  - <https://docs.confluent.io/platform/current/connect/index.html>
  - <https://youtu.be/dXXfkoXXBbs>
- 메모
  - 세상의 모든 데이터가 카프카에 있는 것은 아니기 때문에 다른 데이터 소스에서 카프카로 데이터를 넣거나 또는 카프카에서 다른 곳으로 데이터를 옮겨야 하는 경우가 있다. 그런 역할을 하는 것이 Kafka Connect이다. 예를 들어 카프카에 있는 데이터를 엘라스틱서치에 넣는 것과 같은 케이스이다.
  - Kafka Connect를 사용할 때 복잡하지 않은 경우라면 간단히 설정 파일을 작성하고 실행하는 것을 통해 카프카에서 데이터를 다른 곳에 옮기거나 다른 곳에서 카프카로 옮길 수 있다.
  - Source Connector는 Producer, Sink connect는 Consumer 이다.
  - Confluent hub에서 이미 작성된 다양한 Connector를 볼 수 있다. 영상에서 강조했듯이 대부분의 데이터를 옮기는 작업은 특별할 게 없다. 때문에 가급적이면 직접 코드를 작성하기 전에 검증된 솔루션이 있는지 찾아보고 있으면 그것을 쓰는 것을 권장한다.
    - Confluent는 이와같이 Ecosystem 으로서의 카프카를 강조하는 듯하다.
  - Hands on은 간단한 클릭을 통해 클라우드 환경에서 Kafka Connector - Source 를 작성하는 법을 보여준다. 단순한 경우라면 정말 이렇게 간단하게 될 것 같다.
  - Robin Moffatt 채널의 영상들이 흥미롭다. 퀄리티가 매우 좋은듯하다. Kafka Connector 구조의 핵심은 Connector, Transform, Converter이다.

## Day 9 (20220621)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/kafka-streams/>
  - <https://www.confluent.io/blog/kafka-streams-tables-part-1-event-streaming/>
  - <https://github.com/confluentinc/examples/tree/7.1.1-post/connect-streams-pipeline>
- 메모
  - 카프카 컨슈머의 API는 매우 심플하다. 따라서 단순한 케이스 이외에 무언가를 구현하려면 많은 것들을 직접 해줘야 한다. Time window, 순서에 맞지 않는 메세지, lookup 테이블 등의 케이스, 키를 이용한 집계 등 이다.
    - 많은 경우 stateful operation이다.
  - 카프카 스트림은 Java API이다. Stream Processing에서 필요한 Filtering, Grouping, Aggregating, Joining 과 같은 필수적인 기능들을 제공하기 위한 API 들의 모음이다.
  - Scalable, fault-tolerant, 상태관리 등의 기능도 제공한다.
  - 카프카 스트림은 라이브러리이다. 독립된 컴포넌트가 아니기 때문에 Spring과 같은 프레임워크와 함께 사용해 구현이 가능하다.
  - Streams and Tables in Apache Kafka 글에서 인상 깊었던 부분들:
    - An event stream records the history of what happened in the world as a sequence of events.
    - A table represents the state of the world.
    - Event와 Table의 차이를 위와 같이 설명해 주고 그를 그림으로 도식화해 준 부분이 좋았다.
    - Stream은 immutable data인데 반해 table은 mutable data이다.
    - Stream-table duality는 event sourcing에서 말하는 개념 같았다.
  - Streams and Tables in Apache Kafka 이 시리즈는 전부 읽어 보아야겠다.
  - 100 days of code 과정을 통해 공부하는 동안 이벤트를 기반으로 한 아키텍처에 대해서도 많이 공부할 수 있을 것 같아 좋다.

## Day 10 (20220623)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/ksqldb/>
  - <https://www.confluent.io/blog/how-real-time-stream-processing-safely-scales-with-ksqldb/>
- 메모
  - KsqlDB를 추상적으로 본다면 스트림 프로세싱 어플리케이션에 최적화된 database라고 할 수 있다. 사용자는 SQL로 원하는 동작을 정의할 수 있고 KsqlDB는 SQL을 바탕으로 이벤트 스트림을 계속해서 프로세싱한다. 그리고 어플리케이션은 KsqlDB의 결과를 가져다 쓸 수 있다.
  - 물론 추상적으로 보았을 때 그렇다는 이야기이다. PostgreSQL 같은 RDB를 대체하기 위한 것은 아니다.
  - 두 번째 공부한 자료에 있는 How Real-Time Stream Processing Safely Scales with ksqlDB, Animated 글은 매우 흥미로웠다. 이를 통해 KsqlDB가 어떻게 동작하는지 직관적인 느낌을 받을 수 있었다. 그리고 Fault tolerance, High availability에 대한 부분의 글을 통해 많이 배웠다. 분산 시스템으로서 고민할만한 중요한 포인트들이 잘 설명되어 있어서 좋았다.

## Day 11 (20220624)

- 공부한 자료
  - <https://developer.confluent.io/learn/kraft/>
  - <https://www.confluent.io/blog/kafka-without-zookeeper-a-sneak-peek/#zookeeper-less-kafka>
- 메모
  - Apache Kafka Raft(KRaft)는 주키퍼를 사용하지 않고 카프카 자체적으로 메타 데이터를 관리할 수 있게 하기 위해 도입된 합의 프로토콜이다. 장점은 다음과 같은 것들이다:
    - control plane의 성능 향상으로 카프카 클러스터가 더 많은 파티션을 관리할 수 있다.
    - 카프카의 안정성을 높이고 소프트웨어를 단순하게 한다. 이로 인해 소프트웨어를 더 쉽게 관리할 수 있게 한다.
    - 가벼운 하나의 프로세스로 카프카를 시작할 수 있다.
    - 컨트롤러가 거의 즉시 failover 할 수 있게 한다.
  - 소프트웨어 아키텍처의 단순화로 가져올 수 있는 좋은 점들을 볼 수 있는 것 같다.
  - 내부적으로 event-sourced storage model을 사용한다.

## Day 12 (20220625)

- 공부한 자료
  - <https://developer.confluent.io/get-started/go/>
  - <https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go/kafka>
- 정리한 내용
  - `/example/kafka-examples/golang/` 아래에 Getting Started with Apache Kafka and Go 의 예제를 넣어둠.
- 메모
  - 간단하게 튜토리얼만 해 본 수준이다. <https://github.com/confluentinc/confluent-kafka-go/tree/master/examples> 에 보면 다양한 예제가 많이 있는데 하나하나 읽어 보아야겠다.
  - 잘 사용하려면 라이브러리의 특히 go document를 잘 읽어 보아야겠다.

## Day 13 (20220626)

- 공부한 자료
  - <https://github.com/Shopify/sarama>
  - <https://pkg.go.dev/github.com/Shopify/sarama>
- 메모
  - Go에서 널리 쓰이는 kafka 클라이언트 라이브러리인 sarama에 대해 알아보고 사용해 보았다.
  - API가 뭔가 익숙하지 않았다. Github에 예제가 좀 더 있었으면 좋았을 것 같다.
  - cgo 이슈가 아니라면 confluent-kafka-go를 사용할 것 같다.

## Day 14 (20220628)

- 공부한 자료
  - <https://kafka.apache.org/documentation.html#design>
- 메모
  - sarama github을 보던 중 Kafka protocol 가이드 문서에 좋은 내용이 많다고 해서 훑어보려고 했다. 그런데 protocol 문서를 읽어보려고 하니 design 문서의 내용을 이해하고 있다고 가정한다고 한다. 그래서 design 문서를 먼저 읽기로 했다.
  - 4.1 Motivation, 4.2 Persistence, 4.3 Efficiency 를 읽었다. 내용이 상당히 많았지만 하나하나 중요한 내용들이었다. 디스크에 쓰는 것을 어떻게 최적화했는지, 왜 카프카가 디스크에 쓰는 데도 느리지 않은지를 설명하고 있다. 효율을 위한 제로카피, 압축, 배치 처리 등도 인상깊었다.
    - 사실 잘 이해하지 못 한 부분이 많이 있다. 이 DESIGN 문서들은 여러 번 읽어 보아야겠다.
    - 글에서 소개해 주는 아티클들도 내용이 매우 좋다.

## Day 15 (20220629)

- 공부한 자료
  - <https://kafka.apache.org/documentation.html#design>
  - <https://developer.ibm.com/articles/j-zerocopy/>
  - <https://youtu.be/UNUz1-msbOM>
- 메모
  - 제로카피에 대한 문서는 매우 흥미롭다. 그리고 이해하기 쉽게 쓰여있다.
  - Design 문서의 내용을 이어서 읽었다. 4.4 The Producer 관련 메모
    - 프로듀서의 경우 주요 관심사는 로드 밸런싱과 비동기 전송이다. 클라이언트가 다른 라우팅 계층을 거치지 않고 바로 파티션의 리더를 찾아 데이터를 보낸다. 이를 위해 각 카프카의 노드는 메타데이터를 제공할 수 있어야 한다는 점이 흥미로웠다. 또한 로드밸런싱을 랜덤하게 할 수도 있고 특정 키를 기반으로 할 수도 있으며 특정 키를 기반으로 했을 때 같은 키를 가진 데이터들이 같은 곳에 있다는 가정을 할 수 있고 클라이언트가 이를 이용한 프로그래밍을 할 수 있게 했다.
    - 포로듀서 성능에 있어서 가장 중요한 부분 중 하나는 배치 처리이다. 데이터를 모아 보내기 위한 설정들을 프로듀서에 해줄 수 있다. 쓰루풋과 레이턴시 중 어떤 것을 더 중요시하냐에 따라 설정 값을 조절해서 사용하면 된다.
  - 4.5 The Consumer 관련 메모
    - 컨슈머는 기본적으로 컨슈밍하고 싶은 파티션의 리더인 브로커에게 fetch 요청을 보내는 방식으로 동작한다.
    - 카프카 프로듀서는 Push 방식, 컨슈머는 Pull 방식으로 동작한다. 이는 상대적으로 전통적인 방식이다. 이런 선택에는 장단점이 있지만, 컨슈머가 Push 방식으로 동작할 때는 다양한 컨슈머를 지원하기 어려운 단점이 있다. 또한 데이터가 갑자기 많이 밀려 들어올 때 대처하기도 Pull 방식이 더 유리하다.
    - Pull 방식의 단점인 데이터가 없을 때 pull request가 너무 많아질 수 있다는 점은 컨슈머가 블락할 수 있게 하는 방식을 통해 해결했다.
    - 무엇이 컨슈밍 되었는지를 추적하는 것은 메세징 시스템에서 성능에 매우 중요한 부분이다. 대부분의 경우에 이런 메타데이터는 브로커에서 관리한다. 이는 규모가 작은 시스템에서는 가능하다. 그러나 컨슈머와 브로커가 어디까지 컨슘했는지 성능을 생각하고 합의를 하는 건 매우 복잡한 일이다. 카프카가 이를 해결한 방식을 이해하는 것이 중요하다. 카프카는 성능, 그리고 메세지 유실을 방어하는 걸 목적으로 해결했다. 또한 컨슈머가 rewind 할 수 있게 했다.

## Day 16 (20220701)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/creating-first-apache-kafka-producer-application/kafka.html>
- 정리한 내용
  - `/example/kafka-examples/jvm/` 아래에 How to build your first Apache KafkaProducer application 의 예제를 넣어둠.
- 메모
  - 디자인 문서의 나머지 내용은 카프카 내부를 상세히 공부하는 단계에서 읽어보기로 했다.
  - 심플한 코틀린 어플리케이션 예제를 하나 만들었다. 카프카를 띄우는 설정은 [docker-compose](/example/kafka-examples/docker-compose.yaml) 파일을 사용했다.

## Day 17 (20220702)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/kafka-producer-callback-application/kafka.html>
- 정리한 내용
  - `/example/kafka-examples/jvm/` 아래에 How to build an Apache KafkaProducer application with callbacks 의 예제를 넣어둠.
- 메모
  - 예제를 따라해 보았다.
  - 테스트 예제에 있는 MockProducer가 새로웠고 유용할 것 같다.

## Day 18 (20220703)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/creating-first-apache-kafka-consumer-application/kafka.html>
- 정리한 내용
  - `/example/kafka-examples/jvm/` 아래에 How to build your first Apache KafkaConsumer application 의 예제를 넣어둠.
- 메모
  - 에제를 따라해 보았다.
  - 이번에는 MockConsumer를 활용해서 테스트 코드도 만들어 보았다. 매우 유용할듯 싶다.

## Day 19 (20220704)

- 공부한 자료
  - <https://developer.confluent.io/get-started/spring-boot/>
- 정리한 내용
  - `/example/kafka-examples/spring/` 아래에 Getting Started with Apache Kafka and Spring Boot 의 예제를 넣어둠.
- 메모
  - 간단한 스프링 예제여서 어려움은 없었다. 그런데 @Configuration 도 없이 동작하는 걸 보니 조금 더 자세한 이해가 필요하다고 생각했다.

## Day 20 (20220705)

- 공부한 자료
  - <https://developer.confluent.io/get-started/c/>
- 메모
  - 예제를 따라하기 위해 librdkafka, glibc 와 같은 것들을 WSL 환경에 설치하고 따라해 보았는데 컴파일에 실패했다. 컴파일을 성공하게 하기 위해 이런저런 방법을 찾아보았지만 c에 대한 지식이 부족해서 성공하지 못 하였다. 그래서 그냥 코드를 읽어보는 수준에서 만족하고 다음에 c를 더 공부하고 와서 다시 시도해 보기로 하였다.
  - 바로 직전에 java, spring의 예제를 보다가 c로 짜여진 예제를 보니 확실히 저수준의 코드라는 느낌이 났다. 공부할 때는 이런저런 동작을 더 깊이 알 수 있어서 좋을 것 같긴 한데 프로덕션에서 사용할 일은 없을 것 같다.

## Day 21 (20220707)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/kafka-console-consumer-producer-basics/kafka.html>
  - <https://youtu.be/4Hz46hmwXM4>
- 메모
  - 간단히 카프카 cli를 이용해 메세지를 프로듀싱하고 컨슈밍해 보는 예제들이었다. 특별한 내용은 없었다.

## Day 22 (20220708)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/kafka-console-consumer-producer-avro/kafka.html>
- 메모
  - 카프카 cli 기초 이후에 schema registry를 이용해 (de)serializers를 이용하는 예제들을 보았다. 특별한 내용은 없었다.

## Day 23 (20220709)

- 공부한 자료
  - <https://developer.confluent.io/tutorials/kafka-console-consumer-read-specific-offsets-partitions/kafka.html>
- 메모
  - 카프카 cli 를 다루는 것의 연장선이었다. 특정 오프셋을 입력으로 주고 거기서부터 메세지를 읽어오는 것을 해보았다. 특별한 내용은 없었다.

## Day 24 (20220710)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/kcat.html>
  - <https://github.com/edenhill/kcat>
  - <https://dev.to/de_maric/learn-how-to-use-kafkacat-the-most-versatile-kafka-cli-client-1kb4>
- 메모
  - confluent 문서에 있는 내용은 그다지 도움되지 않았다.
  - apt로 설치했을 때 kcat이라는 이름으로 어플리케이션을 실행할 수는 없었다. kafkacat 로 실행해야 했다.
  - 세 번쨰로 첨부한 블로그의 글이 좋다. cheatsheet 느낌이다. 이 예는 useful-links 에 넣어두고 필요할 때 사용해 보아야겠다.
  - 프로듀서보다는 컨슈머로 사용할 때 도움이 많이 될 것 같다.
  - 간단하게 클러스터와 토픽의 메타데이터를 조회할 수 있는 기능도 좋았다.

## Day 25 (20220714)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/kafka-rest/index.html>
- 메모
  - 카프카 Rest Proxy라는 주제에 대해 간단히 알아보았다. 100 Days of Code with Apache Kafka® 에 있는 Day 25 링크는 동작하지 않는다.
  - Rest Proxy는 카프카 cli나 sdk등을 통해 카프카 프로토콜로 통신하는 게 아니라, Rest API와 같은 인터페이스를 통해 카프카를 사용할 수 있게 해주는 서비스이다.
    - Metadata, Producers, Consumers, Cluster & Load Balancing, Admin operations 등 메세지를 produce/consume 할 수 있고 또한 클러스터의 상태를 보거나 변경하는 등의 액션을 할 수 있다.

## Day 26 (20220716)

- 공부한 자료
  - <https://developer.confluent.io/learn-kafka/apache-kafka/schema-registry/>
  - <https://docs.confluent.io/platform/current/schema-registry/index.html>
- 메모
  - Schema registry에 대해 간략히 소개해 주는 자료들이었다.
  - 영상을 보며 Schema registry의 필요성에 대해서 공감했고 이게 있으면 많은 일이 간단해 질 것이라고 생각한다.
  - 내부 구현에 있어 Kafka를 스토리지 백엔드로 사용하는 점이 흥미로웠다.
  - 이 내용은 confluent cloud를 이용해서 튜토리얼을 진행해 보아야겠다고 생각했다.

## Day 27 (20220718)

- 공부한 자료
  - <https://www.confluent.io/blog/avro-kafka-data/>
- 메모
  - avro에 대한 소개 그리고 데이터를 다루는 데 있어 schema가 있는 것의 장점을 논하는 글이다. Robustness, Clarity and Semantics, Compatibility Schemas are a Conversation, Schemas Eliminate The Manual Labor of Data Science 등을 장점으로 꼽으며 이야기한다. 전적으로 동의한다. Schema가 있는 장점은, serde 에서 효율적인 것도 있지만 여기에서 주로 이야기하는 데이터 퀄리티 및 정책과 관련된 부분에서 오는 게 더 크다고 생각한다. 이런 부분에 있어서 굉장히 잘 설명한 글이라고 생각하고 나중에 이런 주제를 논할 때 근거 자료로 언급하고 싶은 글이다.
  - binary format의 데이터를 정의하고 serde를 하기 위한 도구로 avro와 함께 언급되는 것으로 protobuf가 있다. 깃헙 star-history를 이용해 protobuf와 avro를 비교해보면 protobuf의 인기가 월등하다. avro의 장점에 대해서는 공감하지만 지금 시점에 둘 중 하나를 골라야 한다면 protobuf 를 고를 것 같다.

## Day 28 (20220725)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/java.html?&_ga=2.100125854.1608897575.1658755782-454449210.1645015938#avro-and-confluent-cloud-schema-registry>
  - <https://github.com/confluentinc/examples/blob/7.2.1-post/clients/cloud/java/src/main/java/io/confluent/examples/clients/cloud/ProducerAvroExample.java>
  - <https://github.com/confluentinc/examples/blob/7.2.1-post/clients/cloud/java/src/main/java/io/confluent/examples/clients/cloud/ConsumerAvroExample.java>
- 메모
  - `gitub.com/confluentinc/examples` 에 좋은 예제가 많은 것 같다.
  - 예제 파일들이 모여있는 깃헙에서 Producer, Consumer 예를 보았다. Avro를 이용해 데이터를 serde 하는 데 엄청 특별한 게 필요하거나 코드가 많이 변하진 않았다. 그런데 그걸로 얻는 효과는 커 보인다. 좋은 인상을 받았다.

## Day 29 (20220730)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/python.html#avro-and-confluent-cloud-schema-registry>
  - <https://docs.confluent.io/5.5.0/quickstart/ce-quickstart.html>
- 메모
  - 파이썬 예제 코드들을 읽어 보았다. 특별히 메모할 부분은 없었다.
  - 예제 코드들을 실행해 보기 위해 confluent platform(local)을 설치해 보았다. 이 때 까지 이걸 활용하지 않은 게 아쉬웠다. 이 도구를 설치하고 활용하면 한 번의 커맨드 입력으로 confluent platform을 로컬에 전부 실행시킬 수 있었다. 이 때까지 conflunet platform을 활용하는 실습들은 건너뛴 것들이 있었는데 이걸 활용하면 쉽게 할 수 있는 것들 이었다. 앞으로 이걸 더 많이, 잘 활용해 보아야겠다.
    - 상태도 저장된다.
    - sdkman으로 설치한 자바 버전이 여러 개라 JAVA_HOME이 제대로 설정되어 있지 않으면 카프카를 실행하다가 에러가 발생한다. `sdk use java 11.0.14.10.1-amzn` 이런식으로 하면 JAVA_HOME이 설정되고, 이 이후에 `confluent local services start` 를 하면 된다.

## Day 30 (20220731)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/rest-proxy.html#avro-and-confluent-cloud-schema-registry>
- 메모
  - confluent platform(local)을 이용해 예제들을 실습해 보았다. Rest Proxy 실습은 이전에 하지 못 했던 것이라 Rest Proxy 실습 겸 이것저것 더 해보았다.
  - 특별한 사항이나 어려운 사항은 없었다.

## Day 31 (20220807)

- 공부한 자료
  - <https://docs.confluent.io/platform/current/tutorials/examples/clients/docs/java-springboot.html#avro-and-confluent-cloud-schema-registry>
- 메모
  - confluent platform(local)을 이용해 예제들을 실습해 보았다.
  - spring kafka는 사용해 본 적이 있었기 때문에 특별한 사항이나 어려운 사항은 없었다.

## Day 32 (20220813)

- 공부한 자료
  - <https://www.confluent.io/blog/using-protobuf-in-confluent-cloud>
- 메모
  - 주로 confluent cloud를 이용하는 내용이었다. 거기서 kafka로 데이터를 produce하고 consume할 때 protobuf를 이용해 value를 serde 하는 예제가 담겨 있는 글이다.
  - Schema lifecycle management 라고 해서, schema registry를 이용해 protobuf 파일들을 관리하고 손쉽게 내려받을 수 있는 기능을 제공하는 게 눈길이 갔다.

## Day 33 (20220816)

- 공부한 자료
  - <https://docs.confluent.io/cloud/current/sr/broker-side-schema-validation.html>
- 메모
  - Kafka Broker에서 schema validation을 하는 예이다. Broker side에서 schema validation을 키고 끄는 방법(CLI, Console)을 설명하고 간단한 예를 보여준다.
  - 동작을 간단히 설명하면, 브로커에 설정이 켜져 있다면 프로듀싱 된 메세지는 항상 연관된 스키마의 ID를 갖고 있어야 한다. 그리고 그 스키마는 토픽과 매칭되어야 한다. 이걸 바탕으로 브로커는 validation을 하고 실패한 메세지는 버려진다.
