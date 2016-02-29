package diff

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// float：2^23 = 8388608，一共七位，这意味着最多能有7位有效数字，但绝对能保证的为6位，也即float的精度为6~7位有效数字；
// double：2^52 = 4503599627370496，一共16位，同理，double的精度为15~16位。

type ValueType float32

const (
	IGNORE = iota
	LEFT
	LEFT_TOP
	TOP
	RIGHT_TOP
	RIGHT
	RIGHT_BOTTOM
	BOTTOM
	LEFT_BOTTOM
)

var logger = log.New(os.Stdout, "[CORE]", log.LstdFlags|log.Lshortfile) //log.Logger
var Version = 1

func init() {
	logger.Print("diff.init.")
}

func Log(msg ...interface{}) {
	logger.Println(msg...)
}

// 根据相同元素个数判断
func default_calculate_ratio(matches, length int) float32 {
	return 2.0 * float32(matches) / float32(length)
}

type Valuer interface {
	GetValue(string) float32
	GetShow(string) string
}

//type Differ interface {}

var Calculate_ratio = default_calculate_ratio

func isLinejunk(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
func Default_IsClose(l, r Valuer) bool {
	return math.Abs(float64(l.GetValue(""))-float64(r.GetValue(""))) < 0.0001
}

type LCS_Math struct {
	IsClose func(l, r Valuer) bool
	A       []Valuer
	B       []Valuer
	b2j     [][]int
	direct  [][]int
}

func (sm *LCS_Math) New(a, b []Valuer, isClose func(l, r Valuer) bool) {
	// TODO: check len(a) len(b), maybe too bigger
	sm.A = a
	sm.B = b

	if isClose == nil {
		sm.IsClose = isClose
	}

	sm.b2j = make([][]int, len(b)+1) // [b][a]
	for i := range sm.b2j {
		sm.b2j[i] = make([]int, len(a)+1)
	}
	sm.direct = make([][]int, len(b)+1) // [b][a]
	for i := range sm.direct {
		sm.direct[i] = make([]int, len(a)+1)
	}
}
func (sm *LCS_Math) Calculate() {
	for i := 1; i < len(sm.A)+1; i++ {
		for j := 1; j < len(sm.B)+1; j++ {
			// TODO:
			if sm.IsClose(sm.A[i-1], sm.B[j-1]) {
				sm.b2j[i][j] = sm.b2j[i-1][j-1] + 1
				sm.direct[i][j] = LEFT_TOP
			} else if sm.b2j[i-1][j] >= sm.b2j[i][j-1] {
				sm.b2j[i][j] = sm.b2j[i-1][j]
				sm.direct[i][j] = LEFT
			} else {
				sm.b2j[i][j] = sm.b2j[i][j-1]
				sm.direct[i][j] = TOP
			}
		}
	}
}

func (sm *LCS_Math) PrintLCS(a []Valuer, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	switch sm.direct[i][j] {
	case LEFT_TOP:
		{
			sm.PrintLCS(sm.A, i-1, j-1)
			fmt.Printf("%s ", sm.A[i-1].GetShow(""))
		}
	case TOP:
		sm.PrintLCS(sm.A, i, j-1)
	case LEFT:
		sm.PrintLCS(sm.A, i-1, j)
	default:
		panic("Cannt get here.")
	}
}

func ShowMatrix(matrix [][]int) {
	fmt.Println("====== Show matrix =============")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("====== End show matrix =========")
}
