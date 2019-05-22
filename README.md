# go-chromeStart v0.0.3
Running chrome for for remote control

## Installation

    $ go get github.com/iGiant/go-chromeStart


Usage example:
```Go
package main

import (
	"github.com/iGiant/go-chromeStart"
	"time"
)

func main() {
	c, _ := chrome.New("C:/Program Files (x86)/Google/Chrome/Application/chrome.exe", 9222)
	c.Headless(true)
	c.SetSize(1920, 1024)
	c.Start()
	defer c.Stop()
	time.Sleep(10 * time.Second)
}
```
Chrome will start for 10 seconds in the hidden mode and will be available for remote control through port 9222
