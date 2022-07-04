package bigsky033.spring_examples

import org.slf4j.LoggerFactory
import org.springframework.kafka.annotation.KafkaListener
import org.springframework.kafka.support.KafkaHeaders
import org.springframework.messaging.handler.annotation.Header
import org.springframework.stereotype.Component

@Component
class Consumer {

    companion object {
        private val log = LoggerFactory.getLogger(Consumer::class.java)
    }

    @KafkaListener(id = "myConsumer", topics = ["purchases"], groupId = "spring-boot", autoStartup = "false")
    fun listen(
        value: String,
        @Header(KafkaHeaders.RECEIVED_TOPIC) topic: String,
        @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) key: String
    ) {
        log.info("Consumed event from topic $topic: key = $key value = %value")
    }
}
