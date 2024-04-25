package myjson_test

import (
	"fmt"
	"github.com/misstime/myjson"
	"testing"
)

type Pinyin struct {
	Raw       string `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Pure      string `protobuf:"bytes,2,opt,name=pure,proto3" json:"pure,omitempty"`
	Shengdiao int32  `protobuf:"varint,3,opt,name=shengdiao,proto3,enum=dto.Shengdiao" json:"shengdiao,omitempty"`
}

type EsModelZi struct {
	Id            int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Zi            string    `protobuf:"bytes,2,opt,name=zi,proto3" json:"zi,omitempty"`
	Pinyins       []*Pinyin `protobuf:"bytes,3,rep,name=pinyins,proto3" json:"pinyins,omitempty"`
	FirstLetters  []string  `protobuf:"bytes,20,rep,name=firstLetters,proto3" json:"firstLetters,omitempty"`
	MyDesc        string    `protobuf:"bytes,17,opt,name=myDesc,proto3" json:"myDesc,omitempty"`
	MyPrivateDesc string    `protobuf:"bytes,18,opt,name=myPrivateDesc,proto3" json:"myPrivateDesc,omitempty"`
}

func TestMarshal(t *testing.T) {
	m := &EsModelZi{
		Id:            0,
		Zi:            "",
		Pinyins:       []*Pinyin{{Pure: "hello"}},
		FirstLetters:  nil,
		MyDesc:        "",
		MyPrivateDesc: "",
	}
	jb, err := myjson.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jb))
}
