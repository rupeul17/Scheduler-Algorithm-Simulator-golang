package main

import (
	"Scheduling/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func input_string() string {
	rd := bufio.NewReader(os.Stdin)

	TmpStr, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	TmpStr = strings.TrimSpace(TmpStr)

	return TmpStr
}

func input_number() int {
	rd := bufio.NewReader(os.Stdin)

	TmpInt, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	TmpInt = strings.TrimSpace(TmpInt)
	input_num, err := strconv.Atoi(TmpInt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return input_num
}

func Print_Result(Result [][7]string, number_of_jobs int, total_length int) {

	/*
			close(STDOUT_FILENO);
		    open("./Result.txt", O_CREAT|O_WRONLY|O_TRUNC, 0777);
	*/

	ResultFile, err := os.Create("./Result.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Scheduler Simulation Result\n")
	for q := 0; q < 7; q++ {
		switch q {
		case 0:
			fmt.Fprintf(ResultFile, "FIFO\n")
		case 1:
			fmt.Fprintf(ResultFile, "Round Robin 1\n")
		case 2:
			fmt.Fprintf(ResultFile, "Round Robin 2\n")
		case 3:
			fmt.Fprintf(ResultFile, "Shortest Job First\n")
		case 4:
			fmt.Fprintf(ResultFile, "Shortest To Completion First\n")
		case 5:
			fmt.Fprintf(ResultFile, "MLFQ 1\n")
		case 6:
			fmt.Fprintf(ResultFile, "MLFQ 2\n")
		}

		fmt.Fprintf(ResultFile, "Job\\Time 00")
		for i := 0; i < total_length; i++ {
			if (i + 1) < 10 {
				fmt.Fprintf(ResultFile, "===0%d", i+1)
			} else {
				fmt.Fprintf(ResultFile, "===%d", i+1)
			}
		}

		fmt.Printf("\n")
		for j := 0; j < number_of_jobs; j++ {
			fmt.Fprintf(ResultFile, "%s  :       ", string(65+j))
			for k := 0; k < total_length; k++ {
				//if Result[k][q] == string(65+j) {
				if strings.Compare(Result[k][q], string(65+1)) == 0 {
					fmt.Fprintf(ResultFile, "===  ")
				} else {
					fmt.Fprintf(ResultFile, "     ")
				}
			}
			fmt.Fprintf(ResultFile, "\n")
		}
		fmt.Fprintf(ResultFile, "\n\n")
		q++
	}
}

func FIFO(number_of_jobs int, total_length int, job []lib.Job, Result [][7]string) {
	first_start := job[0].Arrival_Time
	var queue lib.Queue
	lib.InitQueue(&queue)

	for i := 0; i < number_of_jobs; i++ {
		if job[i].Arrival_Time == first_start {
			lib.PushQueue(&queue, &job[i])
		}
	}

	for i := first_start; i < total_length; {
		if lib.IsEmpty(&queue) == 1 {
			i++
			for j := 0; j < number_of_jobs; j++ {
				if job[j].Arrival_Time == i {
					lib.PushQueue(&queue, &job[j])
				}
			}
		} else {
			temp := lib.PopQueue(&queue)
			for lib.IsJobDone(temp) == 1 {
				Result[i][0] = temp.Name
				temp.Service_Time--
				i++
				for j := 0; j < number_of_jobs; j++ {
					if job[j].Arrival_Time == i {
						lib.PushQueue(&queue, &job[j])
					}
				}
			}
		}
	}
}
