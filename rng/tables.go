package rng

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

type RandomTable struct {
	Values []string
	min    int
	max    int
}

func (r *RandomTable) Roll() string {
	roll := rand.Intn(r.max) + r.min
	return r.Values[roll]
}

func (r *RandomTable) Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Parse table data.
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) < 2 {
			continue
		}
		//rollRange := fields[0]
		//value := fields[1]
	}
	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}
