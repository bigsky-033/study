# C언어 기본 문법 2, 빌드 단계

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 함수

- C 언어 함수 맛보기

    ```c
    #include <stdio.h>

    static float s_gas = 500.f;

    void honk(void)
    {
        printf("Honk~ honk~);
    }

    void reduce_gas(float consumed_gas)
    {
        s_gas -= consumed_gas;
    }

    ```

- C는 접근 제어자(public, private 등)가 없다.
- C의 함수는 기본적으로 모두 전역(global) 함수이다.
- 함수 오버로딩도 존재하지 않는다. 따라서 필요할 경우 함수 이름을 다르게 만들어야 한다.
- C는 언제나 위에서 아래로 코드를 본다. ANSI C(C89)에서 함수 정의가 등장하기 전 그 함수를 호출하면 컴파일러가 다음과 같이 가정한다.
  - 반환형은 int
  - 매개 변수는 아무 것이나 올 수 있다.
  - 따라서 나중에 컴파일러가 int가 아닌 다른 것을 반환하는 함수를 찾으면 컴파일 오류를 발생시킨다.

## 함수 정의의 문제

- 위에서 본 문제를 해결하기 위해서 그러면...호출하는 코드 위에 전부 함수를 위치시켜야 할까? 다행히 그렇지는 않다. 아래에 나오는 함수 선언이 이런 문제를 해결해 준다.

## 함수 선언

- 함수의 구현체 없이 함수 원형(prototype)만 선언해 주는 것이다.
  - 원형이란?
    - 함수의 이름
    - 반환형
    - 매개 변수들의 자료형
      - 매개 변수 이름은 없어도 된다. 단, 가독성을 위해 일반적으로 변수 이름도 넣는 게 좋다.
  - 함수 정의(definition)은 실제로 함수를 구현해놓은 것이다.
    - 당연히, 함수 정의는 함수 선언이기도 하다.

```c

/* 함수 선언과 정의를 분리 */

#include <stdio.h>

void foo(void);

int main(void)
{
    foo();
    getchar();
    return 0;
}

void foo(void)
{
    printf("foo called");
}

/* 함수 선언과 정의가 하나 */

#include <stdio.h>

void foo(void)
{
    printf("foo called");
}

int main(void)
{
    foo();
    getchar();
    return 0;
}

```

- 함수를 선언한다는 것은 함수를 사용하기 전에 그 함수를 선언한다는 것이다. 보통은 파일의 제일 위에 선언한다.
  - forward declaration(전방 선언) 이라는 표현으로 많이 쓴다.
- 또는 헤더 파일에 넣을 수도 있고 이게 더 좋은 방법이다.
- 전방 선언의 작동 원리
  - 컴파일 시 실제 어디로 가서 코드를 찾아야 하는지는 모르니 구멍으로 둔다.
  - 컴파일 다음 단계인 링크(link) 단계에서 실제 코드 위치를 찾아서 그 구멍을 채운다.
- C89 표준에서 기본적으로 가정하는 사항을 충족하는 경우(int 를 반환하는 함수)에는 함수 선언을 하지 않아도 될까?
  - 그러지 말자. 함수 선언은 언제나 하도록 하자.
  - C99 표준부터는 int라는 가정을 하지 않는다.
  - 어떤 컴파일러는 경고만 주고, 컴파일은 허용할 수도 있다.
  - 모든 컴파일러가 그렇단 보장이 없다.

## 함수 매개변수 평가 순서, 피연산자 평가 순서

- 실수할 가능성이 높은 부분이다. 주의해서 보도록 하자.

```c
#include <stdio.h>

int add(int op1, int op2)
{
    printf("add()\n");
    return op1 + op2;
}

int subtract(int op1, int op2)
{
    printf("subtract()\n");
    return op1 - op2;
}

int main(void)
{
    int num1 = 128;
    int num2 = 256;

    printf("%d, %d\n", add(num1, num2), subtract(num1, num2));

    return 0;
}

```

