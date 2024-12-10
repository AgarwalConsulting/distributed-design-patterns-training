package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var serverIndex int32

// Round-robin load balancer
func loadBalancer(servers []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index := atomic.AddInt32(&serverIndex, 1) % int32(len(servers))
		target := servers[index]

		// Proxy request to the selected server
		resp, err := http.Get(fmt.Sprintf("http://%s%s", target, r.URL.Path))
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Forward response to client
		w.WriteHeader(resp.StatusCode)
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		_, _ = w.Write([]byte(fmt.Sprintf("Response from server: %s\n", target)))
	}
}

func main() {
	servers := []string{"localhost:8081", "localhost:8082", "localhost:8083"}

	http.HandleFunc("/", loadBalancer(servers))
	fmt.Println("Load balancer running on :8080")
	http.ListenAndServe(":8080", nil)
}
