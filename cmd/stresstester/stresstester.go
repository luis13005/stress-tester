package stresstester

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func Tester(url *string, repeticoes, concorrencias int) {

	inicio := time.Now()
	ctx := context.Background()

	repeticoesExecutadas := make(chan int, 1)
	repeticoesExecutadas <- 0

	var wg sync.WaitGroup

	var mu sync.Mutex
	statusCount := make(map[int]int)
	status200 := 0

	if !strings.HasPrefix(*url, "http://") && !strings.HasPrefix(*url, "https://") {
		*url = "https://" + *url
	}

	for i := 0; i < concorrencias; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {

				rpts := <-repeticoesExecutadas
				if rpts == repeticoes {
					repeticoesExecutadas <- rpts
					return
				}
				rpts++
				repeticoesExecutadas <- rpts

				request, err := http.NewRequestWithContext(ctx, http.MethodGet, *url, nil)
				if err != nil {
					fmt.Println(err)
					continue
				}

				result, err := http.DefaultClient.Do(request)
				if err != nil {
					fmt.Println(err)
					continue
				}

				status := result.StatusCode

				mu.Lock()

				if status == 200 {
					status200++
				} else {
					statusCount[result.StatusCode]++
				}

				mu.Unlock()
				result.Body.Close()
			}
		}()
	}

	wg.Wait()
	resultadoFinalRpts := <-repeticoesExecutadas
	tempoTotal := time.Since(inicio)
	fmt.Println("Tempo total: ", tempoTotal)
	fmt.Println("Total de requisições executadas: ", resultadoFinalRpts)
	fmt.Println("status 200: ", status200)

	fmt.Println("statusCount: ", statusCount)
	for status, quantidade := range statusCount {
		if status != 200 {
			fmt.Printf("status %d: -> %d\n", status, quantidade)
		}
	}
}
