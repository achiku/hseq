package hseq

import "testing"

type testStruct struct {
	ID           int64
	UniqueNumber int
	UniqueString string
}

func TestGetSeq(t *testing.T) {
	cases := []struct {
		Name   string
		Int    int
		Int64  int64
		String string
	}{
		{Name: "user.id", Int: 1, Int64: int64(1), String: "1"},
		{Name: "user.id", Int: 2, Int64: int64(2), String: "2"},
		{Name: "user.id", Int: 3, Int64: int64(3), String: "3"},
		{Name: "user.name", Int: 1, Int64: int64(1), String: "1"},
		{Name: "user.nickname", Int: 1, Int64: int64(1), String: "1"},
	}

	for _, c := range cases {
		s := Get(c.Name)
		if s.String() != c.String {
			t.Errorf("want %s got %s", c.String, s.String())
		}
		if s.Int() != c.Int {
			t.Errorf("want %d got %d", c.Int, s.Int())
		}
		if s.Int64() != c.Int64 {
			t.Errorf("want %d got %d", c.Int64, s.Int64())
		}
	}
}

func TestCreateStruct(t *testing.T) {
	ts1 := testStruct{
		ID:           Get("testStruct.id").Int64(),
		UniqueNumber: Get("testStruct.UniqueNumber").Int(),
		UniqueString: Get("testStruct.uniqueNumber").String(),
	}
	ts2 := testStruct{
		ID:           Get("testStruct.id").Int64(),
		UniqueNumber: Get("testStruct.UniqueNumber").Int(),
		UniqueString: Get("testStruct.uniqueNumber").String(),
	}

	if ts1.ID == ts2.ID {
		t.Errorf("ts1.ID == ts2.ID")
	}
}
