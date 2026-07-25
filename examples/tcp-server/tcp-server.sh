# Пускаме сървър от вида TCP на заден план.
$ go run tcp-server.go &

# Изпращаме данни и прихващаме отговора с помощтта на `netcat`.
$ printf "%s\n" "Здрасти от netcat" | nc localhost 8090
Прието: "ЗДРАСТИ ОТ NETCAT" 

#Убиваме съръвъра
$ killall tcp-server
signal: terminated
[1]+  Изход 1            go run tcp-server.go
