# C언어 기본 문법 2, 빌드 단계

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 포인터 (pointer)

- 하드웨어 처럼 메모리에 있는 변수에 직접 접근할 수 있을까? C에서는 포인터를 통해 가능하다.
- 이는 강력한 기능이기도 하지만 또한 위험한 기능이기도 하다. 따라서 잘 다룰 수 있도록 하자.

## 주소 연산자 &

- 지역 변수의 주소 출력하기

    ```c
    #include <stdio.h>

    void print_address(void)
    {
        int num = 10;
        printf("Address of num: %p\n", (void*)&num);
    }

    int main(void)
    {
        print_address();

        return 0;
    }

    ```

- 주소 연산자 &
  - 비트 연산자 &가 아니다. 비트 연산자는 언제나 피연산자가 두 개이다. 주소 연산자는 피연산자가 하나이다.
  - num이란 변수가 있으면 &num은 그 변수가 위치한 메모리 주소이다.
  - 보통 주소를 보여줄 때는 16진수를 사용한다. 읽기 편하기 때문이다.
  - 참고: 실행할 때마다 주소가 달라질 수 있다. 요즘 운영체제에서는 보안 강화를 위해 실행할 때마다 주소를 바꿔준다. 이를 ASLR 이라 부른다.

## 메모리 주소 저장하기

- 메모리 주소 자체를 어딘가에 저장해두면 더 좋지 않을까?
  - 어떤 데이터를 저장할 때 쓰는 것은? -> 변수
  - 메모리 주소도 변수에 저장할 수 있지 않을까?

- 주소: 값이 저장되어 있는 메모리 상의 위치 라는 표현이다.
- 값: 실제 산술 연산 등을 할 때 사용되는 것.

- 그냥 int 타입 같은 것에 주소를 저장하려고 하니 헷갈린다. 따라서 주소를 저장하기 위한 특별한 변수가 필요하다. 그 특별한 변수가 바로 `포인터`이다.

## 포인터의 의미

- 포인터: 주소를 저장하기 위한 변수형이다. 즉, 변수인데 속에 담긴 내용은 메모리 주소이다.

## 메모리 주소에 저장된 자료형

- 하드웨어는 자료형에 전혀 신경을 쓰지 않는다.
- 그러나 데이터를 읽을 때 해당 주소에서부터 몇 바이트를 읽어야 하는지는 하드웨어에게 알려 줄 필요가 있다.
- 그래서 포인터 변수를 선언할 때 포인터 앞에 자료형을 붙인다.
  - 어떤 메모리 주소가 있는데, 그 메모리 주소에 가서 선언한 자료형으로 자료를 읽어야 한다고 알려주는 것이다.

## 포인터 변수를 선언하는 방법

```c
void save_address(void)
{
    int num = 10;
    int* num_address = &num;
}
```

- 포인터 변수를 선언하려면 자료형 뒤에 별표(*)를 붙인다.
  - `int*`, `char*`, `float*`
- 별표(*) 위치가 어디에 있느냐는 언어에서 강제하지 않는다. `int *address;` 처럼 쓰는 경우도 있다.

## 포인터 변수를 부르는 방법

- 보통 `int* num_address` 에서 num_address를 int 포인터라고 부른다.
- 영어로는 int로의 포인터(pointer to an int)라고도 한다.

## 포인터에 저장된 주소도 바꿀 수 있나요?

- 당연히 바꿀 수 있다. 포인터도 변수이기 때문이다. 즉, 다른 주소로 바꿀 수 있다.

## 역 참조 연산자 *

- 포인터도 변수이다. 따라서 당연히 변수를 쓰는 곳에는 다 쓸 수 있다. 매개 변수로도 쓸 수 있다.

    ```c
    void print_address(int* num)
    {
        /* 주소 출력 */
        printf("address of num: %p\n", (void*)num);
    }

    /* 메인 함수 */
    int score = 88;
    print_address(&score);

    ```

    ```c
    /* 포인터에 저장된 값을 참조하는 예 */

    void print_value(void)
    {
        int score = 100;
        int* pointer = &score;

        printf("score: %d\n", *pointer);
    }

    void print_argument(float* arg)
    {
        printf("argument: %f\n", *arg);
    }

    ```

