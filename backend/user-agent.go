package backend

func (b *backend) RegisterUserAgent(username string, ua UserAgent) error {
	b.userAgentManager.Register(username, ua)

	for _, d := range b.channels {
		if d.isMember(username) {
			d.AddListener(ua.OnMessage)
		}
	}

	//b.events.AddListener(a) // TODO remove userAgent from agent list
	return nil
}
