package bigskypark.grpcspringexamples.lognet.reactive.admin;

import com.google.common.base.Preconditions;
import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.circuitbreaker.CircuitBreakerRegistry;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@Slf4j
@RequestMapping("/admin")
@RestController
public class CircuitBreakerManageController {

  private final CircuitBreakerRegistry circuitBreakerRegistry;

  public CircuitBreakerManageController(final CircuitBreakerRegistry circuitBreakerRegistry) {
    this.circuitBreakerRegistry = Preconditions.checkNotNull(circuitBreakerRegistry);
  }

  @PostMapping("/circuitbreaker/{name}/status/{state}")
  Mono<Void> transit(
      @PathVariable final String name, @PathVariable final CircuitBreaker.State state) {
    for (final var circuitBreaker : circuitBreakerRegistry.getAllCircuitBreakers()) {
      if (name.equals(circuitBreaker.getName())) {
        switch (state) {
          case OPEN:
            circuitBreaker.transitionToOpenState();
            break;
          case CLOSED:
            circuitBreaker.transitionToClosedState();
            break;
          case HALF_OPEN:
            circuitBreaker.transitionToHalfOpenState();
            break;
          case FORCED_OPEN:
            circuitBreaker.transitionToForcedOpenState();
            break;
          case DISABLED:
            circuitBreaker.transitionToDisabledState();
            break;
          case METRICS_ONLY:
            circuitBreaker.transitionToMetricsOnlyState();
            break;
        }
      }
    }
    return Mono.empty();
  }
}
