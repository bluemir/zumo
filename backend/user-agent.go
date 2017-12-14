package backend

func (b *backend) RegisterUserAgent(username string, ua UserAgent) error {
	b.userAgentManager.Register(username, ua)
	return nil
}
