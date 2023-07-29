# Linked list in Go

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/1995parham/linkedlist/ci.yaml?label=ci&logo=github&style=flat-square&branch=main)
[![Codecov](https://img.shields.io/codecov/c/gh/1995parham/linkedlist?logo=codecov&style=flat-square)](https://codecov.io/gh/1995parham/linkedlist)

As you know generics will come to go 1.18 and one of the major drawbacks in go was implementing data structure
because of the lack of generics.
I implemented a small generic linked list in go and I think we can start having brand new data structures in Go.

## ~gotip~

First of all you need to install the master version of golang
and for this you can use `gotip`.

```sh
go install golang.org/dl/gotip@latest
gotip download
```

then you can use the `gotip` command as your normal `go` command.

## Examples

```go
func main() {
        l := list.New[int]()

        l.PushFront(10)
        l.PushFront(20)
        l.PushFront(40)

        fmt.Println(l)
}
```

```go
func main() {
        l := list.New[string]()

        l.PushFront("hello")

        fmt.Println(l)
}
```

```go
func main() {
        l := list.New[int]()

        l.PushFront(10)
        l.PushFront(20)
        l.PushFront(40)
        l.PushFront(42)

        fmt.Println(l)

        s := l.Filter(func(i int) bool {
                return i%10 == 0
        })

        fmt.Println(s)
}
```

## Related Issues

- <https://github.com/golang/go/issues/47896>
