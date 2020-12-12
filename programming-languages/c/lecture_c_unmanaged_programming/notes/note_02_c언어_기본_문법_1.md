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

### int 형

- 표준에 따르면 최소 16bit 그리고 short 크기 이상인 정수형
  - CPU가 아는 크기여야 한다. CPU에게 정수를 처리해 라고 하면 처리할 크기를 말한다. 즉, CPU마다 다를 수도 있다는 의미이다. 그런데 예전에는 16bit CPU가 흔했다. 그래서 최소 16bit이다.
- int는 그냥 정수(integer)라는 의미.
- 32bit 컴퓨터가 나오면서 int의 일반적인 크기는 32bit가 되었다.
- 현재는 64bit 컴퓨터가 나왔다. 그렇지만 여전히 int는 32bit로 표현된다.
  - 너무 오랫동안 32bit로 사용되어 왔다.
  - 캐시 메모리 등의 이유로 인해 64비트로 바꾼다고 해서 딱히 성능에서의 이점이 없다.
  - int가 32bit가 되면 32bit 정수형을 무엇으로 표현할지가 애매해진다.
- int로 표현 가능한 숫자의 범위
  - 포팅에 안전한 범위: short와 같다.
  - (보통) 다음과 같은 것들은 안전하게 생각해도 된다.
    - 크기: 32bit
    - 범위
      - 부호가 없는 경우(unsigned): 0~4,294,967,295
      - 부호가 있는 경우(signed): -2,147,483,648~2,147,483,647
- int의 리터럴

```c
int signed_int = -1024;
unsigned int unsigned_int1 = 394;
unsigned int unsigned_int2 = 2147483648; /* 경고 */
unsigned int unsigned_int3 = 2147483648u; /* 경고 없음, 대문자 U도 됨 */

```

- 리터럴(literal)
  - 'u' 또는 'U': 부호 없는(unsigned) 수를 표현하는 접미사이다. 부호 있는 수의 최댓값보다 큰 값은 unsigned int에 대입할 경우 'u' 혹은 'U'를 붙여야 한다. 그렇지 않으면 경고가 발생한다.

### long 형

- 표준에 따르면 최소 32bit 이고 int 이상의 크기이다.
  - long은 일반적인 컴파일러 기준, C에서 32bit이다.
    - 다른 언어에서는 long이 보통 64bit이다.
- 따라서 보통 long의 크기와 안전한 범위는 int와 같다.
- long의 리터럴

```c
long signed_long = -200000000l;  /* 대문자 L도 된다 */
unsigned long unsigned_long1 = 2147483647;
unsigned long unsigned_long2 = 2147483648;  /* 경고 */
unsigned long unsigned_long3 = 2147483648ul;  /* 경고 없음 */

```

- 리터럴(literal)
  - 'l' 또는 'L': long을 의미하는 접미사
  - 'u' 또는 'U': 부호 없는(unsigned) 수를 표현하는 접미사이다.
  - 두 접미사를 같이 쓸 수도 있다. 이럴 경우 unsigned long이라는 의미가 된다.

## float 형, double 형, long double 형

- 부동 소수점 자료형은 IEEE 754로 대동단결 되었다.
  - float은 IEEE 754 Single(32bit)
  - double은 IEEE 754 Double(64bit)
- 하지만 C는 CPU가 IEEE 754를 지원하는 실수 계산 장치를 장착하기 전부터 쓰였다.

### float 형

- 표준에 따르면 C의 float은 IEEE 754가 아닐수도 있다. 크기는 char 이상이기만 하면 된다.
- unsigned 형이 없다.
- (보통) 다음과 같은 것들은 안전하게 생각해도 된다.
  - 크기: 32bit
  - 범위: IEEE 754 Single과 동일
- 관련 헤더 파일: float.h
- float의 리터럴

```c
float pi1 = 3.14f; /* F도 된다. 그런데 대부분 소문자 f를 쓴다. */
float pi2 = 3.14uf; /* 컴파일 오류, float은 접미사 u를 안 쓴다. */ 

```

### double 형

- 표준에 따르면 CPU가 계산에 사용하는 기본 데이터 크기이다.
  - 크기는 float 이상이기만 하면 된다.
