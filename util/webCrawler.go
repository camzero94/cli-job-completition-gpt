package util

import (
	"fmt"
	"github.com/camzero94/cli_job/types"
	"golang.org/x/net/html"
	"net/http"
	"sync"
)


type CrawlerReq struct {
	Job    string   `json:"job"`
	Skills []string `json:"skills"`
	Page   int      `json:"page"`
}

func NewCrawlerReq(job string, skills []string, page int) *CrawlerReq {
	return &CrawlerReq{
		Job:    job,
		Skills: skills,
		Page:   page,
	}
}

type JobListHandler struct {
	JobsList []types.Job104
}

func (c *CrawlerReq) Crawler() []types.Job104 {
	var numPages = c.Page
	var urlList []string
	var total int = 0
	var wg sync.WaitGroup

	for page := 1; page <= numPages; page++ {
		url := fmt.Sprintf("https://www.104.com.tw/jobs/search/?ro=0&keyword=%s&order=1&asc=0&page=%d&mode=s&langFlag=0&langStatus=0&recommendJob=1&hotJob=1", c.Job, page)
		urlList = append(urlList, url)
	}

	ch := make(chan string)
	chJobs := make(chan types.Job104, 20)

	for _, url := range urlList {
		wg.Add(1)
		go ParseHTML(url, ch, chJobs, &wg)
	}

	// Close the channel when all workers have finished
	go func() {
		wg.Wait()
		close(chJobs)
	}()

	// Recieve the Jobs from the chanel
	var jobsRes []types.Job104
	for job := range chJobs {
		jobsRes = append(jobsRes, job)
	}

	//Recieve from Channel
	total = len(jobsRes)
	fmt.Println("Received:", total)
	fmt.Println("========Here=========")
	fmt.Printf("Num Jobs: %d", total)

  for i,job := range jobsRes{
    fmt.Printf("JobLink %d: %s ",i,job.Link)
  }

	return jobsRes
}

func ParseHTML(url string, ch chan<- string, chListJobs chan<- types.Job104, wg *sync.WaitGroup) {

	defer wg.Done()

	var job *types.Job104 = new(types.Job104)
	var j *JobListHandler = new(JobListHandler)
	var linksJobs []string

  var wg2 sync.WaitGroup
	chDocs := make(chan *html.Node, 20)

  wg2.Add(1)
	// Get Parsed *html.Node to be processed by ExtractInfo
	go GetDocParsed(url, ch, chDocs,&wg2)

	//Extract General Info
	j.ExtractInfo(<-chDocs, 0, job)

	// Getting specific info from each job
	for i, job := range j.JobsList {
		if i == 20 {
			break
		}
		linksJobs = append(linksJobs, job.Link)
	}


	for _, link := range linksJobs {
    wg2.Add(1)
		go GetDocParsed(link, ch ,chDocs,&wg2)
	}
  
  go func (){
    wg2.Wait()
    close(chDocs)
  }()
	// Recieve the Parsed docs *html.Node
  for docs:=range chDocs{
		j.ExtractInfoJob(docs, 0, job)
	}

	// Send 20 Jobs to Channel
	for i, job := range j.JobsList {
		if i == 20 {
			break
		}
		chListJobs <- job
	}

}

// Function to Extract the Job104 List  Datatype: Job , Company, Link and Skills for the job
func (j *JobListHandler) ExtractInfo(n *html.Node, depth int, jobGlobal *types.Job104) {
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

func (j *JobListHandler) ExtractInfoJob(n *html.Node, depth int, jobGlobal *types.Job104) {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		j.ExtractInfoJob(c, depth+1, jobGlobal)
	}
}

//Extract general Info Article Html Node -> Company and name of the job
func extractGeneralInfo(n *html.Node, jobGlobal *types.Job104) {
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
func extractLink(n *html.Node, jobGlobal *types.Job104) {
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
func extractDate(n *html.Node, jobGlobal *types.Job104) {
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
func (j *JobListHandler) extractLocExp(n *html.Node, jobGlobal *types.Job104, count int) {
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

// Handle GET  Request URL
func GetDocParsed(url string, ch chan<- string, chJobsHtml chan <- *html.Node, wg *sync.WaitGroup)  {


  fmt.Println("URL:", url)
  defer wg.Done()
	// Get request to url
	// Start Parsing Main HTML
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //Send to Channel
		return 
	}
	// Handle if status code is not 200
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		ch <- fmt.Sprintf("getting %s: %s", url, resp.Status)
		return 
	}
	// Handle if status code is not 200
	doc, err := html.Parse(resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("Erro reading response:", err)
		return 
	}

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("parsing %s as HTML: %v", url, err)
		return 
	}
  chJobsHtml <- doc

}
