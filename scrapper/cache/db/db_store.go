package db

import (
	"encoding/json"
	"fmt"
	"github.com/camzero94/cli_job/scrapper/types"
	"github.com/camzero94/cli_job/scrapper/util"
	_ "github.com/lib/pq"
)

type SearchQuery struct {
	data  []types.Job104
	cache Cacher
	req   *util.CrawlerReq
}

func SearchEngine(c Cacher, req *util.CrawlerReq) (*SearchQuery, error) {

	var data []types.Job104
	return &SearchQuery{
		data:  data,
		cache: c,
		req:   req,
	}, nil
}

func (db *SearchQuery) Get(key string) (string, error) {

	// If found on the cache
	val, ok := db.cache.Get(key)
	if ok {

		fmt.Println("Key------------------------",key)
		fmt.Println("====================================")
		fmt.Println("\nnFound on the cache\n")
		fmt.Println("====================================")

		return val, nil
	}

	// If not found on the cache, scrappe 104 website
	data, err := db.req.Crawler()
	if err != nil {
		return "", fmt.Errorf("Error Detected", val)
	}

	// Serialize struct to string
	jobsJson, err := json.Marshal(data)

	// Populate cache
	if err := db.cache.Set(key, string(jobsJson)); err != nil {
		return "", err
	}
	// Return value from database
	return string(jobsJson), nil
}
