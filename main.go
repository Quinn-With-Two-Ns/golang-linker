package main

func tlsNext() int32

func main() {
	if got := tlsNext(); got != 7 {
		panic("first TLS read was not 7")
	}
	if got := tlsNext(); got != 8 {
		panic("second TLS read was not 8")
	}
}
