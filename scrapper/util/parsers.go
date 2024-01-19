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

func GetContentSkillsParser(job types.Job104, chContent chan<- string, chJobs chan <- types.Job104 , wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("Job link-------------------->:", job.Link)
	// Create a new slice of skills
	pyCmd := fmt.Sprintf(" && python pyScripts/scrapper.py %s", job.Link)
	cmd := exec.Command("/bin/bash", "-c", "source pyScripts/venv/bin/activate"+pyCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		chContent <- fmt.Sprintf("Error: %v", err)
		return
	}

	skills,content := ExtractSkillsFromContent(string(out))
	job.Skills = skills
	job.Content = content
	fmt.Println("Skills:", job.Skills)
	fmt.Println("Content:", job.Content)
	fmt.Println("====================================================================")
	chJobs <- job
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
	chJobs:= make(chan types.Job104 ,20)
	chContent:= make(chan string )

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
		go GetContentSkillsParser(job,chContent,chJobs, &wg2)
	}
	go func() {
		wg2.Wait()
		fmt.Println("HERRRRRRRRRRRRRRJ Clossing")
		close(chJobs)
	}()
	// counter := 0
	var temp []types.Job104
	for job := range chJobs{
		temp = append(temp, job)
	}
	for i,job:= range temp{
		j.JobsList[i] = job

	}
	// for _,job:= range j.JobsList{
	// 	fmt.Println("Name", job.JobName)
	// 	fmt.Println("Link", job.Link)
	// 	fmt.Println("Content Inside", job.Content)
	// 	fmt.Println("Skills: INside", job.Skills)
	// }
	// for job := range chJobs{
	// 	fmt.Println("HERRRRRRRRRRRRRRJ")
	// 	fmt.Println("Name", job.JobName)
	// 	fmt.Println("Link", job.Link)
	// 	fmt.Println("Content Inside", job.Content)
	// 	fmt.Println("Skills: INside", job.Skills)
	// 	// job.Content = <-chContent
	// 	// job.Skills = <-chSkills
	// 	// j.JobsList[counter].Skills = <-chSkills
	//
	// 	// counter += 1
	// }
	return nil
}

