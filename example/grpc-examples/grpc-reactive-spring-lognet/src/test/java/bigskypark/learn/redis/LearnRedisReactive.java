package bigskypark.learn.redis;

import io.lettuce.core.RedisClient;
import java.util.Map;
import java.util.concurrent.CompletableFuture;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import reactor.core.publisher.Mono;
import reactor.test.StepVerifier;

public class LearnRedisReactive {

  private RedisClient redisClient;

  @BeforeEach
  void init() {
    redisClient = RedisClient.create("redis://localhost");
  }

  @Test
  void testForHSetAndHGet() {
    final var reactiveCommand = redisClient.connect().reactive();
    final var key = "test_key";

    reactiveCommand.hset(key, "hello", "world").block();
    reactiveCommand.hset(key, "bye", "world").block();

    StepVerifier.create(reactiveCommand.hgetall(key))
        .expectSubscription()
        .expectNext(Map.of("hello", "world", "bye", "world"))
        .verifyComplete();

    reactiveCommand.del(key).block();
  }

  @Test
  void testForEmptyCase() throws Exception {
    final var reactiveCommand = redisClient.connect().reactive();
    final var key = "test_key";

    // reactiveCommand.hset(key, "hello", "world").block();
    // reactiveCommand.hset(key, "bye", "world").block();
     reactiveCommand.del(key).block();

    CompletableFuture<String> future = reactiveCommand
        .hgetall(key)
        .name("test")
        .switchIfEmpty(Mono.just(Map.of("empty_key", "empty_value")))
        .map(
            value -> {
              System.out.println("[DAMIAN] Here");
              return value.toString();
            })
        .toFuture();

    System.out.println("[DAMIAN] Future: " + future.get());

    System.out.println("[DAMIAN] Wow");
    // reactiveCommand.del(key).block();
  }
}
