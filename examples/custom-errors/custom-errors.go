// Можем да съставяме нови видове грешки[^custerr] като
// осъщетвим метода `Error()` за вид, създаден от нас. Ето
// една разновидност на предишния пример, който използва
// потребителски[^custom] вид данни, за да представи
// изрично, грешка при подаване на данни.
// [^custerr]: custom error – вид грешка, осъществена от разработчика
// [^custom]: custom – потребителски, на потребителя. В случая потребителят е програмистът.

package main

import (
	"errors"
	"fmt"
)

// Съставените грешки обикновено имат наставката "Error".
type argError struct {
	arg     int
	message string
}

// Като добавим метода `Error`, осъществяваме взаимодействието `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%s %d", e.message, e.arg)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Връщаме съставената от нас грешка.
		return -1, &argError{arg, "не мога да работя с"}
	}
	return arg + 3, nil
}

func main() {

	// Функцията `errors.AsType` е по-сложна разновидност на `errors.Is`. С нея
	// проверяваме дали дадена грешка (или всяка грешка в нейната верига) е от
	// определен вид, като превръща текущата грешка в очаквания вид, връща
	// грешката и `true`. Ако не успее, втората върната стойност е `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.message, ae.arg)
	} else {
		fmt.Println("err не е от вида argError")
		fmt.Println(ae.Error())
	}
}
