# Когато пуснем нашата програма, тя бива заместена от `ls`.
$ go run execing-processes.go
общо 20K
drwxr-xr-x  2 berov berov 4,0K 18 юни 16:01 .
drwxr-xr-x 87 berov berov 4,0K 17 яну 14:27 ..
-rw-r--r--  1 berov berov 2,3K 18 юни 15:47 execing-processes.go
-rw-r--r--  1 berov berov   53 18 юни 16:03 execing-processes.hash
-rw-r--r--  1 berov berov  666 18 юни 16:01 execing-processes.sh

# Забележете, че Го не предлага класическата в Юникс функция `fork`. Обикновено
# това не създава пречки, понеже *гозадачите*, *пораждането* и *изпълнението*
# на процеси покриват повечето от случаите, в които се ползва `fork`.
