package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

func client() {
	var op int64
	var result float64
	var name, _class, score string

	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 3)

	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Call("Server.InitializeServer", op, &op)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Println(" 1) Add student score\n",
			"2) Get student average\n",
			"3) Get general average\n",
			"4) Get class average\n",
			"0) Exit")
		fmt.Scanln(&op)

		switch op {
		//add
		case 1:
			fmt.Println("Name: ")
			scanner.Scan()
			name = scanner.Text()

			fmt.Println("Class: ")
			scanner.Scan()
			_class = scanner.Text()

			fmt.Println("Score: ")
			scanner.Scan()
			score = scanner.Text()

			data[0] = name
			data[1] = _class
			data[2] = score

			err = c.Call("Server.AddStudent", data, &op)
			if err != nil {
				fmt.Println(err)
			}
			err = c.Call("Server.AddClass", data, &op)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Succesfully added")
			}
		//student average
		case 2:
			var name string
			fmt.Println("Name: ")
			scanner.Scan()
			name = scanner.Text()

			err = c.Call("Server.StudentAverage", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(name, " average: ", result)
			}
		//all students average
		case 3:
			err = c.Call("Server.OverallAverage", op, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Overall Average: ", result)
			}
		//class average
		case 4:
			fmt.Println("Class: ")
			scanner.Scan()
			_class = scanner.Text()

			err = c.Call("Server.ClassAverage", _class, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(_class, " average: ", result)
			}
		//exit
		case 0:
			return
		default:
			fmt.Println("Incorrect input")
		}
	}
}

func main() {
	client()
}
