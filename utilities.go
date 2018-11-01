package main

import (
	"os"
	"strings"
	"log"
	"os/exec"
	"bytes"
)

// QuantityResult - data structure that represent comparable value 
// it can be used in tests for easier comparison to values
type QuantityResult struct {
	value float64
	metric string
}

// LowerThan - checks if lower than(want)
func(qr QuantityResult) LowerThan(want float64) bool{
	if qr.value >= want{
		log.Fatalf("ERROR - %s has: %f >= want: %f", qr.metric , qr.value, want)
	}
	log.Printf("OK - %s has %f < want %f", qr.metric , qr.value, want)
	return true
}
// GreaterThan - checks if greater than(want)
func(qr QuantityResult) GreaterThan(want float64) bool{
	if qr.value <= want{
		log.Fatalf("ERROR - %s has: %f <= want: %f", qr.metric, qr.value, want)
	}
	log.Printf("OK - %s has: %f > want: %f", qr.metric, qr.value, want)
	return true
}

// EqualTo - checks if equal to(want)
func(qr QuantityResult) EqualTo(want float64) bool{
	if qr.value !=  want{
		log.Fatalf("ERROR - %s has: %f != want: %f", qr.metric,  qr.value, want)
	}
	log.Printf("OK - %s has: %f == want: %f", qr.metric,  qr.value, want)
	return true
}


func toUnits(units string, value float64) float64 {
	u := strings.ToUpper(units)
	switch u {
		case "KB":
			return  value
		case "MB":
			return  value/1024
		case "GB":
			return value/(1024*1024)
		case "TB":
			return value/(1024*1024*1024)
		default:
			log.Fatal("Not supported units: " + units)
	}
	log.Fatal("Couldn't convert %f to %s", value, units)
	return -1.0
}

func openProcFile(name string) *os.File {
	f, err := os.Open("/proc/" + name)
	if err != nil{
		log.Fatal(err)
	}
	return f
}



func shell(command string) (string, string, error) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil{
		log.Fatalf("ERROR - shell couldn't run command %s, %s", command, err)
	}
    return stdout.String(), stderr.String(), err
}

