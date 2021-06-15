package bigskypark.study.examples.armeria.helloworld.server;

import com.linecorp.armeria.common.HttpResponse;
import com.linecorp.armeria.server.Server;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {

  private static final Logger logger = LoggerFactory.getLogger(Main.class);

  public static void main(String[] args) {
    final var httpPort = 8080;
    final var server = newServer(httpPort);
    Runtime.getRuntime()
        .addShutdownHook(
            new Thread(
                () -> {
                  server.stop().join();
                  logger.info("Server has been stopped");
                }));
    server.start().join();
    logger.info("Server has been started on {} port", httpPort);
  }

  static Server newServer(int httpPort) {
    final var sb = Server.builder();
    sb.http(httpPort);
    sb.service("/hello", (ctx, req) -> HttpResponse.of("Hello"));
    return sb.build();
  }
}
