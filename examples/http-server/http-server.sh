# Пускаме сървъра във фонов режим.
$ go run http-server.go &

# Достъпваме пътя `/hello`.
$ curl localhost:8090/hello
здрасти

# Достъпваме пътя `/headers`.
$ curl localhost:8090/headers
User-Agent: curl/8.14.1
Accept: */*
