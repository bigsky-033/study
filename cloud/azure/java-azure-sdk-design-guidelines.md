# Java Azure SDK Design Guidelines

이 문서는 Java Azure SDK Design Guidelines 문서의 요약을 담는다. *Design Principles*를 제외한 나머지 부분은 내용을 간략히 다룬다.

- References
  - Java Azure SDK Design Guidelines, <https://azure.github.io/azure-sdk/java_introduction.html>
  - Azure-sdk-for-java Github Repository, <https://github.com/Azure/azure-sdk-for-java>
- 문서에서 `*` 를 통해 표시한 내용은 개인적인 생각을 담은 부분이다.
- 내용을 한글로 옮길 때 있어 모호한 부분이 있거나 한글 번역이 어색해 보이는 부분은 괄호 안에 영어를 표기했다.

## Design Principles

Azure SDK가 중요시 여기는 가치는 생산성이다. 완전성, 확장성 그리고 성능과 같은 것들은 물론 중요한 가치이지만 생산성보다 우위에 두기는 어렵다. 고객의 생산성을 달성하기 위해 Azure SDK가 가져야 할 특징들은 다음과 같다.

- **Idiomatic**
  - Azure SDK는 일반적인 자바 개발의 컨벤션을 따라야 한다. 자바 개발자들에게 SDK가 자연스럽게 느껴져야 한다.
  - 생태계와 함께 한다. 생태계의 강점과 결점을 받아들인다. 그리고 개발자를 위한 개선을 위해 생태계와 협력한다.
  - Azure SDK의 버전은 일반적인 자바 라이브러리의 버전과 같은 식이다.

- **Consistent**
  - Azure SDK는 개별적인 라이브러리가 아니라, 단일 팀의 단일 제품 처럼 느껴져야 한다.
    - \* Azure java sdk repository에 보면 각 azure의 각 서비스 별로 클라이언트가 존재한다. 이 많은 클라이언트의 코드들이 일정한 패턴을 따르고 그로 인해 코드를 읽는 사람이 코드 작성 패턴에 익숙하다면 어느 서비스의 클라이언트 코드를 보더라도 마찬가지로 익숙함을 느끼게 해야한다는 말인 것 같다.
  - 사용자는 한 번 전반적인 컨셉을 학습하고 나면 관련된 지식을 모든 SDK 구성 요소에 적용할 수 있어야 한다.
  - 만약 가이드라인과 차이가 발생하는 지점이 있다면 합리적인 이유가 있어야 한다.
    - \* 예를 들어 아래의 *Dependable* 에 보면 100% 하위 호환성을 지켜야 한다는 말이 있다. 때로는 하위 호환성을 지키기 어려운 경우도 있을텐데 그런 경우에 적절한 이유가 있어야 함을 말하는 것 같다.

- **Approachable**
  - 처음 시작하기 위한 단계의 수가 적어야 한다
    - \* 처음 사용하는 사람도 쉽게 시작할 수 있어야 하고 또한 이미 이런 종류의 도구에 익숙한 사람도 쉽게 익숙해질 수 있어야 함을 의미하는 것 같다.
  - 컨셉 그리고 타입의 수가 너무 많지 않아야 한다. 그리고 구성원이 너무 많지 않아야 한다.
    - \* small number of members 는 무엇을 의미하는 지 잘 모르겠다.
  - Azure SDK 구성요소를 설계하는 엔지니어가 아니라 사용자에 의해 접근이 가능해야 한다.
  - *Getting started* 가이드와 예제를 쉽게 찾을 수 있어야 한다.
  - 쉽게 얻을 수 있어야 한다.

- **Dependable**
  - 100% 하위 호환성을 지켜야 한다.
  - 로깅, 트레이싱 그리고 에러 메세지가 great 해야 한다
    - \* 문제를 쉽게 추적할 수 있게 되어 있어야 한다.
  - 라이프 사이클, 기능의 지원 범위와 품질에 있어 예측 가능성을 제공해야 한다.

## Azure SDK API Design

Azure의 서비스는 자바 개발자들에게 하나 이상의 *service client type*과 *supporting types*을 통해 노출된다.

### Service Clients

