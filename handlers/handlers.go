package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/noisersup/dashboard-backend-calendar/models"
	"github.com/noisersup/dashboard-backend-calendar/utils"
)

type Server struct{
	ip 			string
	endpoint 	string
}

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request){
	log.Print("GET!") //TODO: remove
	
	response := models.GetResponse{}
	
	reqRes, err := http.Get("http://"+s.ip+s.endpoint)
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