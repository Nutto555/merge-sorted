// Command merge prints the ascending merge of three collections given as flags.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	merge "github.com/Nutto555/merge-sorted"
)

func main() {
	c1 := flag.String("c1", "9,7,3", "collection_1: comma-separated integers, sorted max to min")
	c2 := flag.String("c2", "1,4,8,10", "collection_2: comma-separated integers, sorted min to max")
	c3 := flag.String("c3", "2,5", "collection_3: comma-separated integers, sorted min to max")
	flag.Parse()

	collections := make([][]int, 3)
	for i, raw := range []string{*c1, *c2, *c3} {
		values, err := parse(raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "collection_%d: %v\n", i+1, err)
			os.Exit(2)
		}
		collections[i] = values
	}

	fmt.Printf("collection_1 (max to min): %v\n", collections[0])
	fmt.Printf("collection_2 (min to max): %v\n", collections[1])
	fmt.Printf("collection_3 (min to max): %v\n", collections[2])
	fmt.Printf("merged (ascending):        %v\n", merge.Merge(collections[0], collections[1], collections[2]))
}

// parse turns "3,1,2" into []int{3, 1, 2}. An empty string is an empty collection.
func parse(raw string) ([]int, error) {
	values := []int{}
	if strings.TrimSpace(raw) == "" {
		return values, nil
	}
	for _, field := range strings.Split(raw, ",") {
		v, err := strconv.Atoi(strings.TrimSpace(field))
		if err != nil {
			return nil, fmt.Errorf("%q is not an integer", field)
		}
		values = append(values, v)
	}
	return values, nil
}
