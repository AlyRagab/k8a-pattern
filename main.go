package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", httpRequest)
	r.HandleFunc("/metrics", metrics).Handler(promhttp.Handler())
	//r.Path("/metrics").HandlerFunc(metrics).Handler(promhttp.Handler())
	go http.ListenAndServe(":8080", r)
	printUsage()
}

// printUsage outputs the Total obtained memory usage abd the current, total and OS memory being used.
func printUsage() {

	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		generateRandomNumbers()

		fmt.Printf("Totall Heap Memory \t= %v MB\n", byteToMb(m.HeapSys))
		fmt.Printf("Used Memory \t\t= %v Byte\n", m.HeapInuse)
		fmt.Printf("Free Memory \t\t= %v MB\n", byteToMb(m.HeapReleased))
		fmt.Printf("Processors \t\t= %v\n", runtime.NumCPU())

		fmt.Printf("=======\n")
		time.Sleep(10 * time.Second)
	}
}

func byteToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func generateRandomNumbers() float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("Generated Number \t= %v \n", r1.Float64())
	return r1.Float64()

}
func httpRequest(w http.ResponseWriter, r *http.Request) {
	gen := generateRandomNumbers()
	fmt.Fprintf(w, "Generated Number is : %v\n", gen)
	fmt.Fprintf(w, "Pattern : Predictable Demands\n")
	//w.Write([]byte(fmt.Sprintf("Generated Number is : %f", gen)))
}

func metrics(w http.ResponseWriter, r *http.Request) {
	// Counting the numer of request made in this API
	if r.URL.Path == "/metrics" {
		requestCounter := promauto.NewCounter(prometheus.CounterOpts{
			Name: "requests_counter",
			Help: "The Number of http request made against this API.",
		})
		requestCounter.Inc()
		fmt.Fprint(w, requestCounter)
	}
}
