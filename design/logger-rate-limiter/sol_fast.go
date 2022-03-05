package loggerratelimiter

type LoggerFast struct {
	m map[string]int
}

func ConstructorFast() LoggerFast {
	return LoggerFast{
		m: make(map[string]int),
	}
}

func (this *LoggerFast) ShouldPrintMessage(timestamp int, message string) bool {
	t, ok := this.m[message]
	if !ok || timestamp-t >= 10 {
		this.m[message] = timestamp
		return true
	}
	return false
}
