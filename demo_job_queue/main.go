package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

// BW background worker
type BW struct {
	id          uint64
	queue       []func()  // list of functions
	concurrency uint      // no concurrent task
	jobs        chan bool // jobs
	stop        chan os.Signal
}

func New(concurrency uint) *BW {
	return &BW{
		queue:       make([]func(), 0),
		concurrency: concurrency,
		jobs:        make(chan bool, concurrency),
		stop:        make(chan os.Signal, 1),
	}
}

func (b *BW) Add(f func()) *BW {
	b.queue = append(b.queue, f)
	return b
}

func (b *BW) Run() {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 34))
	fmt.Println("| Background job worker started...|")
	fmt.Println(strings.Repeat("-", 34))

	signal.Notify(b.stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {

		case <-b.stop:
			fmt.Println()
			fmt.Println(strings.Repeat("-", 40))
			fmt.Println("| Shutting down background job worker! |")
			fmt.Println(strings.Repeat("-", 40))

			close(b.jobs)
			close(b.stop)
			os.Exit(0)

		default:
			// pick first job and assign to worker
			if len(b.queue) > 0 {
				b.jobs <- true
				f := b.queue[0]       // assign first task to worker
				b.queue = b.queue[1:] // remove first task from queue

				go func(b *BW, f func()) {
					atomic.AddUint64(&b.id, 1)
					log.Println("Task ID:", b.id)

					f()

					<-b.jobs
				}(b, f)
			}
			//slow down little bit
			time.Sleep(100 * time.Millisecond)
		}
	}
}

var b = New(2)

func main() {
	http.HandleFunc("/", handle)
	// run background worker
	go b.Run()
	log.Println("Listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	duration := 20 * time.Second
	log.Println("Enqueue task...")
	b.Add(func() {
		name := r.URL.Query().Get("name")
		log.Println("processing:", name)
		time.Sleep(duration)
		log.Println("completed:", name)
	})
	log.Println("Task enqueued...")
	w.Write([]byte("accepted..."))
}
