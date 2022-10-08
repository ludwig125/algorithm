// https://beta.atcoder.jp/contests/abc044/submissions/1959786
import java.util.Scanner
 
object Main {
  def main(args: Array[String]): Unit = {
    val sc = new Scanner(System.in)
    val N, K, X, Y = sc.nextInt()
    val ans = if(N-K>=0) {
      K*X+(N-K)*Y
    } else {
      N*X
    }
    println(ans)
  }
}