- 역참조 연산자 `*`
  - 곱하기 연산자가 아니다. 역 참조 연산자는 피연산자 1개를 가진다(곱하기 연산자는 2개).
  - 포인터 변수 앞에 `*`를 붙이면, 그 주소에 가서 값을 가져오라는 의미이다.

## 참조와 역 참조

- 역 참조라는 게 있으면...참조라는 게 있나?
- 참조라고 하는 것은 포인터가 이미 하고 있는 일을 말한다. 즉 어떤 변수의 값을 직접 가져다 쓰는 게 아니라 그게 어디에 있다고 '참조'하는 것이다. 즉, 값잉 어디에 있는지 가리키고 있는 것이다.
- 역 참조는 주소로 직접 가서 거기 저장되어 있는 값에 접근하는 것이다.
- 실제 데이터에 간접(indirect)적으로 접근한다는 의미이기도 하다. 주소를 이용해 접근한다.

## 역 참조를 이용한 값 변경 예

```c
void update_value(void)
{
    int score = 100;
    int* pointer = &score;

    printf("score: %d\n", *pointer);
    *pointer = 50;
    printf("updated score: %d\n", *pointer);
}

void update_argument(float* arg)
{
    printf("argument: %f\n", *arg);
    *arg = 93485.2f;
    printf("updated argument: %f\n", *arg);
}

```

- 포인터 변수 선언과 역참조를 헷갈리지 말자.
  - 이 헷갈림 때문에, 포인터 변수를 선언할 때 `int *v` 보다 `int* v`가 낫다.

## 포인터로 두 변수의 값 바꾸기

- C에 ref는 없지만 포인터가 있다!

```c
void swap(int* arg1, int* arg2)
{
    int tmp;

    tmp = *arg1;
    *arg1 = *arg2;
    *arg2 = tmp;
}

```

## 값에 의한 전달 vs 참조에 의한 전달

- 엄밀히 말하면 C는 값에 의한 전달이다. 함수를 호출할 때 언제나 변수를 복사한ㄷ. 단, 포인터를 사용해서 참조에 의한 전달을 흉내낼 뿐이다.
- 그런데 뭐...이런게 별로 중요하진 않다. 용어에 너무 의미를 두지 말자. 원본이 바뀌는지 안 바뀌는지가 중요하다. C에서는 포인터를 이용해 원본을 바꿔줄 수 있다.

## 포인터와 함수 반환 값

- 당연히 포인터도 변수니까 함수 반환 값으로 사용이 가능하다.
  - `int* do_something(const int op1, const int op2);`
- 다만 포인터를 반환할 때 조심해야 할 것이 있다.
- 매우 위험한 코드의 예

  ```c
  int* add(const int op1, const int op2)
  {
      int result = op1 + op2;
      return &result;
  }

  int main(void)
  {
      int* result;

      result = add(10, 20);

      return 0;
  }

  ```

  - 왜 위험할까?
    - 함수의 지역 변수는 스택에 저장된다. 함수의 호출이 끝나면 지역 변수도 사라진다. 위의 경우에 &result를 반환하고 함수를 나가면 스택 프레임 자체가 유효하지 않다. 그 값이 덮어쓰여 지기 전까지 값은 유효할 수 있으나 값의 유효함이 보장되지 않는다.

## 댕글링 포인터(dangling pointer)

(위의 경우와 이어지는 내용)

- 지역 변수가 사용한 '주소'자체가 사라지는 것은 아니다. 따라서 컴파일 오류가 나진 않는다. 문제는 포인터가 유효하지 않은 주소를 가리킨다는 것이다. 이러한 포인터를 `댕글링 포인터`라고 한다.
- 포인터를 반환할 경우 댕글링 포인터를 조심해야 한다.
  - 포인터를 반환해도 되는 경우
    - 전역 변수
    - 파일 속 static 전역 변수
    - 함수 내 static 변수
    - 힙 메모리에 생성한 데이터

    ```c
    int* spawn_monster(void)
    {
        static int s_monster_count = 0;

        /* some codes... */

        ++s_monster_count;

        return &s_monster_count;
    }

    ```

