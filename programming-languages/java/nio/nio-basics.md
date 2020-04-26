# Java NIO

## IO vs NIO

| IO                         | NIO          |
| -------------------------- | ------------ |
| Stream oriented processing | Uses buffers |
| Blocking                   | Non Blocking |

- IO is good for high data volume & low simultaneous open file descriptor counts.
- NIO is good for less data volume with high simulataneous open file descriptor counts.

## Buffers

### Definition

- A container for data of a specific primitive type.
- A buffer is a `linear`, `finite` sequence of elements of a specific primitive type.

### Properties

- Capacity: the number of elements it contains. The capacity of a buffer is never negative and never changes.
- Limit: the index of the first element that should not be read or written. A buffer's limit is never negative and is never greater than its capacity.
- Position: index of the next element to be read or written. A buffer's position is never negative and is never grater than its limit.

### Invariants

0 <= mark <= position <= limit <= capacity

- A newly-created buffer always has a position of zero and mark that is undefined.
- The initial limit maybe zero, or it maybe some other value that dependes upon the typ of the buffer and the manner in which it is constructed.
- Each element of a newly-allocated buffer is initialized to zero.

### Basic Operations

- mark(): Sets this buffer's mark at its position.
- clear(): makes a buffer ready for a new sequence of channel-read or relative put operations: It sets the limit to the capacity and the position to zero.
- flip(): makes a buffer ready for a new sequence of channel-write or relative get operations: It sets the limit to the current position and then sets the position to zero.
- rewind(): makes a buffer ready for re-reading the data that it already contains: It leaves the limit unchanged and sets the position to zero
- slice(): creates a subsequence of a buffer: It leaves the limit and the position unchanged.
- duplicate(): creates a shallow of a buffer: It leaves the limit and the position unchanged.

## Channels

- An abstraction for dealing with an open connection to some component that is performing some kind of IO operaiton at a hardware level.
- Channel is a proxy to a component that is responsible for native IO (file or network socket). By acting as a proxy to some native IO component we are able to write and / or read from a Channel.

## Selectors

- A means to work with multiple channels via one abstraction.
- Selecotrs as the name implies, select from multiple SelectableChannel types and notify our program when IO has happened on one of those channels.

## Charsets

- Contains charsets, decoders and encoders for translating between bytes and unicode.
- Charset provides other utility methods for looking up a Charset by name, creating coders (encoder or decoders) and getting the default Charset.

## References

- Java Nio Tutorial for Beginners, https://examples.javacodegeeks.com/core-java/nio/java-nio-tutorial-beginners/
- Buffer, https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/nio/Buffer.html
- Selector, https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/nio/channels/Selector.html
- [NIO] JAVA NIO의 ByteBuffer와 Channel로 File Handling에서 더 좋은 Perfermance 내기!, https://palpit.tistory.com/641
- [NIO] JAVA NIO의 ByteBuffer와 Channel 클래스 사용하기!, http://eincs.com/2009/08/java-nio-bytebuffer-channel/
- 직접 메모리 접근, https://ko.wikipedia.org/wiki/%EC%A7%81%EC%A0%91_%EB%A9%94%EB%AA%A8%EB%A6%AC_%EC%A0%91%EA%B7%BC
