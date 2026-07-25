// Пакетът `math/rand/v2` производство на [псевдослучайни
// числа](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Например, `rand.IntN` връща случайно цяло число – `int` n, `0 <= n <
	// 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` връща десетична дроб – `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Това може да се използва за създаване на случайни
	// числа с плаваща запетая в други порядъци, например
	// `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Ако ви трябва посевка[^seed] с позната стойност,
	// създайте нова стойност на `rand.Source` и я подайте
	// на конструктора `New`. `NewPCG` създава нова
	// стойност от вида
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator).
	// Това е източник, който изисква посев от две числа
	// от вида `uint64`.
	// [^seed]: seed – семе, зърно
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
