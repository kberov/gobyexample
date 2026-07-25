# Като пуснем програмата виждаме стойността на променливата FOO, която сме
# задали, както и че стойността на BAR е празна.
$ go run environment-variables.go
FOO: 1
BAR: 

# Списъкът с ключове от средата може да е различен на всяка машина.
SHELL
SESSION_MANAGER
PERLBREW_ROOT
QT_ACCESSIBILITY
COLORTERM
...
FOO

# Ако пък първо зададем стойност за BAR, програмата ни ще изведе наличната
# стойност.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
