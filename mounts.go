package main

import(
	"bufio"
	"strings"
	"strconv"
	"log"
)

// Mount - data structure represents 
type Mount struct {
	device string
	point string
	fileSystemType string
	options string
	dump int
	fsckOrder int
}

// Mounts - data structure that contains all the mount points
type Mounts []Mount

func (ms *Mounts) readProcDiskMounts() []Mount {
	mounts := Mounts{}
	mf:= openProcFile("mounts")
	defer mf.Close()
	scanner := bufio.NewScanner(mf)
	for scanner.Scan(){
		line := scanner.Text()
		lineFeilds := strings.Fields(line)
		device := strings.Trim(lineFeilds[0],"\t ")
		point:= strings.Trim(lineFeilds[1],"\t ")
		fileSystemType := strings.Trim(lineFeilds[2],"\t ")
		options := strings.Trim(lineFeilds[3],"\t ")
		dumpS := strings.Trim(lineFeilds[4],"\t ")
		d, err := strconv.Atoi(dumpS)
		if err != nil{
			log.Fatal(err)
		}
		fsckOrderS := strings.Trim(lineFeilds[4],"\t ")
		fsckOrder,err := strconv.Atoi(fsckOrderS)
	
		m := Mount{device,point,fileSystemType,options,d,fsckOrder}
		mounts = append(mounts,m)
	}
	return mounts
}

// HasMountPoint - checks if mount point is pointed to separate device
func (ms *Mounts) HasMountPoint(point string) bool {
	for _, m := range ms.readProcDiskMounts(){
		if m.point == point{
		log.Printf("OK - %s is mounted", point)
			return true
		}
	}
	log.Fatalf("ERRROR - %s not mounted ", point)
	return false
}