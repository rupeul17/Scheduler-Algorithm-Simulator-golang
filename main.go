package main

import (
	"Scheduling/lib"
	"fmt"
	"os"
	"strings"
)

func main() {

	var choice string
	//number_of_jobs := 0
	total_length := 0
	//time_slice_1 := 1
	//time_slice_2 := 1
	//time_slice_3 := 1

	// 파라미터 입력
	fmt.Println("How many jobs would like to simulate? : ")

	number_of_jobs := input_number()

	// 구조체 배열 선언
	job := make([]lib.Job, number_of_jobs)

	fmt.Printf("Please insert the jobs in order of their arrivals\n")

	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("Insert Arrival Time of job[%c]:", 65+i)
		job[i].Name = string(65 + i)

		job[i].Arrival_Time = input_number()

		fmt.Printf("Insert Service Time of job[%c]:", 65+i)
		job[i].Service_Time = input_number()
	}

	fmt.Printf("Insert the value of time slice for Round Robin 1: ")
	time_slice_1 := input_number()

	fmt.Printf("Insert the value of time slice for Round Robin 2: ")
	time_slice_2 := input_number()

	fmt.Printf("Insert the value of time slice for MLFQ 2: ")
	time_slice_3 := input_number()

	//fmt.Printf("OK, Simulator Info Setting... \n")
	//time.Sleep(1 * time.Second)

	fmt.Println()

	// 입력받은 값 확인
	fmt.Printf("Job\tArrival Time\tService Time\n")
	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("%s\t%d\t\t%d \n", job[i].Name, job[i].Arrival_Time, job[i].Service_Time)
	}
	fmt.Printf("Round Robin 1 Time Slice: %d\n", time_slice_1)
	fmt.Printf("Round Robin 2 Time Slice: %d\n", time_slice_2)
	fmt.Printf("MLFQ 2 Time Slice: %d\n", time_slice_3)
	fmt.Println()

	fmt.Printf("Are you insertionss accurate? (yes/no): ")
	for {
		choice = input_string()
		if (strings.Compare(choice, "yes") == 0) || (strings.Compare(choice, "no") == 0) {
			break
		}
	}
	if strings.Compare(choice, "no") == 0 {
		fmt.Printf("Ok, Bye\n")
		os.Exit(0)
	} else {
		fmt.Printf("Initiating Scheduler...\n\n")
	}

	total_length = lib.Get_Total_Length(job, number_of_jobs)
	Result := make([][]string, total_length)
	for i := range Result {
		Result[i] = make([]string, 7)
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < total_length; j++ {
			Result[j][i] = string(0)
		}
	}

	// FIFO
	FIFO(number_of_jobs, total_length, job, Result)

	// Round Robin 1
	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_1, 1)

	// Round Robin 2

	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_2, 2)

	// SJF
	Shortest_Job_First(job, Result, total_length, number_of_jobs)
	/*
		// STCF
		Shortest_To_Completion_First(job, Result, total_length, number_of_jobs)
		// MLFQ 1
		MLFQ(job, Result, total_length, number_of_jobs, 1)
		// MLFQ 2
		MLFQ(job, Result, total_length, number_of_jobs, time_slice_3)
		// 결과 출력
	*/
	Print_Result(Result, number_of_jobs, total_length)

}
