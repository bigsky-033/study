package bigskypark.grpcspringexamples.lognet.reactive.simple;

import bigskypark.grpc.examples.simple.GetRequest;
import bigskypark.grpc.examples.simple.GetResponse;
import bigskypark.grpc.examples.simple.ReactorSimpleGrpc;
import com.google.common.base.Preconditions;
import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.reactor.circuitbreaker.operator.CircuitBreakerOperator;
import org.lognet.springboot.grpc.GRpcService;
import reactor.core.publisher.Mono;

@GRpcService
public class SimpleService extends ReactorSimpleGrpc.SimpleImplBase {

  private final SimpleStorage simpleStorage;
  private final CircuitBreaker circuitBreaker;

  public SimpleService(
      final SimpleStorage simpleStorage, final CircuitBreaker simpleServiceCircuitBreaker) {
    this.simpleStorage = Preconditions.checkNotNull(simpleStorage);
    this.circuitBreaker = Preconditions.checkNotNull(simpleServiceCircuitBreaker);
  }

  @Override
  public Mono<GetResponse> get(final Mono<GetRequest> request) {
    Preconditions.checkNotNull(request);
    return request
        .flatMap(r -> simpleStorage.getOrDefault(r.getKey(), "empty"))
        .transformDeferred(CircuitBreakerOperator.of(circuitBreaker))
        .map(v -> GetResponse.newBuilder().setValue(v).build())
        .onErrorReturn(GetResponse.newBuilder().setValue("error").build())
        .log(); // add logging because this application is just example application
  }
}
