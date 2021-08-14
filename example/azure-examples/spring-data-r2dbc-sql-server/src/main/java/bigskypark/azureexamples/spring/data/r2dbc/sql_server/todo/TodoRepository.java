package bigskypark.azureexamples.spring.data.r2dbc.sql_server.todo;

import org.springframework.data.repository.reactive.ReactiveCrudRepository;

public interface TodoRepository extends ReactiveCrudRepository<Todo, Long> {}
