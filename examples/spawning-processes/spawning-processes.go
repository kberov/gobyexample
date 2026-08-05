// Понякога се налага нашите програми да породят[^spawn] други процеси.
// [^spawn]: to spawn a process – пораждам, създавам процес

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Да започнем с проста команда, която не приема податки или вход, а просто
	// отпечатва нещо на стандартния изход. Помощната функция `exec.Command`
	// създава обект, представляващ този външен процес.
	dateCmd := exec.Command("date")

	// Методът `Output` изпълнява командата, чака я да приключи и събира нейния
	// стандартени изход. Ако няма грешки, `dateOut` ще съдържа резен от
	// байтове – изходът на `date`.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` и други методи на `Command` връщат `*exec.Error` ако бъдат
	// възпрепятствани при изпълнението на командата (например грешен пѫт), и
	// `*exec.ExitError`, ако командата се изпълни, но излезе с код различен от
	// нула.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("Провалено изпълнение:", e)

		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("Код на изхода =", exitCode)
		} else {
			panic(err)
		}
	}

	// Сега да разгледаме по-сложен случай, в който подаваме данни на външния
	// процес по неговия стандартен вход и събираме произведеното от процеса от
	// неговия стандартен изход.
	grepCmd := exec.Command("grep", "hello")

	// Тук изрично боравим с входноизходните потоци (тръби)[^pipes], пускаме
	// процеса, пишем някакви данни във входа, четем изхода му и накрая чакаме
	// процеса да приключи.
	// [^pipes]: pipe – тръба. Това е обичайният начин в юникс-подобните уредби за навързване на няколко програми последователно за извършване на някаква определена задача. Изведените изходни данни от една програма са входни данни за друга програма. Данните са просто низове (последователности от байтове.)
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// В горнния пример не проверихме за грешки, но винаги може да се ползва
	// обичайното `if err != nil`. Също така четохме само от `StdoutPipe`, но
	// вие може да ползвате изхода в `StderrPipe` по съвсем същия начин.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Забележете, че когато пускаме команди трябва да предоставим изрично
	// командата отделно и отделно поредица от податки към командата вместо да
	// подадем командата и податките като един цял низ. Ако искате да пуснете
	// команда заедно с податките ѝ в един низ, можете да я податеде на `bash`
	// с флаг `-c`, ето така:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	// Пускане на програма, която не изчакваме да приключи. (бел. прев.)
	// zombie := exec.Commanddd(`bash`, `-c`,
	//	`date >>out.txt; sleep 2; date >> out.txt`)
	// zombie.Start()
}
