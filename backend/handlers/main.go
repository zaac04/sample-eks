package handlers

import (
	models "backend/Models"
	"backend/db"
	"backend/structs"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func AddTodo( w http.ResponseWriter, r *http.Request){
	var data structs.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)

	if err!=nil {
		w.WriteHeader(400)
		w.Write([]byte("error decoding json"))
	}
	id := uuid.New().String()
	response:=db.Pg.Client.Create(models.Todo{
		Id:id,
		Todo:data.Entry,
	})
	if response.Error !=nil{
		w.WriteHeader(500)
		w.Write([]byte("error adding"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("success!"))
}

func GetTodos( w http.ResponseWriter, r *http.Request){
	var Todo []models.Todo
	response:=db.Pg.Client.Find(&Todo)

	if response.Error!=nil{
		w.WriteHeader(500)
		w.Write([]byte("error retrieving data"))
		return
	}
	res,_:=json.Marshal(Todo)
	w.WriteHeader(200)
	w.Write(res)
	return
}


