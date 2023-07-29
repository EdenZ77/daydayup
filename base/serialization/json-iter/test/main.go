package main

import (
	"fmt"
	jsonIter "github.com/json-iterator/go"
)

type Tes struct {
	Protocols *Protocols
}

type Protocols struct {
	Http  *Protocol `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Https *Protocol `protobuf:"bytes,2,opt,name=https,proto3" json:"https,omitempty"`
}

type Protocol struct {
	SiteCats []int32 `protobuf:"varint,1,rep,packed,name=site_cats,json=siteCats,proto3" json:"siteCats,omitempty"`
	Files    []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

type File struct {
	FileCats []int32 `protobuf:"varint,1,rep,packed,name=file_cats,json=fileCats,proto3" json:"fileCats,omitempty"`
	Path     string  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func main() {

	//str := "{\"http\":{\"siteCats\":[1283]},\"https\":{\"siteCats\":[1283]}}"
	//str := "{\"http\":{\"files\":[{\"fileCats\":[1283],\"path\":\"login.php\"}]}}"
	//protocols := &Protocols{}
	//err := jsonIter.UnmarshalFromString(str, protocols)
	//if err != nil {
	//	return
	//}
	//fmt.Println(protocols)
	//fmt.Println(protocols.Http)
	//if len(protocols.Http.SiteCats) > 0 {
	//
	//	fmt.Println(protocols.Http.SiteCats)
	//}
	//if len(protocols.Http.Files) > 0 {
	//	fmt.Println(protocols.Http.Files[0])
	//	if len(protocols.Http.Files[0].FileCats) > 0 {
	//		fmt.Println(protocols.Http.Files[0].FileCats)
	//		fmt.Println(protocols.Http.Files[0].Path)
	//	}
	//}

	//fmt.Println(protocols.Https)

	str := "{\"protocols\":{\"http\":{\"files\":[{\"fileCats\":[1283],\"path\":\"login.php\"}]}}}"

	protocols := &Tes{}

	err := jsonIter.UnmarshalFromString(str, protocols)
	if err != nil {
		return
	}

	fmt.Println(protocols.Protocols.Http.Files[0].FileCats)
	fmt.Println(protocols.Protocols.Http.Files[0].Path)

}
