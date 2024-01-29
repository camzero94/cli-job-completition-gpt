package api

import (
	"encoding/json"
	"fmt"
	"github.com/camzero94/cli_job/scrapper/cache/db"
	"github.com/camzero94/cli_job/scrapper/types"
	"github.com/camzero94/cli_job/scrapper/util"
	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	JobName string   `json:"jobName"`
	Company string   `json:"company"`
	Skills  []string `json:"skills"`
	Link    string   `json:"link"`
	Content string   `json:"content"`
	Exp     string   `json:"exp"`
	Date    string   `json:"date"`
	Loc     string   `json:"loc"`
}

type Server struct {
	listenaddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenaddr: listenAddr,
	}
}

//Start Server endpoint handler functions ->  Return
func (s *Server) Start() error {
	http.HandleFunc("/getJobs", s.handlerGetJobs)
	
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	// Insert the middleware
	handler := cors.Default().Handler(http.DefaultServeMux)
	handler = c.Handler(handler)
	return http.ListenAndServe(s.listenaddr, handler)
}

//Handler Function Middleware
func (s *Server) handlerGetJobs(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	myJob, ok := values["myJob"]
	if !ok || myJob[0] == "" {
		myError := fmt.Sprintf("Missing the Job in the query URL.")
		json.NewEncoder(w).Encode(myError)
		return
	}
	skills, ok := values["skills"]
	if !ok || len(skills) == 0 {
		myError := fmt.Sprintf("Missing the Skills Set of the query URL.")
		json.NewEncoder(w).Encode(myError)
		return
	}
	pages, ok := values["pages"]
	if !ok || pages[0] == "" || pages[0] == "0" {
		myError := fmt.Sprintf("Missing the Pages Depth you want to  the query at 104.")
		json.NewEncoder(w).Encode(myError)
		return
	}

	//Create Crawler from customer variables job, skills, and depth pages
	job := myJob[0]
	pagination, err := strconv.Atoi(pages[0])
	if err != nil {
		myError := fmt.Sprintf("Error COnverting pages to Integer.")
		json.NewEncoder(w).Encode(myError)
		return
	}
	// Create Redis Client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "changeme",
		DB:       0,
	})

	// Create Custom Redis Cache
	ttl := time.Second * 4
	new_redis_cache := db.NewRedisCache(client, ttl)

	// Create Crawler Object
	req := util.NewCrawlerReq(job, skills, pagination)

	// Create Storage Object with custom Redis Cache
	storage, _ := db.SearchEngine(new_redis_cache, req)

	// Create simple redis Key
	key := fmt.Sprintf("%s:%d", job, pagination)
	jobsJsonString, err := storage.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	var jobs []types.Job104
	var answer []Response
	err = json.Unmarshal([]byte(jobsJsonString), &jobs)
	if err != nil {
		log.Fatal(err)
	}
	for _, job := range jobs {
		res := &Response{
			JobName: job.JobName,
			Company: job.Company,
			Link:    job.Link,
			Content: job.Content,
			Skills:  job.Skills,
			Date:    job.Date,
			Exp:     job.Exp,
			Loc:     job.Location,
		}
		answer = append(answer, *res)
	}
	json.NewEncoder(w).Encode(answer)
}
