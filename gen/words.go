package gen

import (
	"../rng"
	"bytes"
	"math/rand"
	"strings"
)

// Adapted from https://worldbuildingworkshop.com/constructing-languages/

type Phonemes struct {
	Consonants string
	Vowels     string
	Finals     string
}

type Language struct {
	*Phonemes
	Structure    string
	MinSyllables int8
	MaxSyllables int8
}

var (
	SyllableStructures = []string{"CV", "VC"}
	ConsonantPresets   = []string{"ptkbdgmnlrsqxʧ", "pbgdzklmnrstʤ",
		"pbgdzklmnrstv"}
	VowelPresets = []string{"aeiou", "aeio", "aeou"}
	FinalPresets = []string{"mnsk", "mnŋsk", "mnʃsk"}
	Orthography  = map[string]string{"ʃ": "sh", "ʧ": "ch",
		"ʤ": "j", "ŋ": "ng", "x": "kh", "ɣ": "gh"}
	Alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func (p *Phonemes) Type(r rune) string {
	switch r {
	case 'C':
		return p.Consonants
	case 'V':
		return p.Vowels
	case 'F':
		return p.Finals
	default:
		return p.Consonants
	}
}

func (l *Language) spell(s string) string {
	var str string
	for _, r := range s {
		s := string(r)
		if _, ok := Orthography[s]; ok {
			str += Orthography[s]
		} else {
			str += s
		}
	}
	return str
}

func (l *Language) Syllable() string {
	var syl string
	for _, r := range l.Structure {
		pType := l.Phonemes.Type(r)
		syl += string(rng.ChooseRune(pType))
	}
	return l.spell(syl)
}

func (l *Language) Word() string {
	var buf bytes.Buffer
	var a, b int
	diff := int(l.MaxSyllables - l.MinSyllables)
	min := int(l.MinSyllables)
	if diff > 0 {
		a = rand.Intn(diff + 1)
	}
	if min > 0 {
		b = min
	}
	numSyllables := a + b
	for i := 0; i < numSyllables; i++ {
		buf.WriteString(l.Syllable())
	}
	return buf.String()
}

func (l *Language) Name() string {
	return capitalize(l.Word())
}

func GenerateLanguage() *Language {
	minSyl := int8(rand.Intn(2) + 2)
	p := &Phonemes{
		Consonants: rng.Choose(ConsonantPresets),
		Vowels:     rng.Choose(VowelPresets),
		Finals:     rng.Choose(FinalPresets),
	}
	return &Language{
		Phonemes:     p,
		Structure:    rng.Choose(SyllableStructures),
		MinSyllables: minSyl,
		MaxSyllables: int8(rand.Intn(2)+1) + minSyl,
	}
}

func capitalize(s string) string {
	if len(s) > 0 {
		o := strings.ToUpper(string(s[0]))
		if len(s) > 1 {
			return o + s[1:]
		}
		return o
	}
	return s
}
