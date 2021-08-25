# linkedlist.go

The first generic linked list in go :dancer:

## gotip

first of all you need to install the master version of golang.

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

- https://github.com/golang/go/issues/47896
