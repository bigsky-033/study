# 콘솔 입력, 파일 입출력, 커맨드 라인 인자

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 입력

- 출력의 반대이다.
- 외부의 데이터를 읽어와서 프로그램에서 사용한다.
- 어떤 데이터가 들어올지 몰라서 이상한 데이터가 종종 들어온다.
  - 사용자가 잘못된 데이터를 키보드에서 입력
  - 예전에 저장해 놓은 파일을 누가 잘못 바꿨거나 일부 데이터가 유실
- 출력에 비해 조심해야 할 일이 많다.
- 데이터 읽기에 실패했는데 제대로 처리하지 않으면 문제가 펑펑 터진다.
- 모든 입력 함수에는 반환 값이 있다. 따라서 어떤 함수가 어떤 값을 반환하는지, 그 반환값의 의미는 무엇인지 문서에서 확실히 읽고 코드에서 검사할 것.
- 입력의 출처는 어디일까? => 어딘가에서 출력을 했다면 거기서 읽어올 수 있다고 생각하자.
  - 스트림
    - 콘솔(키보드), 파일...
  - 문자열
- 입력처리 전략
  - 한 글자씩 읽기
  - 한 줄씩 읽기
  - 한 데이터씩 읽기
  - 한 블록씩 읽기 (이진 데이터)

## 한 글자씩 읽기, 한 글자씩 읽는 알고리즘 1

- 한 글자씩 읽는 알고리즘
  - 한 글자(char)를 읽어 온다.
  - 그 글자를 필요한 곳에 사용한다.
  - 다시 한 글자를 읽어오는 단계로 돌아간다.
  
  ```c
  /* 매우 단순화한 예 */
  #include <stdio.h>

  int c;

  /* #define TRUE (1) */
  while (TRUE) 
  {
      c = getchar();
      putchar(c);
  }

  ```

- getchar()
  - `int getchar(void);`
    - 키보드(stdin)으로부터 문자 하나를 읽어서 int형으로 반환.
    - 반환 값
      - 성공 시, 읽은 문자의 아스키코드를 반환한다.
      - 실패 시, EOF를 반환한다.
    - C에 있는 많은 입출력 함수들이 문자를 읽고 쓸 때 char 대신 int를 쓴다.
  - `int fgetc(FILE* stream);`
    - fgetc(stdin) 는 getchar() 와 같다.

## getchar()와 EOF 키

- 아래의 프로그램은 프로그램이 끝나지 않는 문제가 있다.

  ```c
  #include <stdio.h>

  int c;

  /* #define TRUE (1) */
  while (TRUE) 
  {
      c = getchar();
      putchar(c);
  }

  ```

- 언제 읽는 것을 멈춰야 할까?
  - 방법 1) 특별한 키를 입력 받았을 경우
    - 예) x를 입력 받으면 프로그램을 종료한다.
    - 그런데, 이런 방법은 한계가 있다. 예를 들어 타자 연습 프로그램은 특정 문자를 기준으로 종료하기가 어렵다.
  - 방법 2) 어디에서도 사용 가능한 "입력 끝"을 나타내는 무언가를 이용한다.
    - 레퍼런스, <https://en.cppreference.com/w/c/io/getchar> 에서 힌트를 얻어보자.
      - 레퍼런스에서 반환 값에 대한 설명이 있는데, `The obtained character on success or EOF on failure.`라고 한다.
      - 즉, `EOF`가 반환이 되면 종료하면 된다.
    - 입력의 끝을 나타내는 값을 'EOF'로 많이 쓴다.
      - C 표준에 의하면 EOF는 음수이다.
      - getchar()가 char를 반환했다면 이 EOF를 표현할 방법이 없다. 그래서 int를 반환한다.

## 한 글자씩 읽기 알고리즘 2, EOF 키는 어디 있나요?

