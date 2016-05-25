package fileCommons

import (
	"testing"
)

func TestReadToByte(t *testing.T) {
	//todo
	_, err := ReadToBytes("/etc/services")
	if err != nil {
		t.Error(err)
	}
}

func TestReadToString(t *testing.T) {
	//todo
	_, err := ReadToString("/etc/services")
	if err != nil {
		t.Error(err)
	}
}
