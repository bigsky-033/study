package bigsky033

import java.io.FileInputStream
import java.util.Properties


fun loadProperties(fileName: String): Properties {
    val envProps = Properties()

    val input = FileInputStream(fileName)
    envProps.load(input)
    input.close()

    return envProps
}