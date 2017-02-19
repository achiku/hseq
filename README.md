# hseq

[![Build Status](https://travis-ci.org/achiku/hseq.svg?branch=master)](https://travis-ci.org/achiku/hseq)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/hseq/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/hseq)](https://goreportcard.com/report/github.com/achiku/hseq)

Simply generating sequence values for tests


## Why created

In tests, it is very helpful to have unique sequence values for a name, especially when it interact with RDBMS primary key or unique constraints. This library makes it easier to set up unique sequence values in Go tests.


## Installation

```
go get -u github.com/achiku/hseq
```

## How to use

```go
package main

import (
	"log"

	"github.com/achiku/hseq"
)

type st struct {
	ID           int64
	UniqueNumber int
	UniqueName   string
}

func main() {
	st1 := st{
		ID:           hseq.Get("st.id").Int64(),
		UniqueNumber: hseq.Get("st.unique_number").Int(),
		UniqueName:   hseq.Get("st.unique_name").String(),
	}
	st2 := st{
		ID:           hseq.Get("st.id").Int64(),
		UniqueNumber: hseq.Get("st.unique_number").Int(),
		UniqueName:   hseq.Get("st.unique_name").String(),
	}
	log.Printf("%+v", st1)
	log.Printf("%+v", st2)
}
```
