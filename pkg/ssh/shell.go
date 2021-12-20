package ssh

import (
	"bytes"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/monandkey/ssh/pkg/log"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// interactiveShellCalling is a function for interactive shells.
func interactiveShellCalling(session *ssh.Session) error {
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	fd := int(os.Stdin.Fd())
	state, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, state)

	var w, h int
	w, h, err = term.GetSize(fd)

	switch runtime.GOOS {
	case "windows":
		w, h = 200, 50
	case "linux":
		if err != nil {
			return err
		}
	}

	if err := session.RequestPty("xterm", h, w, modes); err != nil {
		return err
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		return err
	}

	if err := session.Wait(); err != nil {
		return err
	}
	return nil
}

// nonInteractiveShellCalling is a function for non-interactive shells.
func nonInteractiveShellCalling(session *ssh.Session, command string, logger *log.Logger) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	session.Stdout = &stdout
	session.Stderr = &stderr

	if err := session.Run(command); err != nil {
		for _, e := range strings.Split(stderr.String(), "\n") {
			if regexp.MustCompile(`^$`).MatchString(e) {
				continue
			}
			logger.Println(e)
		}
		return
	}

	for _, s := range strings.Split(stdout.String(), "\n") {
		if regexp.MustCompile(`^$`).MatchString(s) {
			continue
		}
		logger.Println(s)
	}
}
