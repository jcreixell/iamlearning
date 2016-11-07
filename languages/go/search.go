package main

import(
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "html/template"
  "time"
  "golang.org/x/talks/2016/applicative/google"
)

func main() {
  http.HandleFunc("/search", handleSearch)
  fmt.Println("serving on http://localhost:7777/search")
  log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
  log.Println("serving", req.URL)

  query := req.FormValue("q")
  if query == "" {
    http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
    return
  }

  start := time.Now()
  results, err := google.Search(query)
  elapsed := time.Since(start)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  type response struct {
    Results []google.Result
    Elapsed time.Duration
  }
  resp := response{results, elapsed}

  var responseTemplate = template.Must(template.New("results").Parse(`
  <html>
  </head>
  <body>
    <ol>
      {{range .Results}}
        <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
      {{end}}
      <p>{{len .Results}} results in {{.Elapsed}}</p>
    </ol>
  </body>
  </html>
  `))

  switch req.FormValue("output") {
    case "json":
      err = json.NewEncoder(w).Encode(resp)
    case "prettyjson":
        var b []byte
        b, err = json.MarshalIndent(resp, "", " ")
        if err == nil {
          _, err = w.Write(b)
        }
    default:
      err = responseTemplate.Execute(w, resp)
  }
}
