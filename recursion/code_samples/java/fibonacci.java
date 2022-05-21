package java;

public class fibonacci {

    static int fibonacciNumber(int n) {
        if (n == 0 || n == 1) {
            return n;
          } else {
            return fibonacciNumber(n-2) + fibonacciNumber(n-1);
          }
    }
    
    public static void main(String[] args) {
        int x = fibonacciNumber(10);
        System.out.println(x);
    }
}
