package util

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

type CrawlerReq struct {
	Job    string   `json:"job"`
	Skills []string `json:"skills"`
	Page   int      `json:"page"`
}

type Job104 struct {
	JobName string   `json:"jobname"`
	Company string   `json:"company"`
	Content string   `json:"content"`
	Link    string   `json:"link"`
	Skills  []string `json:"skills"`
	Exp     int      `json:"exp"`
	Date    string   `json:"date"`
  Location string `json:"location"`
}

type JobListHandler struct {
	JobsList []Job104
}

func NewCrawlerReq(job string, skills []string, page int) *CrawlerReq {
	return &CrawlerReq{
		Job:    job,
		Skills: skills,
		Page:   page,
	}
}

func (c *CrawlerReq) Crawler() []string {

	var numPages = c.Page
	var urlList []string
	var listUrlRet []string
	var total int = 0

	ch := make(chan string)
	chNumJobs := make(chan int)

	for page := 1; page <= numPages; page++ {
		url := fmt.Sprintf("https://www.104.com.tw/jobs/search/?ro=0&keyword=%s&order=1&asc=0&page=%d&mode=s&langFlag=0&langStatus=0&recommendJob=1&hotJob=1", c.Job, page)
		urlList = append(urlList, url)
	}
	for _, url := range urlList {
		go ParseHTML(url, ch, chNumJobs)
	}
	for range urlList {
		total += <-chNumJobs
		fmt.Println("Received:", total)
	}
	//Recieve from Channel
	fmt.Println("========Here=========")
	fmt.Printf("Error: %s", <-ch)
	fmt.Printf("Num Jobs: %d", total)
	return listUrlRet

}

func ParseHTML(url string, ch chan<- string, chNumJobs chan<- int) {
	//Get request to url
	var job *Job104 = new(Job104)
	var j *JobListHandler = new(JobListHandler)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //Send to Channel
		return
	}
	//Handle Status
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		ch <- fmt.Sprintf("getting %s: %s", url, resp.Status)
		return
	}

	doc, err := html.Parse(resp.Body)
	j.ExtractInfo(doc, 0, job)

	printJobs(j.JobsList)
	fmt.Println("===================================================")
	chNumJobs <- len(j.JobsList)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("parsing %s as HTML: %v", url, err)
		return
	}
	ch <- ""

}

// Function to Extract the Job104 List  Datatype: Job , Company, Link and Skills for the job
func (j *JobListHandler) ExtractInfo(n *html.Node, depth int, jobGlobal *Job104) {

//Find Article Html Node
	if n.Type == html.ElementNode && n.Data == "article" {
		for _, attr := range n.Attr {
			if attr.Key == "data-job-name" {
				jobGlobal.JobName = attr.Val
			}
			if attr.Key == "data-cust-name" {
				jobGlobal.Company = attr.Val
			}
		}
	}

//Find Anchor Html Node Link
	if n.Type == html.ElementNode && n.Data == "a" {
		var temp []string
		// fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				temp = append(temp, attr.Val)
			}
			if attr.Val == "js-job-link" {
				jobGlobal.Link = temp[0]
				temp = temp[:0]
				j.JobsList = append(j.JobsList, *jobGlobal)
        resetJobs(jobGlobal)
			}
		}
		// fmt.Println(">")
	}

	// if n.Type == html.ElementNode && n.Data == "span" && n.Attr[0].Val == "b-tit__date" {
 //  }


	for c := n.FirstChild; c != nil; c = c.NextSibling {
		j.ExtractInfo(c, depth+1, jobGlobal)
	}


}


//Zeroed Job
func resetJobs(job *Job104){
    job.Company = ""
    job.JobName = ""
    job.Link = ""
}
func printJobs(jobs []Job104) {
	for _, job := range jobs {
		fmt.Printf("Name Job: %s , Company: %s , Link: %s\n", job.JobName, job.Company, job.Link)
	}
}
