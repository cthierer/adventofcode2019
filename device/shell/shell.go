package shell

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cthierer/advent-of-code/device/filesystem"
)

const (
	commandStart     = "$"
	commandList      = "ls"
	commandChangeDir = "cd"
)

type Shell struct {
	FS  *filesystem.FS
	pwd *filesystem.Dir
}

func NewShell() *Shell {
	fs := filesystem.NewFS()
	pwd := fs.Root()
	return &Shell{FS: fs, pwd: pwd}
}

func isCommandLine(l string) bool {
	return strings.HasPrefix(l, commandStart)
}

type command struct {
	name   string
	arg    string
	output []string
}

type commandIterator struct {
	lines []string
	pos   int
}

func (i *commandIterator) Next() bool {
	return i.pos < len(i.lines) && isCommandLine(i.lines[i.pos])
}

func (i *commandIterator) Get() command {
	l := i.lines[i.pos]
	cmdWithArgs := strings.TrimSpace(strings.TrimPrefix(l, commandStart))

	c := command{}
	c.name = cmdWithArgs[0:2]
	c.arg = strings.TrimSpace(cmdWithArgs[2:])

	for i.pos < len(i.lines)-1 {
		i.pos += 1
		l = i.lines[i.pos]
		if isCommandLine(l) {
			break
		}
		c.output = append(c.output, l)
	}

	return c
}

func processDir(s *Shell, name string) error {
	d := filesystem.Dir{Name: name}
	s.pwd.AddDir(&d)
	return nil
}

func processFile(s *Shell, name, size string) error {
	sizeInt, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return err
	}
	f := filesystem.File{Name: name, Size: sizeInt}
	s.pwd.AddFile(&f)
	return nil
}

func processList(s *Shell, c command) error {
	for _, l := range c.output {
		p := strings.Split(l, " ")
		switch p[0] {
		case "dir":
			if err := processDir(s, p[1]); err != nil {
				return err
			}
		default:
			if err := processFile(s, p[1], p[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

func processChangeDir(s *Shell, c command) error {
	switch c.arg {
	case "..":
		nextPwd := s.pwd.Parent()
		if nextPwd == nil {
			return errors.New("cannot change directory: already at root")
		}

		s.pwd = nextPwd
	case "/":
		s.pwd = s.FS.Root()
	default:
		nextPwd := s.pwd.GetDir(c.arg)
		if nextPwd == nil {
			return errors.New("cannot change directory: no directoy with name exists")
		}

		s.pwd = nextPwd
	}
	return nil
}

func Parse(input string) (*Shell, error) {
	shell := NewShell()
	commands := commandIterator{lines: strings.Split(input, "\n")}

	for commands.Next() {
		cmd := commands.Get()
		switch cmd.name {
		case commandList:
			if err := processList(shell, cmd); err != nil {
				return nil, err
			}
		case commandChangeDir:
			if err := processChangeDir(shell, cmd); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unrecognized command: %v", cmd.name)
		}
	}

	return shell, nil
}
