# gotestinfra

- This is just POC for a library to support testing infrastructure as part of
  ansible playbooks

# Examples

1. Filesystem objects

```go
// FILES/DIRS
	HOME := FileSystemObject{Path: "/home"}	
	HOME.HasMode("drwxr-xr-x")
	HOME.IsDir()
```
2. Venv - python virtual environment

```go

// VENV
	VENV := NewVenv("/data/virtualenv/default")
	VENV.dir.IsDir()
	VENV.dir.HasMode("drwx------")
	VENV.dir.OwnedByUser("deploy")
    VENV.dir.OwnedByGroup("deploy")
    VENV.PackageImportable("chardet", "3.0.4")
```

- RAM 

```go
	// RAM
	RAM := RAM{}
	RAM.Total("GB").GreaterThan(0.9)
	RAM.Free("GB").GreaterThan(0.2)
	RAM.SwapTotal("GB").LowerThan(2.0)
	RAM.SwapTotal("GB").GreaterThan(1.0)
	RAM.SwapFree("GB").GreaterThan(1.0)
	RAM.Buffers("KB").LowerThan(10000)
	RAM.Cached("GB").LowerThan(1)

```

- Diks mount

```go

	// DISK MOUNTS
	MOUNTS := Mounts{}
	MOUNTS.HasMountPoint("/vagrant")

```

- RPM 

```go
	// RPM
	RPM := rpm{}
	RPM.IsInstalled("git")
	RPM.IsInstalledVersion("git","1.8.3.1")
	RPM.IsInstalledRelease("git","14.el7_5")
```

- Performance : Usuaally the tests are running very fast(less than 1sec )
  
  ```go
    real    0m0.672s
    user    0m0.604s
    sys     0m0.238s

  ```

- if you want to execute the tests concurrently you need to add count `wg.Add(1)` to waiting group and execute the normal test with `go` prefix , full example:
  ```go
	// register task to wait for 
	wg.Add(1); 
         go VENV.PackageImportable("chardet", "3.0.4")
    // wait for concurrent go routines to finish before exiting
	wg.Wait()
  ```
	

- 