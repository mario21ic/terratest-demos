Create a env:
```
source gvp
```

Install dependecies:
```
go mod init "github.com/mario21ic/terratest-demos"
go get -u -v -f all
```

Run Tests:
```
go get
go test -v -run TestTerraformHelloWorldExample
go test -v
```
