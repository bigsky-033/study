# C 스타일 문자열, 출력

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 문자열의 표현과 길이

- 문자열과 다른 기본 자료형의 차이
  - 기본 자료형의 크기는 언제나 고정이다.
  - 문자열의 크기는 고정적이지 않다.
- 즉, 컴퓨터에게 '문자열 하나 읽어'라고 명령할 수가 없어서 기본 자료형이 될 수 없다.
- '여러 개'의 문자를 표현하려면 무엇이 좋을까? => 배열
- 따라서 `char str[글자수]` 라고 표현할 수 있을 것 같다.
- 그런데...배열의 길이를 C에서는 구할 수가 없다. 따라서 프로그래머가 그 길이를 따로 기록해 두어야 한다.

## 문자열 관리 시 길이의 문제

- 문자열의 길이를 아래처럼 별도 변수로 두면 문자열을 바꾸었을 때 길이를 같이 바꾸지 않는 실수를 할 수 있다.

```c
void print_string(void)
{
    char chars[] = "Pointer is the best!"; /* 이 문자열이 길이가 다른 문자열로 바뀐다면? */
    const size_t NUM_CHARS = 20;

    size_t i;
    for (i = 0; i < NUM_CHARS; ++i)
    {
        printf("%c", chars[i]);
    }
}

```

## 문자열 길이 문제 해결 방법 1

- 방법 1) 길이를 배열 첫 위치에 저장하기
  - 첫 메모리 위치에 문자열 길이를 저장하고, 실제 문자열이 뒤 따라오게 하면 어떨까?
    - unsigned char로 길이를 저장하기엔 너무 짧다.
      - 그러면...이걸 int로 저장하고 그 뒤에 문자는 char로 저장하면 안 될까?
        - 다른 언어에서는 이런 식으로 많이 구현해 두었다!
    - C에서는 이 방식이 단점이 있다. 글자 하나가 1 바이트인데 여기에 4바이트나 쓰면 낭비일 수 있다. 이건 클래스나 구조체에서 쓰는 방식이다. 또한, 순수 C 코드로 이것을 어떻게 작성해야 할 지가 애매하다. 첫 데이터는 `int*`로 캐스팅해서 읽고 그 다음부터는 `char*`로 읽어야 하나?

## 문자열 길이 문제 해결 방법 2, C 스타일

