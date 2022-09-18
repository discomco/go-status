package status

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatICanBeBorn(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasStatus(me, Unknown))
	// WHEN
	ForceStatus(&me, Born)
	// THEN
	assert.True(t, HasStatus(me, Born))
	assert.True(t, NotHasStatus(me, Unknown))
}

func TestThatICanDie(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasStatus(me, Unknown))
	// WHEN
	SetStatus(&me, Born)
	assert.True(t, HasStatus(me, Born))
	// AND
	SetStatus(&me, Living)
	assert.True(t, HasStatus(me, Living))
	// AND
	UnsetStatus(&me, Born)
	assert.True(t, NotHasStatus(me, Born))
	// AND
	ForceStatus(&me, Dead)
	// THEN
	assert.True(t, NotHasStatus(me, Born))
	assert.True(t, NotHasStatus(me, Living))
	assert.True(t, HasStatus(me, Dead))
}

func TestThatWeCanSetAndUnsetMultipleFlagsAtOnce(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasStatus(me, Unknown))
	// WHEN
	SetMulti(&me, Born, Living, Dead)
	// THEN
	assert.True(t, HasMulti(me, Born, Living, Dead))
	// AND WHEN
	UnsetMulti(&me, Born, Dead)
	assert.True(t, !HasMulti(me, Born, Dead))

}
