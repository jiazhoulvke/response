## response ##

example:

```go
package main

import (
	"github.com/jiazhoulvke/response"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	response.Errors[100001] = "foo must be 1"
	e.GET("/", func(c echo.Context) error {
		if c.FormValue("foo") == "1" {
			return c.JSON(200, response.Succ(map[string]interface{}{
				"data": "bar",
			}))
		}
		return c.JSON(200, response.Fail(100001))
	})
	e.Start(":8888")
}
```

response:

foo=1:

> {"data":"bar","status":"succ"}

foo!=1:

> {"err_code":100001,"err_msg":"foo must be 1","status":"fail"}
