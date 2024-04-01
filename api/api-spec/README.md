# TWELVE Backend

#### Build Docker image

`docker build -f deployments/Dockerfile -t twelve-go .`

#### Run Docker

`docker run -p 8080:8080 twelve-go`

#### Mac only run go (need install go)

go version 1.21.1
`./scripts/start.sh`


go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/cookie
