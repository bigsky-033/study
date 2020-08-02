# Concurrent Programming in Java - Week3

Notes for `Concurrent Programming in Java` class in Coursera.

## 3.1 Actors

- How actor model avoids concurrency errors
  - Forcing all accesses to an object to be isolated by *default*.
  - Object is part of the *local state* of an actor, and cannot be accessed directly by any other actor.
- Actor consists of a *Mailbox*, a set of *Mehtods*, and *Local State*.
- Actor model is *reactive*, in that actors can only execute methods in response to messages.
  - This methods can read/write locals tate and/or send messages to other actors.
  - Only way to modify an object in pure actor model is to send messages to the actor that owns that object as part of its local state.

### Readings for 3.1

- Actor Model, <https://en.wikipedia.org/wiki/Actor_model>
- Akka Actor Library, <https://doc.akka.io/docs/akka/current/typed/guide/index.html?language=java>

## 3.2 Actors Examples

### Readings for 3.2

- Pipeline Parallelism, <https://en.wikipedia.org/wiki/Pipeline_(computing)>

## 3.3 Sieve of Eratosthenes Algorithm

- The power of the Actor Model is reflected in the dynamic nature of this problem, where pieces of the computation (new actors) are created dynamically as needed.

```java
// methods example of Actor's process method
public void process(final Object msg) {
    int candidate = (Integer) msg;

    boolean nonNul = ((candidate & localPrime) != 0);
    if (nonMul) {
        if (nextActor == null) {
            // create & start new actor with candidate as its local prime
        } else {
            nextctor.send(msg);
        }
    }
}

```

### Readings for 3.3

- Sieve of Eratosthenes problem, <https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes>

## 3.4 Producer-Consumer Problem

- Problem
  - How can we satisfy coordinate accesses by multiple producer tasks and multiple consumer tasks to a shared buffer of *unbounded size* without giving up any concurrency?
  - Challenges
    - We cannot assume any a priori knowledge about the rate at which different tasks produce and consume items in the buffer.
    - We can solve this problem by using locks or by using object-based isolation but both approaches will require low-level concurrent programming techniques.
- Actors based solution
  - Think all objects involved in the concurrent program as actors.
    - Producer tasks
    - Consumer tasks
    - Shared buffer
  - Establish communication protocols among the actors.
    - Producer actor can simply send a message to the buffer actor whenever it has an item to producer.
    - Consumer actor
      - Send a message to the buffer actor whenver it is ready to process an item. -> buffer actor receives a message from a producer, it knows which consumers are ready to process items and can forward the produced item to any one of them.

## 3.5 Bounded Buffer Problem

- More realistic version of the producer-consumer problem.
- Producer tasks are not permitted to send items to the buffer when the buffer is full.
- Buffer actor needs to play a *master* role in the protocol by informing producer actors when they are permitted to send data.
  - Producer actor will only send data when requested to do so by the buffer actor.

### Readings for 3.5

- Producer-consumer problem, <https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem>
