package week01.hello_world;

import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class RandomWord {
  public static void main(String[] args) {
    int index = 1;
    String out = null;
    while (!StdIn.isEmpty()) {
      String token = StdIn.readString();
      if (StdRandom.bernoulli(1.0 / index)) {
        out = token;
      }
      index++;
    }
    StdOut.println(out);
  }
}
