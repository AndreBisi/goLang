// Package pkg ...
package pkg

var (
	a = ""  //@ diag(`Unicode control character U+0007`)
	b = "" //@ diag(`Unicode control characters`)
	c = "Test	test"
	d = `T
est`
	e = `ZeroβWidth` //@ diag(`Unicode format character U+200B`)
	f = "\u200b"
	g = "π©π½βπ¬" //@ diag(`Unicode control character U+0007`)
	h = "π©π½βπ¬β" //@ diag(`Unicode format and control characters`)
)
