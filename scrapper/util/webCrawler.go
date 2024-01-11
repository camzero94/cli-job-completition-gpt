package util

import (
	"fmt"
	"github.com/camzero94/cli_job/scrapper/types"
)

type JobListHandler struct {
	JobsList []types.Job104
}

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

func (c *CrawlerReq) Crawler() ([]types.Job104, error) {
	var numPages = c.Page
	var urlList []string
	var total int = 0
	j := new(types.JobListHandler)

	// Pagination for API fetching numPages
	url := fmt.Sprintf("https://www.104.com.tw/jobs/search/?ro=0&keyword=%s&order=1&asc=0&page=%d&mode=s&langFlag=0&langStatus=0&recommendJob=1&hotJob=1", c.Job, numPages)
	urlList = append(urlList, url)

	err := GeneralHtmlHandler(url, j)
	if err != nil {
		return nil, fmt.Errorf("Error in GeneralHtmlHandler: %v", err)
	}

	//Recieve from Channel
	total = len(j.JobsList)
	fmt.Println("Received:", total)
	fmt.Println("========Here=========")
	fmt.Printf("Num Jobs: %d", total)

	return j.JobsList, nil
}
