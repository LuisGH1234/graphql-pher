package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

// Hello ...
func (*query) Hello() string { return "Hello, world!" }

func main() {
	sch, err := readSchema("graphql/schema")
	if err != nil {
		fmt.Println("Error:", err)
	}
	// fmt.Println(sch)

	schema := graphql.MustParseSchema(sch, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readSchema(fileName string) (string, error) {
	ext := ".graphql"
	b, err := ioutil.ReadFile(fileName + ext)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
