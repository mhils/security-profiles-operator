/*
Demo binary to exercise various capabilities that may be restricted by seccomp/apparmor.
*/
package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"
import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var LOGPREFIX_ENV_VAR = "LOGPREFIX"

func main() {
	log.SetPrefix(fmt.Sprintf("%s[pid:%d] ", os.Getenv(LOGPREFIX_ENV_VAR), os.Getpid()))
	log.SetFlags(log.Lshortfile)
	log.Println("⏩", os.Args)

	var fileWrite = flag.String("file-write", "", "write file (e.g. /dev/null)")
	var fileRead = flag.String("file-read", "", "read file (e.g. /dev/null)")
	var fileSymlink = flag.String("file-symlink", "", "Create symlink using the following syntax: OLD:NEW")
	var netTcp = flag.Bool("net-tcp", false, "spawn a tcp server")
	var netUdp = flag.Bool("net-udp", false, "spawn a udp server")
	var netIcmp = flag.Bool("net-icmp", false, "open an icmp socket, exercise NET_RAW capability.")
	var library = flag.String("load-library", "", "load a shared library")
	var crash = flag.Bool("crash", false, "crash instead of exiting.")

	flag.Parse()

	var subprocess = flag.Args()

	if *fileWrite != "" {
		if err := os.WriteFile(*fileWrite, []byte{}, 0666); err != nil {
			log.Fatal("❌ Error creating file:", err)
		} else {
			log.Println("✅ File write successful:", *fileWrite)
		}
		// make file writable for other users so that sudo/non-sudo testing works.
		os.Chmod(*fileWrite, 0666)
	}
	if *fileSymlink != "" {
		old, new, found := strings.Cut(*fileSymlink, ":")
		if !found {
			log.Fatal("❌ Symlink syntax: OLD:NEW")
		}
		if err := os.Symlink(old, new); err != nil {
			log.Fatal("❌ Error creating symlink:", err)
		} else {
			log.Println("✅ Symlink created:", new, "->", old)
		}
	}
	if *fileRead != "" {
		if _, err := os.ReadFile(*fileRead); err != nil {
			log.Fatal("❌ Error reading file:", err)
		} else {
			log.Println("✅ File read successful:", *fileRead)
		}
	}
	if *netTcp {
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatal("❌ Error starting TCP server:", err)
		} else {
			log.Println("✅ TCP server spawned:", listener.Addr())
		}
		defer listener.Close()
	}
	if *netUdp {
		server, err := net.ListenPacket("udp", ":0")
		if err != nil {
			log.Fatal("❌ Error starting UDP server:", err)
		} else {
			log.Println("✅ UDP server spawned:", server.LocalAddr())
		}
		defer server.Close()
	}
	if *netIcmp {
		fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
		if err != nil {
			log.Fatal("❌ Error opening ICMP socket:", err)
		} else {
			log.Println("✅ ICMP socket opened: fd", fd)
		}
		defer syscall.Close(fd)
	}
	if len(subprocess) > 0 {
		cmd := exec.Command(subprocess[0], subprocess[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(), "LOGPREFIX=\t"+os.Getenv(LOGPREFIX_ENV_VAR))
		if err := cmd.Run(); err != nil {
			log.Fatal("❌ Error running subprocess:", err)
		} else {
			log.Println("✅ Subprocess ran successfully:", subprocess)
		}
	}
	if *library != "" {
		if handle := C.dlopen(C.CString(*library), C.RTLD_NOW); handle == nil {
			log.Fatal("❌ Error loading library: ", C.GoString(C.dlerror()))
		} else {
			log.Println("✅ Library loaded successfully:", *library)
		}
	}
	if *crash {
		log.Println("🫡  Terminating with SIGKILL...")
		syscall.Kill(syscall.Getpid(), syscall.SIGKILL)
	}
	log.Println("⭐️ Success.")
}
