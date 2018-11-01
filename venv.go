package main

import(
	"log"
	"fmt"
	"strings"
	"bufio"
)

// Venv - datastructure represents python virtual environment
type Venv struct{
	dir FileSystemObject
}

// NewVenv - Constructs new Venv data structure
func NewVenv(path string) Venv {
	venv := Venv{FileSystemObject{Path: path}}
	return venv
}


func (v *Venv) listRequirements() string {
	pipfreeze := fmt.Sprintf("%s/bin/pip freeze", v.dir.Path)
	stdout,stderr,err := shell(pipfreeze)
	if err != nil {
		log.Fatalf("ERROR - couldn't list venv %s requirements", v.dir.Path)
	}
	
	if stderr != ""{
		log.Fatalf("ERROR - sterr wasn't empty while trying to parse venv requirements: %s", err)
	}

	return stdout
}

// HasPackage - checks if package wantPackageName in virtualenv and has wantPackgeVersion
func(v *Venv) hasPackage(wantPackageName , wantPackgeVersion string) bool{
	requirements := v.listRequirements()
	scanner := bufio.NewScanner(strings.NewReader(requirements))
	for scanner.Scan() {
		line := scanner.Text()
		l:= strings.Split(line,"==")
		hasName := l[0]
		hasVersion  := l[1]

		if  hasName == wantPackageName{
			
			if hasVersion != wantPackgeVersion{
				log.Fatalf("ERROR - package %s is in venv %s but version %s is different than %s ",
							wantPackageName, v.dir.Path , wantPackgeVersion, hasVersion)
			}else{
				log.Printf("OK - Package %s with version %s in %s vitualenv", wantPackageName, wantPackgeVersion, v.dir.Path)
				return true
			}	
		}
	}
	log.Fatalf("ERROR - package %s is not in venv %s ", wantPackageName, v.dir.Path)
	return false
}
// PackageImportable - verifies python package is in venv and is importable
func (v *Venv) PackageImportable(wantName,wantVersion string) bool{
	if v.hasPackage(wantName,wantVersion){
		 shell(fmt.Sprintf("%s/bin/python -c 'import %s'", v.dir.Path, wantName))
	}
	log.Printf("OK - %s,%s is importable in venv %s", wantName,wantVersion, v.dir.Path )
	wg.Done()
	return true


}