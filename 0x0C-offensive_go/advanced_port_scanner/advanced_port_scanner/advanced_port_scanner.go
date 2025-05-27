package advanced_port_scanner

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

func AdvancedPortScanner(numWorkers int, startPort, endPort int, target string) {
    ports := make(chan int, numWorkers)
    results := make(chan int)
    var openPorts []int

    for i := 0; i < numWorkers; i++ {
        go worker(ports, results, target)
    }

    go func() {
        for i := startPort; i <= endPort; i++ {
            ports <- i
        }
    }()

    for i := startPort; i <= endPort; i++ {
        port := <-results
        if port != 0 {
            openPorts = append(openPorts, port)
        }
    }

    close(ports)
    close(results)

    sort.Ints(openPorts)
    fmt.Printf("List of open ports: ")
    fmt.Println(openPorts)
}

func worker(ports, results chan int, target string) {
    for p := range ports {
        conn, err := net.Dial("tcp", target + ":" + strconv.Itoa(p))
        if err != nil {
            results <- 0
            continue
        }
        conn.Close()
        results <- p
    }
}


// steps...
// 1. check if all the command line args are passed correctly
// 2. determine the number of workers (go routines waiting for... well... work)
// 2.1 no. of workers cannot be infinitely large
// 3. use channels to assign a port to  scan to the worker go routines
// 4. once a worker go routine finishes, send result to main thread
// 5. print all the opened pord ports in sorted order
