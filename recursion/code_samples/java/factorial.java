package java;

public class factorial {
    static int factorialNumber(int n) {
        if (n < 2){
            return 1;
          } 
        return (n * factorialNumber(n-1));
    }
    public static void main(String[] args) {
        int x = factorialNumber(10);
        System.out.println(x);
    }
}
