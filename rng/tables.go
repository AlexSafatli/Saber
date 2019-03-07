package rng

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const ()

var ()

type RandomTable struct {
	Values map[int]string
	min    int
	max    int
}

func NewRandomTable(path string) (*RandomTable, error) {
	r := &RandomTable{}
	if err := r.Parse(path); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RandomTable) Roll() string {
	if r.max == r.min && r.min == 0 {
		return ""
	}
	roll := rand.Intn(r.max-r.min) + r.min
	var closest int
	for k := range r.Values {
		if roll > k {
			closest = k
		} else {
			break
		}
	}
	return r.Values[closest]
}

func (r *RandomTable) Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()
	r.Values = make(map[int]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Parse table data.
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) < 3 {
			continue
		}
		rollMin, err := strconv.Atoi(fields[0])
		if err != nil || rollMin < 0 {
			continue
		}
		rollMax, err := strconv.Atoi(fields[1])
		if err != nil || rollMax < 0 {
			continue
		}
		valueIndex := strings.Index(s, fields[2])
		r.min = min(r.min, rollMin)
		r.max = max(r.max, rollMax)
		r.Values[rollMin] = s[valueIndex:]
	}
	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}

func InitRandomTables() {
	//var err error
}
