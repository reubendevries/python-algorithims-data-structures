public class recursion {
    static void recursionExample(int n) {
      if (n == 1) {
          System.out.println(n);
      } else {
          System.out.println(n);
          recursionExample(n-1);
      }
    }
  }