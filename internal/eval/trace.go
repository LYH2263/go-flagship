package eval

// Trace 返回可读评估轨迹。
func Trace(r Result) string {
	return r.Key + ":" + ReasonCode(r)
}
