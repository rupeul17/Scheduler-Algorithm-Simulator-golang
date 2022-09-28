package main

import (
	"Scheduling-Algorithm-Simulator-golang/lib"
	"fmt"
	"os"
	"strings"
)

func main() {

	var choice string
	total_length := 0

	/* 프로세스 (Job) 개수 입력 */
	fmt.Println("How many jobs would like to simulate? : ")

	number_of_jobs := input_number()

	job := make([]lib.Job, number_of_jobs)

	/* 프로세스 별 도착 시간, 실행 시간 입력 */
	fmt.Printf("Please insert the jobs in order of their arrivals\n")

	for i := 0; i < number_of_jobs; i++ {
		fmt.Printf("Insert Arrival Time of job[%c]:", rune(65+i))
		job[i].Name = string(rune(65 + i))

		job[i].Arrival_Time = input_number()

		fmt.Printf("Insert Service Time of job[%c]:", 65+i)
		job[i].Service_Time = input_number()
	}

	/* 스케줄링 알고리즘 설정 입력 */
	fmt.Printf("Insert the value of time slice for Round Robin 1: ")
	time_slice_1 := input_number()

	fmt.Printf("Insert the value of time slice for Round Robin 2: ")
	time_slice_2 := input_number()

	fmt.Printf("Insert the value of time slice for MLFQ 2: ")
	time_slice_3 := input_number()

	fmt.Println()

	/* 시뮬레이터 설정 값 확인 */
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
			Result[j][i] = string(rune(0))
		}
	}

	/* FIFO (First In First Out) */
	FIFO(number_of_jobs, total_length, job, Result)

	/* Round_Robin 1 */
	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_1, 1)

	/* Round Robin 2 */
	Round_Robin(number_of_jobs, total_length, job, Result, time_slice_2, 2)

	/* SJF (Shortest Job First) */
	Shortest_Job_First(job, Result, total_length, number_of_jobs)

	/* STCF (Shortest To Completion First) */
	Shortest_To_Completion_First(job, Result, total_length, number_of_jobs)

	/* MLFQ (Multi Level Feedback Queue) 1 */
	MLFQ(job, Result, total_length, number_of_jobs, 1)
	/* MLFQ (Multi Level Feedback Queue) 2 */
	MLFQ(job, Result, total_length, number_of_jobs, time_slice_3)

	/* 결과 출력 */
	Print_Result(Result, number_of_jobs, total_length)
}
