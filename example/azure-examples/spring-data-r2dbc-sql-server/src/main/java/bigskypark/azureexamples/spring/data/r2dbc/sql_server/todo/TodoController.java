package bigskypark.azureexamples.spring.data.r2dbc.sql_server.todo;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/todo")
@RequiredArgsConstructor
public class TodoController {

  private final TodoRepository todoRepository;

  @PostMapping
  @ResponseStatus(HttpStatus.CREATED)
  public Mono<Todo> create(@RequestBody Todo todo) {
    return todoRepository.save(todo);
  }

  @GetMapping
  public Flux<Todo> listTodos() {
    return todoRepository.findAll();
  }
}
