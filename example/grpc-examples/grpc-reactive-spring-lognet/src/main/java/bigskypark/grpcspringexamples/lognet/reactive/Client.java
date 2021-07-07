package bigskypark.grpcspringexamples.lognet.reactive;

import io.grpc.ManagedChannelBuilder;
import io.grpc.examples.helloworld.HelloRequest;
import io.grpc.examples.helloworld.ReactorGreeterGrpc;

public class Client {
  public static void main(String[] args) {
    final var channel = ManagedChannelBuilder.forTarget("localhost:6565").usePlaintext().build();
    ReactorGreeterGrpc.ReactorGreeterStub stub = ReactorGreeterGrpc.newReactorStub(channel);

    stub.sayHello(HelloRequest.newBuilder().setName("bigsky-park").build())
        .doOnNext(reply -> System.out.println(reply.getMessage()))
        .block();
  }
}
