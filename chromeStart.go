package chrome

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Chrome struct {
	cmd           *exec.Cmd
	path          string
	params        []string
	headless      bool
	width, height int
	port          int
}

func New(path string, port int) (*Chrome, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, errors.New(" the chrome does not exist at the specified path: " + path)
	}
	chrome := &Chrome{path: path,
		port:   port,
		params: []string{"--remote-debugging-port=" + strconv.Itoa(port)},
	}

	return chrome, nil
}

func (c *Chrome) Headless(h bool) error {
	if c.cmd != nil {
		return errors.New("the Chrome already running")
	}
	c.headless = h
	if h {
		err1 := c.AddParam("--headless")
		err2 := c.AddParam("--disable-gpu")
		var err []string
		if err1 != nil {
			err = append(err, err1.Error())
		}
		if err2 != nil {
			err = append(err, err2.Error())
		}
		if len(err) != 0 {
			return errors.New(strings.Join(err, "\n"))
		}
	} else {
		err1 := c.RemoveParam("--headless")
		err2 := c.RemoveParam("--disable-gpu")
		var err []string
		if err1 != nil {
			err = append(err, err1.Error())
		}
		if err2 != nil {
			err = append(err, err2.Error())
		}
		if len(err) != 0 {
			return errors.New(strings.Join(err, "\n"))
		}
	}
	return nil
}

func (c *Chrome) SetSize(width, height int) error {
	if width < 0 || height < 0 {
		return errors.New("window size is negative")
	}
	if c.width != 0 && c.height != 0 && (width == 0 || height == 0) {
		c.width, c.height = 0, 0
		for i := range c.params {
			if strings.HasPrefix(c.params[i], "--window-size=") {
				c.params = append(c.params[:i], c.params[i+1:]...)
				return nil
			}
		}
	}
	if c.width == width && c.height == height {
		return nil
	}
	_ = c.RemoveParam("--window-size=" + strconv.Itoa(c.width) + "," + strconv.Itoa(c.height))
	_ = c.AddParam("--window-size=" + strconv.Itoa(width) + "," + strconv.Itoa(height))
	c.width = width
	c.height = height
	return nil
}

func (c *Chrome) AddParam(p string) error {
	for _, param := range c.params {
		if strings.EqualFold(param, p) {
			return errors.New("param already exist in params")
		}
	}
	c.params = append(c.params, p)
	return nil
}

func (c *Chrome) RemoveParam(p string) error {
	for i := range c.params {
		if strings.EqualFold(c.params[i], p) {
			c.params = append(c.params[:i], c.params[i+1:]...)
			return nil
		}
	}
	return errors.New("param '" + p + "' is not listed")
}

func (c *Chrome) Params(p []string) {
	c.params = p
}

func (c *Chrome) Start() error {
	c.cmd = exec.Command(c.path, c.params...)
	return c.cmd.Start()
}

func (c *Chrome) Stop() error {
	return c.cmd.Process.Kill()
}
