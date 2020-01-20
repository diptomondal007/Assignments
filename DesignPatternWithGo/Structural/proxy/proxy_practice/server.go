package main

import "fmt"

type server interface {
	handleRequest(string, string) (int, string)
}

type nginx struct {
	application *application
	maxAllowedRequest int
	rateLimiter map[string]int
}

type application struct {

}

func newNginxServer() *nginx{
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter: make(map[string]int),
	}
}

func (n *nginx) handleRequest(url string, method string) (int, string){
	allowed := n.checkRateLimiting(url)
	if !allowed{
		return 403, "not allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool{
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func (a *application) handleRequest(url string, method string) (int, string){
	if url == "/app/status" && method == "GET"{
		return 200, "ok"
	}
	if url == "/create/account" && method == "POST"{
		return 200, "user created"
	}
	return 404, "not ok"
}

func main(){
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
}