# 자료구조

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 자료구조 기초

- 자료구조란?
  - 컴퓨터에서 여러 자료들을 조직적, 체계적으로 저장하는 방법.
  - 보통 동일한 자료형을 저장하는 구조를 의미한다.
  - 자료구조에 따라 요소들 사이의 관계를 정의하는 규칙이 있다.
- 자료구조의 효율성
  - 효율성은 주로 시간 복잡도(time complexity)를 말한다.
  - 공간 복잡도(space complexity)를 포함하는 경우도 있지만 최근에는 하드웨어의 발전으로 인해 그렇게 중요하지 않다.

## 배열

- 배열(array)
  - 메모리 한 당어리로 표현 가능한 가장 간단한 자료구조이다.
  - 각 자료는 색인(index)으로 접근한다.
    - 연속된 메모리이니 각 요소의 실제 메모리 상의 위치를 쉽게 찾을 수 있다.
- 배열의 삽입
  - 배열의 제일 뒤에 넣는 경우는 간단히 삽입하고 끝난다.
  - 그러나 그 외의 경우는 삽입하려는 위치의 요소부터 마지막 요소를 모두 뒤로 한 칸씩 민 뒤에 삽입한다.
  - 시간 복잡도는 O(n)이다.
- 배열의 삭제
  - 삭제하는 색인을 기준으로 그 뒤의 값들을 한 칸씩 앞으로 땡겨준다.
  - 시간 복잡도: O(n)
- 배열의 검색
  - 배열 속 요소들을 처음부터 차례대로 방문하며 찾고자 하는 값이 있는지 확인한다.
    - 있으면 해당 색인을 반환하고, 없으면 -1(not found)과 같은 값을 반환한다.
  - 시간 복잡도: O(n)
- 배열의 접근
  - 배열은 색인을 통해 접근하기 때문에 O(1)이다.

## 스택

- 가장 먼저 자료구조에 삽입(push)된 데이터가 제일 마지막에 삭제(pop)된다. 이를 FILO 또는 LIFO 라고한다.
- 스택의 삽입은 보통 push라고 표현한다. 시간 복잡도는 O(1)이다.
- 스택의 제거는 보통 pop라고 표현한다. 시간 복잡도는 O(1)이다.
- 스택의 검색은 제일 위부터 찾을 때까지 뒤져야 한다. 시간 복잡도는 O(n)이다. 보통 push와 pop만 허용하기 때문에 임의의 요소에 접근할 방법이 없다.
  - push와 pop만 있으면 검색은 어떻게 해야할까?
    - 모든 요소를 다 제거했다가 다시 원상복구해야 한다.
- 스택의 용도
  - 일련의 자료들의 순서를 뒤집는데 유용하다.

## 큐

- 스택과 마찬가지로 자료의 삽입과 삭제에 대한 규칙이 있는 자료구조 중 하나이다.
- 가장 먼저 자료구조에 삽입(enqueue)된 데이터가 제일 처음에 삭제(dequeue)된다.
- 선입 선출(First In First Out)이라고도 한다.
- 그냥 배열을 사용해 구현하면 비효율적이다. 내부적으로 배열을 사용하되 원형 버퍼(ring buffer)의 개념을 이용하면 삽입 삭제 모두 O(1)로 구현이 가능하다.
- 큐의 삽입
  - enqueue 라고 한다.
  - 시간 복잡도: O(1)

  ```c
  enum { MAX_NUMS = 8 };
  int s_nums[MAX_NUMS];
  size_t s_front = 0;
  size_t s_back = 0;
  size_t s_num_count = 0;

  /* ... */

  void enqueue(int n)
  {
      assert(s_num_count < MAX_NUMS);

      s_nums[s_back] = n;

      s_back = (s_back + 1) % MAX_NUMS;

      ++s_num_count;
  }

  ```

- 큐의 삭제
  - dequeue 라고 한다.
  - 시간 복잡도: O(1)

  ```c
  enum { MAX_NUMS = 8 };
  int s_nums[MAX_NUMS];
  size_t s_front = 0;
  size_t s_back = 0;
  size_t s_num_count = 0;

  /* ... */

  int is_empty(void)
  {
      return (s_num_count == 0);
  }

  void dequeue(void)
  {
      int ret;

      assert(is_empty() == FALSE);

      ret = s_nums[s_front];

      --s_num_count;
      s_front = (s_front + 1) % MAX_NUMS;

      return ret;
  }

  ```

- 큐의 검색
  - 시간 복잡도: O(n)
  - 제일 처음부터 찾을 때까지 뒤져야 한다.
  - 스택과 마찬가지로 보통 enqueue()와 dequeue()만 허용하므로 임의의 요소에 접근할 방법이 없다.
  - 모든 요소를 다 제거했다가 다시 원상복구해야 한다.

## 연결 리스트: 연결 리스트의 삽입/제거/검색

