// Изявлението _defer_ (отлагане (на изпълнение)) се
// използва за осигуряване извикването на някоя функция
// по-късно в работата на някоя програма. `defer` се
// използва в същия смисъл, в какъвто биха се използвали в
// други езици например `ensure` и `finally`.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Да предположим, че искаме да създадем файл, да напишем
// нещо в него и после да го затворим. Ето как бихме
// направили това, като употребим и `defer`.
func main() {

	// Веднага след като получим обект за файл от
	// функцията `createFile`, ние отлагаме затварянето на
	// файла, като извикваме отложено `closeFile`. това
	// извикване ще се изпълни в края на обгръщатата
	// функция (в случая – `main`), след като `writeFile`
	// е приключила.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	f, err := os.Create(p)
	fmt.Println("създаваме", f.Name())
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("пишем в ", f.Name())
	fmt.Fprintln(f, "данни")
}

func closeFile(f *os.File) {
	fmt.Println("затваряме", f.Name())
	err := f.Close()
	// Важно е да се проверява за грешки, когато затваряме
	// файл дори в отложена функция.
	if err != nil {
		panic(err)
	}
}
