# Purpose

The purpose of this repository is to demonstrate how to implement the
Middleware/Adapter pattern in Golang HTTP servers. I found it confusing
as to what exactly the differences between the "http.Handle" and
"http.HandleFunc" types were, and my misunderstanding of them made it
difficult for me to implement custom middleware for my web applications.

The part that confused me was that developers are implementing middleware
for either Http Handlers or Http HandlerFuncs. It appears that
you can can certainly do both if you wanted
(apply a middleware, such as `RequireAuth` to all `/users/*` endpoints, and a
separate middelware like `IsUser` to `/users/me`). Generally, it seems most
resources on the internet use either `http.Handle` or `http.HandleFunc`, and
not both.

It's important to note that `http.HandleFunc` is wrapped under-the-hood to make
it implement the `http.Handle` interface (`ServeHTTP f(w, r)`). The benefit
of using `http.Handle` over `http.HandleFunc` is the ability to pass additional
parameters from your middleware through to your handler function, since
`http.HandleFunc` requires a function with only two parameters.

See the docs (mentioned below).

# References

Great reads on understanding the differences. Some of the code is sourced
directly from these folks.

1. Http Handle: https://golang.org/pkg/net/http/#Handle
1. Http HandleFunc: https://golang.org/pkg/net/http/#HandleFunc
1. Http Handle vs HandleFunc: https://www.calhoun.io/why-cant-i-pass-this-function-as-an-http-handler/
1. Http Handle and Middleware: https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
1. Http HandleFunc and Middleware: https://medium.com/@chrisgregory_83433/chaining-middleware-in-go-918cfbc5644d
1. More on HttpHandle: https://drstearns.github.io/tutorials/gomiddleware/