- 연결 리스트(linked list)
  - 자료들이 메모리에 산재해 있다.
  - 연결 리스트의 각 자료를 노드(node)라고 부른다.
  - 자료형이 어떻게 산재해 있을 수 있을까?
    - 동적 메모리 할당으로 필요에 따라 각 노드를 할당한다.
  - 노드 사이에 선후 관계를 별도로 지정한다.
    - 어떻게? 다음에 오는 노드의 메모리 주소를 기억한다.
    - 어디에? 노드에 있는 포인터 변수에 기억한다.
    - 제일 마지막 노드는 다음에 올 노드가 없으니 널 포인터를 둔다.
- 연결 리스트 개념
  - 삽입
    - 이미 삽입할 위치를 알면 O(1)
  - 삭제
    - 이미 삭제할 위치를 알면 O(1)
  - 검색
    - 제일 처음 노드부터 찾을 때까지 뒤져야 한다.
    - O(n)
    - 색인으로 접근이 불가능하다.

## 연결 리스트: 전체 출력 예 & 노드 메모리 해제 예, 헤드 노드

- 연결 리스트의 전체 출력

  ```c
  typedef struct node 
  {
      int value;
      node_t* next;
  } node_t;

  /* ... */

  void print_node(const note_t* head)
  {
      note_t* p;

      p = head;
      while (p != NULL)
      {
          /* 출력 */
          p = p -> next;
      }
  }

  ```

- 헤드 노드
  - 연결 리스트의 첫 번째 노드를 가리키는 포인터.
  - 처음 시작할 때의 값은 NULL이다.
- 연결 리스트의 노드는 메모리를 동적 할당한다. 따라서 메모리 해제를 반드시 잘 해주어야 한다.

  ```c
  void destroy(note_t* head)
  {
      node_t* p = head;

      while (p != NULL)
      {
          node_t* next = p->next;
          free(p);
          p = next;
      }

  }

  ```

## 연결 리스트: 삽입하기 예

- 삽입하기 코드 예

  ```c
  /* 가장 앞에 추가하는 코드 */
  void insert_front(node_t** phead, int n)
  {
      node_t* new_node;

      new_node = malloc(sizeof(node_t));
      new_node->value = n;

      new_node->next = *phead;
      *phead = new_node;
  }

  ```

## 연결 리스트: 오름차순으로 삽입하기 예

```c
void insert_sorted(node_t** phead, int n)
{
    node_t** pp;
    node_t* new_node;

    new_node = malloc(sizeof(node_t));
    new_node->value = n;

    pp = phead;
    while (*pp != NULL)
    {
        if ((*pp)->value >= n)
        {
            break;
        }
        pp = &(*pp)->next;
    }

    new_node->next = *pp;
    *pp = new_node;
}

```

## 연결 리스트: 노드 삭제

```c
void remove(node_t** phead, int n)
{
    node_t** pp;

    pp = phead;
    while (*pp = NULL)
    {
        if ((*pp)->value == n)
        {
            node_t* tmp = *pp;
            *pp = (*pp)->next;
            free(tmp);
            break;
        }

        pp = &(*pp)->next;
    }
}

```

## 연결 리스트의 용도

- 길이를 자유롭게 늘리거나 줄일 수 있기 때문에 배열의 한계를 넘으려고 사용하던 자료구조이다.
- 즉, 최대 길이를 미리 특정할 수 없고 삽입/삭제가 빈번할 경우에 사용한다.

## 해시 테이블

## 해시 테이블의 크기는 2배 이상인 소수로

## 중복 색인 문제의 해결

- 방법 1) 나머지 연산으로 구한 색인 위치 이후에 빈 공간을 찾아서 저장한다.
  - 즉, 색인을 1씩 증가해가며 빈 색인 위치를 찾아 거기에 저장한다.
  - 이 방법을 사용하려면 배열에서 비어있는 위치를 알 수 있어야 한다. 이걸 하려면 같은 크기의 bool 배열을 만들어서 사용 여부를 나타내야 한다.
  - 또는 특정한 값을 저장해서 비어있다는 사실을 표시한다. 보통 배열에 들어올 수 없는 값을 사용한다.

  ```c
  #include <limits.h>

  #define TRUE (1)
  #define FALSE (0)
  #define BUCKET_SIZE (23)
  int s_numbers[BUCKET_SIZE];

  void init_hashtable(void)
  {
      size_t i;

      for (i = 0; i < BUCKET_SIZE; ++i)
      {
          s_numbers[i] = INT_MIN;
      }
  }

  int has_value(int value)
  {
      int i;
      int start_index;

      start_index = value % BUCKET_SIZE;
      if (start_index < 0)
      {
          start_index += BUCKET_SIZE;
      }

      i = start_index;

      do {
          if (s_numbers[i] == value)
          {
              return TRUE;
          } else if (s_numbers[i] == INT_MIN) {
              return FALSE;
          }

          i = (i + 1) % BUCKET_SIZE;
      } while (i != start_index);

      return FALSE:
  }

  int add(int value)
  {
      int i;
      int start_index;

      start_index = value % BUCKET_SIZE;
      if (start_index < 0)
      {
          start_index += BUCKET_SIZE;
      }

      i = start_index;

      do {
          if (s_numbers[i] == value || s_numbers[i] == INT_MIN)
          {
              s_numbers[i] = value;
              return TRUE:
          }

          i = (i + 1) % BUCKET_SIZE;
      } while (i != start_index);

      return FALSE:
  }

  ```

