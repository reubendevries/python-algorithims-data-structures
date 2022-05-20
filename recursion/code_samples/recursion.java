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

  static int factorialNumber(int n) {
    if (n <= 0) {
      return -1;
    } else if (n < 2){
      return 1;
    } else {
      return (n * factorialNumber(n-1));
    }
  }

  static int fibonacciNumber(int n) {
    if (n == 0 || n == 1) {
      return n;
    } else {
      return fibonacciNumber(n-2) + fibonacciNumber(n-1);
    }
  }

  public static void main(String[] args) {
    simpleRecursionExample(10); // should count down from 10 to 1
    int a = powerOfTwoRecursive(10);
    System.out.println(a); // should print out recursively return result of 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2
    int b = powerOfTwoIterative(10);
    System.out.println(b); // should print out iteratively return result of 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2
    int c = factorialNumber(5);
    System.out.println(c); // should return 120
    int d = fibonacciNumber(10);
    System.out.println(d); // should return 55
  }
}