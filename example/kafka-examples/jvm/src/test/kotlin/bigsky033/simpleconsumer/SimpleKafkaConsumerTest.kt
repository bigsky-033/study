package bigsky033.simpleconsumer

import bigsky033.loadProperties
import org.apache.kafka.clients.consumer.ConsumerRecord
import org.apache.kafka.clients.consumer.MockConsumer
import org.apache.kafka.clients.consumer.OffsetResetStrategy
import org.apache.kafka.common.TopicPartition
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.nio.file.Files


class SimpleKafkaConsumerTest {

    @Test
    fun `test consumer`() {
        val propertyFilePath = "src/test/resources/config/simpleconsumer/kafka.properties"
        val tempFilePath = Files.createTempFile("test-consumer-output", ".out")
        val recordsHandler = FileWritingRecordsHandler(tempFilePath)
        val properties = loadProperties(propertyFilePath)
        val topic = properties.getProperty("input.topic.name")
        val topicPartition = TopicPartition(topic, 0)
        val mockConsumer = MockConsumer<String, String>(OffsetResetStrategy.EARLIEST)

        val simpleKafkaConsumer = SimpleKafkaConsumer(mockConsumer, recordsHandler)

        mockConsumer.schedulePollTask {
            addTopicPartitionsAssignmentAndAddConsumerRecords(
                topic,
                mockConsumer,
                topicPartition
            )
        }
        mockConsumer.schedulePollTask { simpleKafkaConsumer.shutdown() }
        simpleKafkaConsumer.runConsume(properties)

        val expectedWords = listOf("foo", "bar", "baz")
        val actualRecords = Files.readAllLines(tempFilePath)
        assertEquals(actualRecords, expectedWords)

    }

    private fun addTopicPartitionsAssignmentAndAddConsumerRecords(
        topic: String,
        mockConsumer: MockConsumer<String, String>,
        topicPartition: TopicPartition,
    ) {
        mockConsumer.rebalance(listOf(topicPartition))
        mockConsumer.updateBeginningOffsets(mapOf(topicPartition to 0L))
        addConsumerRecords(mockConsumer, topic)
    }

    private fun addConsumerRecords(mockConsumer: MockConsumer<String, String>, topic: String) {
        mockConsumer.addRecord(ConsumerRecord(topic, 0, 0, null, "foo"))
        mockConsumer.addRecord(ConsumerRecord(topic, 0, 1, null, "bar"))
        mockConsumer.addRecord(ConsumerRecord(topic, 0, 2, null, "baz"))
    }
}
