package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/rental.com/controller"
	"github.com/rental.com/models"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server is working")
}

func HandleKmc(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "this is a post request.",
		})
		return
	}

	var jsonData models.Target

	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Invalid JSON: " + err.Error(),
		})
		return
	}

	if jsonData.T <= 0 {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "please target value must be greater then 0",
		})
		return
	}

	fmt.Println("we are strting go routines this many times ", fmt.Sprint(jsonData.T))
	controller.ConcurrentApiReqest(jsonData.T, w)

}

func BruteForceLeapFrog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("brute force leapfrog function running")
	ch := make(chan struct{}, 50)
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "this is a post request.",
		})
		return
	}
	var jsonData models.Target
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Invalid JSON: " + err.Error(),
		})
		return
	}
	if jsonData.T <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.ErrorMessage{
				Message: "target should be more then 0",
			})
		return
	}

	endPoint := "https://www.mofe.gov.np/"

	leapFrogStatus := leapfrogStatusStructure{}
	var wg sync.WaitGroup
	for i := 0; i <= jsonData.T; i++ {
		wg.Add(1)
		fmt.Println("entering id is ", i)
		go func(id int) {
			defer func() {
				ch <- struct{}{}
			}()
			defer wg.Done()
			res, err := http.Get(endPoint)
			if err != nil {
				fmt.Println("the error is ", err.Error())
			}
			fmt.Println("statuscode of request is  id is ", fmt.Sprint(id), " : is  :", fmt.Sprint(res.StatusCode))

			defer res.Body.Close()
			leapFrogStatus.mu.Lock()
			leapFrogStatus.counter++
			leapFrogStatus.mu.Unlock()
			<-ch
		}(i)
	}
	wg.Wait()

	fmt.Println("the total hit api count is ", leapFrogStatus.counter)
}

type leapfrogStatusStructure struct {
	counter int
	mu      sync.Mutex
}
