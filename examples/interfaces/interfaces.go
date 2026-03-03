// _Взаимодействията_ са именувани набори от заглавия на
// методи.

package main

import (
	"fmt"
	"math"
)

// Ето едно просто взаимодействие за работа с геометрични
// фигури.
type geometry interface {
	area() float64
	perim() float64
}

// За нашия пример ще осъществим[^implement] това
// взаимодействие за видовете `rect` (четириъгълник) и
// `circle` (кръг).
// [^implement]: implement – осъществявам, имплементирам
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// За да осъществим взаимодействие в Го, просто трябва да
// осуществим всички методи, описани във взаимодействието.
// Тук осъществяваме геометрията на четиригиъгълниците.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Осъществяването за кръгове се състои в следното.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Ако дадена променлива е обявена като някакъв вид
// взаимодействие, можем да извикваме върху нея методите,
// описани във взаимодействието. Ето една
// _обобщена_[^generic] за всички променливи от вида
// `geometry` функция.
// [^generic]: generic – обобщен
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Понякога е полезно да знаем фактическия вид на дадена
// променлива по време на изпълнение, която е обявена като
// вид взаимодействие. Един начин да направим това е като
// използваме *потвърждение за вида*[^type_assert]. Друг
// такъв начин е [*превключване според вида*[^type_switch]](switch).
// [^type_assert]: type assertion – потвърждение за вида данни
// [^type_switch]: type switch – превключване според вида данни
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("кръг с радиус:", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Структурните видове `circle` and `rect`
	// осъществяват взаимодействието, та ка че можем да
	// ползваме *инстанции* от тези видове като податки на
	// функцията `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
