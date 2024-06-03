package processor

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func MakeStressTest(url string, requests int, concurrency int64) {

	t := time.Now()

	ctx := context.Background()

	sem := semaphore.NewWeighted(concurrency)
	wg := &sync.WaitGroup{}

	statusCodeCount := map[int]int{}

	for i := 0; i < requests; i++ {

		sem.Acquire(ctx, 1)
		wg.Add(1)

		go func() {

			defer sem.Release(1)
			defer wg.Done()

			resp, err := http.Get(url)

			_, exist := statusCodeCount[resp.StatusCode]
			if !exist {
				statusCodeCount[resp.StatusCode] = 0
			}
			statusCodeCount[resp.StatusCode]++

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
	fmt.Println("Number of requests with HTTP status 200:", statusCodeCount[http.StatusOK])

	delete(statusCodeCount, http.StatusOK)

	for statusCode, qt := range statusCodeCount {
		ms := fmt.Sprintf("Distribution of other HTTP status codes (such as 404, 500, etc.) %d: Quantity:%d", statusCode, qt)
		fmt.Println(ms)
	}
}
