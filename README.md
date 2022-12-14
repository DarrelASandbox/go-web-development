<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#introduction">Introduction</a>
      <ol>
        <li><a href="#01-prereq">01-prereq</a>
        <li><a href="#02-templating">02-templating</a>
      </ol>
    </li>
    <li><a href="#creating-server">Creating Server</a>
      <ol>
        <li><a href="#03-server-tcp">03-server-tcp</a>
        <li><a href="#04-server-net-http">04-server-net-http</a>
        <li><a href="#05-server-servemux">05-server-servemux</a>
        <li><a href="#06-server-serve-files">06-server-serve-files</a>
      </ol>
    </li>
    <li><a href="#creating-state">Creating State</a></li>
    <li><a href="#creating-sessions">Creating Sessions</a></li>
    <li><a href="#mysql">MySQL</a></li>
    <li><a href="#aws">AWS</a></li>
  </ol>
</details>

&nbsp;

## About The Project

- Web Development w/ Google’s Go (golang) Programming Language
- Learn Web Programming from a University Professor in Computer Science with over 20 years of teaching experience.
- [Todd McLeod](https://github.com/GoesToEleven)
- [Original Repo: golang-web-dev](https://github.com/GoesToEleven/golang-web-dev)

&nbsp;

---

&nbsp;

## Introduction

- [GO FAQ - Why did you create a new language?](https://go.dev/doc/faq#creating_a_new_language)

### 01-prereq

- Language review
  - **variables**
    - short variable declaration operator
    - using the var keyword to declare a variable
    - scope
  - **data structures**
    - slice
    - map
    - struct
      - composite literal
  - **functions**
    - func `(receiver) identifier(parameters) (returns) { <code> }`
    - methods
  - **[composition](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)**
    - embedded types
    - interfaces
    - polymorphism

### 02-templating

- [GO Standard Library - text/template](https://pkg.go.dev/text/template#Template)
- [Effective Go - The init function](https://go.dev/doc/effective_go#init)
- [GO Standard Library - text/template - binary comparison operators](https://pkg.go.dev/text/template?utm_source=godoc#hdr-Functions)
- [thoughtworks - Blogs Banner Composition vs. Inheritance: How to Choose?](https://www.thoughtworks.com/insights/blog/composition-vs-inheritance-how-choose)
- [composition with Go here](https://www.goinggo.net/2015/09/composition-with-go.html)

&nbsp;

---

&nbsp;

> **Galen:** Why would you use a FuncMap over a Method?

> **Todd:** A method comes with a VALUE of a certain TYPE when that VALUE is sent to a template. If you need a function in a template NOT ATTACHED TO A TYPE, then you use funcmap. And then, in the template, you can call that function. The similarities, the "two roads", are that they are both functions in templates. Best practice is to do all data massaging before sending into template; though at times, having functions used in templates is a good call.

&nbsp;

---

&nbsp;

## Creating Server

- **Web Programming Synonymous Terms**
  - router
  - request router
  - multiplexer
  - mux
  - servemux
  - server
  - http router
  - http request router
  - http multiplexer
  - http mux
  - http servemux
  - http server
- **Request & response**
  - a request/response line
  - zero or more header lines
  - a blank line (ie, CRLF by itself)
  - an optional message body
- **HTTP request**
  - Request
    - request line
    - headers
    - optional message body
  - Request-Line
    - Method SP Request-URI SP HTTP-Version CRLF
    - example request line: `GET /path/to/file/index.html HTTP/1.1`
- **HTTP response**
  - Response
    - status line
    - headers
    - optional message body
  - Status-Line
    - HTTP-Version SP Status-Code SP Reason-Phrase CRLF
    - example status line: `HTTP/1.1 200 OK`
- **Headers**
  - [List of header fields](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)
- **Inspect**
  - you can use google chrome dev tools / network
  - you can use CURL at the command line: `curl -v golang.org`

### 03-server-tcp

- [GO Standard Library - net](https://pkg.go.dev/net)
- [GO Standard Library - bufio](https://pkg.go.dev/bufio)

```sh
# MacOS
nc localhost 8080

# Windows
telnet localhost 8080
```

### 04-server-net-http

- [GO Standard Library - net/http](https://pkg.go.dev/net/http)
- [**http.handler** (This is one of the most important things to know)](https://pkg.go.dev/net/http#Handler)

```Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

- [**http.ListenAndServe**](https://pkg.go.dev/net/http#ListenAndServe)
  - `func ListenAndServe(addr string, handler Handler) error`
- [**http.ListenAndServeTLS**](https://pkg.go.dev/net/http#ListenAndServeTLS)
  - `func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error`
- _Notice that both of the above functions take a handler_
- [**http.Request**](https://pkg.go.dev/net/http#Request)
- Here it is with _most of the comments and some of the fields_ stripped out:

```go
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging.
    RemoteAddr string
}
```

- Some interesting things you can do with a request:
  - Retrieve URL & Form data
  - `http.Request` is a struct. It has the fields `Form` & `PostForm`.
  - If we read the documentation on these, we see:

```
    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values

```

- If we look at **ParseForm**, `go func (r *Request) ParseForm() error `
- we see that this is a method attached to a \*http.Request.
- If we look at **FormValue\***
- ` go func (r *Request) FormValue(key string) string`
- we see that this is a method attached to a \*http.Request. FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.
- The `http.Request` type is a struct which has a `Method` field.
- The `http.Request` type is a struct which has a `URL` field. Notice that the type is a `*url.URL`
- Take a look at type `url.URL`

```go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```

- The `http.Request` type is a struct which has a `Header` field.
- From the documentation about the `http.Header` struct, we see that: `type Header map[string][]string`
- [**http.ResponseWriter**](https://godoc.org/net/http#ResponseWriter)

```Go
type ResponseWriter interface {
    // Header returns the header map that will be sent by
    // WriteHeader. Changing the header after a call to
    // WriteHeader (or Write) has no effect
    Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
    //
    // If WriteHeader has not yet been called, Write calls
    // WriteHeader(http.StatusOK) before writing the data. If the Header
    // does not contain a Content-Type line, Write adds a Content-Type set
    // to the result of passing the initial 512 bytes of written data to
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with status code.
    // If WriteHeader is not called explicitly, the first call to Write
    // will trigger an implicit WriteHeader(http.StatusOK).
    // Thus explicit calls to WriteHeader are mainly used to
    // send error codes.
    WriteHeader(int)
}
```

- An `http.ResponseWriter` has a method `Header()` which returns a `http.Header`.
- Look at the documentation for `http.Header`: `type Header map[string][]string`
- Look at the methods which are attached to type `http.Header`

```go
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

- We can set headers for a response like this: `res.Header().Set("Content-Type", "text/html; charset=utf-8")`

&nbsp;

---

&nbsp;

> **Amir Hossein:** Why the Handler implementation needs a POINTER to Request? Is this a big value or we need to manipulate that in the Handler?

> **Todd:** They probably made that choice so that they're just passing the address instead of all of the data; not necessarily because they need it to be mutatable.

> **Philippe:** My guess is efficiency. Think about the extra cost of serving millions of users if you make a copy for every request.

&nbsp;

---

&nbsp;

### 05-server-servemux

- In electronics, a [multiplexer (or mux)](https://en.wikipedia.org/wiki/Multiplexer) is a device that selects one of several input signals and forwards the selected input into a single line.
- The term multiplexer has been adopted by web programming to refer to the process of routing requests.
- A web server has requests coming in at different routers and via different HTTP methods. For instance, we might have these requests:
  - REQUEST #1
    - Path: /cat
    - Method: GET
  - REQUEST #2
    - Path: /apply
    - Method: Get
  - Request #3
    - Path: /apply
    - Method: Post
- Based upon the requests coming in, the server needs to determine how to respond to that request - for each request that comes in, different code will be run.
- **ServeMux** is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
- Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash).
- Longer patterns take precedence over shorter ones, so that if there are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will be called for paths beginning "/images/thumbnails/" and the former will receive requests for any other paths in the "/images/" subtree. Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".
- If a subtree has been registered and a request is received naming the subtree root without its trailing slash, ServeMux redirects that request to the subtree root (adding the trailing slash). This behavior can be overridden with a separate registration for the path without the trailing slash. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately.
- Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/".
- ServeMux also takes care of sanitizing the URL request path, redirecting any request containing . or .. elements or repeated slashes to an equivalent, cleaner URL.
- [http.ServeMux](https://pkg.go.dev/net/http#ServeMux)

```Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

- Any value of type `*http.ServeMux` implements the `http.Handler` interface.
- Remember, the `http.Handler` interface requires that a type have the `ServeHTTP` method.

```Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

- What this tells us is that we can pass a value of type `*http.ServeMux` into `http.ListenAndServe`

```Go
func ListenAndServe(addr string, handler Handler) error
```

- You can also see from the documentation of `*http.ServeMux` ...

```Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

- ... that we have a method `Handle` attached the the value of type `*http.ServeMux`. You can see that `Handle` takes a pattern and a `http.Handler`.
- We can use `Handle` like this:

```Go
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/cat", c)
```

- The overall game plan: We will create a NewServeMux, then attach the method `Handle` to it to set routes, then pass our `*http.ServeMux` to `http.ListenAndServe`.
- **DefaultServeMux**
  - ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux: `http.ListenAndServe(":8080", nil)`
- [GO package - julienschmidt/httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter)
  - `go get "github.com/julienschmidt/httprouter"`

&nbsp;

---

&nbsp;

### 06-server-serve-files

- [http.ServeContent](https://pkg.go.dev/net/http#ServeContent)
  - `func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)`
- [http.ServeFile](https://pkg.go.dev/net/http#ServeFile)
  - `func ServeFile(w ResponseWriter, r *Request, name string)`
- [http.FileServer](https://pkg.go.dev/net/http#FileServer)
  - `func FileServer(root FileSystem) Handler`
- [http.StripPrefix](https://pkg.go.dev/net/http#StripPrefix)
  - `func StripPrefix(prefix string, h Handler) Handler`

&nbsp;

---

&nbsp;

## Creating State

- [rfc7231](https://datatracker.ietf.org/doc/html/rfc7231)

&nbsp;

---

&nbsp;

> **Andy:** when you have `.Close()`?
>
> I see you have `defer f.Close()` and `defer dst.Close()` in [this file](https://github.com/GoesToEleven/golang-web-dev/blob/master/027_passing-data/05_form-file/02_store/main.go), I am wondering, how do I know I need to use a `.Close()`, because if I didn't close it, it's not the good practice. but how can I remember to have `.Close()` and write "safe" code? Or is it just all from experience? thanks

> **Todd:** so if you look at [this](https://godoc.org/net/http#Request.FormFile) and then look at [this](https://godoc.org/mime/multipart#File). You see this

```Go
    type File interface {
        io.Reader
        io.ReaderAt
        io.Seeker
        io.Closer
    }
```

> And when something has a method "Closer" that means you probably will want to call it, eg, close it. You can also see that here, as another [example](https://godoc.org/os#File). One of the methods on a `os.File` is `Close()`. Plus, just think about whether or not this is a resource that is being opened which might need to be closed: a file, a connection, etc.

&nbsp;

---

&nbsp;

## Creating Sessions

- This is how we create state:
  - We will store a unique ID in the cookie.
  - On our server, we will associate each user with a unique ID.
  - This will allow us to identify every user visiting our website.
- Security
  - There are two factors which contribute to the security of a session created using a cookie and a unique ID:
    - Uniqueness of the ID
    - Encryption in transit with HTTPS
  - You can use any unique ID that you would like: a Universally unique identifier [(UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier) or a database key. If you're using a database key, make sure it's not the key to your user but to a separate session table.
  - A UUID is very unique. [Wikipedia says this about UUIDs:](https://en.wikipedia.org/wiki/Universally_unique_identifier) " ... only after generating 1 billion UUIDs every second for the next 100 years, the probability of creating just one duplicate would be about 50%."
  - A UUID cannot be intercepted in transit if we are using HTTPS.

&nbsp;

---

&nbsp;

## MySQL

1. install mysql
2. install workbench
3. create mysql db on aws
4. connect workbench to rds mysql db
5. [Easy tutorial for connecting remotely to AWS RDS database using MYSQL Workbench](https://www.youtube.com/watch?v=k68Y-XYapEI)

&nbsp;

---

&nbsp;

## AWS

- **Create EC2 security groups:**
  - A security group acts as a virtual firewall that controls the traffic for one or more instances.
  - When you launch an instance, you associate one or more security groups with the instance. You add rules to each security group that allow traffic to or from its associated instances.
- **Create an ELB load balancer:**
  - Load balancing improves the distribution of workloads across multiple computing resources.
  - Elastic Load Balancing automatically distributes incoming application traffic across multiple Amazon EC2 instances. It enables you to achieve fault tolerance in your applications, seamlessly providing the required amount of load balancing capacity needed to route application traffic.
  - Elastic Load Balancing offers two types of load balancers that both feature high availability, automatic scaling, and robust security.
  - These include the Classic Load Balancer that routes traffic based on either application or network level information, and the Application Load Balancer that routes traffic based on advanced application level information that includes the content of the request.
  - The Classic Load Balancer is ideal for simple load balancing of traffic across multiple EC2 instances, while the Application Load Balancer is ideal for applications needing advanced routing capabilities, microservices, and container-base architectures.
  - Application Load Balancer offers ability to route traffic to multiple services or load balance across multiple ports on the same EC2 instance.
- **Implementing the load balancer:**
  - We will create an image of our web server.
  - We will then create a new instance of a web server from our image.
  - We will make sure both of our web servers are added to the target group being used by the load balancer.
  - We will then test that the load balancer is working correctly.

&nbsp;

---

&nbsp;
