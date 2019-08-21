~ Reminder
```
When checking if a file is wirrten well use the following tools
also sue go fmt to format code
and golint to lint the code
```

# Format of testing
```
BenchmarkFunc(b *testing.B)
ExampleFUnc()
TestFunc(t *testing.T)

# Commands
```
godoc -http=:port
go test
go test -bench .
don’t forget the “.” in the line above
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```