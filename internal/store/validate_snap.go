package store

func ValidateSnapshot(snap *Snapshot) bool {
	if snap == nil {
		return false
	}
	seen := make(map[string]struct{})
	for _, f := range snap.Flags {
		if f.Key == "" {
			return false
		}
		if _, ok := seen[f.Key]; ok {
			return false
		}
		seen[f.Key] = struct{}{}
		if f.Percent < 0 || f.Percent > 100 {
			return false
		}
	}
	return true
}
