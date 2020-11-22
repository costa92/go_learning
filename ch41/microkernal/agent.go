package microkernal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type State int

const (
	Running State = iota
	Waiting
)

var WrongStateError = errors.New("can not take the operation in the current state")


type CollectorError struct {
	CollectorError []error
}

func (ce CollectorError) Error() string{
	var strs []string
	for _,err := range ce.CollectorError {
		strs = append(strs,err.Error())
	}

	return strings.Join(strs, ";")
}


type Event struct {
	Source string
	Content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destory() error

}


type Agent struct {
	collectors  map[string]Collector
	evtBuf      chan Event
	cancel      context.CancelFunc
	ctx         context.Context
	state       State
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}

func (agt *Agent) EventProcessGroutine(){
	var evtSeg [10]Event

	for {
		for i:=0;i<10;i++{
			select {
			case evtSeg[i] = <- agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

func (agt *Agent) RegisterCollector(name string , collector Collector) error{
	if agt.state != Waiting {
		return WrongStateError
	}

	agt.collectors[name] =collector

	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error  {
	var err error
	var errs CollectorError
	var mutex sync.Mutex

	for name,collector := range agt.collectors {
		go func(name string,collector Collector,ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()

			if err!= nil {
				errs.CollectorError = append(errs.CollectorError,errors.New(name+":"+err.Error()))
			}
		}(name,collector,agt.ctx)
	}
	return errs
}


func (agt *Agent) destoryCollection() error {
	var err error
	var errs CollectorError

	for name,collector := range agt.collectors {
		if err = collector.Destory();err != nil {
			errs.CollectorError = append(errs.CollectorError,errors.New(name + ":"+ err.Error()))
		}
	}

	return errs

}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return  WrongStateError
	}
	agt.state = Running
	agt.ctx,agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcessGroutine()
	return agt.startCollectors()
}

func (agt Agent) Stop() error {
	if agt.state != Running {
		return  WrongStateError
	}

	agt.state = Waiting
	agt.cancel()
	return agt.startCollectors()
}

func (agt *Agent) Destory() error {
	if agt.state != Waiting {
		return WrongStateError
	}

	return agt.destoryCollection()
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent {
		collectors: map[string]Collector{},
		evtBuf:		make(chan Event, sizeEvtBuf),
		state:		Waiting,
	}
	return &agt
}