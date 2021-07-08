package bigskypark.grpcspringexamples.lognet.reactive.simple;

import com.github.benmanes.caffeine.cache.AsyncCacheLoader;
import com.google.common.base.Preconditions;
import io.lettuce.core.api.reactive.RedisStringReactiveCommands;
import lombok.extern.slf4j.Slf4j;
import org.checkerframework.checker.nullness.qual.NonNull;
import org.springframework.stereotype.Component;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Executor;

@Slf4j
@Component
public class SimpleAsyncCacheLoader implements AsyncCacheLoader<String, String> {

  private final RedisStringReactiveCommands<String, String> redisReactiveCommands;

  public SimpleAsyncCacheLoader(
      final RedisStringReactiveCommands<String, String> redisReactiveCommands) {
    this.redisReactiveCommands = Preconditions.checkNotNull(redisReactiveCommands);
  }

  @Override
  public @NonNull CompletableFuture<String> asyncLoad(
      @NonNull String key, @NonNull Executor executor) {
    if ("error".equalsIgnoreCase(key)) {
      throw new IllegalArgumentException("Key name is invalid. Input: " + key);
    }

    log.info("Cache load happens for key: {}", key);
    return redisReactiveCommands.get(key).toFuture();
  }
}
