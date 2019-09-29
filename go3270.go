/*
Package go3270 is a Golang interface to the s3270 binary which will allow you to communicate with an IBM mainframe session.
*/
package go3270

import (
	"io"
	"os/exec"
)

type emulator interface {
	// read
	// write
}
type s3270Emulator struct {
	cmd     *exec.Cmd
	pipein  io.WriteCloser
	pipeout io.ReadCloser
}

// startS3270Emulator will start the s3270 executable and set the host connection timeout to the specified value.
// Pipes are created to communicate with the executable via STDIN and STDOUT.
func startS3270Emulator(timeout int) (*s3270Emulator, error) {
	parm := ""
	if timeout > 0 {
		parm = "-connecttimeout " + string(timeout)
	}
	c := exec.Command("s3270", parm)
	in, err := c.StdinPipe()
	if err != nil {
		return nil, err
	}

	out, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}

	e := &s3270Emulator{c, in, out}
	err = e.cmd.Start()
	if err != nil {
		return nil, err
	}
	return e, err
}

func (e *s3270Emulator) terminate() error {
	// TODO: stop S3270: use Kill or command to s3270 followed by Wait()?
	if e.cmd != nil {
		err := e.pipein.Close()
		if err != nil {
			return err
		}

		err = e.pipeout.Close()
		if err != nil {
			return err
		}

		err = e.cmd.Process.Kill()
		if err != nil {
			return err
		}
		e.cmd = nil
	}
	return nil
}
func (e *s3270Emulator) read() string {
	// TODO: finish use scanner := bufio.NewScanner(os.Stdin)
	return ""
}
func (e *s3270Emulator) write() string {
	// TODO: finish
	return ""
}

// Session is the host session.
type Session struct {
	emulator emulator
	Screen   string
	Status   *Status
}

// Status contains the status of the most recent command.
type Status struct {
	// If the keyboard is unlocked, the letter U. If the keyboard is locked waiting for a response from the host, or if not connected to a host, the letter L.
	// If the keyboard is locked because of an operator error (field overflow, protected field, etc.), the letter E.
	KeyboardState string

	// If the screen is formatted, the letter F. If unformatted or in NVT mode, the letter U.
	ScreenFormat string

	// If the field containing the cursor is protected, the letter P. If unprotected or unformatted, the letter U.
	FieldProtection string

	// If connected to a host, the string C(hostname). Otherwise, the letter N.
	ConnectionState string

	// If connected in 3270 mode, the letter I. If connected in NVT line mode, the letter L. If connected in NVT character mode, the letter C.
	// If connected in unnegotiated mode (no BIND active from the host), the letter P. If not connected, the letter N.
	EmulatorMode string

	// Terminal model number 2 to 5.
	ModelNo int

	// The current number of rows defined on the screen. The host can request that the emulator use a 24x80 screen,
	// so this number may be smaller than the maximum number of rows possible with the current model.
	RowCount int

	// The current number of columns defined on the screen, subject to the same difference for rows, above.
	ColumnCount int

	// The current cursor row (zero-origin).
	CursorRow int

	// The current cursor column (zero-origin).
	CursorColumn int

	// The X window identifier for the main x3270 window, in hexadecimal preceded by 0x. For s3270 and c3270, this is zero.
	WindowID string

	// The time that it took for the host to respond to the previous command, in seconds with milliseconds after the decimal.
	// If the previous command did not require a host response, this is a dash.
	ExecTime string
}

// TODO: status struct needs method to convert string to individual fields