- 방법 2) 각 배열의 요소를 연결 리스트를 사용해 저장한다.
  - 배열 안의 각 요소에 연결 리스트를 저장한다.

## 해시란 무엇인가

- 해시 값: 어떤 데이터를 해시 함수에 넣어서 나온 결과.
- 해시 함수: 임의의 크기를 가진 데이터를 고정 크기의 값에 대응하게 하는 함수.
  - 함수이기 때문에, 입력 값이 같으면 출력 같은 언제나 같다.
  - 입력 값이 달라도 출력 값이 같을 수는 있다. 해시 함수에서 이런 것을 해시 충돌(hash collision)이라고 한다.
  - 따라서 출력값으로부터 입력 값을 찾을 수 있다는 보장이 없다.

## 해시 테이블에 문자열 저장하기

- 해시 테이블에 문자열을 저장하기 위해 할 일
  - 문자열을 배열 색인으로 변환한다.
  - 문자열을 정수형으로 반환하는 해시 함수가 있어야 한다.
  - 문자열의 각 요소는 어차피 아스키 코드 값(정수 값)이다.

- 해시 충돌이 빈번해서 잘 쓰이지 않는 함수 예

```c
static int hash_function(const char* str)
{
    int code = 0;
    const char* c = str;
    while (*c != '\0')
    {
        code += *c++;
    }

    return code;
}

```

- 하지만 위 해시 함수는 문자열 크기에 영향을 받는다. O(1)이 아니다. O(문자열 길이) 이다.
  - 최적화를 위해 이 결과를 캐싱해두고 이용할 수 있다.

## 해시 맵

- 어떤 키에 대응하는 어떤 값을 같이 저장해 둔 것을 보통 해시 테이블(해시 맵)이라고 한다.
  - 키(key): 데이터의 위치를 의미하는 데이터
  - 값(value): 실제 저장하는 데이터

## 해시 충돌과 훌륭한 해시 함수

- 키 값이 달라도 같은 해시 값이 나올 수 있다. 따라서 출력 값으로부터 입력 값을 알 수 있다는 보장이 없다.
- 이 문제는 해결해야 하는 문제는 아니다. 그냥 색인 충돌 문제의 해법으로 같이 풀리는 문제이다.
- 훌륭한 해시 함수란?
  - [필수] 어떤 경우에도 고정된 크기의 값으로 변환이 가능하다.
    - 어떤 자료형이든, 데이터의 길이에 상관 없이!
  - 해시 충돌이 거의 없다.

## 해시 충돌 방지와 성능 향상

- 간단하고 나쁘지 않은 해시 함수 예: 65599

```c
size_t hash_65599(const char* string, size_t len)
{
    size_t i;
    size_t hash;

    hash = 0;
    for (i = 0; i < len; ++i)
    {
        hash = 65599 * hash + string[i];
    }
    return hash ^ (hash >> 16);
}

```

- 완전히 충돌을 없앨 수 있나?
  - 기술적으로는 안 된다. 그러나 특정 조건 하에서는 방지가 가능하고 그런 조건이 생각보다 많다.
    - 이게 가능하면 키 배열에 문자열 대신 정수만 저장해도 되어서 성능이 빨라진다.
    - 이 조건이 뭘까?
      - 1) 실행 중에 해시 테이블에 저장될 수 있는 데이터를 다 알고 있다.
      - 2) 개발 도중에 해시 충돌이 없다는 걸 확인하고 보장할 수 있다.

- 충돌까지 고려한 해시 맵 예

```c
int add(const char* key, int value, size_t (*hash_func)(const char*, size_t))
{
    size_t i;
    size_t start_index;
    size_t hash_id;

    hash_id = hash_func(key, strlen(key));
    start_index = hash_id % BUCKET_SIZE;
    i = start_index;

    do 
    {
        if (s_keys[i] == NULL)
        {
            /* 새 키-값을 삽입 */
            return TRUE;
        }

        if (strcmp(s_keys[i], key) == 0)
        {
            /* 그대로 두거나 덮어쓸 수 있다. */
            return TRUE;
        }

        i = (i + 1) % BUCKET_SIZE;
    } while(i != start_index);

    return FALSE;
}
```

- 충돌이 없을 때 해시 맵 예

```c
int add_fast(size_t hash_key, const char* value)
{
    size_t i;
    size_t start_index;

    start_index = hash_key % BUCKET_SIZE;
    i = start_index;

    do
    {
        if (s_keys[i] == INT_MIN)
        {
            /* 새 해시-값 삽입 */
            return TRUE;
        }

        if (s_keys[i] == hash_key)
        {
            return TRUE;
        }

        i = (i + 1) % BUCKET_SIZE;
    } while (i != start_index);

    return FALSE;
}
```
