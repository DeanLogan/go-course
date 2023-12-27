package main

import "fmt"
import "github.com/DeanLogan/go-course/course/12-test-package"

func main(){
	fmt.Println("hello world")
	testpackage.Test()
}

/*

1. go run:
	Description: This command is used to compile and run Go programs. It compiles the Go source code files in the specified package and then executes the compiled binary.
	
	Example: go run filename.go

2. go build: 
	Description: The go build command is used to compile Go source code files into an executable binary. It creates the binary in the same directory as the source code.
	
	Example: go build filename.go

3. go install:
	Description: This command is used to compile and install Go packages and commands. It compiles the Go source code and installs the resulting binary to the 
	$GOPATH/bin directory. If you have your Go workspace properly set up, you can then run the installed binary from any location in your terminal.
	
	Example: go install package-name

4. go get:
	Description: The go get command is used to download and install packages from a remote repository into your Go workspace. 
	It automatically fetches the specified package, its dependencies, and installs them. Additionally, it can also update packages to their latest versions.
	
	Example: go get github.com/example/package
	
	Notes:
		The go get command is often used to acquire third-party packages and libraries for your Go projects.
		It pulls the source code from the specified repository, compiles it, and installs the resulting binaries or libraries to your workspace.
		You can also use go get -u to update packages to their latest versions.
		Example with update: go get -u github.com/example/package
		
		Example with a specific version: go get github.com/example/package@v1.2.3
		
		Example with multiple packages: go get github.com/example/package1 github.com/example/package2

*/
