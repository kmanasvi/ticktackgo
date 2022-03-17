### Documentation:

Instructions on how to build/run your application
(Preferred OS- MacOS) 

-  The repo exists within a standard golang route. (go/src/github.com/)
Therefore unzip this file within a go/src/github folder.

- Under the file go/src/github.com/tictacgo/ execute the following MacOS/Linux command-
    source env/development.env
    go run cmd/tictacgo/main.go 
    or 
    under cmd/matrixoperatemicroservice/ Run
    go run .

  Then Send request with: 


  Test - 
  To run unit test for each operation run

  go test test/playgroundService/playground_test.go

    Versions used -
      go version go1.15.3 darwin/amd64
    The backend will be up on localhost:8080.