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

func Print_Result(Result [][]string, number_of_jobs int, total_length int) {

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

		fmt.Fprintf(ResultFile, "\n")
		for j := 0; j < number_of_jobs; j++ {
			fmt.Fprintf(ResultFile, "%s  :       ", string(65+j))
			for k := 0; k < total_length; k++ {
				if Result[k][q] == string(65+j) {
					//if strings.Compare(Result[k][q], string(65+1)) == 0 {
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

func FIFO(number_of_jobs int, total_length int, job []lib.Job, Result [][]string) {
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
			for lib.IsJobDone(temp) == 0 {
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

/*
func Round_Robin(number_of_jobs int, total_length int, job []lib.Job, Result [][] string, time_slice int, version int) {
	
	first_start := job[0].Arrival_Time
	var queue lib.Queue 
	lib.InitQueue(&queue);

	for j := 0; j < number_of_jobs; j++ {
		if(job[j].Arrival_Time == first_start) {
			lib.PushQueue(&queue, &job[j])
		}
	}

	for i := first_start; i < total_length; {
		if lib.IsEmpty(&queue) == 1 {
			i++
			for k := 0; k < number_of_jobs; k++ {
				if job[k].Arrival_Time == i {
					lib.PushQueue(&queue, &job[k])
				}
			}
		} else {
			temp := lib.PopQueue(&queue);
			for real_time_slice := time_slice; real_time_slice && !lib.IsJobDone(temp); real_time_slice-- {
				Result[i][version] = temp.Name
				temp.Service_Time--
				i++
				for k := 0; k < number_of_jobs; k++ {
					if job[k].Arrival_Time == i{
					  lib.PushQueue(&queue, &job[k])
					}
				}
			}
			if lib.IsJobDone(temp) == 0 {
				lib.PushQueue(&queue, temp);
			}
		}
	}
}
*/

 //SJF
 func Shortest_Job_First(job []lib.Job, Result[][] string, total_length int, number_of_jobs int) {

	var executing lib.Job
	var ready_queue lib.Queue	
	lib.InitQueue(&ready_queue);

	executing.name = RN;
	executing.service_time=-1;

	for sec := 0; sec <= total_length; sec++ {
		tmp := lib.queue_arrive_job(sec,job,number_of_jobs);

		if(tmp!=RN)
		{
			PushQueue(&ready_queue, &job[(int)tmp-65]);
		}

		if(executing.service_time==0) 
		{
			executing.name = RN;
			if(sec==total_length)
			{
				break;
			}
		}

		if(executing.name==RN&&ready_queue.count!=0)
		{	
			executing = priority(&ready_queue);
		}
		
		if(executing.name!=RN)
		{
			executing.service_time--;
		}
		
		Result[sec][3]=executing.name;
	}
    }

    //STCF
    void Shortest_To_Completion_First(Job job[], char Result[][7], int total_length, int number_of_jobs){
    int i, j, sec;
	Job executing, executing2;

	Queue ready_queue;	
	initQueue(&ready_queue);

	executing.name = RN;
	executing.service_time = -1;

	for(sec=0; sec<=total_length; sec++)
	{
		char tmp = queue_arrive_job(sec,job,number_of_jobs);

		if(tmp!=RN)
		{
			PushQueue(&ready_queue, &job[(int)tmp-65]);
		}
		
		if(executing.service_time==0) 
		{
			executing.name = RN;

			if(sec==total_length)
			{
				break;
			}
		}
		
		if(executing.name==RN && ready_queue.count!=0)
		{
			executing = priority(&ready_queue);
		}
		
		if(ready_queue.count!=0 && executing.name!=RN)
		{
			executing2 = priority(&ready_queue);
			if(executing.service_time > executing2.service_time)
			{
				PushQueue(&ready_queue, &executing);
				executing = executing2;
			}

			else if(executing.service_time==executing2.service_time)
			{
				if((int)executing.name<(int)executing2.name)
				{
					PushQueue(&ready_queue, &executing2);
				}

				else
				{
					PushQueue(&ready_queue, &executing);
					executing = executing2;
				}
			}

			else
			{
				PushQueue(&ready_queue, &executing2);
			}
		}
		
		if(executing.name!=RN)
		{
			executing.service_time--;
		}

		Result[sec][4]=executing.name;
	}
    }
