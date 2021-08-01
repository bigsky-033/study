package bigskypark.grpcspringexamples.lognet.reactive.simple;

import com.github.benmanes.caffeine.cache.AsyncLoadingCache;
import com.github.benmanes.caffeine.cache.Caffeine;
import com.google.common.base.Preconditions;
import java.time.Duration;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

@Service
public class SimpleStorage {

  private final AsyncLoadingCache<String, String> cache;

  public SimpleStorage(final SimpleAsyncCacheLoader simpleAsyncCacheLoader) {
    Preconditions.checkNotNull(simpleAsyncCacheLoader);
    cache =
        Caffeine.newBuilder()
            .maximumSize(100L)
            .expireAfterWrite(Duration.ofSeconds(60L))
            .refreshAfterWrite(Duration.ofSeconds(10L))
            .recordStats()
            .buildAsync(simpleAsyncCacheLoader);
  }

  public Mono<String> getOrDefault(final String key, final String defaultValue) {
    return Mono.fromFuture(cache.get(key)).switchIfEmpty(Mono.fromSupplier(() -> defaultValue));
  }
}
