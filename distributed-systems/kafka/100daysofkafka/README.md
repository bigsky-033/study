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
