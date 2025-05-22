package models

type Op int

const (
	Lt  Op = iota // 0
	Lte           // 1
	Gt            // 2
	Gte           // 3
	E             // 4
)
