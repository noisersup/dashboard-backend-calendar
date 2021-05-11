package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetResponse struct{
	Tasks	[]Task `json:"tasks"`
	Events	[]Event `json:"events"`
	Error	string `json:"error"`
}

type Task struct{
	ID			primitive.ObjectID `bson:"_id" json:"id"`
	Title		string `json:"title"`
	Desc		string `json:"desc"`
	Done		bool `json:"done"`
	Order		int `json:"order"`
	Due			int `json:"due"`
}

type Event struct{
	ID			primitive.ObjectID `bson:"_id" json:"id"`
	Title		string `json:"title"`
	Desc		string `json:"desc"`
	Order		int `json:"order"`
	From		int `json:"from"`
	Due			int `json:"due"`
}

type TasksResponse struct{
	Tasks	[]Task `json:"tasks"`
	Error	string `json:"error"`
}