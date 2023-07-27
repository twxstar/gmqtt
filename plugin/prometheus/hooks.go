package prometheus

import (
	"github.com/twxstar/gmqtt/server"
)

func (p *Prometheus) HookWrapper() server.HookWrapper {
	return server.HookWrapper{}
}
