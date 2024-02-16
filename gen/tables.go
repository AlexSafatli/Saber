package gen

import (
	"bufio"
	"bytes"
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

func (rt *RandomTable) Roll() string {
	if rt.max == rt.min && rt.min == 0 {
		return ""
	}
	roll := r.Intn(rt.max-rt.min) + rt.min
	var closest int
	for k := range rt.Values {
		if roll > k {
			closest = k
		} else {
			break
		}
	}
	return rt.Values[closest]
}

func (rt *RandomTable) Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer closeTableFile(f)
	base := filepath.Base(path)
	rt.Name = base[0 : len(base)-len(filepath.Ext(path))]
	rt.Values = make(map[int]string)
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
		rt.min = min(rt.min, rollMin)
		rt.max = max(rt.max, rollMax)
		rt.Values[rollMin] = formatRandomTableString(s[valueIndex:])
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
