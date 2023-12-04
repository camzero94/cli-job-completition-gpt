
package util

import (
  "fmt"
  "github.com/camzero94/cli_job/types"
)
// Reset Job
func ResetJobs(job *types.Job104) {
	job.Company = ""
	job.JobName = ""
	job.Link = ""
	job.Date = ""
	job.Location = ""
	job.Exp = ""
}

func PrintJobs(jobs []types.Job104) {
	for _, job := range jobs {
		fmt.Println(job)
	}
}
