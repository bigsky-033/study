package bigskypark.grpcspringexamples.lognet.reactive.config;

import com.google.common.base.Preconditions;
import io.lettuce.core.RedisClient;
import io.lettuce.core.api.reactive.RedisStringReactiveCommands;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RedisConfiguration {

  @Bean
  public RedisClient redisClient() {
    return RedisClient.create("redis://localhost");
  }

  @Bean
  public RedisStringReactiveCommands<String, String> redisReactiveCommands(
      final RedisClient redisClient) {
    Preconditions.checkNotNull(redisClient);
    return redisClient.connect().reactive();
  }
}
