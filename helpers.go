package flagship

// DefaultFlag 构造默认开启的 100% flag。
func DefaultFlag(key string) FlagDef {
	return FlagDef{Key: key, Enabled: true, Percent: 100}
}

// MustUpsert 测试辅助：失败则 panic。
func MustUpsert(s *Ship, def FlagDef) {
	if err := s.Upsert(def); err != nil {
		panic(err)
	}
}
