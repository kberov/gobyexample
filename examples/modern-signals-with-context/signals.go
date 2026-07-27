// Понякога ни се иска нащите програми да обработват умно [Юникс
// сигнали](https://ru.wikipedia.org/wiki/Сигнал_(Unix\)). Например бихме
// искали сървърът ни да спира, когато получи сигнал `SIGTERM`, или наше
// приложение за командния ред да спре да обработва входни данни, щом получи
// сигнал `SIGINT`. Ето как да обработваме сигнали в Го с помощта на вида
// Context.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` връща обект – текущ _смисъл_, който бива прекъснат, щом
	// някой от сигналите в списъка пристигне.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Програмата ще чака тук, докато някой от настроените сигнали бъде получен.
	fmt.Println("чакам знак")
	<-ctx.Done()

	// `context.Cause` дава отчет защо _смисълът_ е бил прекъснат. При
	// прекъсване, причинено от сигнал, тази стойност включва стойността на
	// сигнала.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("излизам")
}
