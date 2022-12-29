package mutex

import (
	"net/http"
	"sync"
)

type StatusCodes map[string]string

var statusCodes StatusCodes

func statusCode(wg *sync.WaitGroup, mut *sync.Mutex, website string) {
	res, _ := http.Get(website)

	mut.Lock()
	statusCodes[website] = res.Status
	mut.Unlock()

	wg.Done()
}

func WebsiteChecker(websites []string) StatusCodes {
	statusCodes = make(StatusCodes)

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(len(websites))

	for _, website := range websites {
		go statusCode(wg, mut, website)
	}

	wg.Wait()

	return statusCodes
}