- 방법 2) 길이를 기억하진 않는다. 다만, 문자열이 끝나는 위치에 특수한 문자를 두어 끝나는 걸 알려준다.
  - 그런 특수한 문자가 뭐가 있을까? 아스키코드 중에 화면에 출력되지 않는 특수문자들이 있다. 이런 것들을 제어 문자(control character)라고 한다.
    - 제어 문자 중 하나가 바로 0, 널 문자(null character)라고 불린다. 널 포인터하고는 다른 것이다.
    - `char null_char = '\0';`
    - `0`은 숫자 영하고 같다. `\`는 이스케이프 문자이다. 근데 아스키 코드로 0이니 `char null_char = 0;`으로 작성이 가능하기도 하다.
    - 그러나 의미를 명확하게 하기 위해 `'\0'` 으로 써주자.
- C 스타일 문자열이라 하면 널 문자로 끝나는 char 배열을 말한다.
  - `char[]` 로만 구성된다.
  - 문자열이 끝나는 곳에 널 문자를 붙인다.
- 아래의 두 코드는 저장위치 외에는 동일하다. 참고로 문자열 뒤에 별도로 `\0`을 넣지 않아도 컴파일러가 알아서 넣어준다.
  - `char str1[] = "abc"; /* 스택에 "abc 저장" */`
  - `char* str2 = "abc";  /* 데이터 섹션에 "abc" 저장 */`
- 단 이런 경우에는 컴파일러가 `\0`을 넣어주지 않는다.
  - `char str[] = { 'a', 'b', 'c' };` 주의해야 한다.

## C 스타일 문자열의 장단점, 문자열 길이 구하기

- 언제나 문자열의 배열에 널 문자도 포함되어 있다는 걸 있지 말자.
  - 예를 들어, `const str[] = "ABCD"; printf("str length: %d\n", sizeof(str));` 를 하면 널 문자를 포함해 5가 출력된다.
- C 스타일 문자열의 장점
  - 가장 최소한의 메모리를 사용할 수 있다.
  - 한 가지 데이터형으로 문자열과 길이를 다 표현한다.
- C 스타일 문자열의 단점
  - 어떤 문자열의 길이를 알려면 배열을 끝까지 훑어야 한다 => O(N)
- 문자열의 길이 구하기
  - char 배열의 요소를 처음부터 차례대로 읽는다. 널 문자를 만나면 멈춘다. 그때까지의 카운트를 반환한다.

    ```c
    size_t get_string_length(const char* str)
    {
        size_t i;
        for (i = 0; str[i] != '\0'; ++i)
        {
        }
        return i;
    }

    /* 포인터를 이용해 조금 더 효율적으로 구현하는 방법 */
    size_t get_string_length(const char* str)
    {
        const char* p = str;
        while (*p++ != '\0')
        {
        }
        return p - str - 1;
    }


    /* 포인터를 이용해 조금 더 효율적으로 구현하는 방법 */
    size_t get_string_length(const char* str)
    {
        size_t count = 0;
        const char* p = str;

        while (*p++ != '\0')
        {
            ++count;
        }

        return count;
    }

    ```

- 문자열을 구하는 함수는 사실 이미 있다.
  - `<string.h>`를 include 하면 사용할 수 있다.
  - `size_t strlen(const char* str);`
- 문자열 길이와 관련된 실수에 유의하자
  - 배열로 선언해서 마지막에 널 문자를 빠뜨리는 경우!
  - 널 문자가 길이에 포함되어야 하는 걸 깜빡한 경우!
- 외부에서 입출력을 통해 문자열을 읽을 때 유의해야 한다.
  - C11의 strlen_s() 함수가 이런 문제를 해결하기도 한다.
- strlen(str)은 안전하지 않을 수도 있다.
  - '\0' 이 안 나온다면?
  - 하드웨어가 보호하는 메모리를 읽는다면?

## 문자열 조작, 두 문자열의 비교

- 두 문자를 비교하는 함수를 만들어보자.
  - 같으면 0을 반환
  - str0이 str1보다 작으면 < 0
  - str0이 str1보다 크면 > 0
  - 음수도 반환해야 하니 반환형은 int

```c
/* 두 문자를 비교해서 사전식 순서(lexicographical order)로 어떤 문자의 아스키 코드가 더 작냐/같냐/크냐를 판별 */
compare_string(const char* str0, const char* str1)
{

}

```

## 문자열 비교 알고리즘

- 두 문자열에서 문자를 하나씩 읽는다. - c0, c1
- 두 문자를 비교한다.
  - c0이 c1보다 작으면 음수를 반환한다.
  - c0이 c1보다 크면 양수를 반환한다.
  - (c0과 c1이 같고) 널 문자면 0을 반환한다.
- 다음 문자로 이동 후 1번 단계로 돌아간다.

## 더 효율적인 문자열 비교 함수 작성하기 strcmp()와 strncmp()

- 직접 구현하는 방법들

```c
int compare_string(const char* str0, const char* str1)
{
    /* 아직 끝에 도달하지 않았고, 둘이 같을 때까지 포인터를 다음으로 이동 */
    while (*str0 != '\0' && *str0 == *str1)
    {
        ++str0;
        ++str1;
    }

    return *str0 - *str1;
}

```

```c
int compare_string(const char* str0, const char* str1)
{
    /* 아직 끝에 도달하지 않았고, 둘이 같을 때까지 포인터를 다음으로 이동 */
    while (*str0 != '\0' && *str0 == *str1)
    {
        ++str0;
        ++str1;
    }

    if (*str0 == *str1)
    {
        return 0;
    }
    return *str0 > *str1 ? 1 : -1;
}

