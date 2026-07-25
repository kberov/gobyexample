// Често срещано изискване в програмите е да се вземе броя
// на секундите, милисекундите или нанносекундите, изтекли
// от т.нар Епоха на Юникс[^epoch]. Ето как се прави това
// в Го.
// [^epoch]: Епоха на Юникс – времето в невисокосни секунди, започващо в 00:00:00 ч. UTC на 1ви януари 1970 г.  насам. Виж https://en.wikipedia.org/wiki/Unix_time и https://bg.wikipedia.org/wiki/Високосна_секунда

package main

import (
	"fmt"
	"time"
)

func main() {

	// Използвайте `time.Now` с методите `Unix`,
	// `UnixMilli` или `UnixNano`, за да вземете изтеклото
	// (сиреч текущото) време от началото на епохата
	// Юникс съответно в секунди, милисекунди или
	// наносекунди.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Можете също да превърнете целочислените секунди от
	// епохата в съответните стойности от вида `time`.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
