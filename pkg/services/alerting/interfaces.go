package alerting

import "time"

type AlertHandler interface {
	Execute(context *AlertResultContext)
}

type Scheduler interface {
	Tick(time time.Time, execQueue chan *AlertJob)
	Update(rules []*AlertRule)
}

type Notifier interface {
	Notify(alertResult *AlertResultContext)
	GetType() string
}

type AlertCondition interface {
	Eval(result *AlertResultContext)
}