- 위의 코드는 어떤 순서로 출력이 될까? => 모른다.
  - 경우 1) add()가 먼저 평가
  - 경우 2) subtract()가 먼저 평가
- 표준에 따르면, 함수 매개 변수의 평가 순서는 명시되어 있지 않음(unspecificed)
  - undefined 와 unspecificed 는 다르다. unspecified는 문제에 대해 가능한 경우의 수를 정의는 해 놓는다. 단 그 경우의 수 중 어떤 것이 옳은지는 정의하지 않는다.
- 즉, 컴파일러에 따라 평가 순서가 달라질 수 있다.
- `printf()`가 실제로 실행되기 전에 `add()`와 `subtract()`가 실행가 실행되는 것은 보장된다. 단, add와 subtract 중 어떤 것이 먼저 실행될지는 보장하지 않는다. 컴파일러에 따라 다를 수 있다는 말이다.
- 따라서 같은 라인에서 왼편에 있는 함수가 먼저 실행될 것이라 가정하고 코드를 짜는 것은 버그를 유발할 수 있다.
  - 예를 들어, 두 함수가 같은 변수를 사용(read/write)하고 호출 순서에 의존적이라고 할 경우, 컴파일러에 따라 함수 실행 순서가 달라질 수 있고 의도와는 다른 결과를 만들 수 있다.

```c
#include <stdio.h>

float divide(int op1, int op2)
{
    printf("%d / %d = ", op1, op2);
    return op1 / (float) op2;
}

int main(void)
{
    int num = 0;
    float result = divide(++num, ++num);
    printf("%f\n", result);

    return 0;
}

```

- 그렇다면 위의 코드는 실행 결과가 어떻게 될까? => undefined behavior
  - 첫 번째 인자가 먼저 평가될 경우
  - 두 번째 인자가 먼저 평가될 경우
  - 동시에 평가될 수도 있다.

- 기본적으로 한 줄에서 동인한 변수를 여러 번 바꾸면 위험하다.
  - 함수 매개 변수의 평가 순서는 컴파일러마다 다를 수 있다.
  - 한 함수의 매개 변수들이 동일한 변수를 수정할 경우, 결과가 정의되지 않음(undefined behavior)
  
  ```c
  /* 위험한 것들 */
  add(++i, ++i);
  add(i = -1, i = -1);
  add(i, i++);
  ```

- 정의되지 않은 것은 하지말자. 위험하다.
- `+` 연산자는 피연산자의 평가 순서를 강제하지 않는다. `=` 연산자도 마찬가지이다.

## 연산자 우선순위와 평가 순서

```c
int main(void)
{
    int num1 = 10;
    int num2 = 20;

    int result = add(num1, num2) + subtract(num1, num2) * divide(num1, num2);
    printf("result: %d\n", result);

    return 0;
}
```

- 주의하자. 함수의 평가 순서와 연산자 우선 순위는 아무런 관련이 없다. add, subtract, divide 무엇이 먼저 평가될 지는 알 수 없다.

## 평가 순서를 강제하는 연산자

- sequence point, <https://en.wikipedia.org/wiki/Sequence_point> 에 대해 읽어보도록 하자.
- short circuit 평가도 평가 순서를 강제하는 연산자 중 하나이다.
  - 논리 연산자 &&와 ||는 평가 순서를 강제하는 연산자이다.
  - 왼쪽 피연산자의 평가만으로 오른쪽 피연산자를 평가 안 할 수 있다.
- 정리하자면, 한 줄에 있는 여러 피연산자들은 기본적으로 평가 순서가 보장 안 된다고 생각하자.
  - short circuit, 삼항 연산자 정도만 예외적으로 알아두자.

## 범위(scope)

