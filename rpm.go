package main

import (
	"fmt"
	"log"
)

type rpm struct{}

func (r *rpm) IsInstalled(want string) bool{
	shell(fmt.Sprintf("rpm -q %s", want))
	log.Printf("OK - RPM %s is installed", want)
	return true
}

func (r *rpm) IsInstalledVersion(wantPackage, wantVersion string) bool{
	haveVersion, _,_ := shell(
		fmt.Sprintf("rpm -q --queryformat='%%{VERSION}' %s", wantPackage))
	if wantVersion != haveVersion{
		log.Fatalf("ERROR - RPM version mismatch: have %s != want %s", haveVersion, wantVersion)
	}
	log.Printf("OK - RPM %s is installed , version %s", wantPackage, wantVersion)
	return true
}

func (r *rpm) IsInstalledRelease(wantPackage, wantRelease string) bool{
	haveRelease, _,_ := shell(
		fmt.Sprintf("rpm -q --queryformat='%%{RELEASE}' %s", wantPackage))
	if wantRelease != haveRelease{
		log.Fatalf("ERROR - RPM release mismatch: have %s != want %s", haveRelease, wantRelease)
	}
	log.Printf("OK - RPM %s is installed , relesase %s",wantPackage, wantRelease)
	return true
}
