package bigsky033.withcallback

import bigsky033.loadProperties
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.Producer
import java.nio.file.Files
import java.nio.file.Paths


fun main() {
    val propertyFilePath = "src/main/resources/config/withcallback/kafka.properties"
    val dataFilePath = "src/main/resources/data/withcallback/input.txt"

    val properties = loadProperties(propertyFilePath)
    val topic = properties.getProperty("output.topic.name")

    val producer: Producer<String, String> = KafkaProducer(properties)
    val simpleKafkaProducer = WithCallbackProducer(producer, topic)

    val linesToProduce = Files.readAllLines(Paths.get(dataFilePath))
    linesToProduce
        .filter { it.trim().isNotEmpty() }
        .forEach { simpleKafkaProducer.produce(it) }

    producer.close()
}