- 블록 범위
  - 중괄호({}) 안에 선언한 것들은 그 블록 안에서만 사용 가능하다.
  - 블록 안에 또 다른 블록을 넣을 수도 있다. 그러면 안쪽 블록은 바깥 블록에는 접근이 가능하지만 그 반대는 안 된다.
    - c에서는 변수 선언 위치의 제약 때문에 블록이 조금 특별한 의미를 갖는다. 아래의 예를 보자.

    ```c
    int main(void)
    {
        int num1 = 10;
        printf("num: %d\n", num1);

        {
            /* c에서는 변수 선언 위치에 제약이 있다. 이렇게 블록 안에서 선언하면 괜찮다. */
            int num2 = 100;
            int result = num1 + num2;
            printf("result: %d\n", result);
        }

        /* num2, result 접근 못 함 */
    }

    ```

  - 변수 이름 가리기(variable shadowing) 금지. 가독성 뿐만 아니라 버그를 만들 가능성이 높다. 변수 이름 가리기는 블록 안에서 블록 밖에 있는 변수와 같은 이름의 변수를 선언하고 사용하는 것을 말한다.

  ```c
  /* 안 좋은 코드 */
  int main(void)
  {
      int num = 0;
      printf("%d", num);
      {
          int num = 1; /* variable shadowing */
          printf("%d", num);
      }

      return 0;
  }

  ```

- 파일 범위
  - 어떤 블록이나 매개 변수 목록에도 안 속하고 파일 안에 있는 것을 말한다.
  - 엄밀하게 말하면 translation unit이다. 이것에 대해서는 나중에 더 배울 것이다.

  ```c
  # include <stdio.h>

  /*  파일 범위 */
  static int s_num = 1024;

  int add(int op1, int op2);

  int main(void)
  {
      s_num = add(10, 30);

      return 0;
  }
  ```

  - 파일 범위에 있는 변수는,
    - 다른 소스코드 파일에서 링크 가능
    - 프로그램 실행 동안 공간을 차지한다.
      - 즉, 스택 메모리에 들어가는 것이 아니다.
      - 데이터 섹션에 들어간다.
  - 이게 바로 전역 변수이다.
- 함수 범위
  - 유일한 예: 레이블(label)
  - goto 같은 데서 쓰는 것이다.
  - 함수 안에서 선언된 레이블은 함수 어디에서라도 접근 가능하다.
  
  ```c
  int main(int argc, char** argv)
  {
      if (argc != 3) {
        goto exit;
      }

      printf("You have 3 arguments!");
  
  exit:
      return 0;
  }
  ```

- 함수 선언 범위
  - 함수 선언의 매개 변수 목록에 있는 것은 그 목록 안에서 접근이 가능하다.
  - 많이 쓰일 일은 없지만 알아두자.

  ```c
  /* 괜찮은 예제들 */
  void so_something(
      double value,   /* 함수 선언 범위 */
      char array[10 * sizeof(value)] /* value는 첫 번째 매개변수 */
  )
  ```

## const 키워드

## goto 문

## goto 문은 정말로 악마인가요?

## 배열

## 스택 메모리

## 스택 메모리에 대해서 간단히 알아보자

## 스택 메모리 안의 배열, 스택 오버플로

## 배열의 요소 개수 구하는 방법

## 길이가 명시된 매개변수 배열

## 매개변수 배열의 길이, 배열 요소의 초기값

## 다차원 배열

## 소스코드에서 실행파일까지, C 프로그램의 빌드 과정

## .h와 .c 파일

## 헤더 파일이 필요한 이유

## #include <>와 #include ""

## 빌드 과정: 전처리 단계

## 트랜슬레이션 유닛 보는 방법

## 빌드 과정: 컴파일 단계

## 어셈블리어 코드 보는 방법

## 빌드 과정: 어셈블 단계

## 오브젝트 코드 보는 방법

## 빌드 과정: 링크 단계

## 링크 단계가 분리되어 있는 이유

## 라이브러리(library), 정적/동적 라이브러리와 링크

## 분할 컴파일과 전역 변수

## 다른 파일에 있는 전역 변수 사용 시 문제점

## extern 키워드

## 전역 변수의 문제, static 키워드

## .c와 .h 파일 정리, 순환 헤더 인클루드와 해결법

## 인클루드 가드 작동법

## 인클루드 가드 예제

## C 컴파일러의 종류와 특징
