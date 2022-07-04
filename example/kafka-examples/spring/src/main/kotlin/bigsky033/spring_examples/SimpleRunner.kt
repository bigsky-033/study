package bigsky033.spring_examples

import org.springframework.boot.CommandLineRunner
import org.springframework.kafka.config.KafkaListenerEndpointRegistry
import org.springframework.stereotype.Component

@Component
class SimpleRunner(
    private val producer: Producer,
    private val kafkaListenerEndpointRegistry: KafkaListenerEndpointRegistry,
) : CommandLineRunner {

    override fun run(vararg args: String?) {
        args.forEach { arg ->
            when (arg) {
                "--producer" -> {
                    producer.sendMessage("awalther", "t-shirts");
                    producer.sendMessage("htanaka", "t-shirts");
                    producer.sendMessage("htanaka", "batteries");
                    producer.sendMessage("eabara", "t-shirts");
                    producer.sendMessage("htanaka", "t-shirts");
                    producer.sendMessage("jsmith", "book");
                    producer.sendMessage("awalther", "t-shirts");
                    producer.sendMessage("jsmith", "batteries");
                    producer.sendMessage("jsmith", "gift card");
                    producer.sendMessage("eabara", "t-shirts"); }

                "--consumer" -> {
                    kafkaListenerEndpointRegistry.getListenerContainer("myConsumer")?.start()
                }
            }
        }
    }
}
