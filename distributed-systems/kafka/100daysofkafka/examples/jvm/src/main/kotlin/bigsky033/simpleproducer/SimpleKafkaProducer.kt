package bigsky033.simpleproducer

import org.apache.kafka.clients.producer.Producer
import org.apache.kafka.clients.producer.ProducerRecord
import org.apache.kafka.clients.producer.RecordMetadata
import java.util.concurrent.Future

class SimpleKafkaProducer(
    private val producer: Producer<String, String>,
    private val topic: String,
) {

    fun produce(message: String): Future<RecordMetadata> {
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
        return producer.send(producerRecord)
    }
}