- 언제 포인터를 반환할까?
  - 도우미 함수 안에 생성한 변수를 다음 함수에서 사용하고자 할 때. 단, 일반 지역 변수면 안 된다.
  - 함수 안에서 대용량의 데이터를 생성하고 그걸 반환하고자 할 때.

## 널(NULL) 포인터

- 반환할 주소가 없는 경우에는 어떻게 해야할까?

```c
void do_something()
{
    int number;
    int* num_ptr = &number;

    /* some codes... */

    num_ptr = NULL;
}


```

- 아무것도 가리키지 않는 포인터를 널 포인터라고 한다.
- 값이 0인 상수 표현식이다.
- 또는 `(void*)`로 캐스팅된 표현식
- 전용 매크로가 있다.
  - `#define NULL ((void*)0)`
  - 널 포인터를 표현할 때 이 매크로를 사용할 것
- 포인터 변수와 NULL은 비교가 가능하다.

  ```c
  int* ptr;

  if (ptr == NULL)
  {
      /* some codes... */
  }

  if (ptr != NULL)
  {
      /* some codes... */
  }

  ```

- 0을 NULL로 사용하지 말자. NULL을 명시적으로 쓰자.

## NULL이 가지는 문제들

- 매개변수
  - 함수 매개 변수로 포인터가 들어올 때는 언제나 골칫 덩어리이다. 누구나 NULL을 넣을 수 있기 때문이다. 따라서 기본적으로 NULL이 안 들어온다고 가정하고 함수를 작성할 것. NULL이 들어올 수 있는 함수는 매개변수명에서 분명히 밝힐 것.
  - 널 포인터를 허용하는 매개 변수: 매개 변수 이름 끝에 '_or_null' 을 붙인다.

    ```c
    int get_score(const char* student_id_or_null)
    {
        /* some codes... */
    }

    ```

  - NULL이 안 들어온다고 가정한 경우 assert를 사용해 검증 한다.

    ```c
    #include <assert.h>

    #define PRICE (2)

    void increase_proce(int* current_price)
    {
        assert(current_price != NULL);

        *current_proce += PRICE;
    }

    ```

- 반환 값
  - NULL을 반환할 때도 골칫거리이다. 기본적으론 하지 않는다.
  - 반환을 해야 한다면 함수 이름에 NULL을 반환한다는 것을 명시하자.

  ```c
  const const* get_name_or_null(const int id)
  {
      /* 코드 생략 */ 
      return NULL;
  }

  ```

## 널 포인터는 언제 사용하나요?

- 포인터 변수를 초기화하고 싶을 때
  - 아직 참조할 주소가 없을 때.
    - `int* ptr = NULL;`
- 포인터 변수가 유효한 주소를 참조하고 있는지 확인하고 싶을 때
  - 아무것도 가리키지 않는 포인터 변수를 역참조하면? 결과가 정의되지 않음. 어떻게 될지 모른다.
    - `if (ptr != NULL) { ...`
- 댕글링 포인터를 막기 위해서
  - 동적 메모리 할당된 메모리를 더 이상 필요 없어서 해제했는데, 이를 여전히 가리키는 포인터가 있다면? 더 이상 사용할 수 없는 데이터이니 포인터 변수에 저장되어 있는 그 주소를 초기화해야 한다. 이 때 널 포인터를 이용해 리셋한다.

    ```c
    /* 동적 메모리 할당 */
    int* ptr = (int*)malloc(sizeof(int));

    /* 코드... */

    /* 더이상 ptr을 사용하지 않음 */
    free(ptr);
    ptr = NULL;

    ```

- 중요한 것! 존재하지 않는 메모리에서 값을 읽어오려고 하면 문제가 펑펑 터진다.

## 포인터의 비교

- 포인터 변수 자체를 비교하면 주소를 하는 거고 역참조를 통해 비교하면 값을 비교하는 것이다.
- 근데...NULL외의 주소를 왜 비교할까?
  - 변수 하나가 아니라 큰 메모리를 통째로 잡아두고 그 안에 복수의 데이터를 넣어 사용할 때 필요하다.
  - 배열!

