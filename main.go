package main 

import (
  "flag"
  "log"
  "github.com/camzero94/cli_job/api"
)
func main(){
  listeAddr := flag.String("listenaddr",":3000","the server address")
  flag.Parse()

  println("Listening on ",*listeAddr)
  server := api.NewServer(*listeAddr)
  //log.Fata will trigger if there is an error
  log.Fatal(server.Start())



}
