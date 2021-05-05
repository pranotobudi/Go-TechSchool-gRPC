package serializer_test

import (
	"fmt"
	"testing"

	"github.com/pranotobudi/Go-TechSchool-gRPC/sample"
	"github.com/pranotobudi/Go-TechSchool-gRPC/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	if err != nil {
		fmt.Printf("error Test File: %s", err.Error())
	}
	// require.NotError(t, err)
}
