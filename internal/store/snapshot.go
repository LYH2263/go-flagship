package store

import "time"

type RuleSnap struct {
	Attr   string   `json:"attr"`
	Op     string   `json:"op"`
	Values []string `json:"values"`
}

type FlagSnap struct {
	Key         string     `json:"key"`
	Enabled     bool       `json:"enabled"`
	Percent     int        `json:"percent"`
	Rules       []RuleSnap `json:"rules"`
	Description string     `json:"description"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Snapshot struct {
	Flags []FlagSnap `json:"flags"`
}
