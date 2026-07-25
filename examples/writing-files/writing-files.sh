# Изпълнете кода, за да пишете във файлове.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Then check the contents of the written files.
# След това вище съдържанието, записано във файловете.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# А сега ще приложим някои от идеите за Вход/Изход върху потоците `stdin` (стандартен вход) и `stdout` (стандартен изход).
