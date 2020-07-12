package main

import (
	"fmt"
	"github.com/jaypipes/ghw"
	//	"log"
	"net/http"
	//"reflect"
	"runtime"
	"strings"
)

type DeviceType string

const (
	MOBILE DeviceType = "Mobile"
	TABLET DeviceType = "Tab"
	WEB    DeviceType = "Web"
	TV     DeviceType = "TV"
)

func GetType(r *http.Request) DeviceType {

	if isUserAgent(r, "Android", "webOS", "iPhone", "BlackBerry", "Windows Phone") {
		return MOBILE
	}
	if isUserAgent(r, "iPad", "iPod", "tablet", "RX-34", "FOLIO") ||
		(isUserAgent(r, "Kindle", "Mac OS") && isUserAgent(r, "Silk")) ||
		(isUserAgent(r, "AppleWebKit") && isUserAgent(r, "Silk")) {
		return TABLET
	}
	if isUserAgent(r, "TV", "NetCast", "boxee", "Kylo", "Roku", "DLNADOC") {
		return TV
	}

	return WEB
}

func isUserAgent(r *http.Request, userAgents ...string) bool {
	userAgent := r.Header.Get("User-Agent")
	for _, v := range userAgents {
		if strings.Contains(userAgent, v) {
			return true
		}
	}
	return false
}

//return operating system model
func GOOSS() string {
	os := runtime.GOOS
	return os
}

//It will return your System memory
func SystemMemory() (string, int64) {
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}
	phys := memory.TotalPhysicalBytes
	usable := memory.TotalUsableBytes
	return memory.String(), (phys - usable)
}

//it will return system cpu informatation
func SystemCpu() (*ghw.CPUInfo, []*ghw.ProcessorCore) {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}
	var foo []*ghw.ProcessorCore
	for _, proc := range cpu.Processors {
		for _, core := range proc.Cores {
			foo = append(foo, core)
			//fmt.Printf("  %v\n", core)
		}
	}
	return cpu, foo
}

func BlockStorage() (*ghw.BlockInfo, []*ghw.Disk, []*ghw.Partition) {
	var partion []*ghw.Partition
	var Disk []*ghw.Disk
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}
	for _, disk := range block.Disks {
		Disk = append(Disk, disk)
		//log.Println(Disk)
		for _, part := range disk.Partitions {
			partion = append(partion, part)
		}
	}
	return block, Disk, partion
}
