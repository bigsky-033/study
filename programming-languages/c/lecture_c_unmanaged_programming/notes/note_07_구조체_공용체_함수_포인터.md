# 구조체, 공용체, 함수 포인터

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 구조체, 구조체의 필요성

- 구조체는 영어로 structure 이다.
- 데이터의 집합이다. 멤버 함수 같은 건 없다.
- 구조체는 참조형일까 값형일까?
  - C에서는 모든 자료형이 값형이다. 하지만 주소를 전달하면 참조형처럼 쓸 수 있다. 따라서 주소를 전달하지 않는 한 값형이다.
- 구조체의 필요성
  - 구조체를 이용해서 데이터를 사람이 좀 더 이해하기 편한 형태로 이해할 수 있다.
  - 개발 중 실수를 막는 데도 유용할 수 있다.
    - 컴파일러가 알 수 없는 문제를 잡아줄 수 있다. 예를 들어 같은 형의 데이터 여러 개를 매개 변수로 받을 때 순서가 바뀌면 컴파일러가 실수를 찾을 방법이 없다.
      - `float calculate_BMI(float height, float weight);`
    - 묵시적으로 변환 가능한 자료형이 매개 변수에 여러 개일 때도 마찬가지
      - `int save_data(unsigned char level, float money, const char* name);`
    - 매개변수 목록이 길어질수록 더 문제이다.
- 팁) 실수를 줄이려면 원자성을 보장하는 연산(atomic operation)을 사용하는 게 좋다.

## 구조체의 선언 및 사용

- 구조체의 선언

```c
struct date {
    int year;
    int month;
    int day;
};

```

- 위의 구조체 선언은,
  - date란 구조체(새로운 형)을 만든다.
  - 그 안에 들어가 있는 데이터(멤버 변수)는 총 세개이다.
- 마지막에 세미 콜론 잊지 말 것

- 구조체 변수 선언 및 사용하기
  - `int is monday(struct date date);`
  - `struct date date;`
- 지역 변수 선언 시 0으로 초기화되지 않는다.
  - C의 다른 자료형과 마찬가지이다.
  - 스택 위치에 있던 값을 그냥 가져다 쓴다.
- 다른 데이터 형과 달리 좀...장황하다.
  - `struct date date`; 앞에 struct를 굳이 붙여야 하나...? 왜 만든 자료형은 `struct date` 가 자료형인가?

## typedef 이란?, typedef 사용법

- 간결하게 자료형을 표현하기 위해 typedef를 쓴다.
- size_t가 보통 이런 식으로 표현되어 있다.
  - <stddef.h> -> `typedef unsigned int   size_t`;
- typedef 이란?
  - 이미 있는 자료형에 대해 새로운 별명을 지어주는 것이다.
- typedef 사용법 1)
  - 구조체에 typedef를 쓰면 다른 자료형처럼 간결하게 변수 선언이 가능하다.

  ```c
  struct date {
      int year;
      int month;
      int day;
  };

  typedef struct date date_t;

  /* ... */

  date_t date;

  ```

- typedef 사용법 2, 사용법 3)

```c
/* 사용법 2 */
typedef struct date {
    int year;
    int month;
    int day;
} date_t;

/* 사용법 3 */
typedef struct {
    int year;
    int month;
    int day;
} date_t;

```

- 참고) enum도 마찬가지이다.

```c
/* 기본 */
enum game_role {
    GAME_ROLE_MID,
    GAME_ROLE_JUNGLE
};

enum game_role role = GAME_RILE_MID;

/* typedef 1 */
enum game_role {
    GAME_ROLE_MID,
    GAME_ROLE_JUNGLE
};

typedef enum game_role game_role_t;

game_role_t role = GAME_RILE_MID;

/* typedef 2 */
typedef enum game_role {
    GAME_ROLE_MID,
    GAME_ROLE_JUNGLE
} game_role_t;

game_role_t role = GAME_RILE_MID;

/* typedef 3 */
typedef enum {
    GAME_ROLE_MID,
    GAME_ROLE_JUNGLE
} game_role_t;

game_role_t role = GAME_RILE_MID;

```