```

- strcmp()
  - `int strcmp(const char* lhs, const char* rhs);`
  - <string.h> 안에 있다.
- strncmp()
  - `int strncmp(const char* lhs, const char* rhs, size_t count);`
  - <string.h> 안에 있다.
  - 최대 n 문자까지만 비교한다.

## 문자열 복사, strcpy(), strncpy()

- 문자열 복사 하는 코드 예

```c
/* dest는 const가 아니고 src는 const 이다. */
void copy_string(char* dest, const char* src)
{
    while (*src != '\0')
    {
        *dest++ = *src++;
    }

    /* 빠뜨리지 않도록 주의! */
    *dest = '\0';
}

```

- strcpy()
  - `char* strcpy(char* dest, const char* src);`
  - <string.h> 안에 있다.
  - 반환값 char*는 dest를 반환한다. 이걸 잘 쓰지는 않는다.
  - C11에서 나온 strcpy_s()는 errno_t를 반환한다.
    - `errno_t strcpy_s(char *restrict dest, rsize_t destsz, const char *restrict src);`
    - dest 가 src보다 짧으면?
      - 뒤에 실제 소유하지 않은 메모리에도 쓰게 된다. 위험하다.
      - C89에서는 어떻게 해야 그나마 안전할까? => strncpy()
- strncpy()
  - `char* strncpy(char* dest, const char* src, size_t count);`
  - <string.h> 안에 있다.
  - 최대 count만큼 복사한다.
  - 비교적 안전한 문자열 복사 방법이다.
  - src가 count보다 짧거나 같으면 남은걸 다 0('\0')으로 채워준다.
  - src가 count보다 길다면, count만큼 복사한다. 그런데 이렇게 되면 널 문자를 붙일 곳이 없다. 따라서 안 붙여준다;
    - 그래서 이걸 보완하기 위해 보통 이런 방식으로 코드를 쓴다.

    ```c
    strncpy(dest, src, DEST_SIZE);
    dest[DEST_SIZE - 1] = '\0'; /* 추가 */

    ```

## 문자열 합치기, strcat(), strncat()

- strcat()
  - `char* strcat(char* dest, const char* src);`
  - <string.h> 안에 있다.
  - src의 문자열을 dest 뒤에 덧붙이는 함수이다.
    - dest의 널 문자가 들어있는 위치부터 src의 문자열 추가
    - 바꿔 말하면 dest의 널 문자가 src[0]으로 교체
  - dest의 길이가 충분해야 한다. 이 길이를 넘어 쓸 경우 정의되지 않은 결과가 발생할 수 있다.
- strncat()
  - `char* strncat(char* dest, const char* src, size_t count);`
  - <string.h> 안에 있다.
  - 최대 count 개 만큼 src 문자열을 dest 뒤에 덧붙이는 함수이다.
    - dest의 널 문자가 들어있는 위치부터 src의 문자열 추가
    - 바꿔 말하면 dest의 널 문자가 src[0]으로 교체
  - 따라서, 최대 count+1 개의 문자를 덧붙여 준다(+1은 널 문자).
  - dest의 길이보다 길게 쓰면 마찬가지로 정의되지 않은 결과가 발생한다.
    - strcat() 보다는 조금 더 안전하긴 하지만 여전히 문제가 발생할 가능성은 있다.
    - 보통 아래와 같은 식으로 많이 사용한다.

    ```c
    #define DEST_COUNT (20)

    const char* src = "POCU";
    char dest[DEST_COUNT] = "Hi ";

    strncat(dest, src, DEST_COUNT - strlen(dest) - 1);

    ```

- C11에선 이보다 안전한 strcat_s(), strncat_s() 함수가 있다.

## 문자열 찾기

- 없는 문자열을 찾는 경우 => NULL
- 존재하는 문자열을 찾는 경우 => 찾고자 하는 문자열을 찾은 위치부터 시작하는 전체 문자열
  - 원래 있던 문자열에서 위치를 찾아 반환하는 식이다.

```c
#include <stdio.h>
#include <string.h>

int main(void)
{
    char msg[] = "I love string! I love C! I love programming!";

    char* result = strstr(msg, "int");
    printf("result: %s\n", result == NULL ? "(null)" : result);

    return 0;
}

