# 노트 - C Programming A Modern Approach

이 문서는 K.N.KING의 C Programming A Modern Approach를 읽으며 적어두고 싶었던 내용들을 적어놓은 문서이다.

## 1. Introducing C

### *1. Introducing C* 에서 다루는 것들 

  - C언어의 역사
  - C언어의 장단점

### 노트 - *1. Introducing C*

- C를 배우는 게 아직도 유효할까에 대한 작가의 답변:
  - C를 배움으로 인해 C에서 파생된 언어(C++, Java, C#, Perl,...)의 기능들을 더 잘 이해할 수 있게 된다.
  - C로 이미 짜여진 많은 코드들이 있다. 이런 코드를 읽거나 유지보수 해야할 일이 있을 때 도움이 된다.
  - C는 여전히 새로운 소프트웨어를 개발할 때 많이 사용되는 프로그래밍 언어이다. 메모리나 프로세싱 파워가 제한된 상황에서는 여전히 C가 유용하다.

- C언어의 특징
  - C는 로우 레벨 언어이다.
  - C는 가벼운 언어(small language)이다.
  - C는 관대한 언어(permissive)이다.
- C언어의 강점
  - 효율성(Efficiency)
  - 이식성(Portability)
  - 강력함(Power)
  - 유연함(Flexibility)
  - 표준 라이브러리(Standard library)
  - Unix와의 통합(Integration with UNIX)
- C언어의 약점
  - C언어의 약점은 사실 강점이 되는 부분에서 온다. C는 머신과 가깝다는 점이다.
  - 에러가 발생하기 쉽다(error-prone).
  - 이해하기 어려울 수 있다.
  - 수정하기 어려울 수 있다.

- C를 효율적으로 사용하기
  - C의 함정들(pitfalls)을 피하는 법을 배워라.
  - 프로그램을 더 안정적이게(reliable) 하기 위해 소프트웨어 도구들을 사용하라.
  - 이미 존재하는 코드 또는 라이브러리를 사용하라.
  - 괜찮은 코딩 컨벤션을 도입하라.
  - 트릭(tricks)과 과하게 복잡한 코드를 피하라.
  - 표준을 지켜라.

## 2. C Fundamentals

### *2. C Fundamentals* 에서 다루는 것들 

- C 프로그램을 컴파일하고 링크하는 방법.
- 어떻게 프로그램을 일반화 하는지.
- 주석을 다는 방법.
- 변수에 대하여.
- scanf와 같은 함수를 이용해 값을 읽어 변수에 저장하는 방법.

### 노트 - *2. C Fundamentals* 

- C 프로그램을 실행 가능한 상태로 만들기 위해서는 다음의 과정이 필요하다.
  - Preprocessing
  - Compiling
  - Linking

- 프로그램을 일반화 한다면 다음과 같은 것들로 구성된 모양일 수 있다:
  - directives
    - 컴파일 이전에 프로그램의 내용을 변경하기 위해 실행 될 명령들. 
    - 항상 #으로 시작.
    - 기본적으로 한 줄.
    - directive의 끝을 표현하기 위해 세미콜론이나 특별한 마킹이 필요하지 않음.
  - functions
    - 실행 가능한 코드의 블록.
    - 다른 프로그래밍 언어에서 "proceduers" 또는 "subroutines" 와 비슷. 
    - C 프로그램은 많은 함수로 구성될 수 있지만 유일하게 필요한 함수는 main임.
    - main 함수가 반환하는 int는 프로그램이 종료될 때 운영체제에게 전달됨.
  - statements
    - 프로그램이 실행되면 실행될 명령들.

```c

directives

int main(void)
{
  statements
}

```

- 변수가 기본 값을 갖고 있지 않고 프로그램에 의해 값이 할당되지 않은 경우를 초기화되지 않음(uninitialized)라고 한다.
  - 초기화되지 않은 변수에 접근하려고 하면 예상할 수 없는 결과가 나오거나 프로그램이 깨질(crash) 수 있다.
- printf는 값이 필요한 자리에 식을 넣을 수도 있다. 예를 들어 `printf("%d\n", volume);` 대신 `printf("%d\n", height * length * width);`로 쓸 수 있다. 이는 C의 일반적인 원칙을 보여주기도 한다. 이 일반적인 원칙은 값이 필요한 자리에 같은 타입의 식을 쓸 수 있다(Wherever a value is needed, any expression of the same type will do)이다.


## 3. Formatted Input/Output

### *3. Formatted Input/Output* 에서 다루는 것들

- *scanf*와 *printf*에 대해 다룬다.

### 노트 - *3. Formatted Input/Output*

- Conversion specification: `%`뒤에 오는 문자를 통해 출력할 포맷을 지정해줄 수 있는 것과 관련된 스펙, 정의. `%d`, `%f` 등이 여기에 해당한다.
- Conversion specification을 일반화하면 `%m.pX`또는 `%-m.pX`을 가질 수 있다. 여기서 m과 p는 정수 상수이고 X는 문자이다. m과 p는 옵셔널하다. 만약 p가 생략되면 m과 p를 구분하는 구분 점도 드랍된다. 예를 들어 %10.2f에서 m은 10 p는 2 그리고 X는 f이다. %10f에서 m은 10이고 p는 없는 것이다. %.2f에서 m은 없고 p는 2인 것이다.
  - *m* 은 minimum field width를 나타내는데 출력해야 할 문자열의 최소 길이를 의미한다. 만약 출력할 문자열의 길이보다 m이 크다면 출력할 문자열 앞에 공백이 추가된다. 만약 출력할 문자열의 길이보다 m이 작다면 자동으로 확장된다. %4d를 두고 12345를 출력한다고 하면 생략되는 것이 아니라 그대로 12345가 출력된다. %-4d 처럼 마이너스 부호를 두는 것은 공백이 있을 때 왼쪽 정렬이 되게 하는 것을 의미한다. 예를 들어 (%4d, 123)을 하면 (공백)123이 출력 되는데 (%-4d, 123)을 하면 123(공백)이 출력된다.
  - *p*는 정밀도(precision)을 의미한다. 정의가 좀 모호할 수 있다. 왜냐하면 X에 따라 의미가 달라질 수 있기 때문이다.
  - *X*로 주로 사용되는 것들은 다음과 같다.
    - *d*: integer를 10진수로 표현한다. 여기서 p는 화면에 표시해야 할 최소 숫자의 갯수를 나타낸다. 만약 p가 생략되면 1을 기본 값으로 한다. 따라서 %d는 %.1d와 같다.
    - *e*: floating-point number를 exponential format으로 표현한다. p는 decimal point 이후에 표현해야 할 숫자의 갯수를 의미한다. 기본 값은 6이다. 만약 이 값이 0이면 decimal point는 표현되지 않는다.
    - *f*: floating-point number를 fixed decimal format으로 표현한다. exponent가 없다. p의 의미는 e에서와 같다. 
    - *g*: floating-point number를 숫자의 크기에 따라 exponential format 또는 fixed decimal format으로 표현한다. p는 화면에 표시할 maximum number of significant digits을 의미한다. f와 다르게 g는 trailing zeros를 표시하지 않는다. 또한 출력해야 할 값이 decimal point 뒤에 숫자가 없으면 decimal point를 출력하지 않는다.

