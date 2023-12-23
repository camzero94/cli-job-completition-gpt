package util

import (
	"fmt"
	"github.com/camzero94/cli_job/scrapper/types"
	"golang.org/x/net/html"
	"net/http"
	"os/exec"
	"sync"
)

func GetContentSkillsParser(job types.Job104, chContent chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	pyCmd := fmt.Sprintf(" && python pyScripts/scrapper.py %s", job.Link)
	fmt.Println("GetContentSkillsParser ----->", pyCmd)
	cmd := exec.Command("/bin/bash", "-c", "source pyScripts/venv/bin/activate"+pyCmd)

	// Run command
	out, err := cmd.CombinedOutput()
	if err != nil {
		chContent <- fmt.Sprintf("Error: %v", err)
		return
	}
	fmt.Printf("URL", job.Link)
	fmt.Printf("Content", string(out))
	chContent <- string(out)
}

// Handle GET  Request URL
func GetDocHTMLParsed(url string) (*html.Node, error) {
	fmt.Println("URL:", url)
	// Start Parsing Main HTML
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		// ch <- fmt.Sprintf("Erro reading response:", err)
		return nil, err
	}
	return doc, nil
}

func GeneralHtmlHandler(url string, j *types.JobListHandler) error {
	var job *types.Job104 = new(types.Job104)
	var wg2 sync.WaitGroup
	chContent := make(chan string, 40)

	// Get Parsed *html.Node to be processed by ExtractInfo
	doc, err := GetDocHTMLParsed(url )
	if err != nil {
		return err
	}

	//Extract General Info Loc , Name, Company, Link of Given Html
	j.ExtractInfo(doc, 0, job)

	// Iterate over Jobs links to get the content of the job
	for _, job := range j.JobsList {
		wg2.Add(1)
		go GetContentSkillsParser(job, chContent, &wg2)
	}
	go func() {
		wg2.Wait()
		close(chContent)
	}()
	counter := 0
	for content := range chContent {
		j.JobsList[counter].Content = content
		counter += 1
	}
	return nil
}