- 위에서 구현했던 한 글자씩 읽는 알고리즘을 바꿔본다.
- 한 글자씩 읽는 알고리즘
  - 한 글자(char)를 읽어 온다.
  - 글자를 읽어오는 데 실패했다면(반환 값이 EOF라면) 프로그램을 종료한다.
  - 그 글자를 필요한 곳에 사용한다.
  - 다시 한 글자를 읽어오는 단계로 돌아간다.

  ```c
  /* 매우 단순화한 예 */
  #include <stdio.h>

  int c;

  c = getchar();

  while (c != EOF) 
  {
      c = getchar();
      putchar(c);
  }

  ```

- 그런데...EOF는 어떻게 입력하는 걸까?
  - 키보드에도 없고 영어로 EOF로 입력해도 안 된다.
  - 특정한 키를 조합해서 넣어야 한다.
    - 윈도우: `ctrl + z`
    - 유닉스 등의 시스템: `ctrl + d`
- 그런데...위에 보면 getchar()가 두 번 있다.
  - do while로 해결이 가능할까?
    - 안된다. while 검사 전에 c를 사용하기 때문이다.
  - 한 번에 어떻게 안 될까? 가능하다.

    ```c
    int c;

    while ((c = getchar()) != EOF)
    {
        putchar(c);
    }

    ```

  - 위의 경우에서 괄호`(c = getchar())`를 쓰지 않으면 안 된다. != 연산자가 = 연산자보다 우선 순위가 더 높다.
    - 따라서 괄호를 쓰지 않으면, c = (getchar() != EOF) 가 되어버린다. 즉 c는 0 또는 1이 되어버리는 것이다.
    - 이 방법은 실수하기가 쉬워서 잘 쓰지 않는다.

- 한 글자씩 읽는 방법은 언제 유용할까?
  - 가장 간단한 입력 방법이다. 입력이 문자나 문자열일때 좋다.
  - 쓸데 없이 메모리에 입력 값을 저장해두지 않아도 되는 장점이 있다.
  - O(N) 인 알고리즘이다.
    - for 문 딱 한 번만 도는 알고리즘에 적합한 경우가 많다.
  - 그러나...다른 데이터형으로 쓰기는 좀 어렵다. 정수형 숫자를 읽기 같은 것에서는 쓰기 어렵다.

## 한 줄씩 읽기, gets()

- 한 줄씩 읽는 알고리즘
  - 한 줄을 읽어온다.
  - 한 줄을 읽어오는 데 실패하면 프로그램을 종료
  - 성공했다면 한 줄 읽어온 데이터를 필요에 따라 사용한다.
  - 한 줄을 읽어오는 단계로 돌아간다.
- 그러면...한 줄을 어떻게 읽을 수 있을까?
  - 문자들을 읽어와서 어딘가에 잠시 저장해 두어야 한다. 읽어와서 어느 배열인가에 저장한다.
    - 어떻게? 프로그래머가 미리 만든 배열을 함수에 전달한다. 함수는 그 배열에 한 줄을 읽어온다.
- gets() - 위험한 함수
  - `char* gets(char* str);`
  - stdin에서 새 줄 문자('\n') 또는 EOF를 만날 때까지 계속 문자들을 읽어서 str 배열에 저장한다. 그래서 str이 const가 아니다.
  - 마지막 문자 바로 다음에 널 문자('\0')도 넣어준다.
  - stdin에서 새 줄 문자를 제거하지만 버퍼에 저장하지는 않는다.
  - 반환 값
    - 성공 시, str
    - 실패 시, NULL (포인터)

  ```c
  #include <stdio.h>

  /* ... */

  #define LINE_LENGTH (63)

  /* ... */

  char line[LINE_LENGTH];

  while (gets(line) != NULL)
  {
      puts(line);
  }

  ```

  - gets()는 매우 위험한 함수이다. C11에선 아예 제거되었다.
    - 왜 위험할까?
      - 만약 버퍼 이상 입력하면? -> 버퍼 오버플로우가 생긴다. 즉, 올바르지 않은 메모리 주소에 키보드로 입력 받은 값을 써버린 것이다.
        - 잘못 접근하면 버퍼 위치를 넘어 돌아갈 함수 주소 같은 것들을 바꿔버릴 수 있다. 즉, 보안 상 위험한 것이다.
        - 이걸 안전하게 작성할 수 있는 방법이 없다.

