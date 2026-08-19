package flag

import "time"

type Rule struct {
	Attr   string
	Op     string
	Values []string
}

type Record struct {
	Key         string
	Enabled     bool
	Percent     int
	Rules       []Rule
	Description string
	UpdatedAt   time.Time
}
