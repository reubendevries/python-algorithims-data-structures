public class recursion {

  static void simpleRecursionExample(int n) {
    if (n == 1) {
        System.out.println(n);
    } else {
        System.out.println(n);
        simpleRecursionExample(n-1);
    }
  }

  static int powerOfTwoRecursive(int n) {
    if (n == 0) {
      return 1;
    } else {
      return (powerOfTwoRecursive(n-1) * 2);
    }
  }

  static int powerOfTwoIterative(int n) {
    int i = 0;
    int power = 1;
    while (i < n) {
      power = (power * 2);
      i++;
    }
    return power;
  }

  public static void main(String[] args) {
    simpleRecursionExample(10);
    int x = powerOfTwoRecursive(10);
    System.out.println(x);
    int y = powerOfTwoIterative(10);
    System.out.println(y);
  }
}