- 참고로 C에서 _t로 끝나는 자료형은 보통 이렇게 typedef 한 것을 의미한다.

## 구조체 변수 초기화 하기

- 구조체는 지역 변수 여러 개를 따로따로 사용하는 것과 마찬가지라고 생각하자. 사실 기계는 구조체라는 개념을 모른다. 프로그래밍 언어가 개발 편의를 위해 제공해준 개념 정도라 생각하자.

- 구조체 초기화 방법 1
  
  ```c
  date_t date;

  date.yaer = 0;
  date.month = 0;
  date.day = 0;

  ```

- 구조체 초기화 방법 2

  ```c
  date_t date = { 0, };

  ```

  - 나중에 나올 memset() 을 이용해서 할 수도 있다.

- (권장하지 않음) 구조체 초기화 방법 3

  ```c
  date_t date = { 2043, 10, 1};

  ```

## 구조체 매개변수

- 구조체는 값형이다. 참조형이 아니다. 즉 매개변수로 넘길 때 값이 복사되어 넘어간다.
- 함수 안에서 원본의 값을 바꾸고 싶으면 포인터를 전달해주어야 한다.

  ```c
  void increase_year(date_t* date)
  {
      (*date).yaer = (*date).yaer + 1;
  }

  ```

  - ()가 필요한 이유? 연산자 우선순위 때문이다.
    - `.`이 우선순위가 1이고, `*`의 우선순위는 2이다.
    - 그런데...이건 좀 지저분하다. 이걸 합친 간단한 연산자가 있다.
  - `->` 연산자: 우선순위 1순위의 연산자이다.
    - 아래의 세 코드는 모두 같은 의미이다.
      - `(*date).year = (*date).yaer + 1;`
      - `date->year = date->year + 1;`
      - `date->year++;`

- 구조체 매개변수 베스트 프랙티스
  - 값으로 전달 vs 주소로 전달?
    - 기본 자료형일 때는 기본 데이터의 크기가 작으니 원본 값을 바꿀 필요가 있을 때만 주소로 전달했다.
    - 구조체의 경우 데이터의 크기가 클 수도 있다. 이럴 때는 다 복사하는 게 성능이 느릴 수 있다. 이럴 경우엔 주소로 전달하자. 포인터에 `const` 포인터를 붙이면 원본도 못 바꾸니 안전하다.
  - 구조체 매개변수 vs 여러 개의 개별 변수?
    - 엄밀한 규칙이 있지는 않다.
    - 어떤 회사에서는 4개까지는 낱개 변수로 그 이후에는 구조체로 넘기는 규칙을 쓰는 곳도 있다.
      - 실수를 줄이기 위해서 그리고 성능을 빠르게 하기 위해(주소로 전달) 등의 이유가 있다.

## 함수 반환값으로서의 구조체, 구조체 배열

- 아래의 코드를 보자. 코드를 보면 int 형 하나를 반환하는 것과 마찬가지로 복사에 의한 반환이다.
- C언어의 함수는 반환 값이 하나이다. 아래 처럼 반환하면 실질적으로 여러 개의 값을 반환하는 격이다.

```c
date_t get_dday(void)
{
    date_t date;

    date.year = 2043;
    date.month = 10;
    date.day = 1;

    return date;
}

```

- 구조체끼리 서로 대입할 수도 있다.
  - 컴파일러에 memcpy()란 함수를 이용해 데이터를 통채로 복사해주는 경우도 있고, 그냥 하나하나 대입해주는 경우도 있다.

- 구조체로 배열도 만들 수 있다.
- 다음의 코드는 무엇을 출력할까?

```c
typedef struct name 
{
    char* lastname;
    char* firstname;
} name_t;

/* ... */

char firstname[] = "Lulu";
char lastname[] = "Lee";

name_t name;
name_t clone;

name.lastname = lastname;
name.firstname = firstname;

clone = name;
name.lastname[0] = 'N';

printf("origin: %s %s\n", name.firstName, name.lastname);
printf("clone: %s %s\n", clone.firstName, clone.lastname);

```

