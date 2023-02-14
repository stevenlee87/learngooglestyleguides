package main

func sum(values <-chan int) (out int) {
	for v := range values {
		out += v
	}
	// values must already be closed for this code to be reachable, which means
	// a second close triggers a panic.
	return out
}

func main() {
	ch := make(chan int)
	go sum(ch)
}