- unsigned 형이 없다.
- double이 기본이기 때문에 별다른 리터럴이 없다.
- (보통) 다음과 같은 것들은 안전하게 생각해도 된다.
  - 크기: 64bit
  - 범위: IEEE 754 Double과 동일
- 관련 헤더 파일: float.h

### long double 형

- double 보다 정밀도가 높다.
- double 이상의 크기면 된다.
- unsigned 형이 없다.
- 관련 헤더 파일: float.h

## C언어의 bool 형

- C89에 없다.
- C99에서 좀 이상한 형태로 새로 들어왔다.
- 대부분의 C 프로그래머들은 bool을 사용하지 않는다.
- bool 형을 잘 안 쓰는 이유
  - 정수로 대신 쓸 수 있다. 0이면 false, 0이 아니면 true
  - 하드웨어에서도 실제 bool이 없다.
- 따라서 while 문의 조건으로 숫자를 사용 가능하다.

```c
/* 조건식 예 */

int counter = 5;
while (counter--) { /* 나쁜 조건식 */
    printf("%d\n", counter);
}

```

```c
/* C에서 참이나 거짓을 반환해야 하는 함수의 경우는 보통 이렇게 한다.  */

int is_student(const int id)
{
    if (/* 조건 */) {
      return 1; /* true */
    }

    return 0; /* false */
}

```

## 열거형(enum)

- 명시적 캐스팅 없이 int와 섞어서 사용이 가능하다.
  - 따라서 `int -> enum`, `enum -> int`, `enum -> 다른 enum` 과 같은 대입이 다 된다.
- C에서의 열거형은 그냥 정수에 별명을 붙이는 수준이다. 다른 고수준 언어에서 실수를 방지하기 위해 enum을 쓰는 것과는 좀 다른 느낌이다.

```c
enum day { DAY_MONDAY, DAY_TUESDAY, DAY_WEDNESDAY, /* 생략 */ };
enum month { MONTH_JANUARY, MONTH_FEBRUARY, MONTH_MARCH, /* 생략 */ };

enum day hump_day = DAY_WEDNESDAY;
enum month birth_month = hump_day; /* 문제 없이 컴파일 된다. */

```

## 변수 선언

- 변수 선언은 반드시 블럭의 시작에서만 해야 한다.
- 코드 중간에 사용하는 변수는 블록 시작에서 선언만 하고 뒤에 대입한다.

```c
int main(void)
{
    int num1 = 10;
    int num2 = 1234;

    printf("num1 = %d, num2 = %d", num1, num2);  

    result = add(num1, num2);

    printf("%d + %d = %d", num1, num2, result);

    return 0;
}

```

## sizeof 연산자

- 피연산자의 크기를 바이트로 반환해주는 연산자이다. 함수가 아님에 유의하자.
- 함수가 아니라는 의미는, 컴파일 중에 평가된다는 말이다. 바꿔 말하면, 컴파일 할 때 모르는 크기는 찾아줄 수 없다는 말이다.
- char 형을 넣으면 반드시 1이 반환된다. char는 바이트와 같기 때문이다.
- 이 연산자가 반환하는 값은 부호 없는 정수형의 상수로 size_t 형이다.

```c
char ch = 'a';
int num = 100;
char char_array[30];

size_t size_char = sizeof(ch);           /* 1 */
size_t size_int = sizeof(num);           /* 4 */
size_t size_float = sizeof(float);       /* 4 */
size_t size_arary = sizeof(char_array);  /* 30 */

```

## size_t

- 부호 없는 정수형이나 실제 데이터형은 아니다.
- _t 는 typedef를 했다는 힌트이다.
  - typedef는 다른 자료형에 별칭을 붙이는 것을 말한다.
  - 플랫폼에 따라 다른 자료형을 쓰기 위해서 size_t를 typedef 한 것이다.
- size_t의 크기
  - C89 표준은 size_t의 크기를 딱히 명시하지 않았다.
  - 단, 배열을 만들면 그 배열의 바이트 크기를 얻을 수 있다고는 명시했다.
  - 다행히도 C99 표준에서 확실히 최소 16 비트를 요구한다.
  - 보통은 unsigned int를 사용한다.
    - `typedef unsigned int   size_t;`
- size_t의 용도
  - 어떤 것의 크기를 나타내기 위해 사용한다.
  - 반복문이나 배열에 접근할 때 사용한다.

    ```c
    int int_array[30];
    size_t i;

    for (i = 0; i < 30; ++i) {
        int_array[i] = i;
    }
    ```

