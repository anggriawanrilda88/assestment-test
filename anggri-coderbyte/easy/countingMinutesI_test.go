package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingMinutesI(t *testing.T) {
	assert.Equal(t, CountingMinutesI("9:00am-10:00am"), 60)
	assert.Equal(t, CountingMinutesI("1:00pm-11:00am"), 1320)
	assert.Equal(t, CountingMinutesI("9:00am-09:00am"), 0)
	assert.Equal(t, CountingMinutesI("1:00pm-1:00pm"), 0)
}
