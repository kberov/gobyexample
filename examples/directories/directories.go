// Го предоставя няколко полезни функции за работа с *папки* във файловата
// уредба.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Създаваме подпапка в текущата работна папка.
	err := os.Mkdir("подпапка", 0755)
	check(err)

	// Когато създаваме временни директории, е добре да отложим премахването им
	// с помощта на `defer`. `os.RemoveAll` изтрива цялото дърво подобно на
	// командата `rm -rf` в Линукс.
	defer os.RemoveAll("подпапка")

	// Помощна функция за създаване на нов празен файл.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("подпапка/фаил1")

	// Можем да създадем ѝерархия от папки, включително и родителската папка с
	// `MkdirAll`. Действието е като на командата `mkdir -p`.
	err = os.MkdirAll("подпапка/родител/дете", 0755)
	check(err)

	createEmptyFile("подпапка/родител/фаил2")
	createEmptyFile("подпапка/родител/фаил3")
	createEmptyFile("подпапка/родител/дете/фаил4")

	// `ReadDir` извежда списък със съдържанието на папката, като връща резен
	// от обекти от вида `os.DirEntry`.
	c, err := os.ReadDir("подпапка/родител")
	check(err)

	fmt.Println("Списък със съдържанието на подпапка/родител")
	for _, entry := range c {
		вид := ""
		if entry.IsDir() {
			вид = "е папка."
		} else {
			вид = "е файл."
		}
		fmt.Println(" ", entry.Name(), вид)
	}

	// С `Chdir` променяме текущата работна директория – същото като `cd`.
	err = os.Chdir("подпапка/родител/дете")
	check(err)

	// Сега ще видим съдържанието на `подпапка/родител/дете`, като изведем
	// съдържанието на *текущата* папка.
	c, err = os.ReadDir(".")
	check(err)

	cwd, e := os.Getwd()
	check(e)
	fmt.Println("Съдържание на", cwd)
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `cd` обратно, откъдето дойдохме.
	err = os.Chdir("../../..")
	check(err)

	// Можем също да навестим папка чрез *самоизвикване*. Това включва всички
	// подпапки. `WalkDir` приема функция, която да извика, за да обработи
	// всеки файл или папка по пѫтя.
	fmt.Println("Посещаваме подпапка")
	err = filepath.WalkDir("подпапка", посети)
	check(err)
}

// `посети` бива извикана за всеки файл или папка, на които сме се натъкнали при
// обхода чрез `filepath.WalkDir`.
func посети(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	вид := ""
	if d.IsDir() {
		вид = "е папка."
	} else {
		вид = "е файл."
	}
	fmt.Println(" ", path, вид)
	return nil
}
