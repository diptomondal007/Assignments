package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"usercrudwithmocktest/crud"
	"usercrudwithmocktest/crud/models"
)

type CrudHandler interface {
	GetUsers(w http.ResponseWriter,r *http.Request)
	GetUserByUsername(w http.ResponseWriter,r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	crudService crud.Service
}

type Exception struct {
	StatusCode int `json:"code"`
	Error error `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	pageInt,_ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	users, err := h.crudService.Get(pageInt, limitInt)
	if err != nil{
		res , _ := json.Marshal(Exception{Error: err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	responseBody , _ := json.Marshal(users)
	setupResponse(w, responseBody, http.StatusOK)
}

func (h *handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r,"username")
	user, err := h.crudService.GetByUsername(username)
	if err != nil{
		res , _ := json.Marshal(Exception{StatusCode:http.StatusInternalServerError, Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	res , _ := json.Marshal(&user)
	setupResponse(w,res,http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter,r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		res , _ := json.Marshal(Exception{StatusCode:http.StatusInternalServerError, Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil{
		res , _ := json.Marshal(Exception{Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	err = h.crudService.Create(&user)
	if err != nil {
		res , _ := json.Marshal(Exception{Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	res , _ := json.Marshal(Success{Message:"User created"})
	setupResponse(w,res,http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	rowsAffectedCount, err := h.crudService.Delete(username)
	if err != nil {
		res, _ := json.Marshal(Exception{Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	if rowsAffectedCount < 1 {
		res, _ := json.Marshal(Success{Message: "user not found!"})
		setupResponse(w, res, http.StatusOK)
		return
	}
	res, _ := json.Marshal(Success{Message: "user deleted!"})
	setupResponse(w, res, http.StatusOK)
	return
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r,"username")
	requestBody , err := ioutil.ReadAll(r.Body)
	if err != nil{
		res, _ := json.Marshal(Exception{ Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil{
		res, _ := json.Marshal(&Exception{ Error:err})
		setupResponse(w, res, http.StatusInternalServerError)
		return
	}
	err = h.crudService.Update(username,&user)
	if err != nil{
		res, _ := json.Marshal(map[string]string{"error":err.Error()})
		setupResponse(w, res, http.StatusOK)
		return
	}
	res, _ := json.Marshal(&Success{Message: "user updated"})
	setupResponse(w, res, http.StatusOK)
}

func NewHandler(service crud.Service) CrudHandler{
	return &handler{crudService:service}
}

func setupResponse(w http.ResponseWriter, body []byte, statusCode int){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil{
		log.Println(err)
	}
}