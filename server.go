package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

type Server struct {
	classes  map[string]map[string]float64
	students map[string]map[string]float64
}

func (this *Server) InitializeServer(d int64, _d *int64) error {
	this.classes = make(map[string]map[string]float64)
	this.students = make(map[string]map[string]float64)
	return nil
}

func (this *Server) AddStudent(data []string, reply *int64) error {
	student := make(map[string]float64)
	var score float64
	if val, ok := this.classes[data[1]]; ok {
		student = val
	}
	if s, err := strconv.ParseFloat(data[2], 64); err == nil {
		score = s
	}
	student[data[0]] = score
	this.classes[data[1]] = student
	return nil
}

func (this *Server) AddClass(data []string, reply *int64) error {
	class := make(map[string]float64)
	var score float64
	if val, ok := this.students[data[0]]; ok {
		class = val
	}
	if s, err := strconv.ParseFloat(data[2], 64); err == nil {
		score = s
	}
	class[data[1]] = score
	this.students[data[0]] = class

	fmt.Println(this.classes)
	fmt.Println(this.students)
	return nil
}

func (this *Server) StudentAverage(name string, result *float64) error {
	var score float64

	i := 0.0
	if val, ok := this.students[name]; ok {
		for _, score = range val {
			score += score
			i += 1
			fmt.Println(score)
		}
	}

	fmt.Println(score)
	fmt.Println(i)
	*result = score / i
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
