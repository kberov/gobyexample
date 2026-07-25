# Пуснатите програми връщат изход както ако бихме ги били пуснали направо на
# командния ред.
$ go run spawning-processes.go 
> date
17.06.2026 (ср)  0:40:32 EEST

# `date` не приема флаг `-x`, така че ще приключи с ненулев изходен код, което
# означава грешка.
Код на изхода = 1
> grep hello
hello grep

> ls -a -l -h
общо 24K
drwxr-xr-x  2 berov berov 4,0K 17 юни 00:37 .
drwxr-xr-x 87 berov berov 4,0K 17 яну 14:27 ..
-rw-rw-r--  1 berov berov 4,7K 17 юни 00:37 spawning-processes.go
-rw-rw-r--  1 berov berov   53 15 юни 23:32 spawning-processes.hash
-rw-rw-r--  1 berov berov  481 20 май 01:41 spawning-processes.sh
