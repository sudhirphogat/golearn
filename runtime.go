package main

//check regarding the runtime
//godoc package runtime has all the details
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
