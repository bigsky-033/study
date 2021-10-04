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

