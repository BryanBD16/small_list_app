module github.com/BryanBD16/smallListApp //go mod init github.com/BryanBD16/smallListApp

//use the go mod tidy command to remove unused dependencies and add the missing ones
//go get github.com/stretchr/testify/assert is used

go 1.22.2

require github.com/stretchr/testify v1.10.0

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.8.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
