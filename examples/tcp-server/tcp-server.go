// Пакетът `net` предоставя нужните пособия за лесно изграждане на сървъри,
// работещи съгласно TCP протокола.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// Функцията `net.Listen` пуска сървъра в дадената мрежа (на ниво TCP) и порт за
	// адреси (порт 8090 на всички мрежови приставки[^if]).
	// [^if]: network interface – мрежова приставка. Устройството (логическо или физическо), което предоставя мрежа и е част от машината или закачено като външно устройство.
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// Затваряме слушателя, за да освободим порта, когато приложението
	// приключи.
	defer listener.Close()

	// Повтаряме безкрай `listener.Accept()`, за да приемаме заявки за
	// свързване от клиенти.
	for {
		// Чакаме свързване.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Грешка при приемане:", err)
			continue
		}

		// Тук използваме гозадача, за да обработим установената връзка, така
		// че повторението `for` да може да продължи да приема още заявки за
		// свързване.
		go handleConnection(conn)
	}
}

// `handleConnection` обработва отделна заявка от клиент, като чете един ред
// от низа, изпратен от клиента и връща отговор.
func handleConnection(conn net.Conn) {
	// Затваряме установената връзка, за да освободим заетите блага[^res], след
	// като прикючим общуването с клиента.
	// [^res]: resources - блага
	defer conn.Close()

	// Ползваме `bufio.NewReader`, за да прочетем един ред данни от клиента
	// (редът завършва със знак за нов ред).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Грешка при четенето: %v", err)
		return
	}

	// Създаваме и изпращаме отговор обратно към клиента и така показваме
	// двупосочно общуване.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("Прието: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Сървърна грешка при писането: %v", err)
	}
}
