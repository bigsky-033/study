# C언어 기본 문법 1

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## C89 표준

- 1989년에 ANSI에 의해 채택된 첫 번째 C 표준이다. 아직도 대부분의 컴파일러가 지원하는 표준이다.

## Hello World

```c
#include <stdio.h>

int main(void)
{
    printf("Hello World!\n");
    return 0;
}

```

## #include, stdio.h

### #include

- 다른 파일에 구현된 함수나 변수를 사용할 수 있게 해준다.
- #include는 전처리기 지시문 중 하나이다. 전처리기(preprocess)란 컴파일 하기 전에 텍스트를 '복붙'해주는 일을 한다.

### <stdio.h>

- <> 안의 stdio.h 는 실제 디스크 상에 존재하는 파일 이름이다. `#include <stdio.h>` 는 stdio.h 라는 파일을 찾아 그 내용을 복사해 넣어달라는 말이다.
  - include 시점에 `#include <stdio.h>` 가 그 내용으로 대체되는 식이다.
- `#include 'stdio.h'`: 컴파일 오류
- `#include "stdio.h"`: 컴파일은 되지만, 굳이 사용하지 말 것

### C 표준 라이브러리란?

- 다음에 필요한 매크로, 자료형(data type), 함수 등을 모아 놓은 것 
  - 문자열 처리
  - 수학 계산
  - 입출력 처리
  - 메모리 관리
- <stdio.h> 는 libc에서 표준 입출력(Standard Input and Output)을 담당

## main(void) 함수

- 프로그램의 진입점(entry point)
- C 코드를 빌드해서 나온 실행 파일을 실행하면 main(void) 함수가 실행되는 것이다.
- 반드시 int를 반환해야 한다.
  - 0: 프로그램에 문제가 없었다는 뜻이다.
  - $? 를 통해 main 함수의 반환 값을 확인할 수 있다.
- main(void) 함수: (void)
  - 다른 언어와 달리 void를 생략한다고 매개 변수가 없다는 뜻이 아니다.
    - 함수 선언에서 void를 생략하면 매개 변수를 받는다는 의미이다. 단 아직 갯수와 자료형을 모를 뿐이다.
    - 함수 정의에서 void를 생략하면 그 때는 매개 변수가 없다는 뜻이다.
  - 따라서 아무 것도 안 받으면 void를 넣는 습관을 기르자

### 사용할 컴파일 커맨드

```bash
$ clang -std=c89 -W -Wall -pedantic-errors hello.c
```

### main() 함수와 커맨드 라인 인자

나중에 포인터를 배우고 난 뒤 더 자세히 설명할 내용들이다.

```c
// Ex 1)
int main(int argc, char* argv[])
{
    /* codes */
    return 0;
}

// Ex 2)
int main(int argc, char** argv)
{
    /* codes */
    return 0;
}

```

## printf() 함수

- 화면에 데이터를 출력할 때 사용하는 함수
- printf 는 print formatted 라는 뜻이다.
  - %s: 문자열 출력을 위해 쓰인다.
  - %d: 정수 출력을 위해 쓰인다.

## 주석(comment)

- /**/ 만 지원 (최소한 C89에선)
- 주석이 한 줄이든 여러 줄이든 다 /**/ 만 지원한다.

## C언어의 기본 문법

- C는 절차적 언어이다. 즉, C로 작성한 코드는 데이터보다 프로세스에 중점이 맞춰져 있다는 의미이다.
- 클래스 없다.
- 함수는 기본적으로 모든 것이 전역(global)이다.
- 변수
  - 함수 밖에 선언되어 있으면, 전역 변수
  - 함수 안에 선언되어 있으면, 지역(local) 변수

## 자료형, unsigned와 signed

- 기본 자료형
  - char
  - short
  - int
  - long
  - float
  - double
  - long double
- unsigned 를 사용하길 원할 경우, unsigned라는 단어를 자료형 이름 앞에 넣어주어야 한다.
  - ex) unsigned char, unsigned int
- 부호 있음을 명확히 보여주기 위해 signed를 붙일 수도 있다.
  - ex) signed char, signed int
- unsigned/signed를 생략하면 signed이 기본이다.
  - 예외: char

## char 형

- 최소 8비트인 정수형(표준 에서의 정의)
- <limits.h> 헤더를 include 한 뒤, CHAR_BIT를 보면 char의 비트를 알 수 있다.

```c
#include <stdio.h>
#include <limits.h>

int main(void)
{
    char char_size = CHAR_BIT;
    printf("%d\n", char_size);
    return 0;
}

```

- C 표준에서는 기본 자료형의 정확한 바이트 수를 강요하지 않는다. 컴파일러에게 맡긴다.
- 1바이트를 CHAR_BIT 만큼이라고 말한다.
- 정수형이다. 덧셈도 가능하다.
- char는 ASCII 문자를 표현하기에 충분하다. ASCII는 0~127인 숫자이기 때문이다.
- C 표준은 char의 기본을 signed/unsigned 인지 정의하지 않는다. 이것도 컴파일러에게 맡긴다.
  - 따라서 8비트 정수형으로 쓰려고 할 때는 반드시 char 앞에 signed나 unsigned를 넣어주는 게 좋다.
- <limits.h> 헤더 파일에서 CHAR_MIN을 보면 부호 식별자가 없는 char가 signed인지 unsigned인지 알 수 있다.
  - CHAR_MIN이 정의되어 있다. 이게 -128이면 signed이고 0이면 unsigned이다.
- (표준) char로 표현 가능한 숫자의 범위 (포팅에 문제 없는 범위를 말한다)
  - unsigned char: 0~255
  - char: 0~127
  - signed char: -127~127
    - 왜 -128이 아니라 -127일까? 1의 보수를 쓰는 기계가 있을 수도 있다. 때문에 '안전한' 포팅을 위해 표준에선 -128이 아니라 -127으로 정의했다.
- (보통) 다음과 같은 것들은 안전하게 생각해도 된다.
  - 크기: 8bit
  - 부호를 생략할 경우: signed
  - 범위
    - 부호가 없는 경우(unsigned): 0~255
    - 부호가 있는 경우(signed): -128~127

## short 형

- 최소 16비트이고 char의 크기 이상인 정수형
- 포팅 문제 없는 값의 범위
  - 부호 없는 short(unsigned short): 0~65535
  - 부호 있는 short(short, signed short): -32767~32767
- 기본 정수형(보통 int) 보다 짧다. 메모리를 적게 쓰기 위해 사용한다.
- 그러나 int 대신 short를 사용할 경우 성능이 느려질 수도 있다. CPU 계산에서 int가 기본 크기인 경우가 많아 short를 이용해 계산을 하려고 하면 별도의 연산이 필요할 수 있다.
- (보통) 다음과 같은 것들은 안전하게 생각해도 된다.
  - 크기: 16bit
  - 범위
    - 부호가 없는 경우(unsigned): 0~65535
    - 부호가 있는 경우(signed): -32768~32767

## int 형, long 형

## float 형, double 형, long double 형

## C언어의 bool 형

## 열거형(enum)

## 변수 선언

## sizeof 연산자

## size_t

## 역 참조 연산자

## 구조체/공용체 멤버 접근자

## 비교 연산자, 논리 연산자, 조건 연산자

## 조건문과 반복문, if 문, switch/case 문

## for 문, while 문, do while 문
