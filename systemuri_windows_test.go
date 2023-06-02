package systemuri

import (
	winreg "golang.org/x/sys/windows/registry"
	"log"
	"testing"
)

func Test_registerURLHandlerWithRegistry(t *testing.T) {
	err := registerURLHandlerWithRegistry(testCreateFn, "", "", "", "")
	if err != nil {
		log.Fatal(err)
	}
}

func testCreateFn(k winreg.Key, path string, access uint32) (winreg.Key, bool, error) {

	log.Println("test")
	//newk = winreg.Key(syscall.Handle(1))
	return 0, true, nil
}