## 얕은 복사, 깊은 복사

- 얕은 복사: 실제 데이터가 아니라 주소를 복사하는 것
- 깊은 복사: 원본의 값을 복사해 새로운 값을 만드는 것이다.

## 구조체 사용 시 포인터 저장의 문제

- 포인터 변수가 없는 name_t 예

```c
enum { NAME_LEN = 32 };
typedef struct
{
    char firstname[NAME_LEN];
    char lastname[NAME_LEN];
} name_t;

```

## 구조체를 다른 구조체의 멤버로 사용하기, 바이트 정렬

```c
struct user_info
{
    unsigned int id;
    name_t name;
    usingned short height;
    float weight;
    uisngied short age;
} user_info_t;

/* ... */

user_info_t user;

/* ... */
user.name.firstname[NAME_LEN - 1] = '\0';

```

- 위의 구조체에서 실제 저장된 바이트 수를 보면 short가 4바이트를 차지하고 있는 것을 볼 수 있다.
  - 2바이트를 차지하길 기대했다.
- 이건 각 시스템마다 메모리에 접근할 때 사용하는 주소에 대한 요구사항이 다르기 때문이다.
  - 시스템 상의 제약이 있거나 효율성 때문이다.
- 어떤 시스템은 n바이트 배수인 시작 주소에서만 메모리를 접근 가능하다.
- x86 시스템은 4바이트(워드 크기) 경계에서 읽어오는 게 효율적이다. 이걸 4바이트 경계에 정렬된다(aligned)고 말한다.
- 따라서 컴파일러가 알아서 각 멤버의 시작 위치를 그 경계에 맞춰준다. 이렇게 안 쓰는 바이트를 덧붙이는 걸 패딩(padding)을 넣어준다고 한다.
- 따라서 어떤 아키텍처에서 저장한 파일을 다른 아키텍처에서 읽으면 잘못 읽힐 수도 있다.
- 그럼 패딩을 줄일 수 있는 방법은 무엇일까? 변수의 선언 순서를 바꾸면 될 수도 있다.

  ```c
  struct user_info
  {
      unsigned int id;
      name_t name;
      usingned short height;
      uisngied short age;
      float weight;
  } user_info_t;

  ```

  - #pragma pack을 쓰는 방법도 있지만 C 표준이 아니다.

    ```c
    #pragma pack(push, 1)
    struct user_info
    {
        unsigned int id;
        name_t name;
        usingned short height;
        float weight;
        uisngied short age;
    } user_info_t;
    #pragma pack(pop)

    ```

- 구조체 베스트 프랙티스
  - 구조체를 파일이나 다른 데 저장해야 하는데 바이트 크기가 정확히 맞아야 한다면?
    - 보통 assert()를 사용해서 크기를 확인한다.
      - `#include <assert.h>`
      - `assert(sizeof(user_info_t) == 76);`
  - 어쩔 수 없이 패딩이 생기는 경우가 있다. 그럼 구조체에 명시적으로 패딩을 넣기도 한다.

    ```c
    struct user_info
    {
        unsigned int id;
        name_t name;
        float weight;
        usingned char height;
        uisngied char age;
        char unused[2];
    } user_info_t;

    ```

## 비트 필드

