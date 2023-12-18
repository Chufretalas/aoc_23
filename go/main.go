package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMicro()
	D14P2()
	fmt.Printf("Exec time: %d us\n", time.Now().UnixMicro()-start)
}
