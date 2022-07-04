package bigsky033.spring_examples

import org.springframework.boot.SpringApplication
import org.springframework.boot.WebApplicationType
import org.springframework.boot.autoconfigure.SpringBootApplication

@SpringBootApplication
class SpringExamplesApplication

fun main(args: Array<String>) {
    val app = SpringApplication(SpringExamplesApplication::class.java)
    app.webApplicationType = WebApplicationType.NONE
    app.run(*args)
}