- size_t와 -1
  - java에서 String.indexOf() 함수는 문자를 못 찾으면 -1을 반환했다.
  - size_t를 가지고도 마찬가지 일을 할 수 있다.
    - signed int -1과 unsigned int 4,294,967,295는 같은 비트 패턴을 갖는다.

    ```c
    size_t get_students_index(const char* name)
    {
        if (! /* 조건 */) {
            return (size_t)-1;
        }
        return index;
    }
    ```

## 역 참조 연산자, 주소 연산자

```c
int num = 10;
int* p = &num;    /* 여기서 * 는 역 참조 연산자(x). & 는 주소 연산자(o) */
int num1 = *p;    /* 역 참조 연산자(o) */

int result = num1 * num2;   /* 역 참조 연산자(x), 곱셈 연산자(o) */
```

- *의 피연산자가 하나면 역 참조(indirection, dereference) 연산자이다.
- &와 피연산자가 하나일 때 주소 연산자이다.

## 구조체/공용체 멤버 접근자

- . 연산자
  - C에는 클래스가 없으므로 함수 호출에 쓰일 일은 없다.
  - 구조체와 공용체가 있다. 그들의 멤버 변수에 접근할 때 사용한다.
- -> 연산자
  - 2개의 연산자 '.'와 '*'를 합친 것이다.
    - `(*s).data` => `s->data`

## 비교 연산자, 논리 연산자, 조건 연산자

| 비교 연산자     | 논리 연산자 |
| --------------- | :---------: |
| == != < <= > >= |   && !! !   |

- 조건 연산자(삼항 연산자)

```c
int num1 = 10;
int num2 = 56;
int min = num1 < num2 ? num1 : num2;

```

## 조건문과 반복문, if 문, switch/case 문

### if 문

```c
int num1 = 10;
int num2 = 56;
int min;

if (num1 < num2) {
  min = num1;
}
else {
  min = num2;
}

```

- C에서 if 문의 조건식에는 숫자가 오길 기대한다.
  - 모든 비트패턴이 0이 아니면 false 이다.
  - 아니면 true이다.

### switch/case 문

```c

/* 기본적인 생김새 */

int num = 1;

switch (num) {
case 0:
    printf("Hello POCU!\n");
    break;
case 1:
    printf("Hello World!\n");
    break;
default:
    printf("invalid entry\n");
    break;
}

/* C에서 안되는 것 */
const char* name = "Banana";
switch (name) {  /* 컴파일 오류 */
case "Apple":  /* 컴파일 오류 */
    printf("Breakfast\n");
    break;
case "Mango":  /* 컴파일 오류 */
    printf("Lunch\n");
    break;
deafult:
    printf("Unknown\n");
    break;
}

```

- C는 정수형(예: int, char, enum) 만 switch의 조건식에 넣을 수 있다.
- case 안에서 break를 빼먹으면?
  - switch 문을 곧바로 탈출하지 않고 그 아래 있는 코드를 계속해서 실행한다.
  - 다른 곳에서 break를 만나거나 switch 블록의 끝에 도착하면 탈출한다.
  - 이런 것을 fall-through 라고 한다.
    - 따라서 fall-through 를 의도한 경우일 때는 주석으로 `/* intentional fallthrough */` 와 같이 표현해 주는 것이 좋다.
- case 레이블은 반드시 상수만 된다. 배열에 있는 값 같은 것을 넣어줄 수는 없다.

## for 문, while 문, do while 문

### for 문

- 다른 프로그래밍 언어와 다르게 for 문의 초기화 코드에서 인덱스를 위한 변수를 만들고 할당할 수 없다. C에서는 변수 선언은 제일 위에서 해야 하기 때문이다.

```c
int sum = 0;
size_t i;

for (i = 0; i < 10; ++i) {
  sum += i;
}

```

### while 문

- 별로 특별할 게 없다.

```c
int day = 5;
while (day-- > 0) { /* while(day--) 처럼 쓰는 것은 권장하지 않음 */
  printf("%d\n", day);
}

```

### do while 문

- 별로 특별할 게 없다.

```c
int day = 5;
do {
    printf("%d\n", day);
} while (day-- != 0);

```
