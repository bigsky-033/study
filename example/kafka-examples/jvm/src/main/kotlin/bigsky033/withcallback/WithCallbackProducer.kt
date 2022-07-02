package bigsky033.withcallback

import org.apache.kafka.clients.producer.Producer
import org.apache.kafka.clients.producer.ProducerRecord

class WithCallbackProducer(
    private val producer: Producer<String, String>,
    private val topic: String,
) {
    fun produce(message: String) {
        val parts = message.split("-")

        val key: String
        val value: String
        if (parts.size > 1) {
            key = parts[0]
            value = parts[1]
        } else {
            key = "NO-KEY"
            value = parts[0]
        }
        val producerRecord = ProducerRecord(topic, key, value)
        producer.send(producerRecord) { recordMetadata, exception ->
            if (exception == null) {
                println("Record written to offset ${recordMetadata.offset()} timestamp ${recordMetadata.timestamp()}")
            } else {
                println("An error occurred")
                exception.printStackTrace()
            }
        }
    }
}
