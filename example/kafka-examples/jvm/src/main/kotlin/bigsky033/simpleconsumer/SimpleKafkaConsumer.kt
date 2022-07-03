package bigsky033.simpleconsumer

import bigsky033.ConsumerRecordsHandler
import org.apache.kafka.clients.consumer.Consumer
import java.time.Duration
import java.util.Collections
import java.util.Properties

class SimpleKafkaConsumer(
    private val consumer: Consumer<String, String>,
    private val recordsHandler: ConsumerRecordsHandler<String, String>,
) {

    @Volatile
    private var keepConsuming = true

    fun runConsume(consumerProps: Properties) {
        consumer.use { consumer ->
            consumer.subscribe(Collections.singletonList(consumerProps.getProperty("input.topic.name")))
            while (keepConsuming) {
                val consumerRecords = consumer.poll(Duration.ofSeconds(1L))
                recordsHandler.process(consumerRecords)
            }
        }
    }

    fun shutdown() {
        keepConsuming = false
    }
}
