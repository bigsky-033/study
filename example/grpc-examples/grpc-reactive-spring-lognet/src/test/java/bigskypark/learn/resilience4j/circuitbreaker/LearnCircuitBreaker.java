package bigskypark.learn.resilience4j.circuitbreaker;

import io.github.resilience4j.circuitbreaker.CallNotPermittedException;
import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.circuitbreaker.CircuitBreakerRegistry;
import io.github.resilience4j.reactor.circuitbreaker.operator.CircuitBreakerOperator;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import reactor.core.publisher.Mono;
import reactor.test.StepVerifier;

public class LearnCircuitBreaker {

  private CircuitBreaker circuitBreaker;

  @BeforeEach
  void init() {
    final var circuitBreakerRegistry = CircuitBreakerRegistry.ofDefaults();
    circuitBreaker = circuitBreakerRegistry.circuitBreaker("test");
  }

  @Test
  void testForClosedStatus() {
    final var mono =
        Mono.fromSupplier(() -> "hello")
            .map(s -> s + " world")
            .transformDeferred(CircuitBreakerOperator.of(circuitBreaker))
            .log();
    circuitBreaker.transitionToClosedState();

    StepVerifier.create(mono).expectSubscription().expectNext("hello world").verifyComplete();
  }

  @Test
  void testForOpenStatus() {
    final var mono =
        Mono.fromSupplier(() -> "hello")
            .map(s -> s + " world")
            .transformDeferred(CircuitBreakerOperator.of(circuitBreaker))
            .log();
    circuitBreaker.transitionToOpenState();

    StepVerifier.create(mono)
        .expectSubscription()
        .expectError(CallNotPermittedException.class)
        .verify();
  }

  @Test
  void testForDisabledStatus() {
    final var mono =
        Mono.fromSupplier(() -> "hello")
            .map(s -> s + " world")
            .transformDeferred(CircuitBreakerOperator.of(circuitBreaker))
            .log();
    circuitBreaker.transitionToDisabledState();

    StepVerifier.create(mono).expectSubscription().expectNext("hello world").verifyComplete();
  }
}
