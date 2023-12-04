package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMicro()
	D2P2()
	fmt.Printf("Exec time: %d us\n", time.Now().UnixMicro()-start)
}
