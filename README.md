## Example:

### Why?

How to use packages in Go using 1.16 Go version:

When we have `GO111MODULE` set to `on` Go will automatically try to find packages within the project structure.

### What?

1. We have a `main` module which uses two modules `goodbye` and `hello`

2. Make sure the Go Modules is On:

`export GO111MODULE="on"`

3. Start a go module with `go mod init example/main`, you will also see many examples using `go mod init example.com/main/`. The `example.com` note the `.com` sometimes is used as a repository so in the future, we can use `go get <module_name like example.com/subpackage>` to download from the internet(or repository) and if the new sub-package need to be used in the code.

4. `go mod init <Module Name>` will create the Module in the main which is the place where your code execution should start. It needs to be done once in the root path of the Package/Module.

5. Two ways we can Structure the project and packages(As per new standard of 1.16 Go version and as far my knowledge also ;) ):

5.1 Using Just Folder Names

You can see the branch `withoutmod` to see how a structuring can be maintain using folder names as Package paths.

5.2 Using Generic Names and making each subpackage having `mod` files. In this case, you need additional `go.mod` in each package and here the path can be URL as well, which you will see most commonly.

In this Git project `master` branch code shows this type of structure.

6. `go mod tidy` is a useful command to clean up the structuring of the module in the go project if any changes are made.

7. Additional Read: [Go Tutorial from Official Site](https://golang.org/doc/tutorial/call-module-code)

Common Issue: Package **_/_** is not in GOROOT

Reason: Go is not able to understand your structure, do check the spellings on packages and folder properly.
