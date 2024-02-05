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
	"path/filepath"
	"syscall"
)

var TMPFILE = "/dev/null"
var LOGPREFIX_ENV_VAR = "LOGPREFIX"

func main() {
	log.SetPrefix(fmt.Sprintf("%s[pid:%d] ", os.Getenv(LOGPREFIX_ENV_VAR), os.Getpid()))
	log.SetFlags(log.Lshortfile)
	log.Println("⏩", os.Args)
	var SYMLINK_FROM = filepath.Join(os.TempDir(), "demobinary-symlink")
	var SYMLINK_TO = filepath.Join(os.TempDir(), "demobinary-file")

	var fileWrite = flag.Bool("file-write", false, "write to "+TMPFILE)
	var fileRead = flag.Bool("file-read", false, "read from "+TMPFILE)
	var fileTemp = flag.Bool("file-temp", false, "create a random file in "+os.TempDir())
	var fileSymlink = flag.Bool("file-symlink", false, fmt.Sprintf("Symlink %s -> %s and try reading it.", SYMLINK_FROM, SYMLINK_TO))
	var netTcp = flag.Bool("net-tcp", false, "spawn a tcp server")
	var netUdp = flag.Bool("net-udp", false, "spawn a udp server")
	var netIcmp = flag.Bool("net-icmp", false, "open an icmp socket, exercise NET_RAW capability.")
	var library = flag.String("load-library", "", "load a shared library")
	var crash = flag.Bool("crash", false, "crash instead of exiting.")

	flag.Parse()

	var subprocess = flag.Args()

	if *fileWrite {
		if err := os.WriteFile(TMPFILE, []byte{}, 0666); err != nil {
			log.Fatal("❌ Error creating file:", err)
		} else {
			log.Println("✅ File write successful:", TMPFILE)
		}
		// make file writable for other users so that sudo/non-sudo testing works.
		os.Chmod(TMPFILE, 0666)
	}
	if *fileRead {
		if _, err := os.ReadFile(TMPFILE); err != nil {
			log.Fatal("❌ Error reading file:", err)
		} else {
			log.Println("✅ File read successful:", TMPFILE)
		}
	}
	if *fileSymlink {
		if _, err := os.Stat(SYMLINK_TO); err != nil {
			if err := os.WriteFile(SYMLINK_TO, []byte{}, 0644); err != nil {
				log.Fatal("❌ Error creating SYMLINK_TO file:", err)
			}
			if err := os.Chmod(SYMLINK_TO, 0644); err != nil {
				log.Fatal("❌ Error setting permissions for SYMLINK_TO file:", err)
			}
			// defer os.Remove(SYMLINK_TO)
		}
		if _, err := os.Stat(SYMLINK_FROM); err != nil {
			fmt.Println("Cannot stat", err)
			if err := os.Symlink(SYMLINK_TO, SYMLINK_FROM); err != nil {
				log.Fatal("❌ Error creating symlink:", err)
			}
			// defer os.Remove(SYMLINK_FROM)
		}

		if _, err := os.ReadFile(SYMLINK_FROM); err != nil {
			log.Fatal("❌ Error reading symlink:", err)
		} else {
			log.Println("✅ Symlink read successful:", SYMLINK_FROM)
		}
	}
	if *fileTemp {
		f, err := os.CreateTemp("", "demobinary")
		defer f.Close()
		defer os.Remove(f.Name())
		if err != nil {
			log.Fatal("❌ Error creating temp file:", err)
		} else {
			log.Println("✅ Created temp file:", f.Name())
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
