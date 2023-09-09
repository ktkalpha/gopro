package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ttacon/chalk"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage : gopro [user name] [project name]")
		fmt.Println("        gopro")
		fmt.Println("Create Go Project.")
		os.Exit(0)
	} else if len(os.Args) >= 3 {
		name := os.Args[1]
		user := os.Args[2]
		fmt.Println(chalk.Reset)
		cmd := exec.Command("mkdir", name)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("go", "mod", "init", "github.com/"+user+"/"+name)
		cmd.Dir = "./" + name
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(chalk.Green, "✓ Create Project Successfully! | go mod : "+"github.com/"+user+"/"+name, chalk.Reset)
	}

}
