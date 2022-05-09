package library_config

import (
	"github.com/go-bongo/go-dotaccess"
)

func Get(configs map[string]interface{}, config string, defaultValue any) any {
	get, err := dotaccess.Get(configs, config)
	if err == nil {
		return get
	}
	return defaultValue
}
