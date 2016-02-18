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
	// logger.Print("Hello world.")
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
	matching_blocks []int
	opcodes         []int
	fullbcount      int
	b2j             map[Valuer][]int
	bjunk           map[Valuer]int
}

func (sm *SequenceMatcher) New(a, b []Valuer, isJunk func(Valuer) bool) {
	sm.A = a
	sm.B = b
	if isJunk == nil {
		sm.IsJunk = func(Valuer) bool { return false }
	} else {
		sm.IsJunk = isJunk
	}
	sm.b2j = make(map[Valuer][]int, 128)
}
func (sm *SequenceMatcher) chain_b() {
	var b = sm.B
	var b2j = sm.b2j
	for i, e := range b {
		indices, ok := b2j[e]
		if ok {
			b2j[e] = append(indices, int(i))
		} else {
			b2j[e] = []int{int(i)}
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
func (sm *SequenceMatcher) TopN(n int) {

}

func (sm *SequenceMatcher) find_longest_match(alo, ahi, blo, bhi int) (apos, bpos, size int) {
	apos, bpos, size = alo, blo, 0
	var j2len = make(map[Valuer]int)
	var nothing = make([]int, 12)
	for i := alo; i < ahi; i++ {
		var newj2len = make(map[Valuer]int, 10)
		for j := range sm.b2j[sm.A[i]] {
			if j < blo {
				continue
			} else if j >= bhi {
				break
			}
			var v, ok = j2len[j-1]
			if !ok {
				v = 0
			}
			newj2len[j] = v + 1
			var k = newj2len[j]
			if k > size {
				apos, bpos, size = i-k+1, j-k+i, k
			}
		}
		j2len = newj2len
	}
	for apos > alo && bpos > blo && !sm.IsJunk(sm.B[bpos-1]) && sm.A[apos-1] == sm.B[bpos-1] {
		apos, bpos, size = apos-1, bpos-1, size+1
	}
	for apos+size < ahi && bpos+size > bhi && !sm.IsJunk(sm.B[bpos+size]) && sm.A[apos+size] == sm.B[bpos+size] {
		size += 1
	}
	return apos, bpos, size
}
