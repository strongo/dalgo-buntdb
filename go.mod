module github.com/dal-go/dalgo2buntdb

go 1.21

toolchain go1.22.4

require (
	github.com/dal-go/dalgo v0.12.1
	github.com/dal-go/dalgo-end2end-tests v0.0.36
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/buntdb v1.3.1
)

//replace github.com/dal-go/dalgo => ../dalgo
//replace github.com/dal-go/dalgo-end2end-tests => ../dalgo-end2end-tests

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/strongo/random v0.0.1 // indirect
	github.com/strongo/validation v0.0.6 // indirect
	github.com/tidwall/btree v1.6.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
