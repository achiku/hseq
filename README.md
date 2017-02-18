# hseq

[![Build Status](https://travis-ci.org/achiku/hseq?branch=master)](https://travis-ci.org/achiku/hseq)
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
func TestCreateStruct(t *testing.T) {
	ts1 := testStruct{
		ID:           GetSeq("testStruct.id").Int64(),
		UniqueNumber: GetSeq("testStruct.UniqueNumber").Int(),
		UniqueString: GetSeq("testStruct.uniqueNumber").String(),
	}
	ts2 := testStruct{
		ID:           GetSeq("testStruct.id").Int64(),
		UniqueNumber: GetSeq("testStruct.UniqueNumber").Int(),
		UniqueString: GetSeq("testStruct.uniqueNumber").String(),
	}

	if ts1.ID == ts2.ID {
		t.Errorf("ts1.ID == ts2.ID")
	}
}
```