- 비트 플래그: bool 여러 개를 효율적으로 저장(C# 예제)
  - 8개 이하의 bool 값을 하나의 byte에 저장하는 방법

```c#
byte ToBitFlags(bool[] flags)
{
    byte bigflags = 0;

    for (byte i = 0; i < flags.Length; ++i)
    {
        if (flags[i])
        {
            bigflags |= (byte)(1 << i);
        }
    }

    return bigflags;
}

```

- C에서는 구조체를 사용하면 매우 간단히 비트 플래그를 표현할 수 있다.

```c
typedef struct 
{
    unsigned char b0 : 1;
    unsigned char b1 : 1;
    unsigned char b2 : 1;
    unsigned char b3 : 1;
    unsigned char b4 : 1;
    unsigned char b5 : 1;
    unsigned char b6 : 1;
    unsigned char b7 : 1;
} bitflags_t;
```

- `: 1`의 의미는 그 변수 용으로 1비트만 쓰겠다는 것이다.
- `sizeof(bitflags_t)` => 1바이트이다.

- 플래그 전체를 한 번에 체크하려면? 지금 방식은 구조체 전체가 0인지 비교하고 싶은데 안된다.

```c
int is_set = (flags.b1 == 1);   /* OK */
int is_same = (flags.b1 == flags.b7);   /* OK */
int is_all = (flags == 0XFF);   /* 컴파일 오류 */
int is_zero = (flags == 0);     /* 컴파일 오류 */

```

- 그치만 포인터를 이용해 한 번에 비교할 수 있다. 그치만 권장하진 않는다.

```c
char* val;
int is_zero;
bitflags_t flags = { 0, };

flags.b3 = 1;

val = (char*)&flags;
is_zero = = (*val == 0);

```

## 공용체

- 공용체(union)란?
  - 똑같은 메모리 위치를 다른 변수로 접근하는 방법이다.
    - 어떤 경우에는 그 4바이트 메모리를 int로 읽고 싶고 어떤 경우에는 float로 읽고 싶은 경우.
- 코드를 통해 보자.
  - bigflags_t의 val과 bits는 서로 같은 메모리에 접근한다.

```c
typedef union
{
    unsigned char val;
    struct
    {
        unsigned char b0 : 1;
        unsigned char b1 : 1;
        unsigned char b2 : 1;
        unsigned char b3 : 1;
        unsigned char b4 : 1;
        unsigned char b5 : 1;
        unsigned char b6 : 1;
        unsigned char b7 : 1;
    } bits;
} bitflags_t;

/* ... */

int is_same;
int is_zero;
bitflags_t flags = { 0, };

flags.bits.b1 = 1;
flags.bits.b4 = 1;

is_same = (flags.bits.b1 == flags.bits.b7);
is_zero = (flags.val == 0);

```

## 메모리 공유만을 위한 공용체의 예

```c
typedef union
{
    int ivalue;
    double dvalue;
} value_t;

```

## 함수 포인터, 함수를 변수에 저장할 수 있을까?

- 함수를 호출하는 방법을 떠올려보자.
  - 어떤 함수를 호출할 때는 그 함수의 함수 명을 쓴다.
    - `result = sub(op1, op2);`
  - 그러나 어셈블리어로는 그 함수의 시작 주소로 점프한다. 즉 함수를 실행한다는 것 자체가 어떤 메모리 주소로 점프해서 코드를 실행하는 것이다. 음...그치만 이건 컴파일 타임에 결정되는 것이다. 함수 포인터를 이런 식으로 쓰기엔 무리가 있다.
  - 실행 도중에 다음에 실행될 코드 위치를 찾아가는 경우가 있었는데...?
    - 함수에서 반환할 때 돌아가야 하는 호출자의 코드 주소가 그렇다.
    - 돌아갈 주소는 스택 메모리에 저장되어 있다. 실행 도중에 결정된다는 것이다.
  - 즉 실행 도중 조건에 따라 어떤 함수를 실행해 주려면 그 함수의 시작 주소를 알고 있으면 된다.

## 함수를 매개변수로 전달할 때 필요한 것들, 함수 포인터 선언

- 함수 코드의 시작 주소 메모리 주소를 넣어주고 그걸 실행해달라고 할 때 쓰는 게 함수 포인터이다.

```c
result = calculate(op1, op2, operator); /* operator = add() */
result = calculate(op1, op2, &operator); /* operator = add() */

```

- 둘 다 가능한데, 위의 방법을 더 많이 쓴다.

- 함수를 호출하는 함수 선언하기. 그런데...데이터 형은 어떻게 될까?
  - 함수의 매개변수로 전달되는 함수는 다음과 같은 내용을 담고 있어야 한다.
    - 자기 자신이 받아야 하는 매개 변수 목록
    - 자기 자신이 반환하는 자료형

- 올바른 함수 포인터를 선언하고 사용하는 방법

```c
/* 함수 포인터에 대입하고 싶은 함수 */
double add(double x, double y)
{
    return x + y;
}

/* 함수 포인터 변수의 선언과 사용 */
/* func는 변수 이름. func는 포인터를 갖고 있는데, 함수의 주소를 갖고 있다. 그 함수는 (double, double)를 받고 double를 반환한다. */
double (*func)(double, double) = add; /* 함수 포인터에 함수 대입 */
result = func(op1, op2);

/* 함수 포인터 매개변수의 선언과 사용 */
double calculate(double, double, double (*)(double, double));

double calculate(double x, double y, double (*func)double, double))
{
    return func(x, y);
}

/* 사용 */
result = calculate(op1, op2, add);

```

- 함수 포인터 선언
  - `<반환형> (*<변수명>)(<매개 변수 목록>);`
  - 함수의 매개변수 목록과 반환형은 반드시 표기해야 한다.
- 함수 포인터 읽기 (Right-Left Rule)
  - `double (*func)(double, double)
    - 변수 func는 포인터이다. 뭐를 가리키냐면 매개변수 받는 무언가를 가리킨다. (매개변수를 받는 건 함수니까) 무언가는 두 개의 double 형 매개변수를 받는 함수를 가리키는 포인터이다. 그리고 그 함수는 double를 반환한다.
  - `void (*func)(int)`
    - 변수 func는 한 개의 int 형 매개변수를 받고 반환 값이 없는 함수의 포인터이다.
- 함수 포인터 배열
  - `int (*ops[])(int, int) = { add, sub, mul, div };`
    - ops는 배열인데,
      - 배열의 각 요소는 함수 포인터이다. 그 함수 포인터는 두 개의 int형 매개변수를 받고 int형을 반환하는 함수를 가리킨다.

## 함수 포인터 쉽게 읽기

- 복잡한 함수 포인터 읽어보기
  - `void (*bsd_signal(int, void(*)(int)))(int);`
    - bsd_signal 함수의 선언(함수 포인터 아님)
      - 매개 변수
        - int 형
        - 함수 포인터: void (*)(int)
      - 반환형
        - 함수 포인터: void (*)(int)

## 배열의 포인터, 퀵 정렬, void 포인터

### 배열의 포인터

- 배열 전체를 다 가리키는 포인터도 있다. 배열의 각 요소를 포인터로 접근하는 것과는 다른 것이다.

```c
int score[3] = { 80, 90, 100 };
int (*p)[3] = &scores;

```

- `<자료형> (*<변수이름>)[<요소의 수>];`
  - 사실 쓰이는 곳이 많지 않다.
  - 단, 2차원 배열을 매개변수로 받을 때 쓸 수 있다.

## 함수 포인터 예: 퀵 정렬

- `void qsort(void *ptr, size_t count, size_t size, int (*comp)(const void*, const void*));`
  - <stdlib.h> 안에 있다.
  - 데이터에 따라 정렬하는 기준이 다를 수 있다. 그 기준에 맞는 정렬 함수가 세 번째로 넘겨주는 함수 포인터가 가리키는 함수이다.
    - `int (*comp)(const void*, const void*)`

## void*, void 포인터

- 범용적인 포인터.
- 어떤 포인터라도 여기에 대입이 가능하다.
  - 매개변수 형으로 void*를 사용하면 어떤 포인터도 받을 수 있는 함수가 된다.
- 단, 다음과 같은 경우 다른 포인터로 캐스팅 또는 대입해서 써야 한다.
  - 역 참조 (몇 바이트를 읽을지 모르기 때문)
  - 포인터 산술 연산 (몇 바이트 이동할지 모르기 때문)

- void* 예

```c
/* 예 1 */

float pi = 3.14f;
void* p;
float* q;

p = &pi;
q = p;

printf("%f\n", *p); /* 컴파일 오류 발생 */
printf("%f\n", *(float*)p);

/* 예 2 */
int add(void* op1, void* op2)
{
    int result;
    result = *op1 + *op2; /* 컴파일 오류 */
    result = *(int*)op1 + *(int*)op2; /* OK */
}

```
