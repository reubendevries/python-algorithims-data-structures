package java;

public class powerOfTwo {
    
    static int powTwo(int n) {
        if (n == 0) {
          return 1;
        }
        return (powTwo(n-1) * 2);    
    }

    public static void main(String[] args) {
        int x = powTwo(10);
        System.out.println(x); // should print out recursively return result of 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2 x 2
      }
}
