package diff

import (
	"bytes"
	"testing"
	"strings"
	// . "github.com/DanielPWang/wdiff"
)

// type Unit rune
type Unit struct {
	name string
	DefaultValueName string
	attrs map(string)float32
	DefaultShowName  string
	showAttr map(string) string
}

func NewUnit(str string) []Unit {
	bytes := &[]byte(str)
	
}
func (self *Unit) GetValue(name string) float32 {
	if strings.Compare(name, "")==0 {
		name = self.DefaultValueName
	}
	value, ok := self.attrs[name]
	if !ok {
		Log("[WARN] ", self.name, " dont have Attr ", name)
		value = 0.00
	}
	return value
}
func (self *Unit) GetShow(name string) string {
	return self.name;
}

func TestDiff(t *testing.T) {
	t.Log("ok")
	// logger.Print("test")
	Log("test")
}
