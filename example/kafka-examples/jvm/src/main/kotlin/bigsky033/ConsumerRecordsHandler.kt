package bigsky033

import org.apache.kafka.clients.consumer.ConsumerRecords

interface ConsumerRecordsHandler<K, V> {

    fun process(consumerRecords: ConsumerRecords<K, V>)
}
