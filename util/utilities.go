
package util

import (
  "fmt"
)
// Reset Job
func ResetJobs(job *Job104) {
	job.Company = ""
	job.JobName = ""
	job.Link = ""
	job.Date = ""
	job.Location = ""
	job.Exp = ""
}
func PrintJobs(jobs []Job104) {
	for _, job := range jobs {
		fmt.Println(job)
	}
}
