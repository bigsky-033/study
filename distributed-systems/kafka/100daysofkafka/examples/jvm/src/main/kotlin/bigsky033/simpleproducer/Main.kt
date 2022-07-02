package bigsky033.simpleproducer

import bigsky033.loadProperties
import org.apache.kafka.clients.producer.KafkaProducer
import org.apache.kafka.clients.producer.Producer
import org.apache.kafka.clients.producer.RecordMetadata
import java.nio.file.Files
import java.nio.file.Paths
import java.util.concurrent.Future


fun printMetadata(
    metadata: List<Future<RecordMetadata>>,
    fileName: String,
) {
    println("Offsets and timestamps committed in batch from $fileName")
    metadata.forEach { m ->
        val recordMetadata = m.get()
        println("Record written to offset ${recordMetadata.offset()} timestamp ${recordMetadata.timestamp()}")
    }
}


fun main() {
    val propertyFilePath = "src/main/resources/config/simpleproducer/kafka.properties"
    val dataFilePath = "src/main/resources/data/simpleproducer/input.txt"

    val properties = loadProperties(propertyFilePath)
    val topic = properties.getProperty("output.topic.name")

    val producer: Producer<String, String> = KafkaProducer(properties)
    val simpleKafkaProducer = SimpleKafkaProducer(producer, topic)

    val linesToProduce = Files.readAllLines(Paths.get(dataFilePath))
    val metadata = linesToProduce
        .filter { it.trim().isNotEmpty() }
        .map { simpleKafkaProducer.produce(it) }
        .toList()

    printMetadata(metadata, dataFilePath)

    producer.close()
}
