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
      </ol>
    </li>
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

&nbsp;

---

&nbsp;
