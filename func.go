package main

import (
	"Scheduling-Algorithm-Simulator-golang/lib"
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
			fmt.Fprintf(ResultFile, "%s  :       ", string(rune(65+j)))
			for k := 0; k < total_length; k++ {
				if Result[k][q] == string(rune(65+j)) {
					fmt.Fprintf(ResultFile, "===  ")
				} else {
					fmt.Fprintf(ResultFile, "     ")
				}
			}
			fmt.Fprintf(ResultFile, "\n")
		}
		fmt.Fprintf(ResultFile, "\n\n")
	}
}

func FIFO(number_of_jobs int, total_length int, job []lib.Job, Result [][]string) {
	first_start := job[0].Arrival_Time
	var queue lib.Queue
	lib.InitQueue(&queue)

	/*
		입력한 job들 중에서 가장 먼저 도착한
		job을 찾아 queue에 넣는다.
	*/
	for i := 0; i < number_of_jobs; i++ {
		if job[i].Arrival_Time == first_start {
			lib.PushQueue(&queue, &job[i])
		}
	}

	/*
		스케줄링 시작부터 끝까지 반복한다.
	*/
	for i := first_start; i < total_length; {

		/*
			queue가 비어있으면,
		*/
		if lib.IsEmpty(&queue) == 1 {
			i++
			/*
				현재 시간에 도착하는 job 이 있으면
				queue에 삽입한다.
			*/
			for j := 0; j < number_of_jobs; j++ {
				if job[j].Arrival_Time == i {
					lib.PushQueue(&queue, &job[j])
				}
			}
		} else {
			/*
				queue 가 비어있지 않으면, queue에서 job을 꺼낸다.
			*/
			temp := lib.PopQueue(&queue)

			/*
				job의 service 가 끝나지 않았으면,
				결과 배열에 기록 후 service time을 차감한다.
			*/
			for lib.IsJobDone(temp) == 0 {
				Result[i][0] = temp.Name
				temp.Service_Time--

				/*
					현재 시간에 도착하는 job이 있으면
					queue에 삽입한다.
				*/
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

func Round_Robin(number_of_jobs int, total_length int, job []lib.Job, Result [][]string, time_slice int, version int) {

	first_start := job[0].Arrival_Time
	var queue lib.Queue
	lib.InitQueue(&queue)

	/*
		입력한 job들 중에서 가장 먼저 도착한
		job을 찾아 queue에 넣는다.
	*/
	for j := 0; j < number_of_jobs; j++ {
		if job[j].Arrival_Time == first_start {
			lib.PushQueue(&queue, &job[j])
		}
	}

	/*
		스케줄링을 반복한다.
	*/
	for i := first_start; i < total_length; {
		/*
			queue가 비어있으면,
		*/
		if lib.IsEmpty(&queue) == 1 {
			/*
				현재 시간에 도착한 job이 있을 경우
				queue에 삽입한다.
			*/
			i++
			for k := 0; k < number_of_jobs; k++ {
				if job[k].Arrival_Time == i {
					lib.PushQueue(&queue, &job[k])
				}
			}
		} else {
			/*
				queue가 비어있지 않으면, job을 꺼낸다.
			*/
			temp := lib.PopQueue(&queue)

			/*
				time_slice 길이만큼 job을 실행한다.
				(service time을 실행한 시간만큼 차감)
			*/
			for real_time_slice := time_slice; real_time_slice > 0 && (lib.IsJobDone(temp) == 0); real_time_slice-- {
				Result[i][version] = temp.Name
				temp.Service_Time--

				/*
					현재 시간에 도착한 job이 있으면,
					queue에 삽입한다.
				*/
				i++
				for k := 0; k < number_of_jobs; k++ {
					if job[k].Arrival_Time == i {
						lib.PushQueue(&queue, &job[k])
					}
				}
			}

			/*
				꺼내서 수행한 job이 다 안끝났다면 (service_time > 0)
				다시 queue에 넣는다.
			*/
			if lib.IsJobDone(temp) == 0 {
				lib.PushQueue(&queue, temp)
			}
		}
	}
}

//SJF
func Shortest_Job_First(job []lib.Job, Result [][]string, total_length int, number_of_jobs int) {

	var executing lib.Job
	var ready_queue lib.Queue
	var sec int

	/*
		ready queue
	*/
	lib.InitQueue(&ready_queue)

	executing.Name = lib.RN
	executing.Service_Time = -1

	/*
		스케줄링을 반복한다.
	*/
	for sec = 0; sec <= total_length; sec++ {

		/*
			현재 시간에 도착한 job이 있는지 확인하여
			있다면 queue에 넣는다
		*/
		tmp := lib.Queue_arrive_job(sec, job, number_of_jobs)
		if tmp != lib.RN {
			lib.PushQueue(&ready_queue, &job[(int(tmp[0])-65)])
		}

		/*
			현재 실행중인 job이 끝났고, 마지막 job 이었으면 끝낸다.
		*/
		if executing.Service_Time == 0 {
			executing.Name = lib.RN
			if sec == total_length {
				break
			}
		}

		/*
			현재 실행중인 job이 없고 queue에 job이 들어있으면
			우선순위를 따져서 job을 가져온다.
		*/
		if executing.Name == lib.RN && ready_queue.Count != 0 {
			executing = lib.Priority(&ready_queue)
		}

		/*
			현재 실행중인 job이 있으면
			service time을 차감하고 진행 상태를 기록한다.
		*/
		if executing.Name != lib.RN {
			executing.Service_Time--
		}

		Result[sec][3] = executing.Name
	}
}

//STCF
func Shortest_To_Completion_First(job []lib.Job, Result [][]string, total_length int, number_of_jobs int) {

	var ready_queue lib.Queue
	var executing, executing2 lib.Job
	lib.InitQueue(&ready_queue)

	executing.Name = lib.RN
	executing.Service_Time = -1

	/*
		스케줄링을 반복한다.
	*/
	for sec := 0; sec <= total_length; sec++ {

		/*
			현재 시간에 도착한 job이 있는지 확인하여
			있다면 queue에 넣는다.
		*/
		tmp := lib.Queue_arrive_job(sec, job, number_of_jobs)
		if tmp != lib.RN {
			lib.PushQueue(&ready_queue, &job[(int)(tmp[0])-65])
		}

		/*
			현재 실행중인 job이 끝났고, 마지막 job 이었으면 끝낸다.
		*/
		if executing.Service_Time == 0 {
			executing.Name = lib.RN

			if sec == total_length {
				break
			}
		}

		/*
			현재 실행중인 job이 없고 queue에 job이 들어있으면
			우선순위를 따져서 job을 가져온다.
		*/
		if (executing.Name == lib.RN) && (ready_queue.Count != 0) {
			executing = lib.Priority(&ready_queue)
		}

		/*
			현재 실행중인 job이 있고 queue에 job이 들어있으면
			우선순위를 따져서 job을 가져온다.
		*/
		if (ready_queue.Count != 0) && (executing.Name != lib.RN) {
			executing2 = lib.Priority(&ready_queue)
			if executing.Service_Time > executing2.Service_Time {
				/*
					새로 가져온 Job의 실행 시간이 실행 중인 Job보다 작으면
					기존에 실행 중이던 Job을 ready queue 에 넣고
					꺼냈던 Job을 실행한다.
				*/
				lib.PushQueue(&ready_queue, &executing)
				executing = executing2
			} else if executing.Service_Time == executing2.Service_Time {
				/*
					새로 가져온 Job의 실행 시간 == 실행 중인 Job의 실행 시간 이면
					이름이 작은 것부터 실행한다.
				*/
				if executing.Name < executing2.Name {
					lib.PushQueue(&ready_queue, &executing2)
				} else {
					lib.PushQueue(&ready_queue, &executing)
					executing = executing2
				}
			} else {
				/*
					기존 실행 Job의 실행 시간이 크면
					꺼냈던 Job은 다시 넣는다.
				*/
				lib.PushQueue(&ready_queue, &executing2)
			}
		}

		if executing.Name != lib.RN {
			/*
				실행 Job이 있으면 실행 시간을 차감한다.
			*/
			executing.Service_Time--
		}

		Result[sec][4] = executing.Name
	}
}

//MLFQ
func MLFQ(job []lib.Job, Result [][]string, total_length int, number_of_jobs int, time_slice int) {

	first_start := job[0].Arrival_Time
	version := 5
	priority := 2 * lib.Get_Ready_Queue_Length(job, number_of_jobs, time_slice)
	ready_queue := make([]lib.Queue, priority)

	/*
		MLFQ 1 결과 칸이 초기화 상태가 아니면
		MLFQ 2 에 기록한다.
	*/
	for i := 0; i < total_length; i++ {
		if Result[i][5] != string(rune(0)) {
			version = 6
			break
		}
	}

	/*
		우선순위 만큼 큐를 init 한다.
	*/
	for i := 0; i < priority; i++ {
		lib.InitQueue(&ready_queue[i])
	}

	/*
		입력된 job 들 중 가장 우선 실행되는 job의 시작시간을 찾는다.
	*/
	for i := 0; i < number_of_jobs; i++ {
		if first_start > job[i].Arrival_Time {
			first_start = job[i].Arrival_Time
		}
	}

	/*

	 */
	for i := first_start; i < total_length; {

		/*
			job 리스트를 탐색해서 가장 먼저 시작되는 job을 찾으면
			그 job을 ready_queue 에 넣는다.
		*/
		for j := 0; (j < number_of_jobs) && (i == first_start); j++ {
			if job[j].Arrival_Time == i {
				lib.PushQueue(&ready_queue[0], &job[j])
			}
		}

		/*
			ready_queue p에 job이 들어있는지 확인한다.
			(job이 있을 경우 break)
		*/
		var k int
		for k = 0; k < priority; k++ {
			if lib.IsEmpty(&ready_queue[k]) == 0 {
				break
			}
		}

		/*

		 */
		if k < priority {
			temp_job := lib.PopQueue(&ready_queue[k])
			time := lib.Power(time_slice, k)
			for l := 0; (l < time) && (lib.IsJobDone(temp_job) == 0); l++ {
				Result[i][version] = temp_job.Name
				i++
				temp_job.Service_Time--
				for j := 0; j < number_of_jobs; j++ {
					if i == job[j].Arrival_Time {
						lib.PushQueue(&ready_queue[0], &job[j])
					}
				}
			}

			if lib.IsJobDone(temp_job) == 0 {
				a_empty := 1
				for l := 0; l < priority; l++ {
					if lib.IsEmpty(&ready_queue[l]) == 0 {
						a_empty = 0
					}
				}
				if (k == priority-1) || (a_empty == 1) {
					lib.PushQueue(&ready_queue[k], temp_job)
				} else {
					lib.PushQueue(&ready_queue[k+1], temp_job)
				}
			}
		} else {
			i++
			for j := 0; j < number_of_jobs; j++ {
				if i == job[j].Arrival_Time {
					lib.PushQueue(&ready_queue[0], &job[j])
				}
			}
		}
	}
}
