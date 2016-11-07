package google

var (
  Web   = FakeSearch("web", "Something", "http://www.google.com")
  Image = FakeSearch("image", "Something else", "http://www.google.com")
  Video = FakeSearch("video", "Blah", "http://www.google.com")
)

type SearchFunc func(query string) Result

type FakeSearch(kind, title, url string) SearchFunc {
  return func(query string) Result {
    time.Sleep(time.Duration(rand.Intn(100)) * time.Milliseconds)
    return Result{
      Title: fmt.sprintf("%s(%q): %s", kind, query, title)
      URL: url
    }
  }
}