## fgets()로 안전하게 한 줄 읽기

- fgets()
  - `char *fgets(char* str, int count, FILE* stream);`
    - str: gets()와 마찬가지로 입력 받은 한 줄을 저장할 char 배열이다. const가 아님에 유의하자.
    - count: 한 번에 str에 쓰는 최대 문자 수이다. 널 문자를 포함하기 때문에 실제로 읽어오는 문자 수는 count-1 개이다.
    - stream: 데이터를 읽어 올 스트림이다. 키보드 입력을 읽어오고 싶다면 stdin을 넣어주면 된다.
    - 반환 값
      - 성공 시, str
      - 실패 시, NULL (포인터)
  - <stdio.h> 안에 있다.
  - 최대 count-1 개의 문자열을 읽어서 str에 저장한다.
  - 즉 새 줄을 만나지 않아도 함수가 반환할 수 있다.
  - str에 새 줄 문자까지 넣어준다.
    - 왜 새 줄 문자('\n')이 들어올까?
      - 새 줄을 만나서 끝났을 때랑 아닐 때를 구분해야 하기 때문이다.
- FILE* 자료형
  - 스트림을 제어하기 위해 필요한 정보를 담고 있는 자료형이다.
    - 파일 위치 표시자
    - 스트림이 사용하는 버퍼의 포인터
    - 읽기/쓰기 중에 발생한 오류를 기록하는 오류 표시자
    - 파일의 끝에 도달했음을 기록하는 EOF 지시자
  - 플랫폼마다 이 자료형을 구현하는 방식은 다를 수 있다.
  - 입력 및 출력 스트림은 오직 FILE 포인터로만 접근 및 조작이 가능하다.

- 한 줄 읽기 코드 예

```c
#include <stdio.h>

/* ... */

#define LINE_LENGTH (10)

/* ... */

char line[LINE_LENGTH];

while (fgets(line, LINE_LENGTH, stdin) != NULL)
{
    printf("%s", line);
}

```

- fgets()에 쓸 버퍼는 초기화가 필요 없다.

## 한 줄씩 읽는 방법이 유용한 경우

- 단어/글자를 하나씩 읽는 것보단 한 줄씩 읽는 게 빠르다.
- CPU를 벗어나 외부로부터 뭔가를 읽어올 때는 한 번에 많이 읽어오는 게 빠르다.
- 따라서 버퍼의 크기는 충분히 큰 게 좋다.

## 한 데이터씩 읽기, scanf()

- 텍스트로 들어오는 문자를 원하는 데이터형으로 지정해 읽는 방법이다.
- 출력의 `printf` 에 대응하는 입력 함수가 `scanf` 이다.
- 3가지 버전이 있다.
  - scanf()
    - `int scanf(const char* format, ...);`
      - 반환 값
        - 몇 개의 `데이터`를 성공적으로 읽었는지 반환
          - 몇 개의 바이트를 읽었는지가 아니다. 3글자 문자열을 %s 하나로 제대로 읽을 경우 1이 반환된다.
        - 첫 데이터를 읽기 전에 실패했다면 EOF를 반환한다.
    - 키보드(stdin)로부터 입력을 받아 변수에 저장한다.
    - <stdio.h> 에 있다.
    - stdin 에서 정수 읽기

      ```c
      #include <stdio.h>

      int main(void)
      {
          int num;

          printf("Enter a number: ");
          scanf("%d", &num);
          printf("num = %d\n", num);

          return 0;
      }

      ```

    - 참조에 의한 전달을 흉내낸다. &num을 넣는다.

  - fscanf(): 파일 스트림으로 부터 읽는다.
    - `int fscanf(FILE* stream, const char* format, ...);`
  - sscanf(): C스타일 문자열로부터 읽는다.
    - `int sscanf(const char* buffer, const char* format, ...);`

