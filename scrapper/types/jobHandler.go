package types
import (
	"golang.org/x/net/html"
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

func ResetJobs(job *Job104){
  job.JobName = ""
  job.Company = ""
  job.Content = ""
  job.Link = ""
  job.Skills = nil
  job.Exp = ""
  job.Date = ""
  job.Location = ""
}

// Function to Extract the Job104 List  Datatype: Job , Company, Link and Skills for the job
func (j *JobListHandler) ExtractInfo(n *html.Node, depth int, jobGlobal *Job104) {
	var count int = 0
	switch {
	case n.Type == html.ElementNode && n.Data == "article":
		extractGeneralInfo(n, jobGlobal)
	case n.Type == html.ElementNode && n.Data == "a":
		extractLink(n, jobGlobal)
	case n.Type == html.ElementNode && n.Data == "span":
		extractDate(n, jobGlobal)
	case n.Type == html.ElementNode && n.Data == "ul":
		j.ExtractLocExp(n, jobGlobal, count)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		j.ExtractInfo(c, depth+1, jobGlobal)
	}

}

//Extract general Info Article Html Node -> Company and name of the job
func extractGeneralInfo(n *html.Node, jobGlobal *Job104) {
	for _, attr := range n.Attr {
		if attr.Key == "data-job-name" {
			jobGlobal.JobName = attr.Val
		}
		if attr.Key == "data-cust-name" {
			jobGlobal.Company = attr.Val
		}
	}
}

//Extract Anchor Html Node Link
func extractLink(n *html.Node, jobGlobal *Job104) {
	var temp []string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			temp = append(temp, attr.Val)
		}
		if attr.Val == "js-job-link" {
			jobGlobal.Link = "https:" + temp[0]
			temp = temp[:0]

		}
	}

}

//Extract Date Html Node
func extractDate(n *html.Node, jobGlobal *Job104) {
	for _, attr := range n.Attr {
		if attr.Val == "b-tit__date" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					jobGlobal.Date = c.Data
				}
			}
		}
	}
}

// Extract Location and Experience
func (j *JobListHandler) ExtractLocExp(n *html.Node, jobGlobal *Job104, count int) {
	pttLoc := "b-list-inline b-clearfix job-list-intro b-content"
OutterLoop:
	for _, attr := range n.Attr {
		if attr.Val == pttLoc {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				// Extract Location and Years Experience
				if c.Type == html.ElementNode && c.Data == "li" {
					for c := c.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.TextNode {
							if count == 0 {
								jobGlobal.Location = c.Data
								count += 1
							} else {
								jobGlobal.Exp = c.Data
								j.JobsList = append(j.JobsList, *jobGlobal)
								ResetJobs(jobGlobal)
								break OutterLoop
							}

						}
					}
				}
			}
		}
	}
}

