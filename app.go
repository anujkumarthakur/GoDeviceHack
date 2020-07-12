package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	deviceType := GetType(r)
	operSys := GOOSS()
	sysmemory, sysbyte := SystemMemory()
	cpu, foo := SystemCpu()
	block, disk, partion := BlockStorage()

	if operSys == "linux" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}
	if operSys == "windows" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}
	if operSys == "darwin" {
		fmt.Fprintf(w, "<h1>Operating System: %s<h1>", operSys)
	}

	if deviceType == "Mobile" {
		fmt.Fprintf(w, "<h1>Device : Mobile</h1>")
	} else if deviceType == "Web" {
		fmt.Fprintf(w, "<h1>Device : Web</h1>")
	} else if deviceType == "Tab" {
		fmt.Fprintf(w, "<h1>Device : Tablet</h1>")
	}

	fmt.Fprintf(w, "<h1>System Memory: %s<h1>", sysmemory)
	fmt.Fprintf(w, "<h1>Memory Byte Uses: %d<h1>", sysbyte)
	fmt.Fprintf(w, "<h2>CPU Info: %v<h2>", cpu)
	fmt.Fprintf(w, "<h2>Processor Threads: %v<h2>", foo)
	fmt.Fprintf(w, "<h2>Storage : %v<h2>", block)
	fmt.Fprintf(w, "<h2>Disk : %v<h2>", disk)
	fmt.Fprintf(w, "<h2>Partion : %v<h2>", partion)
	elapsed := time.Since(start)
	fmt.Printf("Time take to serve static file %s\n\n", elapsed)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	http.HandleFunc("/", RouteHandler)
	error := http.ListenAndServe(":"+port, nil)
	fmt.Println(error)
}
