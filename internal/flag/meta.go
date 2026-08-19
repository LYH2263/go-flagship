package flag

func Summarize(r Record) string {
	on := "off"
	if r.Enabled {
		on = "on"
	}
	return r.Key + ":" + on
}
