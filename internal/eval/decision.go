package eval

func ReasonCode(r Result) string {
	switch r.Reason {
	case "disabled", "rules_miss", "rollout_on", "rollout_off":
		return r.Reason
	default:
		if r.On {
			return "on"
		}
		return "off"
	}
}
