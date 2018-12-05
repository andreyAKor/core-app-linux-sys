# core-app-linux-sys
Core of golang app for linux system service

[Go](http://golang.org) package of сore app for linux system service

This library includes some packages and architecture skeleton for writing go-app as linux system service.

Install
-------

    go get github.com/andreyAKor/core-app-linux-sys

Example
-------

```go
package main

import (
	core "github.com/andreyAKor/core-app-linux-sys"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init service
	service := core.NewService("app.conf", new(config.Configuration), new(App))

	// Start service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```
