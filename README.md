# Dependency injection tool

[![GoDoc](https://godoc.org/github.com/romanyx/shelf?status.svg)](https://godoc.org/github.com/romanyx/shelf)

Shelf allows to easily access dependencies across codebase.

* **Example**.

``` go
func main() {
  shelf.Put[*greeter](&greeter{w: "Hello"}) 
  shelf.Put[*greeter](&greeter{w: "こんにちは"}, "jp") 

  d := shelf.Take[*greeter]()
  d.Hello()
  jp := shelf.Take[*greeter]("jp")
  jp.Hello()

  //
  // Hello
  // こんにちは
}

type greeter struct{
  w string
}

func(g greeter) Hello() {
  fmt.Println(g.w)
}
```

