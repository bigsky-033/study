package bigskypark.grpcspringexamples.lognet.reactive.config;

import com.google.common.base.Preconditions;
import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.circuitbreaker.CircuitBreakerRegistry;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class CircuitBreakerConfiguration {

  @Bean
  public CircuitBreakerRegistry circuitBreakerRegistry() {
    return CircuitBreakerRegistry.ofDefaults();
  }

  @Bean
  public CircuitBreaker simpleServiceCircuitBreaker(
      final CircuitBreakerRegistry circuitBreakerRegistry) {
    Preconditions.checkNotNull(circuitBreakerRegistry);
    return circuitBreakerRegistry.circuitBreaker("simpleService");
  }
}
