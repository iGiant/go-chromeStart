# go-chromeStart
Running chrome for for remote control

Usage example:
```Go
package main

import (
	"github.com/iGiant/go-chromeStart"
	"time"
)

func main() {
	r, _ := chrome.New("C:/Program Files (x86)/Google/Chrome/Application/chrome.exe", 9245)
	r.Start()
	defer r.Stop()
	time.Sleep(10 * time.Second)
}
```
