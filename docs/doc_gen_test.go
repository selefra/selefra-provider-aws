package main

import (
	"fmt"
	"github.com/selefra/selefra-provider-aws/provider"
	"github.com/selefra/selefra-provider-sdk/doc_gen"
	"os"
	"testing"
)

func Test(t *testing.T) {

	fmt.Println("begin...")
	docOutputDirectory := os.Getenv("SELEFRA_DOC_OUTPUT_DIRECTORY")
	if docOutputDirectory == "" {
		docOutputDirectory = "./tables"
	}
	fmt.Println(docOutputDirectory)
	err := doc_gen.New(provider.GetProvider(), docOutputDirectory).Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("done...")

}