## 포인터의 크기

- 모든 포인터는 동일한 크기를 가진다.
- 포인터의 크기는 코드를 컴파일하는 시스템 아키텍쳐에 따라 결정된다.
  - 보통 CPU가 한 번에 처리할 수 있는 데이터의 크기(=워드, word)와 동일하다.
- 포인터의 크기 예

```c
void print_pointer_size()
{
    char ch = 'c';
    int number = 934563;
    float pi = 3.1415f;

    char *char_ptr = &ch;
    int *int_ptr = &number;
    float *float_ptr = &pi;

    printf("char size: %d, char* size: %d\n", sizeof(*char_ptr), sizeof(char_ptr));
    printf("int size: %d, int* size: %d\n", sizeof(*int_ptr), sizeof(int_ptr));
    printf("float size: %d, float* size: %d\n", sizeof(*float_ptr), sizeof(float_ptr));
}

```

## 배열 포인터에 대입하기

- 가능하다.

## 배열 속 각 요소의 위치, 각 요소의 위치 계산하기

- 배열에서 각 요소 사이의 바이트 간격은 일정하다.
- 각 요소의 위치를 어떻게 계산할 수 있을까?
  - 첫 번째 요소의 주소와 자료형의 크기만 안나면 n번째 요소의 위치를 알 수 있다.
    - n 번째 요소 위치 = 첫 번째 요소 주소 + 자료형의 크기 * (n-1)
- 다음 요소 위치를 구할 때 주의할 점
  - 포인터에 정수 1을 더하는 것과 1바이트를 더하는 것을 헷갈리지 말자. 1바이트를 더하는 게 아니다.
    - 예를 들어 `int* ptr = nums; ptr = ptr + 3;`을 하면, int의 사이즈를 세 번 더한 것이다.

## 배열 요소에 포인터로 접근하기

- 다음과 같은 것들은 다 같은 것이다.
  - 재밌는 것은 배열의 첨자 연산자(`[]`)도 포인터에 쓸 수 있다.
  - `nums[1] == ptr[1] == *(ptr + 1)`
  - 컴파일러에게 모두 같은 의미이다.

```c
int nums[3] = { 10, 20, 30 };
int* ptr = nums;

printf("%d, %d, %d\n", nums[1], ptr[1], *(ptr + 1));

```

- 배열의 모든 요소 더하기 예제

```c
int sum(int* data, const size_t length)
{
    int result = 0;
    size_t i;
    for (i = 0; i < length; i++)
    {
        /* 방법 1 */
        result += data[i];
        /* 방법 2 */
        result += *(data + i);
    }
    return result;
}

```

## 포인터의 중간 정리, 포인터의 캐스팅

- C에서는 주소를 얻을 수 있는 방법이 딱 두가지이다.
  - 주소 연산자(&)
  - 배열의 이름: 배열의 시작 주소를 알려준다.
- 포인터에 정수 n을 더하거나 빼면 언제나 `(sizeof) * n`한 만큼 메모리 주소를 이동한다.
- 정말 딱 한 바이트만 옮기고 싶다면?
  - 한 바이트짜리 포인터로 캐스팅: `int_ptr = (char*)int_ptr + 1`;
- `int*` -> `char*` 캐스팅은 무엇을 캐스팅 하는 걸까?

## 딱 '한' 바이트만 옮기기

## 두 주소 간의 사칙 연산

## 자바와 C#에서는 모든 것이 포인터다

## 포인터를 사용한 안전하지 않은 코드

## 포인터와 배열의 차이

## 다시 만나는 연산자 결합 법칙

## 포인터와 연산자 우선순위 및 결합 법칙

## 동일한 우선순위를 갖는 연산자들

## 조금 더 빠른 배열의 요소 더하기 함수

## 왜 *p++이 더 빠르죠?

## 포인터와 const

## 주소를 보호하는 const 포인터

## 값을 보호하는 const를 가리키는 포인터

## 두 const의 정리와 예

## 주소와 값 모두 지키는 const

## const 포인터 읽는 방법 정리

## const는 절대 제거하지 말자

## 포인터의 용도

## 포인터 배열

## 2차원 포인터 배열
