package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type idsFlag []string

func (ids idsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type Person struct {
	name string
	born time.Time
}

func (p Person) String() string {
	return fmt.Sprintf("Person: %q (%v year old)", p.name, p.born.String())
}

func (p *Person) Set(input string) error {
	p.name = input
	p.born = time.Now()
	return nil
}

func main() {
	var ids idsFlag
	var p1 Person

	flag.Var(&ids, "id", "the id to be appended to the list")
	flag.Var(&p1, "name", "the name of the person")
	flag.Parse()

	fmt.Println(ids)
	fmt.Println(p1)
}
