# go_url
URL shortener in Go

## Structure of the project

```
go_url/
├── cmd/ - where we define commands to be run from the command line
│   └── url_shortener.go - only existing command for now, just starts the url_shortener app
└── server/ (server package)
    ├── server.go - defines the server bahavior (start server, routes, etc)
    └── server_test.go - tests for the server
```

## Building and running the app

Building and re-running the app is still not automated, so every time you change something, to see the changes on your running app you'll need to stop the server and do

```
$ go clean github.com/tuttiq/go_url/cmd/url_shortener
$ go install github.com/tuttiq/go_url/cmd/url_shortener
$ url_shortener
```

Assuming you have your `$GOPATH/bin` folder added to your `$PATH`. If not, then replace the last command by `$GOPATH/bin/url_shortener`.

Access the app from `localhost:8080`on your browser.
