package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", httpRequest)
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
		time.Sleep(2 * time.Second)
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