```

- `char* strstr(const char* str, const char* substr);`
  - <string.h> 안에 있다.
  - 반환 값: char 포인터
    - substr이 str에 있다면: 해당 substr이 시작하는 주소
    - substr이 str에 없다면: 널 포인터(NULL)
  - 주의할 점이 있다. 인풋은 const인데 반환 값은 const가 아니다. 따라서 원래 const를 쓰던 걸 저기에 넣으면, 받아온 포인터는 const가 아니기 때문에 받아온 포인터로는 값을 바꿀 수 있게 된다. 주의해야 한다.

## 문자열 찾기 함수가 메모리 주소를 반환하는 이유

- 새로운 문자열을 만들어 반환할 경우 메모리 관리 측면에서 효율적이지 못하고 실수할 수 있기 때문이다.
  - 새 문자열을 반환하려면 메모리 '어딘가'에 그 문자열을 복사해야 한다.
    - 복사하는 위치가 만약 스택이라면, 함수가 끝나면 사라진다. 따라서 반환 값이 더이상 유효하지 않은 메모리 주소가 된다.
    - 복사하는 위치가 힙이라면, 메모리 할당을 운영체제에게 요청해야 하기 때문에 느리다. 그리고 더 이상 사용하지 않을 경우 프로그래머가 직접 메모리 해제 함수를 호출해야 하는데, 깜빡 하고 하지 않을 수 있다. 메모리 누수가 발생할 수 있다.

## 문자열 토큰화

- C에서는 메모리를 추가로 할당하지 않으며 토큰화를 할 수 있는 방법이 있다.
- delimiter를 널 문자로 바꾸는 식이다.

```c
char msg[] = "Hi, there. Hello. Bye";
const char delims[] = ",. ";

char* token = strtok(msg, delims);
while (token != NULL)
{
    printf("%s\n", token);
    token = strtok(NULL, delims); /* NULL 이 매개변수로 들어간다. */
}

