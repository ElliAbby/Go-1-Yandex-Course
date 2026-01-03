// Мемоизация вместо рекурсии

// Теперь вместо того, чтобы каждый раз пересчитывать одни и те же значения,
// функция будет искать значение в словаре memo и возвращать его, если оно уже было вычислено.
// Это значительно понижает сложность функции.

package main

import "fmt"

var memo = map[int]int{}

func fib(n int) int {
    if n < 2 {
        return n
    }
    if val, ok := memo[n]; ok {
        return val
    }
    memo[n] = fib(n-1) + fib(n-2)
    return memo[n]
}

func main() {
	fmt.Println(fib(8))
}
