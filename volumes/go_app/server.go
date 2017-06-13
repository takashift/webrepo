package main

import (
    "net"
    "net/http"
//    "net/http/fcgi"
    "os"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "おっぱいぼいんぼいん@Echo")
    })
//    e.Logger.Fatal(e.Start(":1323"))

    os.Remove("/usock/domain.sock");
    uni, err := net.Listen("unix", "/usock/domain.sock")
    if err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chmod("/usock/domain.sock", 0600); err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chown("/usock/domain.sock", 1000, 1000); err != nil {
      e.Logger.Fatal(err)
    }
    e.Listener = uni
    e.Logger.Fatal(e.Start(""))


//    tcp, err := net.Listen("tcp", ":1323")
//    if err != nil {
//      e.Logger.Fatal(err)
//    }
//    e.Listener = tcp
//    e.Logger.Fatal(e.Start(""))

//    fcgi.Serve(tcp, nil)
}
