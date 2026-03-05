package main

import (
	"encoding/json"
	"sync"
	"testing"
)

type student struct {
	name   string
	age    int32
	remark [1024]byte
}

var buf, _ = json.Marshal(student{name: "lxy", age: 18})

var studentpool = sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

func Test_benchmarkunmarshal(t *testing.T) {
	for n := 0; n < 10000; n++ {
		stu := &student{}
		json.Unmarshal(buf, stu)
	}
}

func Test_benchmarkunmarshalwithpool(t *testing.T) {
	for n := 0; n < 10000; n++ {
		stu := studentpool.Get().(*student)
		json.Unmarshal(buf, stu)
		studentpool.Put(stu)
		//s := studentpool.Get().(*student)
		//fmt.Println(s)
	}
}

// 注意命名规范 Benchmark+首字母大写的方法名 参数固定
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "aaaaaaaaa　aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		go a(str, nil)
	}
}
func BenchmarkB(b *testing.B) {
	p := new(sync.Pool)
	t := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	p.Put(t)
	for i := 0; i < b.N; i++ {
		go a(nil, p)
	}
}

func a(s interface{}, p *sync.Pool) interface{} {
	if s == nil {
		return p.Get()
	} else {
		return s
	}
}
