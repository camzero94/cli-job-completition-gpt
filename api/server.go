package api

import(
  "net/http"
  "fmt"
  "encoding/json"
  "github.com/camzero94/cli_job/util"
  "strconv"
)

type Server struct{
  listenaddr string
}

func NewServer( listenAddr string) (*Server) {
  return &Server{
    listenaddr: listenAddr,
  }
}

//Start Server endpoint handler functions ->  Return  
func (s *Server) Start() error{
  http.HandleFunc("/getJobs", s.handlerGetJobs) 
  return http.ListenAndServe(s.listenaddr,nil)
}

//Handler Function Middleware
func (s *Server) handlerGetJobs (w http.ResponseWriter, r *http.Request){

  values:= r.URL.Query()

  myJob,ok:= values["myJob"]
  if !ok || myJob[0] =="" {
    myError := fmt.Sprintf("Missing the Job in the query URL.")
    json.NewEncoder(w).Encode(myError)
    return 
  }
  skills,ok:= values["skills"]
  if !ok || len(skills) == 0{
    myError := fmt.Sprintf("Missing the Skills Set of the query URL.")
    json.NewEncoder(w).Encode(myError)
    return 
  }
  pages,ok:= values["pages"]
  if !ok || pages[0] == "" || pages[0] == "0" {
    myError := fmt.Sprintf("Missing the Pages Depth you want to  the query at 104.")
    json.NewEncoder(w).Encode(myError)
    return 
  }

  //Create Crawler from customer variables job, skills, and depth pages
  job := myJob[0]
  depth,err := strconv.Atoi(pages[0])
  if err != nil{
    myError := fmt.Sprintf("Error COnverting pages to Integer.")
    json.NewEncoder(w).Encode(myError)
    return
  }
  c:= util.NewCrawlerReq(job,skills,depth)
  fmt.Println(c)
  list := c.Crawler()
  fmt.Println(list,"Welcome Data")

  json.NewEncoder(w).Encode(list)



  // errorMssg := strings.Join(myError,",")
  // response := &types.ResponseReq{
  //   Mssg:  mssg,
  //   Error: errorMssg,
  // }
}
