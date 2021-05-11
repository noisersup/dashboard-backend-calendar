package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noisersup/dashboard-backend-calendar/handlers"
)

func main(){
	r := mux.NewRouter()
	config := getVars()
	s := handlers.NewServer(config.TasksIP,config.TasksPort,config.TasksEndpoint)

	r.HandleFunc("/",s.GetTasks).Methods("GET")

	log.Printf("Starting http server on port :%d...",config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Port),r))
}

type Config struct{
	TasksIP 		string
	TasksPort 		int
	TasksEndpoint 	string
	Port 			int
}

func getVars() *Config {	
	var config Config

	config.TasksIP = os.Getenv("TASKS_IP")
	config.TasksEndpoint = os.Getenv("TASKS_ENDPOINT")
	config.TasksPort,_ = strconv.Atoi(os.Getenv("TASKS_PORT"))
	config.Port,_ = strconv.Atoi(os.Getenv("PORT")) //default: 8000

	if config.TasksIP=="" || config.TasksPort== 0 || config.TasksEndpoint == ""{ 
		log.Fatal("ENV variables did not set")
	}

	if config.Port == 0 {
		config.Port = 8000
		log.Printf("PORT variable is not set. loading defaults: %d",config.Port)
	}

	return &config
}