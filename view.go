package flagship

// ViewFromDef 转视图。
func ViewFromDef(d FlagDef) FlagView {
	return FlagView{
		Key:         d.Key,
		Enabled:     d.Enabled,
		Percent:     d.Percent,
		Rules:       cloneRules(d.Rules),
		Description: d.Description,
		UpdatedAt:   d.UpdatedAt,
	}
}
