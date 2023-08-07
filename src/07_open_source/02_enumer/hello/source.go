/*
@Time: 2023/8/8 0:14
@Author: wxw
@File: source
*/
package hello

//go:generate go run github.com/dmarkham/enumer -type=YOURTYPE
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
