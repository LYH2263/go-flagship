package flagship

import "time"

type RuleOp string

const (
	OpEQ      RuleOp = "eq"
	OpNE      RuleOp = "ne"
	OpIn      RuleOp = "in"
	OpNotIn   RuleOp = "notin"
	OpExists  RuleOp = "exists"
	OpMissing RuleOp = "missing"
)

type Rule struct {
	Attr   string
	Op     RuleOp
	Values []string
}

type FlagDef struct {
	Key         string
	Enabled     bool
	Percent     int
	Rules       []Rule
	Description string
	UpdatedAt   time.Time
}

type FlagView struct {
	Key         string
	Enabled     bool
	Percent     int
	Rules       []Rule
	Description string
	UpdatedAt   time.Time
}

type Decision struct {
	Key     string
	On      bool
	Reason  string
	Bucket  int
	Matched bool
}

type Stats struct {
	Flags     int
	Upserts   uint64
	Evaluates uint64
	OnCount   uint64
	OffCount  uint64
	Closed    bool
}

type EvalAttrs map[string]string
