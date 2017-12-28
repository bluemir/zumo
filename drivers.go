package main

import (
	// store
	_ "github.com/bluemir/zumo/backend/store/bunt"
	_ "github.com/bluemir/zumo/backend/store/etcd"
	// bot
	_ "github.com/bluemir/zumo/bots/glados"
	_ "github.com/bluemir/zumo/bots/http-checker"
	_ "github.com/bluemir/zumo/bots/todo"
	// server plugins
)
