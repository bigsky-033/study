package bigskypark.azureexamples.spring.data.r2dbc.sql_server.todo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Todo {

  @Id private Long id;

  private String description;

  private String details;

  private boolean done;
}
