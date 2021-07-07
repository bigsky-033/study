package bigskypark.grpcspringexamples.lognet;

import io.grpc.ManagedChannelBuilder;
import io.grpc.examples.helloworld.GreeterGrpc;
import io.grpc.examples.helloworld.HelloRequest;

public class Client {
  public static void main(String[] args) {
    final var channel = ManagedChannelBuilder.forTarget("localhost:6565").usePlaintext().build();
    final var blockingStub = GreeterGrpc.newBlockingStub(channel);
    final var reply =
        blockingStub.sayHello(HelloRequest.newBuilder().setName("bigsky-park").build());
    System.out.println(reply.getMessage());
  }
}