개발자들이 *Azure의 서비스*(\* 이하 서비스)을 Azure SDK를 통해 접근하게 되면 진입점(main starting points)이 되는 곳이 *Service Clients*(\* 이하 서비스 클라이언트 또는 클라이언트) 이다. 때문에 이 클래스들은 찾기 쉬워야 한다. 각 서비스의 클라이언트 라이브러리는 메인 네임스페이스에 클라이언트를 하나 이상 가지고 있어야 한다. 가이드라인의 이 섹션에서는 service client를 디자인 하기 위한 패턴들을 다룬다. 자바에서는 synchronous 클라이언트와 asynchronous 클라이언트가 별도로 각각 필요하다. 때문에 이 섹션에서는 일반적인 service client 에 대한 가이드, sync 클라이언트에 대한 가이드 그리고 async 클라이언트에 대한 가이드 순서로 가이드를 진행한다.

서비스에게 HTTP(또는 기타의 방법)로 요청을 보낸다고 해서 전부 서비스 클라이언트가 되는 것은 아니다. 서비스 클라이언트의 지정은 서비스에 고유하게 표시된다. 때문에 서비스 클라이언트로 지정이 되는 클래스들은 직접 구성될 수 있는(directly constructed) 클래스여야 한다. 또한 서비스 클라이언트 지정은 클라이언트를 직접 생성하는 것이 적절한 경우에만 적용된다. 리소스를 고유하게 식별할 수 없거나 타입을 직접 작성할 필요가 없는 경우에는 서비스 클라이언트 지정을 적용하면 안 된다.

- 서비스 클라이언트의 이름은 *Client*로 끝나야 한다. 예를 들어 `ConfigurationClient`와 같은 식이다.
- 서비스 클라이언트 클래스에는 항상 `@ServiceClient` 라는 어노테이션이 붙어야 한다.
- 사용자가 상호 작용(interact)할 가능성이 가장 높은 서비스 클라이언트 타입을 클라이언트 라이브러리의 루트 패키지 내에 배치해야 한다. 예를 들어 `com.azure.<group>.servicebus` 와 같은 식이다. 특화된 서비스 클라이언트들은 서브 패키지에 두어야 한다.
  - \* `com.azure.messaging.servicebus` 를 보면, 이 패키지 바로 아래에 `ServiceBusClientBuilder` 와 이 빌더가 생성하는 클래스들이 있다. 그리고 서브 패키지에 `administration(com.azure.messaging.servicebus.administration)` 이 있고 이 밑에 `ServiceBusAdministrationClientBuilder`와 이 빌더가 생성하는 클래스들이 있다. 이런 구조를 말하는 게 아닐까 싶다.
- 모든 서비스 클라이언트 클래스들은 불변(immutable)인 것 그리고 상태가 없어야(stateless) 한다.
- Sync API들을 제공하는 클라이언트 클래스와 Async API들을 제공하는 클라이언트 클래스는 별도의 클래스로 분리되어 있어야 한다.

#### Sync Service Clients

- Sync 서비스 클라이언트 클래스의 이름은 `<ServiceName>Cleint`와 같은 식이어야 한다. 단일 서비스에 대해 둘 이상의 서비스 클라이언트를 제공할 수 있다.
- **상세 구현 관련**
  - 생성자의 접근 제어자는 package-private 이어야 한다. 외부에서 인스턴스를 만들 때는 builder를 이용하도록 해야 한다.
  - 내부에 async client를 포함하고 있는 식으로 구현할 수 있다. 예를 들어 클래스 내부에 인스턴스 변수로 `<service_name>AsyncClient client` 와 같은 변수를 포함하는 식이다. 그리고 제공하는 public API에서는 `return client.<service_operation>(<parameters>).block();` 와 같은 식으로 호출해 결과를 돌려주는 식으로 구현할 수 있다.
  - List API에는 페이징을 지원하는 방식과 지원하지 않는 방식이 있다. 페이징을 지원하지 않는 방식에서는 `IterableStream<<model>>` 를 리턴하게 하고 페이징을 지원하는 방식에서는 `PagedIterable<<model>>`을 리턴하게 구현한다.

#### Async Service Clients

- Async 서비스 클라이언트 클래스의 이름은 `<ServiceName>AsyncClient`와 같은 식이어야 한다. 단일 서비스에 대해 둘 이상의 서비스 클라이언트를 제공할 수 있다.
- 비동기 API를 만들 때 *Project Reactor* 를 이용해서 구현해야 한다. `CompletableFuture` 또는 `RxJava`와 같은 API를 사용하지 않는다.
  - \* 비동기 API의 모든 리턴 타입이 Publisher 이어야 하고 Mono 또는 Flux의 형태여아 한다는 말이다. 흥미로운 부분이었다.
