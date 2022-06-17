package main

import "fmt"

func main() {

	number_of_jobs := 0

	// 파라미터 입력
	fmt.Println("How many jobs would like to simulate? : ")
	fmt.Scanf("%d", &number_of_jobs)

	fmt.Printf("Please insert the jobs in order of their arrivals\n")

	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("Insert Arrival Time & Service Time of job[%c]:", 65+i)
	}
}
