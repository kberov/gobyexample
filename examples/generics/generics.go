// От издание 1.18 насам Го поддържа _обобщения_[^generics], познати още като _податки за вид_[^type_p].
// [^generics]: generics – обобщения
// [^type_p]: type parameter – податка за очакван вид данни

package main

import "fmt"

// `SlicesIndex` е пример за обобщена функция. Приема отрязък от всеки сравним (`comparable`) вид данни и един член от такъв вид, и връща показалец към първото съвпадение с `v`  в `s`, или -1, ако не го намери.
// _Ограничението_[^constraint] `comparable` (сравним)
// означава, че можем да сравняваме данните от такъв вид
// чрез действията[^operators] `==` и `!=`. Много по-задълбочено обяснение ще намерите в [тази статия](https://go.dev/blog/deconstructing-type-parameters).
// Обърнете внимание, че тази функция съществува в
// стандартната библиотека под името
// [slices.Index](https://pkg.go.dev/slices#Index).
// [^constraint]: generic constraint – обобщaващо ограничение (за определен набор от видове)
// [^operators]: operators – действия
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Ето един примерен обобщен вид. `List` е единично
// свързан списък, който може да съдържа стойности от
// _всякакъв_[^any] вид.
// [^any]: any – всякакъв. Това е друго име за `interface{}`, което в Го означава данна от _всякакъв вид_.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Можем да създаваме методи за обобщени видове, но трябва
// да обявяваме податките за вида. Вида е `List[T]`, не
// `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// `AllElements` връща всички членове на `List` като
// отрязък. В следващия пример ще видим
// по-изразителен[^idiomatic] начин да обикаляме по
// членовете на потребителски[^custom] видове.
// [^custom]: custom – потребителски. Собственоръчно направен от  програмиста – потребителя на езика.
// [^idiomatic]: idiomatic – изразителен, присъщ. Изразителен, според начина на мислене, изказване, слововобаразуване, изразяване в съответния език.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"ала", "бала", "ница"}

	// Когато извикваме обобщени функции, често можем да
	// разчитаме, че Го ще _отгатне_[^tinf] вида.
	// Забележете, че не е нужно да указваме вида на `S` и
	// `E`, когато извикваме `SlicesIndex`. Компилаторът
	// отгатва вида им сам.
	// [^tinf]: type inference – отгатване на вида. To infer – правя заключение за, загатвам, подсказвам, а в случая – отгатвам.
	fmt.Printf("на кое място е %q:  %d\n",
		s[0], SlicesIndex(s, "ала"))

	// ... макар че бихме могли да ги зададем и изрично.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("списък:", lst.AllElements())
}