## scanf()의 일반적인 서식 문자열 형식

- `%[*][너비][길이]서식_지정자`
  - 일반적으로 '%' 뒤에 최대 4개의 지정자를 가질 수 있다.
    - `*` (선택)
    - 너비 (선택)
    - 길이 수정자 (선택)
    - 서식 지정자 (필수)
  - 반드시 순서를 지켜서 작성해야 한다.

| 서식 지정자 | 용도                                             | 사용 예                                |
| ----------- | ------------------------------------------------ | -------------------------------------- |
| %           | %를 순수하게 문자로 인식. %를 읽고 싶을 때 사용. | `scanf("%%d", &num); /* 컴파일 경고*/` |
| c           | 문자(char)                                       | `scanf("%c", &ch);`                    |
| s           | 한 단어                                          | `scanf("%s", str);`                    |
| d           | 부호 있는 10진수 수                              | `scanf("%d", &num);`                   |
| x           | 부호 없는 16진수 수                              | `scanf("%x", &num);`                   |
| f           | 부동 소수점(float, double)                       | `scanf("%f", &num);`                   |

- 모든 데이터는 한 단어씩 (공백 문자로 구분) 또는 가능 할 떄가지 읽는다.
- 공백 문자는 버린다(예외: %c).
- 대입 생략 문자 - `*`
  - assignment-suppressing character
  - 이 문자를 쓸 경우 받은 입력을 변수에 저장하지 않는다.
  - 필요할 경우 찾아보자. 쓸 일이 많지 않을 것 같다.
- 너비 지정

| 서식 지정자 | 용도              | 사용                                        |
| ----------- | ----------------- | ------------------------------------------- |
| 정수        | 읽을 최대 문자 수 | `scanf("%5s", str); scanf("%10d\n", &num);` |

- %s 의 경우 너비를 지정하지 않으면 버퍼 오버플로가 날 수 있다.
  - 그러나 잘 쓰이는 방법은 아니다.
  - 너비 지정 시 주의할 점
    - 너비를 지정한 후에 여러 데이터를 한 번에 읽을 경우 문제가 발생할 수 있다.

- 길이 수정자(length modifier)
  - 이런 게 있기는 한데...필요할 경우에 찾아보도록 하자.

## scanf() 사용 예

- 예 1) "123p"를 입력했을 때
  - `result = scanf("%d", &num1);`
    - result에는 1이 들어간다.
    - num1에는 123이 들어간다.
    - p는 여전히 스트림에 남아 있게 된다.
- 예 2) "pocu"를 입력했을 때
  - `result = scanf("%d", &num1);`
    - result에는 0이 들어간다. 못 읽은 것이다.
- 예 3) " 12 34"를 입력했을 때 - 12와 34 앞에 공백이 두 칸씩 있는 경우
  - `result = scanf("%d %d, &num1, &num2);`
    - result에는 2가 들어간다. 2개 읽은 것이다.
    - num1에는 12가 num2에는 34가 들어간다.
- 예 4) "12p34"를 입력했을 때
  - `result = scanf("%d %d", &num1, &num2);`
    - result에는 1이 들어간다. 하나를 읽고 p에서 막혀서 다음 정수는 읽지 못 했다.
    - num1에는 12가 들어간다. num2에는 값을 읽어 저장하지 못 했다.
- 예 5) "12 34.56"를 입력했을 때
  - `result = scanf("%d %f", &num1, &fnum);`
    - result에는 2가 들어간다. 2개 읽은 것이다.
    - num1에는 12가 fnum에는 34.56이 들어간다.