- 스트리밍이나 비동기 동작의 구현을 위해 custom API를 작성하지 않는다. Azure core library에 있는 기능을 이용한다.
  - \* 서비스 클라이언트의 역할과 책임을 명확히 분리한 부분인 것 같다. 서비스 클라이언트는 사용자와의 상호 작용을 위한 부분에 최대한 신경쓰고 그것을 잘 해야 하는 것이고 실제 통신과 관련된 부분은 core library가 담당하는 것 같다.
- **상세 구현 관련**
  - 생성자의 접근 제어자는 package-private 이어야 한다. 외부에서 인스턴스를 만들 때는 builder를 이용하도록 해야 한다.
  - `@ServiceClient` 어노테이션의 값 중 `isAsync`가 true 여야 한다.

#### Service Client Creation

- 서비스 클라이언트에는 `public`이나 `protected`와 같은 접근 제어자를 가진 생성자가 존재해서는 안 된다. package-private 수준의 접근 제어자를 제공하도록 한다(같은 패키지 내에서는 생성자를 직접 호출할 수 있게 한다). 그리고 service client builder를 통해서 인스턴스를 만들어 사용하게 한다.
- 서비스 클라이언트 빌더 클래스의 이름은 `<ServiceName>CleintBuilder`와 같은 식이어야 한다.
- 서비스 클라이언트 빌더는 서비스 클라이언트를 만들기 위해 유연한(fluent) API를 제공해야 한다. 그리고 sync 서비스 클라이언트 인스턴스를 만들기 위한 `buildClient()` API와 async 클라이언트 인스턴스를 만들기 위한 `buildAsyncClient()` API를 각각 제공해야 한다.
- **상세 구현 관련**
  - 만약 서비스 클라이언트 빌더가 여러 타입의 서비스 클라이언트를 생성할 수 있는 경우 다음과 같은 패턴으로 build 메서드를 이름지어야 한다: `build<client>Client()`, `build<client>AsyncClient()`.
  - 서비스 클라이언트 빌더에는 항상 `@ServiceClientBuilder` 라는 어노테이션이 붙어야 한다. 이 어노테이션에는 serviceClient 라는 값을 통해 이 클라이언트 빌더를 통해 생성할 수 있는 서비스 클라이언트를 명시할 수 있다. 예를 들어 다음과 같은 식이다: `@ServiceClientBuilder(serviceClients = {<service_name>Client.class, <service_name>AsyncClient.class})`
  - API의 consistency를 위해 빌더의 fluent API 에서는 일반적으로 통용되는 함수의 이름을 사용하는 것을 권장한다.
    - \* HTTP와 AMQP based 클라이언트에 대해, 널리 쓰이는 함수의 이름들이 공식 가이드 문서에 정의되어 있다. 이 목록을 참고하자.
  - 빌더 함수가 상호 배타적인(mutually exclusive)한 인자를 받았을 때 `IllegalStateException`을 던지게 하라. 그리고 에러 메세지를 통해 문제가 무엇인지 명확히 전달해야 한다.
  - 사용자가 서비스를 연결하고 인증하는 데 필요한 최소한의 정보로 서비스 클라이언트를 만들 수 있도록 한다.
    - \* 디폴트 값들을 잘 세팅해주어야 한다는 의미인 것 같다.
  - 빌더가 만든 서비스 클라이언트는 항상 유효한(valid) 상태여야 한다. 만약 사용자가 `build*()` 를 호출했을 때 설정 값들이 불충분하거나 유효하지 않으면 `IllegalStateException`이 던져져야 한다.

##### Service Versions

- 별도의 설정이 없을 경우 지원되는 API의 버전 중 가장 높은 버전의 API를 호출하게 한다. 그리고 이것에 대한 문서화가 잘 되어 있어야 한다.
- 서비스 클라이언트를 만들 때 사용자가 버전을 지정해줄 수 있게 한다.
  - 빌더의 퍼러미터 중 `serviceVersion => ex) .serviceVersion(<ServiceName>Version.V1_0);` 를 이용해서 지정할 수 있게 하도록 한다
- 서비스의 버전은 `ServiceVersion` 인터페이스를 구현하는 enum 을 통해 표현한다.
  - \* `com.azure.core.util.ServiceVersion` 를 보면 getVersion 함수를 가지는 인터페이스이다.

#### Service Methods

- 서비스 메서드는 서비스에 대한 어떤 작업(operation)을 호출하는 메서드이다. 일반적으로 서비스 메서드들은 Client로 끝나는 클래스들에 있지만 다른 리소스 클래스에서도 찾을 수 있다.

