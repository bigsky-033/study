package week01.hello_world;

public class HelloGoodbye {
  public static void main(String[] args) {
    if (args == null || args.length != 2) {
      return;
    }
    String firstName = args[0];
    String secondName = args[1];

    System.out.printf("Hello %s and %s.%n", firstName, secondName);
    System.out.printf("Goodbye %s and %s.%n", secondName, firstName);
  }
}
