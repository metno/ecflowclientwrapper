package goecflowclientwrapper

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
)

// No swig . Just call ecflow_client and get over with it.
var ecflow_client_bin = ""

func init() {
	
	
	cmd := exec.Command("lsb_release", "-sc")
	osreleaseb, err := cmd.CombinedOutput()
	osrelease := strings.TrimSuffix(string(osreleaseb), "\n")
	if err != nil {
			log.Fatalf("lsb_release: %v\n", err)
	}

	if osrelease == "Core" {
		ecflow_client_bin = "/modules/centos7/user-apps/ecflow/5.3.0/bin/ecflow_client"
	} else if osrelease == "bionic" {
		ecflow_client_bin = "/modules/bionic/user-apps/ecflow/5.5.2-ssl/bin/ecflow_client"
	} else if osrelease == "xenial" {
		ecflow_client_bin = "/modules/xenial/user-apps/ecflow/4.14.0/bin/ecflow_client"
	} else {
		log.Fatalf("os relase %s not supported", osrelease)
	}
}

// Ecflowinit - Run ecflow_client --init ... Use env JOB_ID for rfid if set (=> Running on gridengine)
func Ecflowinit(host string, port string) error {

	rfid := fmt.Sprintf("%d", os.Getpid())
	if os.Getenv("JOB_ID") != "" { // Running on gridengine
		rfid = os.Getenv("JOB_ID")
	}

	stderr, stdout, err := EcflowClient("--host", host, "--port", port, "--init", rfid)
	if err != nil {
		return fmt.Errorf("ecflow_client --init failed with %v, stdout: %s, stderr: %s", err, stdout, stderr)
	}

	return nil
}

// Ecflowcomplete - Run ecflow_client --complete  ...
func Ecflowcomplete(host string, port string) error {

	stderr, stdout, err := EcflowClient("--host", host, "--port", port, "--complete")
	if err != nil {
		return fmt.Errorf("ecflow_client --complete failed with %v, stdout: %s, stderr: %s", err, stdout, stderr)
	}
	return nil
}

// EcflowClient - Call ecflow_client binary. Returns teh command stdout, stderr, and go's cmd.Run() error
func EcflowClient(args ...string) (string, string, error) {
	
	cmd := exec.Command(ecflow_client_bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	if err != nil {
		return outStr, errStr, err
	}

	return outStr, errStr, err
}
