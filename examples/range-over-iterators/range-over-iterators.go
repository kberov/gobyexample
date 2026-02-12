// В издание 1.23 Го добави поддръжка на
// [повторители](https://go.dev/blog/range-functions)[^iterators],
// които ни позволяват да обхождаме почти всичко!
// [^iterators]: iterators – повторители – функции, които се ползват с `range`

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Да разгледаме отново вида `List` от [предишния
// пример](generics). Там имахме метод `AllElements`,
// който връщаше отрязък с всички членове на списъка. С
// повторителите в Го можем да го направим по-добре, както
// е показано долу.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// `All` връща _повторител_. В Го това е функция със [особено заглавие](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Повтарящата функция приема друга функция като
		// данна, наречена обичайно `yield` (но името ѝ
		// може да е произволно). Повтарящата функция ще
		// извика `yield` за всеки член върху които искаме
		// да извършим някакво действие и ще следи
		// връщаната от `yield` стойност, ако е нужно да
		// приключи незабавно.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Повторението не изисква някаквакви данни, които да
// обхождаме, и дори може да е безкрайно! Ето една
// функция, връщаща повторител върху числата на Фибоначи.
// Тя продължава да връща число, докато `yield` връща
// `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Тъй като `List.All` връща повторител, можем да я
	// ползваме както обикновено с `range`
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Пакети като [slices](https://pkg.go.dev/slices)
	// имат множество полезни функции за работа с
	// повторители. Например `Collect` приема всякакъв
	// повторител и събира всичките му стойности в
	// отрязък.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Ако някое повторение се натъкне на `break` или
		// `return`, функцията `yeld`, подадена на
		// повторителя, ще върне `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