- **상세 구현 관련**
  - 비동기 서비스 메서드에 `Async`를 붙이는 방식으로 구현하지 마라. 대신 비동기 호출의 경우에는 별도로 존재하는 Async 서비스 클라이언트를 사용하게 하라.
  - *CRUD operation*의 경우에는 일반적인 네이밍 규칙이 존재한다. 이 내용을 참조해서 구현하자.
    - \* 공식 문서에 있는 표를 참고하자.
  - 중요한 것은 개발자 경험에 가장 적합한 이름을 사용해야 한다는 것이다. 이에 대해서는 유연함을 가져야 한다. 일반적인 자바 개발자의 경험에 비추어 보았을 때 관용적이지 않은 네이밍이 발생하지 않도록 한다. 예를 들어 자바 개발자는 복수의 값을 가져올 때 `getAll` 보다 `list` 라는 이름을 선호한다.
  - Sync 클라이언트에 대해서는 `com.azure.core.util`을 사용하는 오버로딩 메서드를 제공하자. 중요한 점은 sync 클라이언트에 대해서만 제공한다는 점이다. `Context`는 varargs를 사용하는 경우를 제외하고 서비스 메서드의 마지막 인자여야 한다. 서비스 메서드에 여러 오버로딩이 있는 경우에는, 가장 많은 인자를 갖고 있는 함수에만 `Context` 인자를 추가로 넣는 오버로딩을 구현해야 한다.
    - \* 이를 표현하기 위해 *maximal overLoad* 라는 표현을 사용한다.
  - Async 클라이언트에 대해서는 `Context`를 포함하는 오버로딩을 제공해서는 안 된다. Async 클라이언트는 Reactor에서 제공하는 `Subscriber Context`를 사용해야 한다.
    - \* 쓰레딩 모델에서 오는 차이 때문인 것 같다. `ThreadLocal`의 workaround 인 것 같다.

#### Non-Service Methods

서비스 클라이언트들은 종종 서비스 메서드 외에 다른 메서드들을 갖고 있는 경우가 있습니다. 예를 들어 service version, http pipeline 과 같은 것들의 상세(detail)에 접근하는 경우에 서비스 메서드 외의 다른 메서드들을 사용합니다. 또는 사용자에게 특수한 목적을 가진 서브 클라이언트를 만들 수 있는 기능을 제공하는 API도 있을 수 있습니다. 이런 서비스 클라이언트들은:

- 서비스 메서드가 아닌 모든 메서드에 표준 자바 빈(JavaBean)의 접두사(prefix)를 사용하라.
- Sync 서비스 클라이언트에서 *create or vend* 서비스 클라이언스를 하는 경우엔 `get`이라는 접두사를 붙이고 `Client` 라는 접미사(suffix)를 붙여라. 예를 들어 `container.getBlobClient()`와 같은 식이다. 이와 유사하게 async 서비스 클라이언트에서는 `get`이라는 접두사를 붙이고 `AsyncClient` 라는 접미사를 붙여라. 예를 들어 `container.getBlobAsyncClient()`와 같은 식이다. `Client` 라는 접미사를 사용할 때는 주의해야 한다. `Client`라는 접미사는 서비스 클라이언트에게만 사용될 수 있기 때문에, 타입이 클라이언트가 아닌 경우 메서드의 이름을 `get*Client`로 지으면 안 된다.

##### Cancellation

- Sync 그리고 async 클라이언트 모두에 cancellation token을 받는 API를 만들지 마라. Cancellation은 자바에서 흔한 패턴이 아니다. 대신 리퀘스트를 cancel 하는 API가 필요한 경우 사용자가 async 클라이언트를 사용하게 하라. Async 클라이언트는 unsucribe 하거나 리퀘스트를 cancel 할 수 있다.
  - \* Async API에서 제공되는 이런 기능은 reactive programming 그리고 Reactor와 관련된 내용이다.

##### Return Types

서비스에 대한 요청은 기본적으로 두 종류가 있다. 단일 논리적 요청(single logical request)를 만드는 경우와 결정적인 요청 시퀀스(deterministic sequence requests)를 만드는 경우이다. 단일 논리적 요청의 예는 그 작업 내부에서 리트라이를 할 수 있는 요청이 있다. 결정적인 요청 시퀀스의 예로는 페이징 작업이 있다.

