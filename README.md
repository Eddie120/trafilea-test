# trafilea-test

To execute Go code, you need to follow these steps:

1. Install Go: If you haven't already, you need to install Go on your system. You can download the official distribution of Go from the Go website (https://golang.org/) and follow the installation instructions specific to your operating system.


2. Set up the Go workspace: Go requires a specific directory structure called the Go workspace. The workspace contains the Go source files, packages, and binaries. By default, the Go workspace is located in the `GOPATH` environment variable. Create a directory for your Go workspace, if you haven't already, and set the `GOPATH` environment variable to point to that directory.

3.Installation:
```shell
$ git clone https://github.com/Eddie120/trafilea-test.git
$ cd project
$ go mod tidy
```

4. Build and run the code: Use the `go run` command to compile and run the Go code. In the terminal, run the following command inside the project:

```shell
cd cmd/trafilea-server
go run main.go
```

The `go run` command compiles and executes the Go code in the specified file (`main.go` in this case).

6. View the output: After executing the `go run` command, you should see ip address printed in the terminal [http://127.0.0.1:8080/docs#/](http://127.0.0.1:8080/docs#/).
7. Some products id that you can use with the API: 
```shell
[
  "PROD-9eab9df2-6d54-11ee-b962-0242ac120002",
  "PROD-9eab9zf2-6d54-11ee-b962-0242ac120002",
  "PROD-9eab9lf2-6d54-11ee-b962-0242ac120002",
  "PROD-9eaba09a-6d54-11ee-b962-0242ac120002",
  "PROD-9eaba1da-6d54-11ee-b962-0242ac120002",
  "PROD-9eaba842-6d54-11ee-b962-0242ac120002",
  "PROD-9eabae28-6d54-11ee-b962-0242ac120002",
  "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"
]
```
**Note**: If you want to see the full catalog please visit: https://github.com/Eddie120/trafilea-test/blob/main/services/models/models.go#L20
