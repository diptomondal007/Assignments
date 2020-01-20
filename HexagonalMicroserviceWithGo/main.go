package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	h "hexagonalmicroservicewithgo/api"
	mr "hexagonalmicroservicewithgo/repository/mongo"
	//rr "hexagonalmicroservicewithgo/repository/redis"

	"hexagonalmicroservicewithgo/shortener"
)

// https://www.google.com -> 98sj1-293
// http://localhost:8000/98sj1-293 -> https://www.google.com

// repo <- service -> serializer  -> http

func main() {
	repo := chooseRepo("mongo")
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "8000"
	//if os.Getenv("PORT") != "" {
	//	port = os.Getenv("PORT")
	//}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo(dbType string) shortener.RedirectRepository {
	switch dbType {
	//case "redis":
	//	redisURL := os.Getenv("REDIS_URL")
	//	repo, err := rr.NewRedisRepository(redisURL)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	return repo
	case "mongo":
		mongoURL := "mongodb://localhost:27017"
		mongodb := "shortener"
		mongoTimeout, _ := strconv.Atoi("200")
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
