package processor

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type statusCodeCounter struct {
	sync.Mutex
	statusCodeCount map[int]int
}

func MakeStressTest(url string, requests int, concurrency int64) {

	t := time.Now()

	ctx := context.Background()

	sem := semaphore.NewWeighted(concurrency)
	wg := &sync.WaitGroup{}

	scc := &statusCodeCounter{
		statusCodeCount: map[int]int{},
	}

	for i := 0; i < requests; i++ {
		sem.Acquire(ctx, 1)
		wg.Add(1)

		go func() {
			defer sem.Release(1)
			defer wg.Done()

			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}

			client := http.Client{
				Transport: tr,
			}

			resp, err := client.Get(url)
			if err != nil {
				slog.Error("unable to make request", "error", err)
				return
			}

			scc.Lock()
			_, exist := scc.statusCodeCount[resp.StatusCode]
			if !exist {
				scc.statusCodeCount[resp.StatusCode] = 0
			}
			scc.statusCodeCount[resp.StatusCode]++
			scc.Unlock()

			if err != nil {
				slog.Error("could not make request", "url", url, "error", err)
				return
			}
		}()

	}
	wg.Wait()

	total := time.Since(t)

	fmt.Println("Total time spent on execution:", total)
	fmt.Println("Total number of requests:", requests)
	fmt.Println("Number of requests with HTTP status 200:", scc.statusCodeCount[http.StatusOK])

	delete(scc.statusCodeCount, http.StatusOK)

	for statusCode, qt := range scc.statusCodeCount {
		ms := fmt.Sprintf("Distribution of other HTTP status codes (such as 404, 500, etc.) %d: Quantity:%d", statusCode, qt)
		fmt.Println(ms)
	}
}
