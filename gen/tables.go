package gen

import (
	"bufio"
	"bytes"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const (
	TablesPath                   = "tables"
	TableBuildingsFileName       = "building.txt"
	TableCharacteristicsFileName = "characteristic.txt"
	TableProfessionsFileName     = "profession.txt"
)

var (
	Tables = []*RandomTable{
		&TableBuildings,
		&TableCharacteristics,
		&TableProfessions,
	}
	TableBuildings       RandomTable
	TableCharacteristics RandomTable
	TableProfessions     RandomTable
)

type RandomTable struct {
	Name   string
	Values map[int]string
	min    int
	max    int
}

func NewRandomTable(path string) (RandomTable, error) {
	r := RandomTable{}
	if err := r.Parse(path); err != nil {
		return r, err
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
	defer closeTableFile(f)
	base := filepath.Base(path)
	r.Name = base[0 : len(base)-len(filepath.Ext(path))]
	r.Values = make(map[int]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Parse table data.
		s := scanner.Text()
		if strings.HasPrefix(s, "//") {
			continue // comment
		}
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
		r.Values[rollMin] = formatRandomTableString(s[valueIndex:])
	}
	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}

func rollAllTables() map[string]string {
	m := make(map[string]string)
	for _, table := range Tables {
		if table != nil {
			m[table.Name] = table.Roll()
		}
	}
	return m
}

func formatRandomTableString(s string) string {
	buf := bytes.Buffer{}
	t := template.Must(template.New("").Parse(s))
	if err := t.Execute(&buf, rollAllTables()); err != nil {
		return s
	}
	return buf.String()
}

func randomTablePath(fileName string) string {
	return filepath.Join(TablesPath, fileName)
}

func closeTableFile(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func initRandomTable(fileName string) RandomTable {
	table, err := NewRandomTable(randomTablePath(fileName))
	if err != nil {
		panic(err)
	}
	return table
}

func InitRandomTables() {
	TableBuildings = initRandomTable(TableBuildingsFileName)
	TableProfessions = initRandomTable(TableProfessionsFileName)
}