- 예 6) 1234.56를 입력했을 때
  - `result = scanf("%d %f", &num1, &fnum);`
    - result에는 2가 들어간다. 2개 읽은 것이다.
    - num1에는 1234가 fnum에는 0.56이 들어간다.
- 예 7) "12 34.56p"를 입력했을 때
  - `result = scanf("%d %f", &num1, &fnum);`
    - result에는 2가 들어간다. 2개 읽은 것이다.
    - num1에는 12가 fnum에는 34.56이 들어간다.
    - p는 여전히 스트림에 남아 있음에 유의하자.
- 예 8) "12 po 34"를 입력했을 때
  - `result = scanf("%d %s %d", &num1, str, &num2)`;
    - result에는 3이 들어간다. 3개 읽은 것이다.
    - num1에는 12가 str에는 po가 num2에는 34가 들어간다.
- 예 9) "12.34p"를 입력했을 때
  - `result = scanf("%f %s", &fnum, str);`
    - result에는 2가 들어간다.
    - fnum에는 12.34가 str에는 p가 들어간다.
- 예 10) "12.34p"를 입력했을 때
  - `result = scanf("%d %s", &num1, str);`
    - result에는 2가 들어간다.
    - num1에는 12가 str에는 .34p가 들어간다.
    - str은 입력을 문자열로 다 읽을 수 있다.

## 문자를 읽을 때 scanf()의 문제점과 해결책, clearerr()

- %s를 쓸 때 배열 크기보다 큰 문자열이 들어오면 버퍼 오버플로가 발생한다.
- 따라서 scanf()는 문자열을 읽을 때 쓰면 별로다.
- 뿐만 아니라 다른 자료형을 읽을 때도 툭하면 무한 루프에 빠질 위험이 있다.

  ```c
  int num;
  int sum = 0;

  while (TRUE)
  {
      if (scanf("%d", &num) == 0)
      {
          printf("Error!\n");
          continue;
      }

      if (num == 0)
      {
          break;
      }

      sum += num;
  }

  printf("Sum: %d\n", sum);

  ```

  - 위의 코드는 어느 부분에서 위험할까? 잘못 해서 문자를 넣으면 계속해서 실패한다. 버퍼에 빼지 않은 문자가 계속 남아있기 때문이다.
    - 해결법? fgets()와 sscanf() 함수를 같이 쓰는 게 좋다.
    - 안전하게(무한 루프 없이) 숫자를 읽는 예

      ```c
      #define LINE_LENGTH (1024)

      /* ... */

      int sum = 0;
      int num;
      char line[LINE_LENGTH];

      while (TRUE)
      {
          if (fgets(line, LINE_LENGTH, stdin) == NULL)
          {
              clearerr(stdin);
              break;
          }

          if (sscanf(line, "%d", &num) == 1)
          {
              sum += num;
          }
      }

      ```

      - 버퍼 오버플로는 발생하지 않는다. 버퍼보다 긴 문자열이 들어오면 그냥 짤리는 게 전부이다.
    - 버퍼 오버플로 문제 없이 문자열 읽기

      ```c
      #define LENGTH (4096)

      /* ... */

      char line[LENGTH];
      char word[LENGTH];

      while (TRUE)
      {
          if (fgets(line, LINE_LENGTH, stdin) == NULL)
          {
              clearerr(stdin);
              break;
          }

          if (sscanf(line, "%d", word) == 1)
          {
              printf("%s\n", word);
          }
      }

      ```

- clearerr()
  - clear error를 말한다.
  - 스트림을 읽거나 쓸 때 EOF를 만나면 그 스트림의 EOF 표시자(indicator)가 세팅된다.
  - 그 외의 이유로 실패하면 오류 표시자(error indicator)를 세팅한다.
  - 그게 잘 안 지워져서 다음에 읽거나 쓸 때 계속 실패할 수도 있다. 그래서 그 오류를 지워주는 것이다.
  - 저 표시자의 세팅 여부를 보고 싶으면 `feof()` 나 `ferror()` 함수를 사용하면 된다.
  - 음...크게 중요하진 않다.