- \* 용어가 조금 낯선데, 쉽게 구분할 수 있는 방법은 서비스 클라이언트의 리턴 타입이다. 단일 값을 리턴하는지 또는 여러 개의 값을 리턴하는지로 구분할 수 있을 것 같다.

논리적 엔티티(logical entity)는 프로토콜 중립적인(protocol neutral) 응답의 표현이다. 논리적 엔티티의 데이터는 헤더, 바디 그리고 status line을 합해서 구성할 수 있다. 예를 들어 `ETag` 헤더와 같은 것을 논리적 엔티티의 프로퍼티로 노출할 수 있다. `Response<T>` 같은 경우를 *완전한 리스폰스(complete resonse)*라고 한다. 이 객체는 HTTP 헤더들, status code 그리고 리스폰스 바디로부터 만들어진 `T` 객체를 포함한다. 여기서 T 객체가 logical entity 될 수도 있다.

- \* 쉽게 말해 HTTP의 경우, 논리적 엔티티는 리스폰스에서 필요한 것들을 모아 사용자가 사용하기 쉽게 만들어진, 사용자 수준의 객체인 것 같다.

- **상세 구현 관련**
  - **Synchronous** 한 서비스 메서드의 경우 logical entity(`T`)를 리턴해야 한다.
  - 네트워크 요청을 만드는 **asyncronous** 서비스 메서드의 경우 `Mono`로 감싼 logical eneity(`T`)를 리턴해야 한다.
  - 서비스 메서드 중 Maximal overload 메서드가 `Response<T>`를 리턴할 수 있게 해라.
  - 여러 요청을 단일 호출로 결합하는 메서드의 경우:
    - 그 메서드가 특정 호출의 메타데이터를 요청하는 게 분명하지 않은 경우에는 개별 요청 단위의 HTTP 메타 데이터(헤더 등)을 전달하지 마라.
    - 호출 실패 시 개발자가 적절한 조치를 취할 수 있는 충분한 정보를 제공해야 한다. 여기에는 무엇이 잘못되었는지를 설명하는 메세지와 취해야 할 조치에 대한 세부 정보가 포함되어야 한다.

#### Service Method Parameters

##### Option Parameters

서비스 메서드의 경우에는 퍼러미터의 수와 복잡성에 따라 두 가지 그룹으로 분류할 수 있다:

- 단순한 인풋을 사용하는 경우: *simple methods*(심플 메서드)
  - 심플 메서드는 최대 여섯개 까지의 퍼러미터를 받을 수 있고 대부분의 경우가 단순한 원시 타입(prmitive types)인 경우를 의미한다.
  - 심플 메서드는 퍼러미터 리스트와 오버로딩 디자인에 관련해 일반적인 자바의 베스트 프랙티스를 따른다.
- 복잡한 인풋을 사용하는 경우: *complex methods*(컴플렉스 메서드)
  - 컴플렉스 메서드의 경우에는 많은 수의 퍼러미터를 받는 경우인데 일반적으로 복잡한 요청 페이로드의 REST API에 해당한다.
  - 컴플렉스 메서드는 리퀘스트의 페이로드를 표현하기 위해 *option parameter*를 도입해야 한다. 이 문서에서는 *option pattern*이라고 불린다.

- **상세 구현 관련**
  - Options에 대하여:
    - Option parameters를 이름 지을 때 서비스 메서드의 이름을 고려해서 짓는다. `<operaiton>Options`과 같은 식이다. 예를 들어 Option parameter가 사용될 메서드가 `createBlob`이라면, `CreateBlobOptions`가 될 수 있다.
    - Options parameter 패턴은 컴플렉스 메서드에서 사용한다.
      - 미래에 퍼러미터가 늘어날 것이라고 예상되는 경우 심플 메서드에서도 options parameter 패턴을 사용할 수 있다.
      - Options parameter 패턴을 사용하여 메서드의 단순한 오버로드 메서드를 추가할 수 있다.
    - 일반적인 시나리오에서 사용자는 options parameter가 나타내는 작은 부분 집합만 전달할 가능성이 있는 경우라면 이 부분 집합만을 나타내는 매개 변수 목록으로 오버로드를 추가하는 것이 좋을 수 있다.
  - 클라이언트 측에서만 사용하는 매개 변수(`Context`, timeout 등)를 제외하고 option parameters 뿐만 아니라 일부 매개 변수를 사용하는 메서드 오버로드를 도입하지 마라.
  - Options parameter가 있다면 `*WithResponse` 메서드에 대해 사용하라. 만약 없다면, 단순히 `*WithResponse` 메서드를 위해 만들지는 마라.
  - Options 타입들을 루트 레벨의 `models` 패키지에 둬라. 이는 root-level 에 너무 많은 너무 많은 패키지가 생기는 걸 방지하기 위함이다.
  - Model class type에 적용된 것과 마찬가지로 options type에 대해서도 일반적인 JavaBean의 네이밍 컨벤션을 적용하라(`get*`, `set*`, `is*` 등). 또한 required 퍼러미터를 지원하기 위한 생성자 오버로딩을 지원할 수 있다.

