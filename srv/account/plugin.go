package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/micro/go-plugins/transport/nats"
	// _ "github.com/micro/go-plugins/broker/nats"

	// _ "github.com/micro/go-plugins/client/selector/static"
	_ "github.com/xmlking/micro-starter-kit/shared/micro/client/selector/static"
	// _ "github.com/micro/go-plugins/broker/googlepubsub"
	_ "github.com/xmlking/micro-starter-kit/shared/micro/broker/googlepubsub"
)
