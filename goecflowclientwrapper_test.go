package goecflowclientwrapper

import (
	"os"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	_, stderr, err := EcflowClient("--host", "dontexist", "--port", "3141", "--ping")

	if err == nil {
		t.Errorf("got err == nil; want err == %v", err)
	}

	if !strings.Contains(stderr, "Failed to connect to dontexist:3141. After 2 attempts. Is the server running ?") {
		t.Errorf("stderr == %s; want stderr == Failed to connect to dontexist:3141. After 2 attempts. Is the server running ?", stderr)
	}

}

func TestClientNOECF(t *testing.T) {

	os.Setenv("NO_ECF", "1")
	stdout, stderr, err := EcflowClient("--host", "dontexist", "--port", "3141", "--ping")

	if err != nil {
		t.Errorf("want err != nil; got err == %v", err)
	}

	if !strings.Contains(stdout, "NO_ECF") {
		t.Errorf("got stdout==%s; want stdout==NO_ECF", stdout)
	}

	if stderr != "" {
		t.Errorf("got stderr==%s; want stderr==\"\"", stderr)
	}

}