```

- 토큰화하는 문자열은 const가 아니다! 원본이 바뀐다.
- 함수 매개변수로 NULL이 들어올 때 그전에 받았던 msg를 사용한다. 이건 어디에 저장되어 있어야 한다. 어디일까?
  - 함수 내에서 기억할 수 있는 장치가 필요하다. 함수 내 정적 변수가 제일 적합하지 않을까?
- `char* strtok(char* str, const char* delims);`

## 출력, 서식 지정(formatted) 출력, 서식 문자열(format string)

### 출력

- 프로그램에서 프로그램 외부로 데이터를 보여주는 행위이다.
- 서식 지정(formatted) 출력
  - C에서 출력을 논할 때 가장 기본이 되는 함수이다.
    - printf(): 콘솔창(stdout)에 출력한다.
    - fprintf(): 스트림에 출력한다.
    - sprintf(): 문자열에 출력한다.
  - fprint()와 sprintf()는 printf()하고 작동법이 동일하다.

  ```c
  const char* msg = "Hello World!";

  /* 아래의 두 함수는 동일하게 동작한다. */ 
  fprintf(stdout, "%s\n", msg);
  printf("%s\n", msg);

  ```

  - printf()의 첫 번째 매개변수는 언제나 문자열이어야 한다. int형 변수를 넣는다고 int를 출력해주지 않는다. C는 함수 오버로딩 없다.

### 서식 문자열(format string)

- %로 시작하는 문자열을 말한다.
- 소수점 이하 자리수, 자리수 정렬, 어떤 데이터를 출력할지 등을 알려주는 문자열이다.

## 일반적인 서식 문자열 형식, 서식 지정자(format specifier)

### 일반적인 서식 문자열 형식

- `%[플래그][너비][.숫자 정밀도 | .문자열 최소/최대 출력 개수][길이]서식 지정자`
- 일반적으로 '%' 뒤에 최대 4개의 지정자를 가질 수 있다.
  - 1) 플래그 (선택)
  - 2) 너비 (선택)
  - 3) 정밀도 (선택)
    - 숫자 정밀도
    - 문자열 최소/최대 출력 개수
  - 4) 길이 (선택)
  - 5) 서식 지정자 (`필수`)
- 반드시 순서를 지켜서 작성해야 한다.

### 서식 지정자(format specifier)

| 서식 지정자 | 내용                                                                  | 사용 방법                                     |
| ----------- | --------------------------------------------------------------------- | --------------------------------------------- |
| %           | '%'를 출력                                                            | `printf("%%\n");`                             |
| c           | 문자(char) 출력                                                       | `printf("%c\n", 'D');`                        |
| s           | 문자열(char[]) 출력                                                   | `printf("%s\n", "Hello");`                    |
| d           | 부호 있는 정수 출력                                                   | `printf("%d\n", -10);`                        |
| u           | 부호 없는 정수 출력                                                   | `printf("%u\n", 10);`                         |
| o           | 부호 없는 정수를 8진수로 출력. 숫자 없에 '0'은 안 붙여준다.           | `printf("%0\n", 10");`                        |
| x           | 부호 없는 정수를 16진수(소문자)로 출력. 숫자 앞에 '0x'는 안 붙여준다. | `printf("%x\n", 10);`                         |
| X           | 부호 없는 정수를 16진수(대문자)로 출력. 숫자 앞에 '0X'는 안 붙여준다. | `printf("%X\n", 10);`                         |
| f           | 부동소수점 출력, double, float 모두 가능                              | `printf(%f\n", 3.14);`                        |
| e/E         | 부동소수점을 지수표기법으로 출력                                      | `printf("%e\n", 3.14); printf("%E\n", 3.14);` |
| p           | 포인터 값을 출력                                                      | `printf("%p\n", (void*)name);`                |

- `%u`에 부호 있는 수를 넣을 경우, 해당 수의 비트 패턴에 해당하는 부호 없는 수가 출력된다.
- `%p`는 주소를 출력하는 데 (void*)만 받는다.
  - 모든 주소는 어차피 길이가 같다. 따라서 어떤 포인터를 (void*)로 캐스팅해도 안전하다.

- 너비
  - 출력 너비
    - 정수: 출력 너비(출력할 값이 너비보다 작으면 공백)
      - 사용법: `printf("%5d\n", number);`

- 플래그

| 플래그 | 내용                                                                  | 기본                    | 사용 방법                   |
| ------ | --------------------------------------------------------------------- | ----------------------- | --------------------------- |
| -      | 왼쪽 정렬                                                             | 오른쪽 정렬             | `printf("%-5d\n", number);` |
| 0      | 빈 공백을 0으로 체워준다.                                             | '-' 가 있을 경우 무시   | `printf("%05d\n", number);` |
| +      | 항상 부호(+,-)를 표시한다.                                            | 기본은 움수 기호만 출력 | `printf("%+5d\n", number);` |
| 공백   | 양수인 경우에도 부호 칸을 비워 둔다                                   | '+' 가 있을 경우 무시   | `printf("% d\n", number);`  |
| #o     | 부호 없는 정수를 8진수로 출력. 숫자 앞에 항상 '0'이 붙는다.           |                         | `printf("%#o\n", number);`  |
| #x     | 부호 없는 정수를 16진수로 출력(소문자). 숫자 앞에 항상 '0x'가 붙는다. |                         | `printf("%#x\n", number);`  |
| #X     | 부호 없는 정수를 16진수로 출력(대문자), 숫자 앞에 항상 '0X'가 붙는다. |                         | `printf("%#X\n", number);`  |

- `#`은 다른데에도 붙일 수 있다. 그나마 유용한 곳은 x나 X에 붙일 때이다.

- 정밀도 - 서식 지정자 'f'와 사용
  - `전체_최소_너비.소수점_아랫_자리_수`
  - (소수점 포함) 원래 숫자의 너비보다 최소너비가 크면 공백으로 채운다.
  - (소수점 포함하지 않음) 원래 숫자의 소수점 아래자리 수보다 소수점 아랫자리수가 크면 0으로 채워준다.
  - 기본 소수점 아랫 자리 수: 6
- 정밀도 - 서식 지정자 's'와 사용
  - `최소_너비.최대_너비`
  - 출력할 문자열의 길이가 최소 너비보다 작으면 왼쪽을 공백으로 채운다.
  - 출력할 문자열의 길이가 최대 너비보다 크면 자른다.

- 길이 수정자(length modifier)

| 길이 수정자 | 서식 지정자 |             | 사용법                     |
| ----------- | ----------- | ----------- | -------------------------- |
|             | d           | int         | `printf("%d\n", number);`  |
| l           | d           | long int    | `printf("%ld\n", number);` |
|             | f           | double      | `printf("%f\n", number);`  |
| L           | f           | long double | `printf("lf", number);`    |

- 인자의 바이트 크기를 지정해준다. 하지만 최근 플랫폼에선 별 의미가 없다.

## 서식 문자열이 필요한 이유, fprintf(), stdout, 버퍼링, sprintf()

- C에는 함수 오버로딩이 없다. `printf(int), printf(char)` 불가능.
- 임시 문자열 등을 자동으로 생성해주지 않는다.
  - 컴파일 오류: `printf("Hello, " + name);`
- fprintf()도 똑같다. 단, 여기에서는 스트림을 사용한다.
  - 보통 프로그램이 실행할 때 기본적으로 3개의 스트림을 준다. 이 3개의 스트림을 표준 스트림(standard stream)이라고 한다.
    - stdout (콘솔 출력)
      - 우리가 계속 봐왔던 그 콘솔 스트림이다.
      - stdout은 보통 라인 버퍼링(line buffering)을 사용한다.
      - 버퍼링
        - 출력할 내용이 있어도 곧바로 출력하지 않고 쌓아둔다.
        - 어느 정도 버퍼가 차면 그제서야 철력한다.
      - 라인 버퍼링: 다음과 같은 경우에 버퍼를 비운다.
        - 버퍼가 꽉 차거나
        - 버퍼에 '\n'이 들어올 때
      - 강제로 버퍼를 비우고 싶다면 `fflush(stdout);`을 호출하면 된다.
      - 버퍼링의 종류
        - 풀 버퍼링(full buffering): 버퍼가 가득 차면 비운다. 라인 버퍼링과 마찬가지로 fflush()로 강제로 비울 수 있다.
        - 라인 버퍼링(line buffering): 버퍼가 꽉 차거나, 버퍼에 '\n'이 들어오면 버퍼를 비운다.
        - 버퍼링 없음(no buffering): 버퍼를 사용하지 않음.
        - 표준에서는 stdout, stderr, stdin의 버퍼링 종류를 지정하지 않는다. 따라서 구현마다 다를 수 있다.
    - stdin (콘솔 입력)
    - stderr (콘솔 출력, 하지만 오류 메세지를 출력하는 스트림)
- sprintf()도 있다.
  - `int springf(char* buffer, const char* format, ...);`
  - char 배열에 출력한다.
  - 충분히 큰 버퍼를 잡아주지 않으면 위험할 수 있는 함수이기도 하다.

  ```c
  char buffer[100];
  int score = 100;
  const char* name = "Rachel";

  sprintf(buffer, "%s: %d", name, score);

  printf("%s\n", buffer);

  ```

## 출력 함수의 안정성, 기타 출력 함수

- 안전하지 않은 경우

```c
char buffer[20];
int score = 100;
const char* name = "Carterina Hassinger";

sprintf(buffer, "%s: %d", name, score);

printf("%s\n", buffer);

```

- 안전하지 않은 이유: 실제 소유하지 않은 메모리에 쓸 수 있다.
- C99에 snprintf()가 들어온다.

- 이 외의 출력 함수
  - `int puts(const char* str);`
    - 문자열을 stdout에 출력한다.
    - 마지막에 줄도 바꿔준다: '\n'
    - 결과적으로 `fputs(str, stdout);` 과 같다.
  - `putchar()`
    - 문자열을 stdout에 출력한다.
    - 결과적으로 `fputc(ch, stdout);` 과 같다.
