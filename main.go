package main

import proto "code.google.com/p/goprotobuf/proto"
import (
	"fmt"
	"log"
)

func main() {
	user := &entry.User{
		Uid:  proto.Int32(1),
		Name: proto.String("blackbeans"),
	}
	encObj, err := proto.Marshal(user)
	if nil == err {
		fmt.Println("length:", len(encObj))
		tobj := &entry.User{}
		e := proto.Unmarshal(encObj, tobj)
		if nil == e {
			fmt.Println(tobj.GetName())
		} else {
			log.Fatalln("decode fail ", e)
		}
	} else {
		log.Fatalln("encode fail", err)
	}

}
