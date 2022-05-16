package kafka

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/yunyandz/tiktok-demo-backend/internal/config"
)

var (
	comsumer     sarama.Consumer
	comsumerOnce sync.Once
)

func NewComsumer(cfg *config.Config) sarama.Consumer {
	comsumerOnce.Do(func() {
		var err error
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
		comsumer, err = sarama.NewConsumer(cfg.Kafka.Brokers, config)
		if err != nil {
			panic(err)
		}
	})
	return comsumer
}
