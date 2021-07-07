package bigskypark.grpcspringexamples.lognet.reactive.helloworld;

import com.google.common.base.Preconditions;
import io.grpc.examples.helloworld.HelloReply;
import io.grpc.examples.helloworld.HelloRequest;
import io.grpc.examples.helloworld.ReactorGreeterGrpc;
import lombok.extern.slf4j.Slf4j;
import org.lognet.springboot.grpc.GRpcService;
import reactor.core.publisher.Mono;

import javax.annotation.PostConstruct;

@Slf4j
@GRpcService
public class HelloWorldService extends ReactorGreeterGrpc.GreeterImplBase {

  @Override
  public Mono<HelloReply> sayHello(final Mono<HelloRequest> request) {
    Preconditions.checkNotNull(request);

    return request
        .map(HelloRequest::getName)
        .map(name -> "Hello " + name)
        .map(message -> HelloReply.newBuilder().setMessage(message).build());
  }

  @PostConstruct
  public void post() {
    log.info("Constructed");
  }
}
