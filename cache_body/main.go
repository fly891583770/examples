package main

import (
	"time"

	"github.com/kataras/iris"
)

var testMarkdownContents = `## Hello Markdown from Iris

This is an example of Markdown with Iris



Features
--------

All features of Sundown are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Thread safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.

	[this is a link](https://github.com/kataras/iris) `

func main() {
	// if this is not setted then iris set this duration to the lowest expiration entry from the cache + 5 seconds
	// recommentation is to left as it's or
	// iris.Config.CacheGCDuration = time.Duration(5) * time.Minute

	iris.Get("/hi", iris.Cache(func(c *iris.Context) {
		c.WriteString("Hi this is a big content, do not try cache on small content it will not make any significant difference!")
	}, time.Duration(10)*time.Second))

	bodyHandler := func(ctx *iris.Context) {
		ctx.Markdown(iris.StatusOK, testMarkdownContents)
	}

	expiration := time.Duration(5 * time.Second)

	iris.Get("/", iris.Cache(bodyHandler, expiration))

	// if expiration is <=time.Second then the cache tries to set the expiration from the "cache-control" maxage header's value(in seconds)
	// // if this header doesn't founds then the default is 5 minutes
	iris.Get("/cache_control", iris.Cache(func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1>Hello!</h1>")
	}, -1))

	iris.Listen(":8080")
}
