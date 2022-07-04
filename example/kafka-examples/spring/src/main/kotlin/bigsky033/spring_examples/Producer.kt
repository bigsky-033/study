package bigsky033.spring_examples

import org.slf4j.LoggerFactory
import org.springframework.kafka.core.KafkaTemplate
import org.springframework.kafka.support.SendResult
import org.springframework.stereotype.Component
import org.springframework.util.concurrent.ListenableFutureCallback

@Component
class Producer(
    private val kafkaTemplate: KafkaTemplate<String, String>,
) {

    companion object {
        private val log = LoggerFactory.getLogger(Producer::class.java)
        private const val TOPIC = "purchases"
    }

    fun sendMessage(key: String, value: String) {
        val future = kafkaTemplate.send(TOPIC, key, value)
        future.addCallback(object : ListenableFutureCallback<SendResult<String, String>> {
            override fun onSuccess(result: SendResult<String, String>?) {
                log.info("Produced event to topic $TOPIC: key = $key value=$value")
            }

            override fun onFailure(ex: Throwable) {
                ex.printStackTrace()
            }
        })
    }
}
