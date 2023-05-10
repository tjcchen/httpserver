## httpserver
A simple http server implemented with programming language GO

## Makefile
```sh
# run the root target
make

# run the root target
make root

# build your application & generate binary
make build

# containerize your application
make release

# push your application to mirror warehouse
make push
```

## Tree
```sh
# tree display with level 3
tree -L 3

# directory hierarchy
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── bin
│   └── amd64
│       └── httpserver
├── build.sh
├── go.mod
├── go.sum
├── main.go
└── test
    └── pprof.go
```

## License
Apache 2.0