- 한 데이터씩 읽는 방법이 유용한 경우?
  - 텍스트를 다른 자료형으로 곧바로 읽어오는 가장 간단한 방법.
  - 사용자 입력을 받을 때(또는 여러 데이터가 혼용된 텍스트 파일을 읽어올 때) 사용하는 가장 흔한 방법이다.

## 한 블록씩 읽기

- 앞에서 본 방법은 텍스트 데이터를 읽는 방법이었다. 그러면 바이너리 데이터는 어떻게 읽을까? 그게 바로 한 블록씩 읽는 것이다.
- fread()
  - `size_t fread(void* buffer, size_t size, size_t count, FILE* stream);`
    - stream으로부터 size 바이트짜리 데이터를 총 count 개수 만큼 읽어서 buffer에 저장하는 것이다.
    - EOF를 만나면 당연히 멈춘다.
      - 그러면 count보다 적은 수를 읽을 수도 있다는 이야기이다. 그러면 실제로 읽은 개수(바이트 아님)를 반환한다.

- fwrite()
  - `size_t fwrite(const void* buffer, size_t size, size_t count, FILE* stream);`

```c
int nums[64];
size_t num_read;
FILE* fstream;

num_read = fread(nums, sizeof(nums[0]), 64, fstream);
fwrite(nums, sizeof(nums[0]), 64, fstream);

```

- int 블록을 총 64개 저장한다.
  - 총 바이트는 `64 * sizeof(int)` 만큼이다.

- 한 블록씩 읽는 방법이 유용한 경우?
  - 바이너리 데이터를 읽기 위해 필요하다. 바이너리 데이터를 하나씩 읽을 수도 있지만 한꺼번에 읽으면 성능이 향상된다.
- 한 블록씩 읽을 때 주의할 점
  - 기본 데이터형의 크기는 시스템마다 다르다.
  - 파일에 저장할 데이터 크기를 저장해두는 게 좋다.

## 파일 입출력

- C에서의 파일 관련 연산은 이런 식이다.
  - 1) 파일을 열어서 파일 스트림을 가져온다.
  - 2) 그 파일 스트림을 사용해서 하고 싶은 걸 한다.
  - 3) 그 파일을 닫아준다.

## 파일 열기

```c
#include <stdio.h>

#define LENGTH (1024)

FILE* stream;
char list[LENGTH];

stream = fopen("hello.txt", "r");

if (fgets(list, LENGTH, stream) != NULL)
{
    printf("%s", list);
}

```

- 파일 열기
  - `FILE* fopen(const char* filename, const char* mode);`
    - filename으로 지정된 파일을 연다.
    - 열 때 사용하는 모드(읽기 전용, 이진 파일 등)는 mode로 지정한다.
    - 반환 값은 파일 스트림 포인터이다.
- 파일 열기 모드

| 모드 | 설명                                            | 파일이 이미 있다면        | 파일이 없다면      |
| ---- | ----------------------------------------------- | ------------------------- | ------------------ |
| r    | read, 파일을 읽기 전용으로 연다                 | 파일의 첫 부분부터 읽는다 | 열기에 실패한다    |
| w    | write, 파일을 쓰기 전용으로 생성한다            | 내용을 모두 없앤다        | 새 파일을 생성한다 |
| a    | append, 파일에 이어 쓴다                        | 파일의 끝부분부터 읽는다  | 새 파일을 생성한다 |
| r+   | read extended, 읽기/쓰기용으로 파일을 연다      | 파일의 첫 부분부터 읽는다 | 오류               |
| w+   | write extended, 일기/쓰기용으로 파일을 생성한다 | 파일의 내용을 모두 없앤다 | 새 파일을 생성한다 |
| a+   | append extended, 읽기/쓰기용으로 파일을 연다    | 파일의 끝부분부터 읽는다  | 새 파일을 생성한다 |

