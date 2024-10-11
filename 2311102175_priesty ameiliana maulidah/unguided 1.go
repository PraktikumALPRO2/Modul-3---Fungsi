
package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func permutation(n, r int) *big.Int {
	return new(big.Int).Div(factorial(n), factorial(n-r))
}

func combination(n, r int) *big.Int {
	return new(big.Int).Div(factorial(n), new(big.Int).Mul(factorial(r), factorial(n-r)))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(permutation(a, c))
		fmt.Println(combination(b, d))
	}
}

