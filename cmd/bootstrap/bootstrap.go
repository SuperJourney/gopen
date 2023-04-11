package bootstrap

import "github.com/SuperJourney/gopen/repo/model"

var GetTables = func() []interface{} {
	return []interface{}{
		model.App{},
		// model.AppAttr{},
		model.Attr{},
	}
}
