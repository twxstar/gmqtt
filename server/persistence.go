package server

import (
	"github.com/twxstar/gmqtt/config"
	"github.com/twxstar/gmqtt/persistence/queue"
	"github.com/twxstar/gmqtt/persistence/session"
	"github.com/twxstar/gmqtt/persistence/subscription"
	"github.com/twxstar/gmqtt/persistence/unack"
)

type NewPersistence func(config config.Config) (Persistence, error)

type Persistence interface {
	Open() error
	NewQueueStore(config config.Config, defaultNotifier queue.Notifier, clientID string) (queue.Store, error)
	NewSubscriptionStore(config config.Config) (subscription.Store, error)
	NewSessionStore(config config.Config) (session.Store, error)
	NewUnackStore(config config.Config, clientID string) (unack.Store, error)
	Close() error
}
