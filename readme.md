# Create Go Package
To Create a Go Lang Package `go mod init todo`. (in JS like npm init).

### Note
Import but not used `_ "github.com/lib/pq"`

### Middleware
go to `http://localhost:8081/api/v1/users`

We need to restart the server each time a change happen by 
`https://github.com/codegangsta/gin` gin library

*Incase error try `GO111MODULE=off go get github.com/codegangsta/gin` from CLI make sure go is path as below

```
export GOPATH=$HOME
export PATH=$PATH:$GOPATH/bin
```



Run the go server by `gin -p 8080 -a 8081`