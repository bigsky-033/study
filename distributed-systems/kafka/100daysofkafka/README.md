# 100 Days of Code with Apache Kafka

<https://developer.confluent.io/100-days-of-code/> 를 따라 공부한 내용을 정리한 문서입니다. 각 일자 별 공부한 자료들, 정리한 내용을 적고 필요하다면 간략한 메모를 남깁니다.

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
  - 카프카를 사용하는 어플리케이션 개발이란 프로듀서와 컨슈머를 개발하는 것을 의미한다. 토픽에서 메세지를 읽고 쓰기 위한 코드를 작성하는 것이다.
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
