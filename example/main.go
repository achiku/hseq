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
