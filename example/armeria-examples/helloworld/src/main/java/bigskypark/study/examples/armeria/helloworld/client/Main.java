package bigskypark.study.examples.armeria.helloworld.client;

import com.linecorp.armeria.client.WebClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger logger = LoggerFactory.getLogger(Main.class);

  public static void main(String[] args) {
    final var webClient = WebClient.of("http://localhost:8080");
    final var response = webClient.get("/hello").aggregate().join();
    logger.info("Response: {}", response);
  }
}
