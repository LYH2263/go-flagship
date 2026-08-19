package rollout

func SubjectFromAttrs(attrs map[string]string) string {
	if attrs == nil {
		return ""
	}
	for _, k := range []string{"user", "user_id", "uid", "org", "org_id", "subject"} {
		if v := attrs[k]; v != "" {
			return v
		}
	}
	return ""
}
