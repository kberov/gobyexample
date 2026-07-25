// Писането във файлове следва подобни похвати като тези при четенето.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// За начало, да видим как да хвърлим някакъв низ (или просто байтове)
	// във файл.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// При нужда от писане „на парче”, трябва първо да отворим файла за
	// писане.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// Обичайно, веднага след отварянето отлагаме `Close`.
	defer f.Close()

	// С `Write`(пиши) можете да пишете и последователности от байтове.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("записах %d байта\n", n2)

	// Също така има и метод `WriteString` (ПишиНиз).
	n3, err := f.WriteString("пише\n")
	check(err)
	fmt.Printf("записах %d байта\n", n3)

	// Изпълнете `Sync`, за да отиде всичко писано (досега в подвижната
	// памет) на диска.
	f.Sync()

	// Пакетът `bufio` предоставя освен складиращи четци, така и складиращи
	// писци.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("складирано\n")
	check(err)
	fmt.Printf("записах %d байта\n", n4)

	// Използвайте `Flush`, за да се уверите, че всички действия са
	// приложени върху основния писец.
	w.Flush()

}
