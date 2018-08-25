package main

import ( 
	"fmt" 
	"os/exec" 
)

func main() {
	output, err := exec.Command("/bin/ls").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
}