- 위의 파일 열기 모드는 텍스트 모드이다.
  - b를 붙이면 이진 모드로 파일을 연다.
    - 사실 유닉스 계열에서는 아무런 차이가 없다.
    - 윈도우에서는 새 줄 문자 처리하는 것만 달라진다.

## 파일에 쓰기/읽기 예, fflush(), 파일에 이어 쓰기 예

- 파일에 쓰기 예(미완성)

```c
#include <stdio.h>
#include <string.h>

#define LENGTH (5)

/* ... */

FILE* stream;
int scores[LENGTH] = { 100, 34, 95, 56, 72 };

stream = fopen(filename, "wb");

fwrite(scores, sizeof(scores[0]), LENGTH, stream);

/* ... */

write_file("hello.txt");

```

- 그런데 위의 경우는 바로 파일에 써지지 않는다. 보통 쓰기는 버퍼링 때문에 바로 파일에 저장되지 않는다. 이런 경우엔,
  - '\n'이 들어오거나
  - fflush()를 호출해야 한다
- 그런데...fwrite()는 '\n'을 '\n'으로 인식을 못 한다.
  - fwrite() 의 첫 번째 인자로 받는 buffer가 `char*`도 아니고 `int*`도 아니고 `float*` 도 아니고 그냥 `void*` 이다.
  - 즉 fwrite() 입장에서는 그냥 비트 패턴이 들어온다는 것이다. 숫자에 의미가 없다.
    - 0X0A가 fwrite() 입장에서는 '\n'을 의미하는 건지 정수 10을 의미하는 건지 뭔지 알 수가 없다.
  - 따라서 fflush() 만이 유일한 해결법이다.

- 파일에 쓰기 예(완성)

```c
#include <stdio.h>
#include <string.h>

#define LENGTH (5)

/* ... */

FILE* stream;
int scores[LENGTH] = { 100, 34, 95, 56, 72 };

stream = fopen(filename, "wb");

fwrite(scores, sizeof(scores[0]), LENGTH, stream);

/* ... */

write_file("hello.txt");
fflush(stream);

```

- 파일에 읽기 예(미완성)

```c
#include <stdio.h>
#include <string.h>

#define LENGTH (6)

/* ... */

void read_file(const char* filename)
{
    FILE* stream;
    char data[LENGTH];

    stream = fopen(filename, "rb");

    while (TRUE)
    {
        if (fgets(data, LENGTH, stream) == NULL)
        {
            break;
        }
        printf("%s\n", data);
    }
}

```

- 파일에 이어쓰기 예 (미완성)

```c
#include <stdio.h>
#include <string.h>

#define LENGTH (1024)

/* ... */

void append_file(const char* filename)
{
    FILE* stream;
    char data[LENGTH];

    stream = fopen(filename, "ab");

    if (fgets(data, LENGTH, stdin) != NULL)
    {
        fwrite(data, strlen(data), 1, stream);
    }
}

/* ... */

append_file("hello.txt");

```

## 파일 닫기, 파일 오류처리, stderr, strerror(), perror()

- 위에서 한 것을 보면...뭔가 빠진 느낌이다. 파일을 닫는 부분이 없다. 파일을 열기만 하고 닫는 부분이 없다.
- 파일은 내가 열었으면 내가 닫아야 한다.
  - 파일은 운영체제가 열어주는 것이다. 운영체제는 우리가 언제 파일을 다 써서 필요 없는지 모른다. 따라서 직접 말해주지 않으면 알아서 닫아주지 않는다.
  - 계속 파일을 일기만 하면 어느 순간 운영체제가 더 이상은 파일을 열 수 없다고 하며 뻗을 수 있다.
