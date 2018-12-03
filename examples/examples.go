/*
Package examples shows how to use the go-sanitize package.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.

Author: MrZ
*/
package examples

import (
	"fmt"
	"github.com/mrz1836/go-sanitize"
)

func main() {
	//Run and display
	fmt.Println("Result:", gosanitize.IPAddress(" 192.168.0.1 "))

	// 192.168.0.1
}
