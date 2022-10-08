// https://beta.atcoder.jp/contests/abc042/submissions/1944572
object Main extends App {
 
  def solve(arr: Array[Int]): Boolean = {
    val go = if (arr.count(_ == 5) == 2) true else false
    val nana = if (arr.count(_ == 7) == 1) true else false
    go & nana
  }
 
  val sc = new java.util.Scanner(Console.in)
  val a, b, c = sc.nextInt()
 
  val result = solve(Array(a, b, c))
  if (result) println("YES") else println("NO")
}
