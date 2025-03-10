package applicationError

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultErrorMessage(t *testing.T) {
	t.Run("각 status에 알맞은 default message를 반환해야한다.", func(t *testing.T) {
		for key, v := range defaultMessageMap {
			assert.Equal(t, v, getDefaultErrorMessage(key, ""))
		}
	})
	t.Run("status에 해당하지 않는 경우 msg를 그대로 반환한다.", func(t *testing.T) {
		assert.Equal(t, "msg", getDefaultErrorMessage(0, "msg"))
	})
}
