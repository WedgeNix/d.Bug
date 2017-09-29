package d

const (
	expiredSessions = "expired sessions"
)

var (
	settings = map[string]bool{
		expiredSessions: true,
	}

	debugr = bufio.NewReader(os.Stdin)
)

func debug(setting string) bool {
	b, found := settings[setting]
	if !found {
		panic(setting + " not found")
	}
	if !b {
		return false
	}
	for {
		println("WARNING: " + setting + " ON. Continue? (y/n)")
		input, _ := debugr.ReadString('\n')
		if strings.HasPrefix(input, "n") {
			panic(setting + " rejected")
		} else if strings.HasPrefix(input, "y") {
			return true
		}
	}
}
