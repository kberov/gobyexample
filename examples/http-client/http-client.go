// Стандартната библиотека на Го включва отлична поддръжка за създаване на
// клиенти и сървъри за HTTP[^http] чрез пакета `net/http`. В този пример ще
// ползваме пакета, за да направим прости заявки по HTTP.
// [^http]: HTTP: Hypertext Transfer Protocol – Протокол за пренос на хипертекст (свръх-текст -> надсловен плет). HTTP е протокол на приложно ниво за предаване на хипермедийни документи, като например HTML.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Правим заявка чрез HTTP метода GET. Функцията `http.Get` е удобна
	// обвивка за последователността от действия – създаване на обект от вида
	// `http.Client` и извикване на метода му `Get`. Тя ползва готовия обект
	// `http.DefaultClient`, с полезни предварителни настройки.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Отпечатваме кода за състоянието на отговора.
	fmt.Println("Състояние на отговора:", resp.Status)

	// Отпечатваме първите пет реда от тялото на отговора.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
