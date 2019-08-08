# golang-learn

### Core Basic Commands:

- If you want to see what type of file,which we had build. Just type "file: with compiled file.

    -  EX: file first

- Want to know what is run command dose. 

    - EX : go run -x first.go

- How to run no of files with in the package.

    - Ex: go run *.go / go run learn_package_Ex.go first.go

- To access methods from different file within the package.
    - EX: go run first.go packagescope.go

### create custom library 
 - define your package name as your folder name
 - then build your go file.
 - then install your comiled file.
    EX: 
    -    go run customlib.go
    -    go run: cannot run non-main package
    -    go build
    -    go install
    -    EX: pkg\windows_amd64\github.com\sathishkumar64\golang-learn   
#### Using custom lib,Please refer Printer foder.

#### Package scope always use keyword, like package,var and func. this is the reason we cannot decalre short variable at package scope.


##### to debug race
go run - race hello.go