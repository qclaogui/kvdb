<div align="center">
  <h1>kvDB</h1>
</div>
<p align="center">
<a href="https://travis-ci.org/qclaogui/kvdb"><img src="https://travis-ci.org/qclaogui/kvdb.svg?branch=master"></a>
<a href="https://goreportcard.com/report/github.com/qclaogui/kvdb"><img src="https://goreportcard.com/badge/github.com/qclaogui/kvdb?v=1" /></a>
<a href="https://godoc.org/github.com/qclaogui/kvdb"><img src="https://godoc.org/github.com/qclaogui/kvdb?status.svg"></a>
<a href="https://github.com/qclaogui/kvdb/blob/master/LICENSE"><img src="https://img.shields.io/github/license/qclaogui/kvdb.svg" alt="License"></a>
</p>
key value DB


## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/qclaogui/kvdb"
)

func main() {
	m := kvdb.NewMem()
	m.Put("/app/redis/username", "qclaogui")
	m.Put("/app/redis/password", "123456789")
	m.Put("/app/port", "80")
	v, err := m.Get("/app/redis/username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %s\n", v)

	if ks, err := m.GetMany("/app/*/*"); err == nil {
		for _, v := range ks {
			fmt.Printf("Value: %s\n", v)
		}
	}
	// Output:
	// Value: qclaogui
	// Value: 123456789
	// Value: qclaogui
}
```