<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#introduction">Introduction</a>
      <ol>
        <li><a href="#01_prereq">01_prereq</a>
        <li><a href="#02_templating">02_templating</a>
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

### 01_prereq

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

### 02_templating

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
