package util

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"sync"
)

type CrawlerReq struct {
	Job    string   `json:"job"`
	Skills []string `json:"skills"`
	Page   int      `json:"page"`
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

	var wg sync.WaitGroup

	for page := 1; page <= numPages; page++ {
		url := fmt.Sprintf("https://www.104.com.tw/jobs/search/?ro=0&keyword=%s&order=1&asc=0&page=%d&mode=s&langFlag=0&langStatus=0&recommendJob=1&hotJob=1", c.Job, page)
		fmt.Println("URL:\t", url)
		urlList = append(urlList, url)
	}

	ch := make(chan string)
	chJobs := make(chan Job104, 20)

	for _, url := range urlList {
		wg.Add(1)
		go ParseHTML(url, ch, chJobs, &wg)
	}

  // Close the channel when all workers have finished
	go func() {
		wg.Wait()
		close(chJobs) 
	}()

	var jobsRes []Job104
	for job := range chJobs {
		fmt.Println("Here")
		jobsRes = append(jobsRes, job)
	}
	total = len(jobsRes)
	fmt.Println("Received:", total)

	//Recieve from Channel
	fmt.Println("========Here=========")
	PrintJobs(jobsRes)
	fmt.Printf("Num Jobs: %d", total)
	return listUrlRet
}

func ParseHTML(url string, ch chan<- string, chListJobs chan<- Job104, wg *sync.WaitGroup) {

	defer wg.Done()
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

	// htmlContent , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Erro reading response:", err)
		return
	}
	doc, err := html.Parse(resp.Body)
	j.ExtractInfo(doc, 0, job)
	for i, job := range j.JobsList {
		if i == 20 {
			break
		}
		fmt.Println("here")
		chListJobs <- job
	}
	fmt.Println(len(j.JobsList))
	fmt.Println("===================================================")
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("parsing %s as HTML: %v", url, err)
		return
	}

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
		j.extractLocExp(n, jobGlobal, count)
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
			jobGlobal.Link = temp[0]
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
func (j *JobListHandler) extractLocExp(n *html.Node, jobGlobal *Job104, count int) {
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
