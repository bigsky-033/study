package bigskypark.grpcspringexamples.lognet.reactive.simple;

import bigskypark.grpc.examples.simple.GetRequest;
import bigskypark.grpc.examples.simple.GetResponse;
import bigskypark.grpc.examples.simple.ReactorSimpleGrpc;
import bigskypark.grpc.examples.simple.SetRequest;
import com.google.common.base.Preconditions;
import com.google.protobuf.Empty;
import lombok.extern.slf4j.Slf4j;
import org.lognet.springboot.grpc.GRpcService;
import reactor.core.publisher.Mono;

@Slf4j
@GRpcService
public class SimpleService extends ReactorSimpleGrpc.SimpleImplBase {

  private final SimpleStorage simpleStorage;

  public SimpleService(final SimpleStorage simpleStorage) {
    this.simpleStorage = Preconditions.checkNotNull(simpleStorage);
  }

  @Override
  public Mono<GetResponse> get(final Mono<GetRequest> request) {
    Preconditions.checkNotNull(request);
    return request
        .flatMap(r -> simpleStorage.getBy(r.getKey()))
        .map(v -> GetResponse.newBuilder().setValue(v).build());
  }

  @Override
  public Mono<Empty> set(final Mono<SetRequest> request) {
    Preconditions.checkNotNull(request);
    return request
        .flatMap(r -> simpleStorage.set(r.getKey(), r.getValue()))
        .flatMap(
            v -> {
              if (!"OK".equals(v)) {
                Mono.error(new IllegalStateException("Set failed"));
              }
              return Mono.just(Empty.newBuilder().build());
            });
  }
}
