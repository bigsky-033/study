package bigskypark.grpcspringexamples.lognet.helloworld;

import com.google.common.base.Preconditions;
import io.grpc.examples.helloworld.GreeterGrpc;
import io.grpc.examples.helloworld.HelloReply;
import io.grpc.examples.helloworld.HelloRequest;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import org.lognet.springboot.grpc.GRpcService;

import javax.annotation.PostConstruct;

@Slf4j
@GRpcService
public class HelloWorldService extends GreeterGrpc.GreeterImplBase {

  @Override
  public void sayHello(
      final HelloRequest request, final StreamObserver<HelloReply> responseObserver) {
    Preconditions.checkNotNull(request);
    Preconditions.checkNotNull(responseObserver);

    final var reply = HelloReply.newBuilder().setMessage("Hello " + request.getName()).build();
    responseObserver.onNext(reply);
    responseObserver.onCompleted();
  }

  @PostConstruct
  public void post() {
    log.info("Constructed");
  }
}
