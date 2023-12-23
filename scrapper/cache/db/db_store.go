package db

import (
	_ "github.com/lib/pq"
	"fmt"
	"github.com/camzero94/cli_job/scrapper/types"
	"github.com/camzero94/cli_job/scrapper/util"
)

type SearchQuery struct {

	data []types.Job104
	cache    Cacher
	req *util.CrawlerReq
}

func SearchEngine(c Cacher, req *util.CrawlerReq) (*SearchQuery, error) {

	var data []types.Job104
	return &SearchQuery{
		data: data,
		cache:    c,
		req: req,
	}, nil
}

func (db *SearchQuery) Get(key string)(string, error) {

	// If found on the cache
	val, ok:= db.cache.Get(key)			
	if ok{
		// Bursting Cache
		if err := db.cache.Remove(key); err != nil{
			fmt.Println(err)
		}

		fmt.Println("Returned value")
		return val,nil
	}

	// If not found on the cache, scrappe 104 website
	data,err := db.req.Crawler()
	if err != nil{
		return "", fmt.Errorf("Key not found %s ",val) 
	}

	// Populate cache   
	if err := db.cache.Set(key, val); err != nil{
		return "", err
	}

	// Return value from database  
	fmt.Println("Returning val from database")
	return val, nil
}
