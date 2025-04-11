package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	todoist = "todoist"
)

func execute_add(message string) {

	cmd := exec.Command("todoist", "add", "--project-name", "mac", message)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Print(string(output))
}
func main() {
	var cmd, arg string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}
	switch cmd {
	case "0":
		if arg == "" {
			fmt.Println("Error: argument is not defined")
		} else {
			fmt.Print(arg)
			exec.Command("todoist", "c", arg)
		}
	case "1":
		cmd := exec.Command("todoist", "l")
		output, err := cmd.Output()

		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		fmt.Print(string(output))
	default:
		execute_add(cmd)

	}
	// switch cmd {
	// case "a":
	// 	execute("add", arg)

	// // case "d":
	// // 	exec.Command(todoist, "close")
	// // case "l":
	// // 	exec.Command(todoist, "l")
	// default:
	// 	fmt.Println("Unknown command")
	// }
}
