package reioc

import (
	"github.com/remicro/api/ioc"
	"go.uber.org/dig"
)

func New() ioc.Container {
	return &container{
		digContainer: dig.New(),
	}
}

type container struct {
	invokers     []interface{}
	digContainer *dig.Container
}

func (ctr *container) Provide(fn interface{}) ioc.Container {
	ctr.digContainer.Provide(fn)
	return ctr
}

func (ctr *container) Invoke(fn interface{}) ioc.Container {
	ctr.invokers = append(ctr.invokers, fn)
	return ctr
}

func (ctr *container) Inject() (err error) {
	for i := range ctr.invokers {
		err = ctr.digContainer.Invoke(ctr.invokers[i])
		if err != nil {
			break
		}
	}
	return
}
