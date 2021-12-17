package ssh

import (
	"os"
	"runtime"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func interactiveShellCalling(session *ssh.Session) error {
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	var w, h int
	switch runtime.GOOS {

	case "windows":
		w, h = 1000, 1000
	case "linux":
		fd := int(os.Stdin.Fd())
		state, err := term.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer term.Restore(fd, state)
		w, h, err = term.GetSize(fd)
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

func nonInteractiveShellCalling(session *ssh.Session, command string) error {
	session.Stdout = os.Stdout
	if err := session.Run(command); err != nil {
		return err
	}
	return nil
}
