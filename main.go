package main


func main() {

	// VENV
	VENV := NewVenv("/data/virtualenv/default")
	VENV.dir.IsDir()
	VENV.dir.HasMode("drwx------")
	VENV.dir.OwnedByUser("deploy")
	VENV.dir.OwnedByGroup("deploy")
	// concurrent execution example
	wg.Add(1); go 
		VENV.PackageImportable("chardet", "3.0.4")

	// FILES/DIRS
	HOME := FileSystemObject{Path: "/home"}	
	HOME.HasMode("drwxr-xr-x")
	HOME.IsDir()

	// RAM
	RAM := RAM{}
	RAM.Total("GB").GreaterThan(0.9)
	RAM.Free("GB").GreaterThan(0.2)
	RAM.SwapTotal("GB").LowerThan(2.0)
	RAM.SwapTotal("GB").GreaterThan(1.0)
	RAM.SwapFree("GB").GreaterThan(1.0)
	RAM.Buffers("KB").LowerThan(10000)
	RAM.Cached("GB").LowerThan(1)

	// DISK MOUNTS
	MOUNTS := Mounts{}
	MOUNTS.HasMountPoint("/vagrant")

	// RPM
	RPM := rpm{}
	RPM.IsInstalled("git")
	RPM.IsInstalledVersion("git","1.8.3.1")
	RPM.IsInstalledRelease("git","14.el7_5")



	// wait for concurrent go routines to finish before exiting
	wg.Wait()
}

