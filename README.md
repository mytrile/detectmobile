# detectmobile [![wercker status](https://app.wercker.com/status/abbd5829c488bad4489a52e2a1474ff4/s "wercker status")](https://app.wercker.com/project/bykey/abbd5829c488bad4489a52e2a1474ff4)

Martini middleware/handler for detecting mobile device by HTTP headers. 
This middleware will check User-Agent and Accept headers and will set X-Mobile-Device header to true/false.

## Usage

~~~ go
package main

import (
    "github.com/go-martini/martini"
    "github.com/mytrile/detectmobile"
)

func main() {
    m := martini.Classic()
    m.Use(detectmobile.DetectMobile())
    m.Get("/", func(mobileDevice string) string {
      // mobileDevice will be true/false depending on detection
      return "Hello world!"
    })
    m.Run()
}

~~~

## Meta

* Author  : Dimitar Kostov
* Email   : mitko.kostov@gmail.com
* Website : [http://mytrile.github.com](http://mytrile.github.com)
* Twitter : [http://twitter.com/mytrile](http://twitter.com/mytrile)