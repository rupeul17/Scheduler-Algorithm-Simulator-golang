package lib

import (
	"fmt"
)

const (
	RN       = "@"
	MAX_TIME = 1000
)

/* 스케줄링 정보 */
type Job struct {
	Name         string
	Arrival_Time int
	Service_Time int
	Next         *Job
}

/* 큐 */
type Queue struct {
	First *Job
	Last  *Job
	Count int
}

func Power(a int, b int) int {
	c := 1
	for i := 1; i <= b; i++ {
		c *= a
	}
	return c
}

/* 큐 초기화 */
func InitQueue(queue *Queue) {
	queue.First = nil
	queue.Last = nil
	queue.Count = 0
}

/* 큐가 비었는지 확인 */
func IsEmpty(queue *Queue) int {
	if queue.Count == 0 {
		return 1
	} else {
		return 0
	}
}

/* 큐에 삽입 */
func PushQueue(queue *Queue, job *Job) {
	new_job := new(Job)
	new_job.Name = job.Name
	new_job.Arrival_Time = job.Arrival_Time
	new_job.Service_Time = job.Service_Time
	new_job.Next = nil

	if IsEmpty(queue) == 1 {
		queue.First = new_job
	} else {
		queue.Last.Next = new_job
	}

	queue.Last = new_job
	queue.Count++
}

/* 큐에서 삭제 */
func PopQueue(queue *Queue) *Job {
	if IsEmpty(queue) == 1 {
		return nil
	} else {
		returning_job := new(Job)
		returning_job = queue.First
		queue.First = queue.First.Next
		queue.Count--
		return returning_job
	}
}

func Queue_arrive_job(sec int, job []Job, size int) string {

	for i := 0; i < size; i++ {
		if job[i].Arrival_Time == sec {
			return job[i].Name
		}
	}

	return "@"
}

func Priority(queue *Queue) Job {
	pri_name := RN
	time := MAX_TIME

	pri_job_scan := queue.First

	for pri_job_scan != nil {
		if time > pri_job_scan.Service_Time {
			time = pri_job_scan.Service_Time
			pri_name = pri_job_scan.Name
		}

		if pri_job_scan.Next == nil {
			break
		}

		pri_job_scan = pri_job_scan.Next
	}

	pri_job_scan = queue.First
	var back *Job = nil

	for pri_job_scan != nil {
		if pri_name == pri_job_scan.Name {
			if back == nil {
				queue.First = queue.First.Next
			} else {
				back.Next = pri_job_scan.Next
			}

			break
		}

		back = pri_job_scan
		pri_job_scan = pri_job_scan.Next
	}

	if pri_job_scan.Next == nil {
		queue.Last = back
	}

	queue.Count--
	var return_job Job = *pri_job_scan

	return return_job
}

/* */
func IsJobDone(job *Job) int {
	if job.Service_Time == 0 {
		return 1
	} else {
		return 0
	}
}

func Get_Total_Length(job []Job, number_of_jobs int) int {

	length := job[0].Arrival_Time

	for i := 0; i < number_of_jobs; i++ {
		if (i + 1) < number_of_jobs {
			if (length + job[i].Service_Time) < (job[i+1].Arrival_Time) {
				length = job[i+1].Arrival_Time
			} else {
				length += job[i].Service_Time
			}
		} else {
			length += job[i].Service_Time
		}
	}

	return length
}

func Get_Ready_Queue_Length(job []Job, number_of_jobs int, time_slice int) int {

	longest_job := job[0].Service_Time
	length := 0

	for i := 0; i < number_of_jobs; i++ {
		if job[i].Service_Time > longest_job {
			longest_job = job[i].Service_Time
		}
	}

	for i := 0; longest_job >= 0; {
		length++
		longest_job -= Power(time_slice, i)
		i++
	}

	return length
}

func init() {
	fmt.Println("lib package > init start!")
}
