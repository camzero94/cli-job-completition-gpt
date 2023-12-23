package api

import (
	"encoding/json"
	"github.com/camzero94/cli_job/scrapper/util"
	"net/http"
	"strconv"
	"github.com/camzero94/cli_job/scrapper/cache/db"
	"fmt"
	"github.com/go-redis/redis/v8"
)

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
	return http.ListenAndServe(s.listenaddr, nil)
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
	new_redis_cache := db.NewRedisCache(client)

	// Create Database
	req := util.NewCrawlerReq(job, skills, pagination)
	storage, _ := db.SearchEngine(new_redis_cache,req)

	// json.NewEncoder(w).Encode(data)
}

// func main() {
// 	for i := 0; i < 4; i++ {
// 		val, err := s.Get("1")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(val)
// 	}
//
// }
