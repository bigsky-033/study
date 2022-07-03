package bigsky033.simpleconsumer

import org.apache.kafka.clients.consumer.ConsumerRecord
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.apache.kafka.common.TopicPartition
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.nio.file.Files

class FileWritingRecordsHandlerTest {

    @Test
    fun `test process`() {
        val tempFilePath = Files.createTempFile("test-handler", ".out")
        try {
            val recordsHandler = FileWritingRecordsHandler(tempFilePath)
            recordsHandler.process(createConsumerRecords())
            val expectedData = listOf("it's but", "a flesh wound", "come back")
            val actualRecords = Files.readAllLines(tempFilePath)

            assertEquals(actualRecords, expectedData)
        } finally {
            Files.deleteIfExists(tempFilePath)
        }
    }

    private fun createConsumerRecords(): ConsumerRecords<String, String> {
        val topic = "test"
        val partition = 0
        val topicPartition = TopicPartition(topic, partition)
        val consumerRecordsList = mutableListOf<ConsumerRecord<String, String>>()
        consumerRecordsList.add(ConsumerRecord(topic, partition, 0, null, "it's but"))
        consumerRecordsList.add(ConsumerRecord(topic, partition, 0, null, "a flesh wound"))
        consumerRecordsList.add(ConsumerRecord(topic, partition, 0, null, "come back"))
        val recordsMap = mapOf(
            topicPartition to consumerRecordsList
        )

        return ConsumerRecords(recordsMap)
    }
}