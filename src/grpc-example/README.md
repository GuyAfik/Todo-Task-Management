# gRPC example
This is an example of a simple gRPC server/clinet archtiecture.
The client is in python while the server is in golang. 
This server has a single service which knows how to calucalte the average of a test scores of a user.

## Running the server 
```
cd $HOME/go/src/github.com/Todo-Task-Management/src/grpc-example
go run server/main.go
```

## Running the client
```
cd $HOME/go/src/github.com/Todo-Task-Management/src/grpc-example
python client/main.py
```

## Go project location
Should be under ```$HOME/go/src/github.com```, for example: ```$HOME/go/src/github.com/Todo-Task-Management/src/grpc-example```

## go.mod & go.sum
A go.mod file is like the package.json file for your Go application.
**The go.mod file contains information like the module name of your project, the version of Go you’re using, and a list of third-party Go modules.**

```go mod init github.com/Todo-Task-Management/src/grpc-example``` - the go module name here will be ```github.com/Todo-Task-Management/src/grpc-example```. 
This will also be the relative path to importing local packages.

[How to import local files packages in go-lang](https://linguinecode.com/post/how-to-import-local-files-packages-in-golang)

## Creating go server code using proto CLI
```protoc --go_out=server users.proto``` - creates the objects in which the gRPC will use.
```protoc users.proto --go-grpc_out=server/``` - creates the actual service interface (API)

## Creating python client code using proto CLI
```python -m grpc_tools.protoc -I. --python_out=client --grpc_python_out=client users.proto```


[gRPC with go-lang and python](https://medium.com/@andersonborges_70700/grpc-with-golang-and-python-f5b7aa602d74)