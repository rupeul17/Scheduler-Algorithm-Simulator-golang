package lib

import (
	"fmt"
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

/*
func queue_arrive_job(sec int, job []Job, size int) string {

	for i := 0; i < size; i++ {
		if job[i].Arrival_Time == sec {
			return job[i].Name
		}
	}

}
*/

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

/*
func Get_Ready_Queue_Length(job[] Job, number_of_jobs int, time_slice int) int {

	longest_job := job[0].Service_Time

	for i := 0; i < number_of_jobs; i++ {
		if job[i].Service_Time > longest_job {
			longest_job = job[i].Service_Time
		}
	}

	for i := 0; longest_job >= 0; longest_job -= power(time_slice, i++) {
		length++
	}

	return length
}
*/

func init() {
	fmt.Println("lib package > init start!")
}

/*
func Get_Total_Length(Job job, int number_of_jobs) {
	length := Job[0].Arrival_Time

	for i := 0; i < number_of_jobs; i++ {
		if i+1 < number_of_jobs {
			if length+job[i].Service_Time < job[i+1].Arrival_Time {
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
*/
