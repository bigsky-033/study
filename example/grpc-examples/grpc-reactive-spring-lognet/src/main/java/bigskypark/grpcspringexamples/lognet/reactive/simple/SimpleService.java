package bigskypark.grpcspringexamples.lognet.reactive.simple;

import bigskypark.grpc.examples.simple.GetRequest;
import bigskypark.grpc.examples.simple.GetResponse;
import bigskypark.grpc.examples.simple.ReactorSimpleGrpc;
import com.google.common.base.Preconditions;
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
        .flatMap(r -> simpleStorage.getOrDefault(r.getKey(), "empty"))
        .map(v -> GetResponse.newBuilder().setValue(v).build())
        .onErrorReturn(GetResponse.newBuilder().setValue("error").build())
        .log();
  }
}
