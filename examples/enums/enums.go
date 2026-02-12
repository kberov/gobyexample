// _Изброяващите видове_ (броители)[^enums] са частен
// случай на сумарен вид от алгебрични данни – [sum
// types](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Броителят е вид данни, който има определен брой
// възможни стойности, всяка от които има различно име. В
// Го няма отделен вид данни за броители като част от
// езика, но те са прости за осъществяване чрез
// съществуващите вече в езика изразни средства.
// [^enums]: enumerated types, enums – изброяващи видове, броители

package main

import "fmt"

// Нашият вид `ServerState` (СъстояниеНаСървъра) е основан
// на вида `int`.
type ServerState int

// Възможните стойности `ServerState` са зададени като
// непроменливи. Ключовата дума
// [`iota`](https://go.dev/ref/spec#Iota) създава
// автоматично последователни непроменливи стойности; в
// този случай 0, 1, 2 и т.н.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Като осъществим взаимодействието [fmt.Stringer](https://pkg.go.dev/fmt#Stringer) можем да извеждаме стойностите на `ServerState` или да ги превръщаме в низове.
// Това може да се окаже тежичко, ако имаме много възможни
// стойности. В такъв случай пособието[^tool] за командния ред
// [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// може да бъде ползвано с `go:generate` за да
// автоматизираме задачата. В [тази
// статия](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// ще видите по-подробно обяснение.
// [^tool]: tool – пособие, оръдие (на труда), инструмент
var stateName = map[ServerState]string{
	StateIdle:      "бездеен",
	StateConnected: "свързан",
	StateError:     "грешка",
	StateRetrying:  "нов опит",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Ако имаме стойност от вида `int`, не можем да я
	// подадем на `transition`. Компилаторът ще се оплаче
	// за несъответствие на вида. Това в някаква степен ни
	// предоставя безопасност по отношение на броителите
	// по време на компилация.
	ns2 := transition(ns)
	fmt.Println(ns2)
}

// `transition` наподобява прехода към различни състояния
// на сървър. Приема текущо състояние и връща ново
// състояние.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Да си представим тук, че проверяваме някакви
		// предпоставки, за да определим следващото
		// състояние.
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("незнайно състояние: %s", s))
	}
}
