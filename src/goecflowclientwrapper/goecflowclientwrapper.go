package goecflowclientwrapper

import (
	"log"
	"os"
	"os/exec"
)

// Ecflowinit - Run ecflow_client --init ...
func Ecflowinit() {
	rfid := os.Getenv("JOB_ID")
	cmdinit := exec.Command("ecflow_client",
		"--host", os.Getenv("ECF_HOST"),
		"--port", os.Getenv("ECF_PORT"),
		"--init",
		rfid,
	)
	stdoutStderr, errinit := cmdinit.CombinedOutput()
	if errinit != nil {
		log.Printf("ERROR: Command %v failed with %s, output: %s\n", cmdinit, errinit, stdoutStderr)
	}
}

// Ecflowcomplete - Run ecflow_client --complete  ...
func Ecflowcomplete() {
	cmdcomplete := exec.Command("ecflow_client",
		"--host", os.Getenv("ECF_HOST"),
		"--port", os.Getenv("ECF_PORT"),
		"--complete",
	)
	stdoutStderr, errcomplete := cmdcomplete.CombinedOutput()
	if errcomplete != nil {
		log.Printf("ERROR: Command %v failed with %s, output: %s\n", cmdcomplete, errcomplete, stdoutStderr)
	}

}
