package util

import (
	"fmt"
	"github.com/camzero94/cli_job/scrapper/types"
	"github.com/camzero94/cli_job/scrapper/constants"
	"golang.org/x/net/html"
	"net/http"
	"os/exec"
	"sync"
	"strings"
)

func ExtractSkillsFromContent(content string) ([]string, string){
	commonSkillsMap:= constants.CreateNewSkills()
	var jobSkills []string
	replacements := map[string]string{
		"\\n": " ",
		"\\n\\n": " ",
		"\\n\\n\\n": " ",
		"/": " ",
		"\\r":  " ",
		",":  " ",
		")":  " ",
		"(":  " ",
		"\\t": " ",
	}
	tokenizeContent:= content
	for oldTok, newTok := range replacements {
		tokenizeContent = strings.ReplaceAll(tokenizeContent, oldTok, newTok)
	}
	fmt.Println("Here tokenizeContent", tokenizeContent)

	tokenizeContentArr := strings.Split(tokenizeContent, " ")
	for _,token:= range tokenizeContentArr{
		token = strings.ToLower(token)
		if val, ok :=commonSkillsMap[token];ok{
			if val == 0 {
				commonSkillsMap[token] = 1
				jobSkills = append(jobSkills, token)		
			} 	
		}
	}
	return jobSkills, tokenizeContent
}

func GetContentSkillsParser(job types.Job104, chContent chan<- string, chSkills chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a new slice of skills
	pyCmd := fmt.Sprintf(" && python pyScripts/scrapper.py %s", job.Link)
	cmd := exec.Command("/bin/bash", "-c", "source pyScripts/venv/bin/activate"+pyCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		chContent <- fmt.Sprintf("Error: %v", err)
		return
	}
	skills,content := ExtractSkillsFromContent(string(out))
	chSkills <- skills	
	chContent <- content
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
	chSkills:= make(chan []string, 40)

	// Get Parsed *html.Node to be processed by ExtractInfo
	doc, err := GetDocHTMLParsed(url)
	if err != nil {
		return err
	}

	//Extract General Info Loc , Name, Company, Link of Given Html
	j.ExtractInfo(doc, 0, job)

	// Iterate over Jobs links to get the content of the job
	for _, job := range j.JobsList {
		wg2.Add(1)
		go GetContentSkillsParser(job, chContent,chSkills, &wg2)
	}
	go func() {
		wg2.Wait()
		close(chContent)
		close(chSkills)
	}()
	counter := 0

	for content := range chContent {
		j.JobsList[counter].Content = content
		counter += 1
	}

	counterIdx := 0
	for skills := range (chSkills){
		fmt.Println(counterIdx, skills)
		j.JobsList[counterIdx].Skills = skills
		counterIdx += 1
	}

	return nil
}

