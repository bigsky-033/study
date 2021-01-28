# 전처리기

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 전처리기로 할 수 있는 일들

- 다른 파일을 인클루드
  - 전처리기 지시문 `#include`를 사용
- 매크로를 다른 텍스트로 대체
  - `#define`, `#undef` 와 전처리기 연산자 #, ##를 사용
- 소스 파일의 일부를 조건부로 컴파일
  - 전치리기 지시문 `#if`, `#ifdef`, `#ifndef`, `#else`, `#elif`, `#endif`
- 일부러 오류를 발생시킨다.
  - 전처리기 지시문 `#error`를 사용

## 매크로 대체와 조건부 컴파일

### 매크로 대체: #define

- `#define A (10)`
  - 전처리기가 소스 코드 뒤지다가 A가 보이면 모두 (10)으로 바꿔준다.
- `#define A`
  - 이것도 가능하다. 하지만 바꿔줄 내용이 없다.
  - 그 대신 다른 전처리기 지시어로 A가 정의되어 있는지 판단이 가능하다.
- `#define TRUE(1) #define FALSE(0)`
  - TRUE / FALSE 로 많이 보았다.
- `#define 식별자(매겨변수) 대체_목록`
  - 함수처럼 쓰는 것도 가능하다. 이럴 경우는 매크로 함수라고 한다.

### 매크로 대체: #undef

- `#undef 식별자`
  - 이미 정의된 식별자를 없애는 것이다.
  - 해당 식별자로 정의된 매크로가 없다면 이 지시문은 무시된다.

```c
#define NUMBER (20)
#undef NUMBER

/* 메인 함수 */
printf("%d\n", NUMBER); /* 컴파일 오류 */

```

### 매크로 대체: 미리 정의되어 있는 #define

- 모든 C 구현이 정의하는 것들
  - `__FILE__`: 현재 파일의 이름을 문자열로 표시
  - `__LINE__`: 현재 코드의 줄 번호를 정수형으로 표시
  - 오류 출력 시 사용하는 예
    - `fprintf(stderr, "internal error: %s, line %d.\n", __FILE__, __LINE__);`
  - (C95부터 지원) `__STDC_VERSION__`: 현재 컴파일에 사용 중인 C 표준

### 조건부 컴파일

- 조건에 따라 특정 부분의 코드를 컴파일에 포함 또는 배제
  - `#if 표현식`
  - `#ifdef 식별자` 혹은 `#if defined 식별자`
  - `#ifndef 식별자` 혹은 `#if !defined 식별자`
  - `#elif 표현식`
  - `#else`
  - `#endif`

### 조건부 컴파일: 인클루드 가드

- 어떤 상수를 `#define`으로 정의한다. 그 후 컴파일러에게 조건적으로 코드를 컴파일하라고 지시한다.

```c
#ifndef FOO_H
#define FOO_H

/* codes... */

#endif /* FOO_H */

```

### 조건부 컴파일 예

```c
#ifndef NULL
#define NULL (0)
#endif

#if !defined(NULL)
#define NULL (0)
#endif

#if defined(NULL)
#undef NULL
#endif

#define NULL (0)

```

### 조건부 컴파일에서 주의할 점

```c
#define A

#if defined(A) /* 참 */
#define LENGTH (10)
#endif

#if A /* 거짓 */
#define LENGTH (10)
#endif

```

### 조건부 컴파일: 버전 관리

```c
int spawn_monster(...)
{
    get_monster_skin();
    get_monster_stat();

#if define(FILE_VERSION_2)
    use_custom_skin(...);
#elif defined(FILE_VERSION_3)
    use_custom_voice(...);
#else
    use_default_skin(...);
    use_default_voice(...);
#endif
    calculate_spawn_location();
    return TRUE;
}

```

### 조건부 컴파일: 주석 처리 편하게

```c
#if 0

/* 주석 처리 하고 싶은 코드 */

#end if

```

## 컴파일 오류 만들기, 컴파일 중에 매크로 정의하기

