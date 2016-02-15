package diff

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "[CORE]", log.LstdFlags|log.Lshortfile) //log.Logger
var Version = 1

func init() {
	logger.Print("Hello world.")
}

func Display() {
	fmt.Println("dir test")
}

// 根据相同元素个数判断
func default_calculate_ratio(matches, length int) float32 {
	return 2.0 * float32(matches) / float32(length)
}

type Valuer interface {
	GetValue(string) int
	// -int.Max
	Distance(Valuer) int
}

//type Differ interface {}

var Calculate_ratio = default_calculate_ratio

func isLinejunk(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

type SequenceMatcher struct {
	IsJunk          func(Valuer) bool
	A               []Valuer
	B               []Valuer
	matching_blocks []uint
	opcodes         []uint
	fullbcount      uint
	b2j             map[Valuer][]uint
	bjunk           map[Valuer]uint
}

func (sm *SequenceMatcher) New(a, b []Valuer, isJunk func(Valuer) bool) {
	sm.A = a
	sm.B = b
	if isJunk == nil {
		sm.IsJunk = func(Valuer) bool { return false }
	} else {
		sm.IsJunk = isJunk
	}
	sm.b2j = make(map[Valuer][]uint, 128)
}
func (sm *SequenceMatcher) chain_b() {
	var b = sm.B
	var b2j = sm.b2j
	for i, e := range b {
		indices, ok := b2j[e]
		if ok {
			b2j[e] = append(indices, uint(i))
		} else {
			b2j[e] = []uint{uint(i)}
		}
	}
	for k, _ := range sm.b2j {
		if sm.IsJunk(k) {
			sm.bjunk[k] = 1
		}
	}
	for k, _ := range sm.bjunk {
		delete(sm.b2j, k)
	}
}
func (sm *SequenceMatcher) TopN() {

}

func (sm *SequenceMatcher) find_longest_match(alo, ahi, blo, bhi uint) (apos, bpos, size uint) {
	apos, bpos, size = alo, blo, 0
	var j2len = make(map[Valuer]uint)
	var nothing = make([]uint, 12)
	for i := range ahi - alo {
		i = alo + i

	}
	return apos, bpos, size
}
