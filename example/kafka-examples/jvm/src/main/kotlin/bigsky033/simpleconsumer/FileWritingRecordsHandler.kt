package bigsky033.simpleconsumer

import bigsky033.ConsumerRecordsHandler
import org.apache.kafka.clients.consumer.ConsumerRecords
import java.nio.file.Files
import java.nio.file.Path
import java.nio.file.StandardOpenOption

class FileWritingRecordsHandler(
    private val path: Path
) : ConsumerRecordsHandler<String, String> {

    override fun process(consumerRecords: ConsumerRecords<String, String>) {
        val values = mutableListOf<String>()
        consumerRecords.forEach { values.add(it.value()) }
        if (values.isNotEmpty()) {
            Files.write(path, values,
                StandardOpenOption.CREATE, StandardOpenOption.WRITE, StandardOpenOption.APPEND)
        }
    }
}