- `#error 메세지`
  - 컴파일 도중에 강제로 오류를 발생시키는 매크로
  - 메세지를 꼭 따옴표로 감쌀 필요는 없다.

  ```c
  /* version.h */
  #define VERSION 10

  /* builder.h */
  #if VERSION != 11
  #error "unsupported version"
  #endif

  ```

### 컴파일 중에 매크로 정의하기

- 컴파일 도중에 -D 옵션으로 전달 가능
  - `clang -std=c89 -W -Wall -pendatic-errors -DA *.c`
  - `#define A (1)`과 똑같은 결과 (`#define A`가 아님)
  - 직접 대체할 값을 지정할 수도 있다.
    - `clang -std=c89 -W -Wall -pendatic-errors -DA=52 *.c`
      - `#define A (52)`와 똑같은 결과이다.
- 배포용으로 컴파일하기: `-DNDEBUG`
  - `clang -std=c89 -W -Wall -pendatic-errors -DNDEBUG *.c`
  - 배포(release) 모드로 실행파일을 컴파일하라고 알려주는 매크로
    - NDEBUG: '디버그가 아니다(not debug)'라는 뜻이다.
    - assert()가 사라진다.
    - 디버그 모드에서만 실행될 코드는 `#if !defined(NDEBUG)` 속에 넣을 것
  - 이 대신 다음과 같은 매크로를 직접 정의해서 사용하는 프로젝트도 많다.
    - DEBUG: 디버그용 빌드
    - RELEASE: 배포용 빌드
    - 기타: 필요에 따라 다양한 빌드를 지정

## 매크로 함수

- #define을 할 때 `대체 가능한 매개변수 목록`을 받아 함수처럼 쓴다.
  - `#define SQUARE(a) a * a`
  - `#define ADD(a,b) a + b`

  ```c
  /* main.c */
  int num1;
  int num2;
  int result;

  num1 = 10;
  num2 = 20;
  result = ADD(num1, num2); /* 30 */
  result = SQUARE(num1); /* 100 */

  /* 주의 */
  result = 10 * ADD(num1, num2); /* 300이 아니라 120이 나온다. 매크로는 복붙되는 개념이기 때문! */

  ```

  - 위의 마지막 경우와 같은 실수를 방지하기 위해, 매크로 함수의 구현은 소괄호로 감싸주는 게 좋다.
    - `#define SQUARE(a) (a * a)`
    - `#define ADD(a,b) (a + b)`

## 매크로 함수의 활용

- 매크로가 여러 줄이면? => \ 를 사용하면 된다.

  ```c
  #define POW(n,p,i,r) r = 1;                     \
                        for (i = 0; i < p; ++i) { \
                          r *= n;                 \
                        }                         \
  ```

- 매크로 함수의 활용: 어서트 재정의

  ```c
  #define ASSERT(condition, msg)                                \
    if (!(condition)) {                                         \
      fprintf(stderr, "%s(%s: %d)\n", msg, __FILE__, __LINE__); \
      __asm { int 3 }                                           \
    }                                                           \

  ```

### 전처리기 명령어: # 명령어

- `#define 식별자(매개변수) 대체_목록`
  - 매개변수 자체를 문자열로 바꾸어 준다.
  - 매개변수를 쌍따옴표로 감싸는 것
    - 예: `#define str(s) #s`

### 전처리기 명령어: ## 명령어

- `#define 식별자(매개변수) 대체_목록`
  - 대체 목록 안에 있는 두 단어를 합쳐서 새로운 텍스트로 바꾼다. 말 그대로 복붙이다.
    - 단어는 매개변수 일 수도 아닐 수도 있다.
  - `#`와 달리 문자열 데이터를 만들어 주는 게 아니다.
    - 예: `#define print(n) printf( "%d\n", g_id_##n)`

### 매크로 함수의 장점과 단점

- 장점
  - 함수 호출이 아니라 곧바로 코드를 복붙하는 개념이다. 따라서 함수 호출에 따른 과부하가 없다.
  - 다른 언어에서는 되었지만 안 되는 기능들을 매크로를 이용해 해결 가능 할 수도 있다.
- 단점
  - 디버깅이 아주 어렵다. 브레이크 포인트 같은 게 동작하지 않는다.
