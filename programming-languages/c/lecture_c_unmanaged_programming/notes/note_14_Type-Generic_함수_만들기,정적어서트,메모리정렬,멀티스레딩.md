# Type-Generic 함수 만들기, 정적어서트, 메모리 정렬, 멀티스레딩

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## Type-Generic 함수 만들기

- `<tgmath.h>`와 제네릭 선택
  - `<tgmath.h>`
    - 매개변수 형에 알맞는 수학 함수를 찾아서 호출해주는 매크로 함수
    - 컴파일러가 알아서 구현해 준 것
    - 프로그래머가 이런 매크로를 직접 만들 방법이 없었다. 그러나 C11 에서는 가능해졌다.
  - C11
    - 제네릭 선택(generic selection)이라 부른다.
    - _Generic 키워드를 사용한다.
    - 이제 `<tgmath.h>`도 이 키워드를 사용해서 직접 구현이 가능하다.
- _Generic 키워드
  - `_Generic(<제어 표현식>, <연관 목록>)`
    - 컴파일 도중에 여러 가지 표현식 중 하나를 선택하는 방법
      - 실행 중에 선택하는 게 아니다.
    - 매크로 함수의 대체 목록으로 사용하는 게 일반적이다.
- 예

```c
#include <stdio.h>
#include <math.h>

#define ceil(X) _Generic((X),           \
                long double: ceill,     \
                    default: ceil,      \
                    float: ceilf)(X)

int main(void)
{
    float num1 = 3.1415;
    double num2 = 697429.9748;

    printf("ceil(%f) = %f\n", num1, ceil(num1));
    printf("ceil(%f) = %f\n", num2, ceil(num2));
}

```

- 연관 목록
  - 연관 목록의 각 홍목은 다음과 같은 형태를 가진다.
    - `<자료형>: <호출할 함수 이름>` : 여러 개 사용 가능
    - `default: <호출할 함수 이름>`: 딱 하나만. 혹은 생략 가능
  - 대체 규칙
    - 연관 목록에 그 형이 있으면 그에 대응하는 표현식으로 대체
    - 연관 목록에 형이 없고 default: 가 있다면 그에 대응하는 표현식으로 대체
    - default: 도 없다면 컴파일 되지 않음

## 정적 어서트

- 어서트는 코드 작성 시 프로그래머가 세운 가정 또는 선 조건이 올바른지 검사할 때 사용하는 도구이다.
- 어떤 가정(예: int 크기는 4바이트)이 깨질 경우 아예 컴파일이 안 되게 하는 방법이다.
- `_Static_assert(<표현식>, <메시지>)`
- `<assert.h>`를 인클루드 하면 static_assert가 정의되어 있다.
  - `#define static_assert _Static_assert`
- 예

```c
_Static_assert(sizeof(status_t) == 8, "status_t size mismatch");

// 혹은 

static_assert(sizeof(status_t) == 8, "status_t size mismatch");
```

- 표현식이 0으로 평가되면 컴파일 오류가 발생한다.

## _Noreturn

- `_Noreturn <반환형> <함수 이름>(<매개 변수>)`
- `<stdnoreturn.h>`를 인클루드 하면 noreturn을 대신 사용 가능
- void와의 차이점
  - void는 반환 값이 없다는 의미이다. 여전히 함수에서 원래 호출자로 return 하긴 한다. noreturn은 return 하지 않는다는 의미이다.
  - 호출자로 돌아가지 않겠다는 의미이다.
- 다음의 표준 라이브러리 함수들은 호출자로 돌아오지 않는다.
  - abort()
  - exit()
  - _Exit()
  - quick_exit()
  - thrd_exit()
  - longjmp()

## 메모리 정렬

- 시작 주소가 4의 배수로 나눠 떨어지는 메모리를 4바이트로 정렬된(aligned) 메모리라고 한다.
- 프로그래머가 직접 메모리를 정렬할 일이 있을까?
  - 하드웨어에 따라 직접 지정해줘야 하는 경우가 있다.
    - 특정 바이트로 정렬을 하면 성능 향상이 되는 경우
    - 정렬을 하지 않으면 동작하지 않는 경우
- aligned_alloc()
  - `void* aligned_alloc(size_t alignment, size_t size);`
    - alignment: 메모리 시작 주소가 정렬돼야 하는 바이트
    - size: 할당할 바이트의 크기. 반드시 alignment의 배수여야 한다.
    - 반환 값: 성공 시 할당된 메모리 주소, 실패시 널 포인터
    - 실패 조건
      - alignment가 구현에서 유효하지 않거나 지원하지 않는 크기일 때
      - size가 alignment의 배수가 아닐 때
    - 단, C11 첫 버전에서는 널 포인터 반환이 아니다(결과가 정의되지 않는다).
    - 수정 버전(DR 460)부터 널 포인터 반환.
- _Alignas
  - 변수 선언을 할 때 사용하면 스택 메모리에 생기는 변수들도 정렬할 수 있다.
  - `_Alignas(<표현식>)`
  - `<stdalign.h>`를 인클루드 해서 alignas 로 사용할 수 있다.
- 구조체 정렬
  - 두 가지 방법이 있다. 각 멤버를 따로 정렬 하거나 모든 멤버를 일괄적으로 같은 크기로 정렬하는 방법이 있다.
  - 그 외에 구조체 변수를 선언할 때, 그 변수의 시작 주소도 정렬이 가능하다.
  - 구조체 정렬 예1) 멤버 변수 별 정렬

  ```c
  typedef struct data {
      alignas(4096) int num;
      alignas(1024) int dummy[10];
  } data_t;
  ```

  - 구조체 정렬 예2) 모든 멤버 변수 동일 정렬

  ```c
  typedef struct data {
      int num;
      int dummy[10];
  } data_t;

  typedef struct aligned_data {
      int num;
      alignas(4096) int dummy[10];
  } aligned_data_t;
  ```

  - 구조체 변수 선언 시 메모리 정렬

  ```c
  typedef struct data {
      int num;
      int dummy[10];
  } data_t;

  /* ... */ 

  alignas(4006) data_t data;

  ```

  - 동적 할당을 할 때는 주의해야 한다. malloc() 함수는 구조체용 메모리를 할당한다는 사실을 모른다. 그 대신 aligned_alloc()을 사용해야 한다.

- _Alignof()
  - `_Alignof(<자료형>)`
  - `<stdalign.h>`를 인클루드 하면 alignof()로 사용 가능
