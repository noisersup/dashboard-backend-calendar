package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/noisersup/dashboard-backend-calendar/models"
	"github.com/noisersup/dashboard-backend-calendar/utils"
)

type Server struct{
	client		*http.Client
	tasksUri 	*url.URL
}

func NewServer(tasksIp string, tasksPort int, tasksEndpoint string) *Server {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	
	uri,err := url.Parse(fmt.Sprintf("http://%s:%d%s",tasksIp,tasksPort,tasksEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Url: %s",uri)
	return &Server{client: &client, tasksUri: uri}
}

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request){
	log.Print("GET!") //TODO: remove
	
	response := models.GetResponse{}
	
	

	reqRes, err := s.client.Get(s.tasksUri.String())
	if err != nil { 
		log.Print(err) 
		utils.SendResponse(w,response,http.StatusInternalServerError)
		return
	}
	defer reqRes.Body.Close()
	
	tasksResponse := models.TasksResponse{}

	json.NewDecoder(reqRes.Body).Decode(&tasksResponse)
	if tasksResponse.Error != "" {
		log.Print(tasksResponse.Error)
		response.Error = tasksResponse.Error
		utils.SendResponse(w,response,reqRes.StatusCode)
		return
	}
	
	response.Tasks = tasksResponse.Tasks
	utils.SendResponse(w,response,http.StatusOK)
}