package reioc

import (
	"context"
	"github.com/remicro/trifle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	ctr := New()
	assert.NotNil(t, ctr)
}

func TestContainer_Provide(t *testing.T) {
	t.Run("expect provide in container", func(t *testing.T) {
		assert.NotPanics(t, func() {
			New().Provide(context.Background)
		})
	})
	t.Run("call providers", func(t *testing.T) {
		exp := trifle.Int()
		result := 0
		ctr := New().Provide(func() int {
			return exp
		})
		require.NoError(t, ctr.(*container).digContainer.Invoke(func(val int) {
			result = val
		}))
		assert.Equal(t, exp, result)
	})
}

func TestContainer_Invoke(t *testing.T) {
	err := trifle.UnexpectedError()
	ctr := New().
		Invoke(func() error {
			return err
		})
	invokers := ctr.(*container).invokers
	require.Len(t, invokers, 1)
	assert.Equal(t, err, invokers[0].(func() error)())
}

func TestContainer_Inject(t *testing.T) {
	t.Run("expect error on inject", func(t *testing.T) {
		exp := trifle.UnexpectedError()
		err := New().
			Invoke(func() error { return exp }).
			Inject()
		require.Error(t, err)
		assert.Equal(t, err, exp)
	})
	t.Run("expect provider inject to invoker", func(t *testing.T) {
		result := -1
		exp := trifle.Int()
		err := New().Provide(func() (int, error) {
			return exp, nil
		}).Invoke(func(value int) error {
			result = exp
			return nil
		}).Inject()
		require.NoError(t, err)
		assert.Equal(t, exp, result)
	})
}
