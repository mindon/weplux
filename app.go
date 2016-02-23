package weplux

import (
    "fmt"
    "time"
    "net/http"
)

func init() {
  http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  var now = time.Now();
  fmt.Fprintf(w, "Hello, Wellcome to WePlux - by Mindon!\n%v, %v/365 = %.2f%% ",
    now.Format("2006年01月02日 时区 Z07:00 15:04:05.000  Mon"),
    now.YearDay(),
    float64( 100 * now.YearDay()) / 365.0)

  fmt.Fprint(w, "\n\nhttp://www.pinqic.com/")
}
