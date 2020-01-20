package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"quiethackernews/hn"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	http.HandleFunc("/", handler(numStories))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int) http.HandlerFunc {
	sc := storyCache{
		numStories: numStories,
		duration:   6 * time.Second,
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			<-ticker.C
			temp := storyCache{
				numStories: numStories,
				duration:   6 * time.Second,
			}
			temp.stories()
			sc.mutex.Lock()
			sc.cache = temp.cache
			sc.expiration = temp.expiration
			sc.mutex.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories: stories,
			Time:    time.Since(start),
		}
		log.Printf("%s", time.Since(start))
		err = json.NewEncoder(w).Encode(&data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
		log.Printf("%s", time.Since(start))
	})
}

var (
	cache           []item
	cacheExpiration time.Time
)

type storyCache struct {
	numStories int
	cache      []item
	expiration time.Time
	duration   time.Duration
	mutex      sync.Mutex
}

func (sc *storyCache) stories() ([]item, error) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if time.Now().Sub(sc.expiration) < 0 {
		return sc.cache, nil
	}
	stories, err := getTopItems(sc.numStories)
	if err != nil {
		return nil, err
	}
	sc.expiration = time.Now().Add(sc.duration)
	sc.cache = stories
	return sc.cache, nil
}

func getTopItems(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load up top stories")
	}
	type result struct {
		idx  int
		item item
		err  error
	}
	resultCH := make(chan result)
	for i := 0; i < numStories; i++ {
		go func(idx, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCH <- result{idx: idx, err: err}
			}
			resultCH <- result{idx: idx, item: parseHNItem(hnItem)}
		}(i, ids[i])

	}
	var results []result
	for i := 0; i < numStories; i++ {
		results = append(results, <-resultCH)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})
	var stories []item
	for _, res := range results {
		if res.err != nil {
			continue
		}
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
			if len(stories) >= numStories {
				break
			}
		}
	}
	return stories, nil
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
