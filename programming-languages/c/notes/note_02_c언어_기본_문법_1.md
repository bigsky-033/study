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

## 주석(comment)

## C언어의 기본 문법

## 자료형, unsigned와 signed

## char 형

## short 형

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
