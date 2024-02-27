# Cancel context with ctrl+C in Go

To run the examples:

```shell
$ go run main.go
...
$ go run signalcontext/main.go
...
```

The code is commented to explain what is going on. Run the programs and try different scenarios to see how they handle ctrl+c. Some suggestions:

- ctrl+c at any moment
- ctrl+c right after seeing a "starting to work"
- two ctrl+c at any moment
- two ctrl+c right after seeing a "starting to work"
- a ctrl+c at any moment, another after "doStuff: context done, cleaning up

The example in `signalcontext/main.go` uses the `signal.NotifyContext` function, added in Go 1.16. It makes handling one ctrl+c simpler, but it does not handle the double ctrl+c.

Sources:

- https://pace.dev/blog/2020/02/17/repond-to-ctrl-c-interrupt-signals-gracefully-with-context-in-golang-by-mat-ryer.html
- https://medium.com/@matryer/make-ctrl-c-cancel-the-context-context-bd006a8ad6ff