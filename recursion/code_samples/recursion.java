public class recursion {
  static void simpleRecursionExample(int n) {
    if (n == 1) {
        System.out.println(n);
    } else {
        System.out.println(n);
        simpleRecursionExample(n-1);
    }
  }

  public static void Recursion(int[] args) {
    simpleRecursionExample(10);
  }
}