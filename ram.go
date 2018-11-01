package main

import (
	"strconv"
	"log"
	"bufio"
	"os"
	"strings"
	"errors"
)

// RAM - data structure to represent memory stats
type RAM struct{}



func (r *RAM) readProcMeminfo(attr string) (float64, error) {
	memoryReader, err := os.Open("/proc/meminfo")
	if err != nil{
		log.Fatal(err)
	}
	defer memoryReader.Close()
	scanner := bufio.NewScanner(memoryReader)
	for scanner.Scan(){
		line := scanner.Text()
		lineFeilds := strings.Fields(line)
		key:= strings.Trim(lineFeilds[0],"\t :")
		if key == attr {
			value := strings.Trim(lineFeilds[1],"\t :")
			result,err := strconv.ParseFloat(value,64)
			if err != nil{
				log.Fatal(err)
			}
		return result, nil
		}
	}

	err = errors.New("Couldn't find attribute " + attr)
	return -1, err
}

// Total - get total memory in specific untits (KB,MB,GB,TB)
func(r RAM) Total(u string) QuantityResult {
	res, err := r.readProcMeminfo("MemTotal")
	if err != nil{
		log.Fatal(err)
	}
	resInUnits:= toUnits(u, res)
	return QuantityResult{ 
		value: resInUnits,
		metric: "MemTotal"}

}
// Free - get free memory in specific untits (KB,MB,GB,TB)
func(r RAM) Free(u string) QuantityResult {
	res, err := r.readProcMeminfo("MemFree")
	if err != nil{
		log.Fatal(err)
	}
	freeInUnits := toUnits(u, res)
	return QuantityResult{ 
		value: freeInUnits,
		metric: "Memfree"}
}
// SwapTotal - get swap total size in specific untits (KB,MB,GB,TB)
func(r RAM) SwapTotal(u string) QuantityResult {
	res, err := r.readProcMeminfo("SwapTotal")
	if err != nil{
		log.Fatal(err)
	}
	swapInUnits := toUnits(u, res)
	return QuantityResult{ 
		value: swapInUnits,
		metric: "SwapTotal" }
}

// SwapFree - get swap total size in specific untits (KB,MB,GB,TB)
func(r RAM) SwapFree(u string) QuantityResult {
	res, err := r.readProcMeminfo("SwapFree")
	if err != nil{
		log.Fatal(err)
	}
	swapFreeInUnits := toUnits(u, res)
	return QuantityResult{ 
		value: swapFreeInUnits,
		metric: "SwapFree" }
}

// Buffers - get buffers total size in specific untits (KB,MB,GB,TB)
// A buffer is something that has yet to be "written" to disk.
func(r RAM) Buffers(u string) QuantityResult {
	res, err := r.readProcMeminfo("Buffers")
	if err != nil{
		log.Fatal(err)
	}
	buffersInUnits := toUnits(u, res)
	return QuantityResult {
		value: buffersInUnits,
		metric: "Buffers"}
}

// Cached - get cached size in specific untits (KB,MB,GB,TB)
// A cache is something that has been "read" from the disk and stored in memory for later use.
func(r RAM) Cached(u string) QuantityResult {
	res, err := r.readProcMeminfo("Cached")
	if err != nil{
		log.Fatal(err)
	}
	cachedInUnits :=  toUnits(u, res)
	return QuantityResult{ 
		value: cachedInUnits,
		metric: "Cached"}
}

