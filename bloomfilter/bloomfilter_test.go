package bloomfilter

import (
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedisBF_Check(t *testing.T) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	bf := NewRedisBF(conn, 50, 3, "testbf")
	bf.Add([]byte("test"))
	expected := true
	if expected != bf.Check([]byte("test")) {
		t.Errorf("Test should be exist but not exist in BF")
	}

}
