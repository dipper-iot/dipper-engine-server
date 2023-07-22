package rule_engine

import (
	"context"
	"github.com/dipper-iot/dipper-engine/bus"
	"github.com/dipper-iot/dipper-engine/core"
	"github.com/dipper-iot/dipper-engine/data"
	"github.com/dipper-iot/dipper-engine/queue"
	"github.com/dipper-iot/dipper-engine/rules/arithmetic"
	"github.com/dipper-iot/dipper-engine/rules/conditional"
	"github.com/dipper-iot/dipper-engine/rules/fork"
	"github.com/dipper-iot/dipper-engine/rules/input_redis_queue"
	"github.com/dipper-iot/dipper-engine/rules/input_redis_queue_extend"
	logRule "github.com/dipper-iot/dipper-engine/rules/log"
	"github.com/dipper-iot/dipper-engine/rules/output_redis_queue"
	output_redis_extend_queue "github.com/dipper-iot/dipper-engine/rules/output_redis_queue_extend"
	_switch "github.com/dipper-iot/dipper-engine/rules/switch"
	"github.com/dipper-iot/dipper-engine/store"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	BusData      bus.Bus
	InputSession queue.QueueEngine[*data.Session]
	store        store.Store
	config       *core.ConfigEngine
	dipper       *core.DipperEngine
	ctx          context.Context
}

func NewEngine(
	config *core.ConfigEngine,
	store store.Store,
	inputSession queue.QueueEngine[*data.Session],
	busData bus.Bus,
	ctx context.Context,
) *Engine {
	return &Engine{
		store:        store,
		config:       config,
		InputSession: inputSession,
		BusData:      busData,
		ctx:          ctx,
	}
}

func (e *Engine) ListControl() []string {
	return e.dipper.ListControl()
}

func (e *Engine) ControlSession(ruleName string) []uint64 {
	return e.dipper.ControlSession(ruleName)

}

func (e *Engine) ControlGetRule(sessionId uint64) []string {
	return e.dipper.ControlGetRule(sessionId)
}

func (e *Engine) ControlStopSession(ruleName string, session uint64) {
	e.dipper.ControlStopSession(ruleName, session)
	return
}

func (e *Engine) StartSession(session uint64) error {
	return e.dipper.StartSession(context.TODO(), session)
}

func (e *Engine) Start() error {

	factoryQueue := core.FactoryQueueDefault[*data.InputEngine]()
	factoryQueueName := core.FactoryQueueNameDefault[*data.OutputEngine]()

	e.dipper = core.NewDipperEngine(
		e.config,
		factoryQueue,
		factoryQueueName,
		e.store,
		e.BusData,
	)

	e.dipper.AddRule(
		input_redis_queue.NewInputRedisQueueRule(),
		input_redis_queue_extend.NewInputRedisQueueExtendRule(),
		output_redis_queue.NewOutputRedisQueueRule(),
		output_redis_extend_queue.NewOutputRedisQueueExtendRule(),
		logRule.NewLogRule(),
		arithmetic.NewArithmetic(),
		fork.NewForkRule(),
		conditional.NewConditionalRule(),
		_switch.NewSwitchRule(),
	)

	e.dipper.LoadPlugin()

	e.InputSession.Subscribe(e.ctx, func(deliver *queue.Deliver[*data.Session]) {
		err := e.dipper.Add(e.ctx, deliver.Data)
		if err != nil {
			logrus.Error(err)
			deliver.Reject()
			return
		}
		deliver.Ack()
	})

	return e.dipper.Start()
}

func (e *Engine) Stop() error {
	return e.dipper.Stop()
}
