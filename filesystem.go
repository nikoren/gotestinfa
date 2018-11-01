package main

import(
	"os"
	"path/filepath"
	"log"
	"syscall"
	"os/user"
	"strconv"
	// "net"
	// "fmt"
)
// FileSystem - data structure to represent file system objects(dir/file)
type FileSystemObject struct{
	Path string
}

func(fs *FileSystemObject) getFileInfo() os.FileInfo {
	if ! filepath.IsAbs(fs.Path){
		log.Fatal("%s is not is absolute path", fs.Path)
	}
	fsInfo, err := os.Stat(fs.Path)
	if err != nil{
		log.Fatalf("Something went wrong - couldn't get file sysemt stats of %s, seems like the path is not exist ", fs.Path)
	}
	return fsInfo
}

// IsFile - check it path is a a regular file
func(fs *FileSystemObject) IsFile() bool{
	fsInfo := fs.getFileInfo()
	if fsInfo.IsDir(){
		log.Fatalf("%s is not file, it is directory ", fs.Path)
	}
	log.Printf("OK - %s is file", fs.Path) 
	return true
}

// IsDir - check if path is directory
func(fs *FileSystemObject) IsDir() bool{
	fsInfo := fs.getFileInfo()
	if ! fsInfo.IsDir(){
		log.Fatal("%s is not directory", fs.Path)
	}
	
	log.Printf("OK - %s is directory",fs.Path)
	return true
}

// HasMode - check if path is directory
func(fs *FileSystemObject) HasMode(want string) bool{
	fsInfo := fs.getFileInfo()
	fsMode := fsInfo.Mode()
	// fmt.Println("isDir " , fsMode.IsDir(), 
	// 	"isRegular " , fsMode.IsRegular(), 
	// 	"Perm: " , fsMode.Perm, "String: ", fsMode.String())

	if fsMode.String() != want {
		
		log.Fatalf("'%s' permissions '%s' != '%s'", 
					fs.Path, fsMode.String(), want)
	}
	log.Printf("OK - %s permissions %s", fs.Path, fsMode.String())	
	return true
}

// OwnedByUser - verifies file system object owned by specific user
func( fs *FileSystemObject ) OwnedByUser(want string) bool {

	hasUID := fs.getFileInfo().Sys().(*syscall.Stat_t).Uid
	wantUser,err := user.Lookup(want)
	if err != nil{
		log.Fatalf("ERROR - Couldn't find user %s for path %s", want, fs.Path)
	}
	wantUID, err:= strconv.Atoi(wantUser.Uid)
	if err != nil{
		log.Fatalf("ERROR - Failed to parse %s uid for path %s", want , fs.Path )
	}
	if uint32(wantUID) != hasUID{
		log.Fatalf("ERROR - Has UID %d != want %d UID for path %s",hasUID, wantUID, fs.Path)
	} 
	log.Printf("OK - %s owned by user %s", fs.Path, want)
	return true

}
// OwnedByGroup - verifies file system object owned by specific user
func( fs *FileSystemObject ) OwnedByGroup(want string) bool {

	hasGID := fs.getFileInfo().Sys().(*syscall.Stat_t).Gid
	wantGroup,err := user.LookupGroup(want)
	if err != nil{
		log.Fatalf("ERROR - Couldn't find group %s for path %s", want, fs.Path)
	}
	wantGID, err:= strconv.Atoi(wantGroup.Gid)
	if err != nil{
		log.Fatalf("ERROR - Failed to parse %s gid for path %s", want , fs.Path)
	}
	if uint32(wantGID) != hasGID{
		log.Fatalf("ERROR - Has GID %d != want %d GID for path %s",hasGID, wantGID, fs.Path)
	} 
	log.Printf("OK - %s owned by group %s", fs.Path, want)
	return true
}

// IsTCPPortAvailable returns a flag indicating whether or not a TCP port is
// available.
// func IsTCPPortAvailable(port int) bool {
// 	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
// 	if err != nil {
// 		return false
// 	}
// 	conn.Close()
// 	return true
// }
