# go_url
URL shortener in Go

## Setup your environment

Install go following [these instructions](https://golang.org/doc/install).

It's [recommended](https://golang.org/doc/code.html#Organization) to organize your go code under a path correspondent to where the package will be published. So create the folder `tuttiq` under `$GOPATH/src/github.com/` if you don't have it yet.

```
cd $GOPATH/src/github.com/tuttiq
git clone git@github.com:tuttiq/go_url.git
```

Build and run the server

```
go build cmd/server.go
./server
```

Access `localhost:8080` on your browser and voilá!