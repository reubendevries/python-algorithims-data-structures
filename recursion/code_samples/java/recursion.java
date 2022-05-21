package java;

public class recursion {

  static int simpleRecursion(int n) {
    if (n == 1) {
        return n;
    }
    System.out.println(n);
    simpleRecursion(n-1);
    return 0;
  }

  

  public static void main(String[] args) {
    int x = simpleRecursion(10); // should count down from 10 to 1
    System.out.println(x);
  }

}