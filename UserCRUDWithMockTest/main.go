package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"usercrudwithmocktest/api"
	"usercrudwithmocktest/conn"
	"usercrudwithmocktest/crud"
	"usercrudwithmocktest/repository/psql"
)

func main(){

	repo, err := psql.NewPsqlRepository(conn.DB())
	if err != nil{
		log.Println(err)
	}
	service := crud.NewCrudService(repo)
	handler := api.NewHandler(service)
	r := chi.NewRouter()
	r.Get("/{page}/{limit}",handler.GetUsers)
	r.Get("/{username}",handler.GetUserByUsername)
	r.Post("/",handler.CreateUser)
	r.Put("/{username}",handler.UpdateUser)
	r.Delete("/{username}",handler.DeleteUser)
	log.Fatalln(http.ListenAndServe(":8000",r))
}
