package main

import (
	"testing"

	"github.com/davidswisa/protobuf-vs-json/pbjson"
	proto "github.com/golang/protobuf/proto"
)

func BenchmarkProtoMarshall(b *testing.B) {
	// m := &Large{
	// 	Id:   "id",
	// 	Name: "name",
	// 	Lit:  true,
	// 	Info: &Large_Info{
	// 		Id:   "id",
	// 		Name: "name",
	// 		Lit:  true,
	// 	},
	// }
	s := &pbjson.Small{
		Name:   "name",
		Height: 1.0,
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(s)
	}
}
