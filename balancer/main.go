package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {

	serverEndpoints := []string{"http://localhost:8001/", "http://localhost:8002/", "http://localhost:8003/"}

	fmt.Println("the server are started")

	for i := 0; i < len(serverEndpoints); i++ {
		fmt.Println("server is started at port ", serverEndpoints[i])
	}

	loadBalancerInstance := &LoadBalancer{
		serverEndpoints: serverEndpoints,
		counter:         0,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("the current  counter is ", loadBalancerInstance.counter)

		targetServer := loadBalancerInstance.counter

		if targetServer == 0 {
			hitServer(targetServer, w, serverEndpoints)

		} else if targetServer == 1 {
			hitServer(targetServer, w, serverEndpoints)

		} else if targetServer == 2 {
			hitServer(targetServer, w, serverEndpoints)

		} else {
			fmt.Println("the required server is not found ")
		}
		setCounter(loadBalancerInstance)
	})

	startServer("8000")

}

func setCounter(l *LoadBalancer) {
	if l.counter == 0 {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.counter++
	} else if l.counter == 1 {
		l.mu.Lock()
		defer l.mu.Unlock()

		l.counter++
	} else if l.counter == 2 {
		l.mu.Lock()
		defer l.mu.Unlock()

		l.counter = 0
	} else {
		fmt.Println("sorry the counter is something else")
	}

}

func hitServer(i int, w http.ResponseWriter, endpoints []string) {
	res, err := http.Get(endpoints[i])
	if err != nil {
		fmt.Println("server occured at this", endpoints[i])
		fmt.Println("error form server is ", err.Error())
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("sorry error reading response body", err.Error())
	}
	fmt.Println("the response from server is ", string(data))

	dataToSend := string(data)

	fmt.Fprint(w, dataToSend)

}

func startServer(port string) {
	fmt.Println("our load balancer is starting at port 8000")
	fullPort := ":" + port
	http.ListenAndServe(fullPort, nil)
}

type LoadBalancer struct {
	serverEndpoints []string
	counter         int
	mu              sync.Mutex
}
