// Го предлага вградена поддръжка за превръщане към и от
// JSON[^json] както за вградени така и за потребителски
// видове данни.
// [^json]: JSON – JavaScript Object Notation – описване на данни като обекти в JavaScript

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Ще използваме тези две структури по-долу, за да покажем
// превръщане към и от потребителски видове.
type response1 struct {
	Страница int
	Fruits   []string
}

// Само изнесените полета ще бъдат превърнати от и в JSON.
// За да са _изнесени_, полетата трябва да започват с
// главна буква.
type response2 struct {
	Страница int      `json:"страница"`
	Fruits   []string `json:"fruits"`
}

func main() {

	// Първо ще видим как се превръщат първични видове
	// данни в JSON низове. Ето някои примери с
	// неразложими стойности.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// А ето и такива с отрязъци и карти, които биват
	// превърнати в поредици и обекти в JSON, както се
	// предполага.
	slcD := []string{"ябълка", "праскова", "круша"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"ябълка": 5, "маруля": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Пакетът JSON може да превръща автоматично
	// потребителски видове данни. Само изнесените полета
	// ще бъдат включени в изхода и имената им ще бъдат
	// ключове в JSON.
	res1D := &response1{
		Страница: 1,
		Fruits:   []string{"ябълка", "праскова", "круша"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Можете да отбелязвате с какво име да бъдат
	// превръщани полетата, като използвате бележки,
	// познати като тагове (tags). Вижте обявлението на
	// структурата `response2` горе. Там са използвани
	// тагове.
	res2D := &response2{
		Страница: 1,
		Fruits:   []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Сега да видим как се превръщат данни от JSON в Го.
	// Ето пример с призволна структура в JSON.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Трябва да предоставим променлива, в която да
	// запишем разбраните данни. Тази
	// `map[string]interface{}` ще бъде карта с ключове
	// низове и стойности в произволен вид.
	var dat map[string]interface{}

	// Ето превръщането в действие. При превръщането
	// проверяваме за грешки.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// За да използваме данните, трябва да ги превръщаме
	// (всяко поотделно) в подходящ вид. Например тук
	// превръщаме стойността, прикрепена към `num` в
	// очаквания вид `float64`[^assertion].
	// [^assertion]: Забележете, че използвахме потвърждение за вида данни, без да проверяваме възможно ли е това. TODO: Да се добави отделен пример за type assertion!!!
	num := dat["num"].(float64)
	fmt.Println(num)

	// За да достъпим вгнездени данни, трябва да правим
	// ново потвърждение на всяко ниво.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Можем също да превръщаме JSON в потребителски
	// видове данни. Така добавяме сигурност за вида
	// данни, който очакваме и няма нужда да потвърждаваме
	// вида им при достъп до тях.
	str := `{"страница": 1, "fruits": ["ябълка", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%+v\n", res)
	fmt.Println(res.Fruits[0])

	// В примерите горе винаги използвахме байтове и
	// низове, като посредници за данните между
	// JSON и стандартния изход. Можем също да точим
	// данните в двете посоки към и от всеки обект
	// осъществяващ `os.Writer`, например `os.Stdout`, и
	// дори тялото на отговора в HTTP[^HTTP].
	// [^HTTP]: HTTP (Hypertext Transfer Protocol) – Протокол за пренос на хипертекст
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// Четене на поток от данни чрез взаимодействия
	// `os.Reader` като `os.Stdin` или от тялото на
	// отговори в HTTP, се прави с `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
