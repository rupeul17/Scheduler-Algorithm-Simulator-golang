package lib

import "fmt"

type Job struct {
	name         string
	arrival_time int
	service_time int
}

func init() {
	fmt.Println("lib package > init start!")
}

/*
func Get_Total_Length(Job job, int number_of_jobs) {
	length := Job[0].arrival_time

	for i := 0; i < number_of_jobs; i++ {
		if i+1 < number_of_jobs {
			if length+job[i].service_time < job[i+1].arrival_time {
				length = job[i+1].arrival_time
			} else {
				length += job[i].service_time
			}
		} else {
			length += job[i].service_time
		}
	}

	return length
}
*/
