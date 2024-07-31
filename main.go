package main

import (
	"encoding/hex"
	"fmt"
	"test/test/pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	example1 := &pb.Test{Dummy: "dummy"}
	fmt.Printf("example: plain: empty\n%s\n", hex.Dump(mustMarshal(example1)))

	example2 := &pb.Test{V: false, Dummy: "dummy"}
	fmt.Printf("example: plain: false\n%s\n", hex.Dump(mustMarshal(example2)))

	example3 := &pb.Test{V: true, Dummy: "dummy"}
	fmt.Printf("example: plain: true\n%s\n", hex.Dump(mustMarshal(example3)))

	example4 := &pb.Test2{Dummy: "dummy"}
	fmt.Printf("example: optional: empty\n%s\n", hex.Dump(mustMarshal(example4)))

	b1 := false
	example5 := &pb.Test2{V: &b1, Dummy: "dummy"}
	fmt.Printf("example: optional: false\n%s\n", hex.Dump(mustMarshal(example5)))

	b2 := true
	example6 := &pb.Test2{V: &b2, Dummy: "dummy"}
	fmt.Printf("example: optional: true\n%s\n", hex.Dump(mustMarshal(example6)))

	example7 := &pb.Test3{Dummy: "dummy"}
	fmt.Printf("example: struct: empty\n%s\n", hex.Dump(mustMarshal(example7)))

	example8 := &pb.Test3{V: &pb.BoolValue{}, Dummy: "dummy"}
	fmt.Printf("example: struct: value empty\n%s\n", hex.Dump(mustMarshal(example8)))

	example9 := &pb.Test3{V: &pb.BoolValue{V: false}, Dummy: "dummy"}
	fmt.Printf("example: struct: value false\n%s\n", hex.Dump(mustMarshal(example9)))

	example10 := &pb.Test3{V: &pb.BoolValue{V: true}, Dummy: "dummy"}
	fmt.Printf("example: struct: value true\n%s\n", hex.Dump(mustMarshal(example10)))
}

func mustMarshal(e protoreflect.ProtoMessage) []byte {
	out, err := proto.Marshal(e)
	if err != nil {
		panic(err)
	}
	return out
}