##### Parameter Validation

서비스는 보통 두 종류의 퍼러미터들을 받는다: service parameters와 client parameters이다.

- *Service parameters*는 서비스에 직접 전달되는 퍼러미터들이다.
- *Client parameters*는 클라이언트 라이브러리에서 사용되는 퍼러미터들이다. 이것들은 서비스에게 직접 전달되지 않는다.

- **상세 구현 관련**
  - Client parameters 를 검증하라.
  - Service parameters 는 검증하지 마라. Null checks, empty string check 그리고 일반적인 검증 조건들 모두를 검증하지 마라. 서비스가 이를 검증하게 두어라.
  - 적절하지 않은 service parameters가 전달되었을 때의 개발자 경험을 테스트해 보아야 한다. 적절한 것은 클라이언트로부터 깔끔한 에러메세지가 전달되는 것이다. 만약 개발자 경험이 좋지 못 하다면, 서비스 팀과 상의하여 문제를 고쳐라.

#### Methods Returning Collections (Paging)

많은 Azure REST API가 배치 또는 페이징을 통해 데이터 컬렉션을 반환한다. 클라이언트 라이브러리는 그런 API를 호출하는 메서드를 `PagedIterable<T>`, `PagedFlux<T>` 또는 그것들의 부모 타입들을 반환하는 함수를 통해 노출한다. 이러한 타입들은 azure-core library 에 정의되어 있다.

- **상세 구현 관련**
  - 서비스 메서드 중 syncronous한 방식으로 데이터 컬렉션을 반환할 수 있는 경우엔 `PagedIterable<T>`를 반환하게 하라. 예를 들어 다음과 같은 식이다. `public PagedIterable<ConfigurationSetting> listSettings(...) {`
    - `PagedIterable`를 사용하면 사용자는 일반적인 `Iterable`를 사용하는 것 처럼 for loop syntax를 사용하는 방식을 통해 코드를 작성할 수 있다. 또한 자바의 `Stream`을 사용할 수도 있다. 사용자는 또한 `streamByPage()` 및 `iterableByPage()`를 통해 페이지간 경계를 나눠 데이터를 사용할 수도 있다. `<serviceName>PagedIterable` 또는 `<operation>PagedFlux` 와 같은 이름 규칙을 따르는 한 서브 클래스 또한 허용될 수 있다.
  - `List`, `Stream`, `Iterable` 또는 `Iterator` 와 같이 다른 컬렉션의 타입을 리턴하게 하지 마라.
  - 서비스 메서드 중 asyncronous한 방식으로 데이터 컬렉션을 반환할 수 있는 경우엔 `PagedFlux<T>`를 반환하게 하라. 이 서비스가 페이징을 지원하지 않는다고 하더라도 이 타입을 반환하게 하라. 사용자가 일관된 방식으로 사용할 수 있게 하기 위함이다.
    - `PagedFlux.byPage()`는 `continuationToken`를 받을 수 있는 오버로딩을 제공한다. 이 함수는 이 토큰을 통해 알 수 있는 페이지부터의 데이터를 가져온다.
  - 적절한 경우라면, 사용자에게 서비스 별 추가 API를 제공하기 위해 azure-core에 있는 클래스들에 대한 서브 클래스를 만들 수도 있다. 예를 들어 `SearchPagedFlux` 또는 `SearchPagedIterable`와 같은 식이다. 이름에서 보이는 것 처럼 서비스의 이름을 접두사로 사용해야 한다. 그리고 이런 서브 클래스들은 일반적으로 서비스의 root package 아래 `util` 패키지 아래에 둔다.

#### Methods Invoking Long-Running Operations

#### Conditional Request Methods

##### Safe conditional requests (e.g. GET)

##### Unsafe conditional requests (e.g. POST, PUT, or DELETE)

#### Hierachical Clients

문서에 TODO로 되어 있음.
