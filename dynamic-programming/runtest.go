package main

import "fmt"
import "time"

func main() {
	word := "aaaabbbba"
	fmt.Printf("%d\n", lpsRec(word, 0, len(word)-1))
	fmt.Printf("%d\n", lpsDp(word))
	fmt.Printf("%d\n", lpsDp2(word))

	fmt.Println("矩阵连乘问题-----")
	D := []int{2, 2, 2, 2, 2} // 4 matrices
	fmt.Print(matrixChainRec(D, 1, 4), "\n")
	fmt.Print(matrixChainDp(D), "\n")

	fmt.Println("rod-cutting-----")
	length := 10
	price := []int{0, 1, 5, 8, 9, 17, 17, 17, 20, 24, 30}
	// price := []int{0, 10, 5, 8, 9, 17, 17, 17, 20, 24, 30}

	// fmt.Print(price[5]+price[length-5], "\n")

	fmt.Print(cutRodRec(price, length), "\n")
	fmt.Print(cutRodDp(price, length), "\n")
	// fib数量在多于100的时候递归方式已经很难出结果，递归层级太多了
	fmt.Println("fib-----")
	t := time.Now()
	fmt.Println(fibMemo(100))
	fmt.Println("fibMemo耗时：", time.Since(t))
	fmt.Println(fibBotUp(100))
	fmt.Println(fibarray(100)[100-1])
	//res := fibonacci(40)
	//fmt.Println(res)
	fmt.Println(fibTail(1, 1, 98))
}
