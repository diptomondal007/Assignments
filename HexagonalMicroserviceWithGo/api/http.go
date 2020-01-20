package api

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"hexagonalmicroservicewithgo/json"
	"hexagonalmicroservicewithgo/shortener"
	"io/ioutil"
	"log"
	"net/http"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortener.RedirectService
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	log.Println(code)
	redirect, err := h.redirectService.Find(code)
	log.Println(redirect)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	redirect, err := h.serializer(contentType).Decode(requestBody)
	log.Println(redirect)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = h.redirectService.Store(redirect)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(redirect)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)

}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func(h *handler) serializer(contentType string) shortener.RedirectSerializer {
	return &json.Redirect{}
}
