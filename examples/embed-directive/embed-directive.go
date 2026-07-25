// `//go:embed` (влагам) е директива (указание) за компилатора, която позволява програмите
// да включват в изпълнимия си файл произволни файлове и папки по време на
// изграждането му. Можете да прочетете повече за директивата в документацията ѝ -
// (https://pkg.go.dev/embed).
package main

// Внесете пакета `embed`. Ако не ползвате никакви изнесени имена, можете да
// внесете пакета без името му, като напишете `_ "embed"`.
import (
	"embed"
)

// Указанията за `embed` приемат относителни пътища спрямо папката, в която се
// намира изходния файл, в който е съответното указание. Следното указание
// влага съдържанието на файла в низовата променлива, обявена веднага под него.
//
//go:embed folder/single.txt
var fileString string

// Съдържанието на файла може да се вложи и в променлива от вид `[]byte`.
//
//go:embed folder/single.txt
var fileByte []byte

// Също така можем да вложим множество файлове или цели папки чрез заместващи
// знаци (wildcards). Следнното изявление ползва променлива от вида embed.FS
// (https://pkg.go.dev/embed#FS).
//
//go:embed folder/single.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Извеждаме съдържанието на `single.txt`.
	print(fileString)
	print(string(fileByte))

	// Извличаме съдържанието на файлове от вложената папка.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))

	content3, _ := folder.ReadFile("folder/single.txt")
	print(string(content3))
}
