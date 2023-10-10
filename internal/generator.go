package internal

import (
	"flag"
	"fmt"

	"github.com/ninomae42/ulid_gen/internal/id"
)

func GenerateULIDs() {
	count, name := parseArgs()
	ids := generate(count, name)
	for _, v := range ids {
		fmt.Println(v)
	}
}

func parseArgs() (count int, name string) {
	flag.IntVar(&count, "c", 1, "number of ulid to generate")
	flag.StringVar(&name, "n", "", "prefix to the ulid. [name]+[ulid]")
	flag.Parse()
	return
}

func generate(count int, name string) []id.ID {
	var ids []id.ID
	for i := 0; i < count; i++ {
		ids = append(ids, id.NewID(name))
	}
	return ids
}
