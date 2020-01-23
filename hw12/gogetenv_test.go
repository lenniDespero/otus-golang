package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		t.Fatalf("Fail to create dir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()
	tmpfile, err := os.Create(dir + "/my_env")
	if err != nil {
		t.Fatalf("Fail to create file: %s", err)
	}
	defer func() {
		_ = tmpfile.Close()
	}()
	_, err = tmpfile.Write([]byte("my_value"))
	if err != nil {
		t.Fatalf("Fail to write in file: %s", err)
	}
	result, err := ReadDir(dir)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(result) != 1 {
		t.Errorf("Wait %d envs, get %d", 1, len(result))
	}
	if result["my_env"] != "my_value" {
		t.Errorf("Wait '%s' , get '%s'", "my_value", result["my_env"])
	}
}

func TestReadDirEmpty(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		t.Fatalf("Fail to create dir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()
	tmpfile, err := os.Create(dir + "/my_env")
	if err != nil {
		t.Fatalf("Fail to create file: %s", err)
	}
	defer func() {
		_ = tmpfile.Close()
	}()
	result, err := ReadDir(dir)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(result) > 0 {
		t.Errorf("Wait %d envs, get %d", 0, len(result))
	}
}

func TestReadDirEmptyFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		t.Fatalf("Fail to create dir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	result, err := ReadDir(dir)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if len(result) > 0 {
		t.Errorf("Wait %d envs, get %d", 0, len(result))
	}
}

func TestReadDirWrong(t *testing.T) {

	_, err := ReadDir("SomeSrangeNotExistDir")
	if err == nil {
		t.Errorf("Wait Error but not get it")
	}
}

func ExampleRunCmd() {
	env := map[string]string{"BIG_STRANGE_VAR": "test"}
	cmd := []string{"env", "|", "grep", "BIG_STRANGE_VAR"}
	code := RunCmd(cmd, env)
	if code != 0 {
		log.Fatalf("Except Exit code - 0, get - %d", code)
	}
	// Output:
	// BIG_STRANGE_VAR=test
}
