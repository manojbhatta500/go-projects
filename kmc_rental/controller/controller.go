package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/rental.com/models"
)

type stateManager struct {
	count int
	mu    sync.Mutex
}

func ConcurrentApiReqest(targetCount int, w http.ResponseWriter) {

	ch := make(chan struct{}, 10000)

	var wg sync.WaitGroup
	manager := stateManager{}

	sendingData := models.LoginRksPostBody{
		Mobile:   "9822706345",
		Password: "123456",
	}
	for i := 0; i <= targetCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- struct{}{}

			jsonData, err := json.Marshal(sendingData)
			if err != nil {
				log.Printf("Error creating request: %v", err)
			}
			var endpoint string = "http://182.93.65.165:28366/api/v1/auth/login"

			req, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))

			if err != nil {
				log.Printf("Error creating request: %v", err)
				return
			}
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				log.Printf("Error reading response body: %v", err)
			} else {
				fmt.Printf("id : %d : statuscode %d : response body: %s\n", i, req.StatusCode, string(bodyBytes))
			}
			defer req.Body.Close()
			fmt.Println(req.Body)
			fmt.Println("id : " + fmt.Sprint(i) + " : statuscode" + fmt.Sprint(req.StatusCode))
			manager.mu.Lock()
			manager.count++
			manager.mu.Unlock()
			<-ch
		}(i)
	}
	wg.Wait()

	if err := json.NewEncoder(w).Encode(models.ErrorMessage{
		Message: "we have hitted api in this many times " + fmt.Sprint(targetCount),
	}); err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "something went wrong please try again" + err.Error(),
		})
		return
	}

	fmt.Println("concurrent api request completed  count : " + fmt.Sprint(targetCount))

}