- 파일을 열었으면 언제나 닫자!
- fclose()
  - `int fclose(FILE* stream);`
    - 파일을 닫는 함수이다.
    - 성공하면 0, 실패하면 EOF를 반환한다.
    - 버퍼링 중인 스트림은 이렇게 작동한다.
      - 출력 스트림: 버퍼에 남아있는 데이터는 파일로 다 내보낸다.
      - 입력 스트림: 무시하고 버린다.
- C는 예외 처리가 없다. 그러면 fopen() 함수는 실패하면 무엇을 반환할까?
  - `FILE* fopen(const char* filename, const char* mode);`
    - 실패하면 NULL 포인터를 반환한다.

    ```c
    #include <stdio.h>

    /* ... */

    #define LENGTH (1024)

    /* ... */

    void open_file(const char* filename)
    {
        FILE* stream;
        char data[LENGTH];

        stream = fopen(filename "rb");
        if (stream == NULL)
        {
            fprintf(stderr, "error while opening %s", filename);
            return;
        }

        if (fgets(data, LENGTH, stream) != NULL)
        {
            printf("%s", data);
        }

        if (fclose(stream) != 0)
        {
            fprintf(stderr, "error while closing");
        }
    }

    /* ... */

    int main(void)
    {
        open_file("hello.txt");
        return 0;
    }

    ```

- stderr
  - 프로그램이 실행될 때 자동으로 3개의 스트림을 만들어 준다. stdin, stdout 그리고 stderr이다.
  - stderr은 stdout 하고 비슷하다.
  - 다만, stderr는 오류 관련 메세지를 출력하는 전용 스트림이다.
  - stderr은 보통 버퍼링을 사용하지 않는다.

- 파일 열기에 실패할 경우의 코드

```c
#include <stdio.h>

/* ... */

void open_file(const char* filename)
{
    FILE* stream = fopen(filename, "rb");
    if (!stream)
    {
        fprintf(stderr, "error while opening %s", filename);
        return;
    }

    /* ... */

}

```

- 실패한 이유를 알고 싶을 경우엔 어떻게 해야할까?
  - 몇몇 표준 라이브러리 함수들이 실패할 때 그 이유를 오류코드(숫자)로 어딘가에 저장해 둔다.
  - 그게 errno라는 매크로(#define)
    - <error.h> 안에 있다.

  - 오류 코드를 보여주는 코드

  ```c
  #include <stdio.h>
  #include <errno.h>

  void open_file(const char* filename)
  {
      FILE* stream = fopen(filename, "rb");
      if (!stream)
      {
          fprintf(stderr, "[%d] error while opening %s", errno, filename);
          return;
      }

      /* ... */
  }

  ```

  - 오류 코드를 말로 설명해주는 함수가 있다.
    - `char* strerror(int errnum);`
      - <string.h>에 정의된다.
      - errno를 넣으면 문자열로 된 친절한 설명을 돌려준다.

      ```c
      #include <stdio.h>
      #include <errno.h>

      void open_file(const char* filename)
      {
          FILE* stream = fopen(filename, "rb");
          if (!stream)
          {
              fprintf(stderr, "%s - %s", filename, strerror(errno));
              return;
          }

          /* 코드 생략 */
      }

      ```

      - perror()도 있다.
        - `void perror(const char* s);`

      ```c
      #include <stdio.h>
      #include <errno.h>

      void open_file(const char* filename)
      {
          FILE* stream = fopen(filename, "rb");
          if (!stream)
          {
              perror("error while opening");
              return;
          }

          /* 코드 생략 */
      }

      ```

- C에서의 오류처리는 보통 이런 식이다.
  - 함수가 곧바로 오류 코드를 반환한다.
  - 내부적으로 오류 코드를 전역 변수로 들고 있다가 검사한다.

## 파일 탐색

## 입출력 리디렉션

## 커맨드 라인 인자

## 커맨드 라인 인자 메모리 뷰
