package main

import (
	"fmt"
	"os"
)

func main() {

	var choice string
	number_of_jobs := 0
	total_length := 0
	time_slice_1 := 1
	time_slice_2 := 1
	time_slice_3 := 1

	// 파라미터 입력
	fmt.Println("How many jobs would like to simulate? : ")
	fmt.Scanf("%d", &number_of_jobs)

	// 구조체 배열 선언
	var job [number_of_jobs]Job

	fmt.Printf("Please insert the jobs in order of their arrivals\n")

	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("Insert Arrival Time & Service Time of job[%c]:", 65+i)
		job[i].name = 65 + i
		fmt.Scanf("%d %d", &job[i].arrival_time, &job[i].service_time)
	}

	fmt.Printf("Insert the value of time slice for Round Robin 1: ")
	fmt.Scanf("%d", &time_slice_1)

	fmt.Printf("Insert the value of time slice for Round Robin 2: ")
	fmt.Scanf("%d", &time_slice_2)

	fmt.Printf("Insert the value of time slice for MLFQ 2: ")
	fmt.Scanf("%d", &time_slice_3)
	fmt.Println()

	// 입력받은 값 확인
	fmt.Printf("Job\tArrival Time\tService Time\n")
	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("%c\t%d\t\t%d \n", job[i].name, job[i].arrival_time, job[i].service_time)
	}
	fmt.Printf("Round Robin 1 Time Slice: %d\n", time_slice_1)
	fmt.Printf("Round Robin 2 Time Slice: %d\n", time_slice_2)
	fmt.Printf("MLFQ 2 Time Slice: %d\n", time_slice_3)
	fmt.Println()

	fmt.Printf("Are you insertionss accurate? (y/n): ")
	for {
		fmt.Scanf("%s", &choice)
		if choice[0] == 'y' || choice[0] == 'n' {
			break
		}
	}
	if choice[0] == 'n' {
		fmt.Printf("Ok, Bye\n")
		os.Exit(0)
	} else {
		fmt.Printf("Initiating Scheduler...\n\n")
	}

	total_length = Get_Total_Length(job, number_of_jobs)

	// FIFO
	FIFO(number_of_jobs, total_length, job, Result)
	// Round Robin 1
	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_1)
	// Round Robin 2
	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_2)
	// SJF
	Shortest_Job_First(job, Result, total_length, number_of_jobs)
	// STCF
	Shortest_To_Completion_First(job, Result, total_length, number_of_jobs)
	// MLFQ 1
	MLFQ(job, Result, total_length, number_of_jobs, 1)
	// MLFQ 2
	MLFQ(job, Result, total_length, number_of_jobs, time_slice_3)
	// 결과 출력
	Print_Result(Result, number_of_jobs, total_length)
}
