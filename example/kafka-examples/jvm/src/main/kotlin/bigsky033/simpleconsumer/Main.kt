package bigsky033.simpleconsumer

import bigsky033.loadProperties
import org.apache.kafka.clients.consumer.KafkaConsumer
import java.nio.file.Paths

fun main() {
    val propertyFilePath = "src/main/resources/config/simpleconsumer/kafka.properties"

    val properties = loadProperties(propertyFilePath)
    val filePath = properties.getProperty("file.path")
    val consumer = KafkaConsumer<String, String>(properties)
    val recordsHandler = FileWritingRecordsHandler(Paths.get(filePath))
    val simpleKafkaConsumer = SimpleKafkaConsumer(consumer, recordsHandler)

    Runtime.getRuntime().addShutdownHook(Thread(simpleKafkaConsumer::shutdown))

    simpleKafkaConsumer.runConsume(properties)
}
