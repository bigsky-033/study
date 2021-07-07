package bigskypark.grpcspringexamples.lognet.reactive.simple;

import com.google.common.base.Preconditions;
import io.lettuce.core.api.reactive.RedisStringReactiveCommands;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

@Service
public class SimpleStorage {

  private final RedisStringReactiveCommands<String, String> command;

  public SimpleStorage(
      final RedisStringReactiveCommands<String, String> redisStringReactiveCommands) {
    this.command = Preconditions.checkNotNull(redisStringReactiveCommands);
  }

  public Mono<String> getBy(final String key) {
    Preconditions.checkNotNull(key);
    return command.get(key);
  }

  public Mono<String> set(final String key, final String value) {
    Preconditions.checkNotNull(key);
    Preconditions.checkNotNull(value);
    return command.set(key, value);
  }
}
