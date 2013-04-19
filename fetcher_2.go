package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (result *fakeResult, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, resCh chan<- *resultWithUrl, errCh chan<- error) {
    // TODO: Don't fetch the same URL twice.
    if depth <= 0 {
        return
    }
    result, err := fetcher.Fetch(url)    
    if err != nil {
        errCh <- err
        return
    }
    
    resCh <- &resultWithUrl {url, result}    

    for _, u := range result.urls {
        go Crawl(u, depth-1, fetcher, resCh, errCh)
    }
    return
}

func printer(resCh <-chan *resultWithUrl, errCh <-chan error) {
    for {
        select {
            case res := <- resCh:  
                fmt.Printf("found: %s %q\n", res.url, res.result.body)  
            case err := <- errCh:
                fmt.Println(err)       
        }
    }
}

func main() {
    // TODO: how to avoid fetching the same url twice? - pass list of seen along? what about sync? could be done by ref, but would not be thread-safe

    resCh := make(chan *resultWithUrl, 100)
    errCh := make(chan error, 100)

    go Crawl("http://golang.org/", 4, fetcher, resCh, errCh)
    go printer(resCh, errCh)

    var input string
    fmt.Scanln(&input)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type resultWithUrl struct {
    url string
    result *fakeResult
}

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (*fakeResult, error) {
    if res, ok := (*f)[url]; ok {
        return res, nil
    }
    return nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
