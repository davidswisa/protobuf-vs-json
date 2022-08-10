package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davidswisa/protobuf-vs-json/pbjson"
	proto "github.com/golang/protobuf/proto"
)

var (
	ipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce posuere nisl eu orci dignissim malesuada. Integer nec dapibus felis. Cras sed mollis urna. Curabitur pretium, ipsum id viverra dictum, neque ex vehicula metus, eget placerat ligula metus a velit. Quisque efficitur lectus magna, eu fermentum mi commodo ac. Pellentesque massa ipsum, accumsan porttitor ultrices id, dictum ac sapien. Suspendisse ut magna ac ante mollis vestibulum. Ut turpis tellus, dictum in congue vel, suscipit id risus. Donec sed massa a leo volutpat iaculis. Nullam ut ultrices nisi, ac venenatis dolor. Nam in felis quis augue bibendum vestibulum nec id ligula. In lacinia vel ante vel lobortis. Praesent rutrum ligula quis dapibus bibendum. Aliquam erat volutpat. Sed mollis, velit in elementum dignissim, eros orci tincidunt risus, vitae pharetra metus libero in orci. Aliquam convallis, lacus quis volutpat malesuada, lectus nibh congue diam, at dictum quam dui vitae metus.

	Nam pharetra consequat ex. Suspendisse ut convallis enim. Phasellus auctor enim vel augue blandit fermentum. Proin euismod congue tellus, eu mollis lectus facilisis sed. Quisque auctor, sem non malesuada lobortis, lacus magna pellentesque dolor, eu maximus augue felis vel tortor. Sed iaculis magna dolor, sit amet dapibus orci tristique sed. Duis id metus id tellus convallis tincidunt nec eget tellus. Proin id ipsum tempus massa pellentesque sagittis eu eleifend libero. Donec pulvinar congue diam, vel hendrerit ligula elementum vel. Nullam in ex nibh. Curabitur interdum augue non neque vestibulum, facilisis vestibulum quam tempus.
	
	Maecenas tincidunt ultricies libero, in consequat mauris tristique sed. Aliquam mattis nunc nec est vulputate auctor. Duis sit amet convallis enim. Suspendisse tempor venenatis laoreet. Quisque condimentum vulputate eros, sit amet iaculis nunc euismod eu. Suspendisse tempor lobortis finibus. Ut dapibus gravida nunc, in vestibulum felis interdum et. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Curabitur sem metus, sodales convallis dui ac, sollicitudin maximus ex. Curabitur lobortis elementum ex, eget pellentesque justo placerat quis. Duis dictum risus at velit imperdiet, quis dignissim lacus vehicula. Maecenas id lobortis odio.
	
	Donec vitae lacinia erat, at malesuada dolor. Pellentesque nec urna justo. Nullam in porta nisi. Phasellus vitae lacinia odio, eu tempus nisl. Nunc commodo, leo a feugiat luctus, diam felis pharetra ligula, sed dapibus mauris enim sed magna. Morbi ut tincidunt mi, aliquet iaculis ante. Nunc volutpat vel ipsum vitae imperdiet. Integer eu dignissim nulla, sed fermentum augue. Suspendisse pulvinar eu sapien sed tincidunt. Nullam pharetra dapibus orci vitae scelerisque. Praesent at quam commodo, mattis nunc sed, commodo tortor. Fusce quam nibh, aliquam ut nibh non, imperdiet rutrum magna. Donec vitae augue luctus, mattis tellus vel, rhoncus purus.
	
	Praesent in laoreet leo. Donec viverra interdum lorem vitae lacinia. Duis placerat erat eget molestie scelerisque. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Proin sem nisl, pretium a vehicula in, dictum non enim. Vestibulum vehicula dolor ac euismod tempus. Curabitur vitae mauris ut lectus varius lacinia. Quisque nec dui et quam consectetur sagittis vel sed arcu. Fusce feugiat, augue eget tempus malesuada, urna felis laoreet ante, a faucibus tellus ligula sit amet turpis. Mauris vehicula tellus a metus sodales ultrices.
	
	Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Suspendisse ultrices ultricies enim, at cursus nisl. Nunc quis nibh ac nisl malesuada commodo a quis tellus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent sed aliquam dolor. Donec pretium augue quis augue eleifend, ornare mollis turpis finibus. Fusce viverra metus ac finibus lobortis. Sed scelerisque tellus a nisl ultrices gravida. Vestibulum ante leo, viverra quis mauris vel, pretium euismod dolor. Sed euismod non enim sed consequat. Ut varius tempus elit vel blandit. Nulla tempus eu urna non porttitor. Proin est nisi, efficitur in lobortis eget, molestie at dolor. Ut quis lacinia justo, ac venenatis mauris. Aenean eget euismod mauris, efficitur gravida risus.
	
	Fusce est mi, consequat ut ligula quis, egestas iaculis diam. Donec erat enim, commodo in viverra non, finibus ac augue. Nunc non auctor libero. Nunc ante diam, euismod eu condimentum sit amet, laoreet rutrum massa. Vestibulum ac imperdiet erat, in feugiat leo. Fusce tincidunt faucibus eros, nec bibendum sem vehicula sit amet. Nam ac dictum nisl. Nullam feugiat tellus ut enim varius dictum.
	
	Aenean nunc est, consequat vitae dolor id, auctor dignissim lectus. Mauris accumsan neque nec consectetur porta. Etiam porttitor, justo sed posuere ornare, metus ipsum ultrices ipsum, quis sollicitudin lacus ligula id risus. Aenean vulputate non leo id mattis. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec et volutpat dolor. Phasellus molestie felis a consectetur ullamcorper. Praesent in enim velit. Mauris consequat nibh vel dictum condimentum. Duis in congue odio. Donec est odio, blandit eget dolor ac, blandit semper mauris.
	
	Aliquam non tempor turpis. Fusce odio est, commodo quis faucibus at, fringilla ac metus. Vestibulum eu urna magna. Nam aliquet ornare pellentesque. Nulla vulputate pretium hendrerit. Sed vel velit sed velit faucibus vehicula. Sed suscipit, velit in egestas sagittis, augue nisl varius nulla, vitae lacinia lectus velit a dolor. Integer et mi nec nulla eleifend accumsan. Aliquam egestas eros vel eros semper faucibus. Ut facilisis dictum pellentesque. Aliquam cursus convallis lacinia. Morbi id massa eu arcu interdum tincidunt quis vitae leo. Duis vitae felis sed magna vulputate maximus eget nec urna. Nullam elementum ultrices velit, ut laoreet est malesuada fringilla. Nullam sed nisl nec neque faucibus vestibulum at ac quam.
	
	Aenean aliquam elit vitae ultricies vehicula. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus consequat neque velit, in lacinia mauris pharetra eget. Fusce et ante mattis, finibus nulla id, auctor enim. Phasellus suscipit nisi ut tellus auctor facilisis. Sed mollis luctus faucibus. Suspendisse volutpat neque sapien, id congue nulla elementum nec. Donec sed lectus vitae lacus suscipit gravida nec in lorem. Praesent vehicula eu lacus eu interdum. Sed viverra cursus tellus vehicula ornare. Curabitur eleifend eu purus a tincidunt. Mauris dui risus, ultrices ac magna in, consectetur mattis mi. Quisque ac consectetur nulla, id volutpat lacus. Aenean non tortor vitae magna dapibus sodales.`

	small = &pbjson.Small{
		Name:   "name",
		Height: 1.0,
	}

	medium = &pbjson.Medium{
		Name:        "name",
		Height:      1.0,
		Description: []byte(ipsum),
		Num:         -3.4e+38,
		Tru:         true,
	}

	large = &pbjson.Large{
		Name:        "name",
		Height:      1.0,
		Description: []byte(ipsum),
		Num:         -3.4e+38,
		Tru:         true,
		Data:        ExplodingData(),
	}
)

func ExplodingData() []*pbjson.Large_Info {
	var data []*pbjson.Large_Info
	for i := 0; i < 100; i++ {
		data = append(data, &pbjson.Large_Info{
			Name: fmt.Sprintf("name%d", i),
			Id:   fmt.Sprintf("id%d", i),
			Lit:  true,
		})
	}
	return data
}

func BenchmarkProtoMarshall(b *testing.B) {

	s := small
	m := medium
	l := large

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := proto.Marshal(s)
			_ = d
		}
	})

	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := proto.Marshal(m)
			_ = d
		}
	})

	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := proto.Marshal(l)
			_ = d
		}
	})
}

func BenchmarkJSONMarshall(b *testing.B) {

	s := small
	m := medium
	l := large

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := json.Marshal(s)
			_ = d
		}
	})

	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := json.Marshal(m)
			_ = d
		}
	})

	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			d, _ := json.Marshal(l)
			_ = d
		}
	})
}

func BenchmarkProtoUnmarshall(b *testing.B) {

	s := small
	m := medium
	l := large

	sMarsh, _ := proto.Marshal(s)
	mMarsh, _ := proto.Marshal(m)
	lMarsh, _ := proto.Marshal(l)

	var sUnmarsh pbjson.Small
	var mUnmarsh pbjson.Medium
	var lUnmarsh pbjson.Large

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = proto.Unmarshal(sMarsh, &sUnmarsh)
		}
	})

	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = proto.Unmarshal(mMarsh, &mUnmarsh)
		}
	})

	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = proto.Unmarshal(lMarsh, &lUnmarsh)
		}
	})
}

func BenchmarkJSONUnmarshall(b *testing.B) {

	s := small
	m := medium
	l := large

	sMarsh, _ := json.Marshal(s)
	mMarsh, _ := json.Marshal(m)
	lMarsh, _ := json.Marshal(l)

	var sUnmarsh pbjson.Small
	var mUnmarsh pbjson.Medium
	var lUnmarsh pbjson.Large

	b.Run("Small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(sMarsh, &sUnmarsh)
		}
	})

	b.Run("Medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(mMarsh, &mUnmarsh)
		}
	})

	b.Run("Large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(lMarsh, &lUnmarsh)
		}
	})
}
