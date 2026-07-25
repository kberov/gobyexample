// Го предлага вградена поддръжка за XML  и подобни на XML
// формати чрез пакета `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Структурата `Plant` ще бъде превърната в XML низ.
// Подобно на примерите с JSON, към полетата са прикрепени
// бележки, които са указания за превръщане към и от XML.
// Тук ползваме някои особени възможности на пакета XML.
// Полето `XMLName` има бележка (таг в Го), която указва
// как ще се обозначава структурата в низовия вид XML –
// като таг (начина за отбелязване в XML), докато полето
// `Id` ще се отбелязва като _атрибут_ (нещо присъщо,
// прикачено към таг, негово свойство (на тага)), а не като
// вложена част (таг), както ще се отбелязват другите
// полета.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Име     string   `xml:"име"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Име, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Име: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Превръщаме в XML като ползваме `MarshalIndent`, за
	// да произведем по-четим изход.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// За да добавим по-обичайна за XML заглавка към
	// изхода, трябва да я добавим изрично.
	fmt.Println(xml.Header + string(out))

	// Ползваме `Unmarshal`, за да превърнем поток от
	// байтове от XML в структура от данни. Ако в потока
	// има зле форматиран XML или не се вмества в `Plant`,
	// функцията ще върне описателна грешка.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Име: "Домат"}
	tomato.Origin = []string{"Mexico", "California"}

	// Бележката `parent>child>plant` за полето `Plants`
	// казва на превръщача да вгнезди всички структури от
	// вида `plant` в `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
