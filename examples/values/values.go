// Go  има различни видове[^types] стойности, като низове,
// цели числа, числа с плаваща запетая, булеви и т.н. Ето
// някои основополагащи примери.
// [^types]: types – типове, видове

package main

import "fmt"

func main() {

	// Низовете може да бъдат съединявани със знака `+`.
	fmt.Println("go" + "lang")

	// Цели и дробни числа.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Булеви с булеви действия, както бихте очаквали.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
