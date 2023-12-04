package types

import (
	"fmt"
	// "golang.org/x/net/html"
	// "net/http"
	// "sync"
)

type ResponseReq struct {
	Mssg  string `json:"mssg"`
	Error string `json:"error"`
}

type Job104 struct {
	JobName  string   `json:"jobname"`
	Company  string   `json:"company"`
	Content  string   `json:"content"`
	Link     string   `json:"link"`
	Skills   []string `json:"skills"`
	Exp      string   `json:"exp"`
	Date     string   `json:"date"`
	Location string   `json:"location"`
}

type JobListHandler struct {
	JobsList []Job104
}

// //====Strategy Pattern for different job sites====//
//
// //Interface Parent for all job sites Parse()
// type IParser interface {
// 	Parse()
// }
// type Siteparser struct {
// 	Parser IParser
// }
//
// func (prs Siteparser) Parse() {
// 	prs.Parser.Parse()
// }
//
// //Subclass for 104 job site
// //Parse General Info about the jobs
// type Job104MainSite struct {
// 	Jobsite    string
// 	Url        string
// 	Ch         chan<- string
// 	ChListJobs chan<- Job104
// 	Wg         *sync.WaitGroup
// }
//
// func (j *Job104MainSite) Parse() {
// 	defer j.Wg.Done()
// 	j.Jobsite = "Here"
// 	fmt.Println("Parsing 104 Main" + j.Jobsite)
// 	//Get request to url
// 	var job *Job104 = new(Job104)
// 	var jh *JobListHandler = new(JobListHandler)
// 	resp, err := http.Get(j.Url)
// 	if err != nil {
// 		j.Ch <- fmt.Sprint(err) //Send to Channel
// 		return
// 	}
// 	// Handle if status code is not 200
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		j.Ch <- fmt.Sprintf("getting %s: %s", j.Url, resp.Status)
// 		return
// 	}
// 	// Handle if status code is not 200
// 	doc, err := html.Parse(resp.Body)
//
// 	if err != nil {
// 		j.Ch <- fmt.Sprintf("Erro reading response:", err)
// 		return
// 	}
//
// 	// j.ExtractInfo(doc, 0, job)
//
// 	// Send 20 Jobs to Channel
// 	for i, job := range jh.JobsList {
// 		if i == 20 {
// 			break
// 		}
// 		fmt.Println("here")
// 		j.ChListJobs <- job
// 	}
//
// 	resp.Body.Close()
//
// 	if err != nil {
// 		j.Ch <- fmt.Sprintf("parsing %s as HTML: %v", j.Url, err)
// 		return
// 	}
//
// }

//Parse Specific Info about the jobs
type Job104SpecificSite struct {
	Jobsite string
}

func (job Job104SpecificSite) Parse() {
	job.Jobsite = "Inside"
	fmt.Println("Parsing 104 Specific" + job.Jobsite)
}
