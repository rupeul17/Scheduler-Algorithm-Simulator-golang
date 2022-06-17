package lib

import "fmt"

type Job struct {
	name         string
	arrival_time int
	service_time int
	next
}

func init() {
	fmt.Println("lib package > init start!")
}
