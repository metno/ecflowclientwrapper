package goecflowclientwrapper

import (
	"fmt"
	"os"
	"os/exec"
)

// No swig . Just call ecflow_client and get over with it.

// Ecflowinit - Run ecflow_client --init ...
func Ecflowinit(host string, port string) error {

	if os.Getenv("NO_ECF") != "" && os.Getenv("NO_ECF") != "0" {
		return nil
	}

	rfid := fmt.Sprintf("%d", os.Getpid())
	if os.Getenv("JOB_ID") != "" { // Running on gridengine
		rfid = os.Getenv("JOB_ID")
	}

	cmdinit := exec.Command("ecflow_client",
		"--host", host,
		"--port", port,
		"--init",
		rfid,
	)
	stdoutStderr, errinit := cmdinit.CombinedOutput()
	if errinit != nil {
		return fmt.Errorf("Ecflowinit: command %v failed with %s, output: %s",
			cmdinit, errinit, stdoutStderr)
	}
	return nil
}

// Ecflowcomplete - Run ecflow_client --complete  ...
func Ecflowcomplete(host string, port string) error {
	if os.Getenv("NO_ECF") != "" && os.Getenv("NO_ECF") != "0" {
		return nil
	}

	cmdcomplete := exec.Command("ecflow_client",
		"--host", host,
		"--port", port,
		"--complete",
	)
	stdoutStderr, errcomplete := cmdcomplete.CombinedOutput()
	if errcomplete != nil {
		return fmt.Errorf("Ecflowcomplete: Command %v failed with %s, output: %s", cmdcomplete, errcomplete, stdoutStderr)
	}
	return nil
}
