# За да придобием опит с флаговете на командния ред, най-добре е първо да
# компилираме програмата и после да изпълним направо произведения изпълним
# файл.
$ go build command-line-flags.go

# Пробваме изградената програма, като първо ѝ подадем стойности за всички
# флагове.
$ ./command-line-flags -дума=opt -numb=7 -fork -svar=flag
дума: opt
numb: 7
fork: true
svar: flag
tail: []

# Забележете, че ако пропуснете да подадете стойност за някой от флаговете, той
# ще има стойността по подразбиране.
$ ./command-line-flags -дума=opt
дума: opt
numb: 42
fork: false
svar: bar
tail: []

# Допълнителни податки, могат да се предоставят след флаговете.
$ ./command-line-flags -дума=opt a1 a2 a3
дума: opt
...
tail: [a1 a2 a3]

# Забележете, че пакетът `flag` изисква всички флагове да се намират преди
# допълнителните податки (иначе флаговете ще бъдат разбрани като допълнителни
# податки).
$ ./command-line-flags -дума=opt a1 a2 a3 -numb=7
дума: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Използвайте флаговете `-h` или `--help` за показване на автоматично породена
# помощ за програмата. 
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -дума="foo": a string

# Ако подадете необявен в програмата флаг, тя ще изведе грешка  и ще покаже
# помощта отново.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
