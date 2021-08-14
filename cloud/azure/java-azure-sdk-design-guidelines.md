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
- 서비스 클라이언트 클래스에는 항상 @ServiceClient 라는 어노테이션이 붙어야 한다.
- 사용자가 상호 작용(interact)할 가능성이 가장 높은 서비스 클라이언트 타입을 클라이언트 라이브러리의 루트 패키지 내에 배치해야 한다. 예를 들어 `com.azure.<group>.servicebus` 와 같은 식이다. 특화된 서비스 클라이언트들은 서브 패키지에 두어야 한다.
  - \* `com.azure.messaging.servicebus` 를 보면, 이 패키지 바로 아래에 `ServiceBusClientBuilder` 와 이 빌더가 생성하는 클래스들이 있다. 그리고 서브 패키지에 `administration(com.azure.messaging.servicebus.administration)` 이 있고 이 밑에 `ServiceBusAdministrationClientBuilder`와 이 빌더가 생성하는 클래스들이 있다. 이런 구조를 말하는 게 아닐까 싶다.
- 모든 서비스 클라이언트 클래스들은 불변(immutable)인 것 그리고 상태가 없어야(stateless) 한다.
- Sync API들을 제공하는 클라이언트 클래스와 Async API들을 제공하는 클라이언트 클래스는 별도의 클래스로 분리되어 있어야 한다.

#### Sync Service Clients

#### Async Service Clients

#### Service Client Creation

#### Service Methods

#### Non-Service Methods

#### Methods Returning Collections (Paging)

#### Methods Invoking Long-Running Operations

#### Conditional Request Methods

#### Hierachical Clients
