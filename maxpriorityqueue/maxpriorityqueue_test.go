package maxpriorityqueue

import (
	"testing"
)

// This test shows a very simple, illustrative use for the MaxPriorityQueue.
func TestDemonstratesTypicalUsage(t *testing.T) {
	var queue MaxPriorityQueue

	queue.Insert(SimpleQItem{42, "foo"})
	queue.Insert(SimpleQItem{43, "bar"})
	queue.Insert(SimpleQItem{41, "baz"})

	maxPriorityItem := queue.Pop().(SimpleQItem)
	if maxPriorityItem.Priority() != 43 {
		t.Errorf("Expected 43, got <%d>", maxPriorityItem.Priority())
	}
	maxPriorityItem = queue.Pop().(SimpleQItem)
	if maxPriorityItem.Priority() != 42 {
		t.Errorf("Expected 42, got <%d>", maxPriorityItem.Priority())
	}
}

// This test provides a sufficient number of insertions to exercise more fully
// the priority-based insertion logic. And it ensures that repeated pop
// operations can be programmed to stop properly when the queue is exhausted, and
// that the items are provided in the expected priority order.
func TestMoreItemsAndNilReturn(t *testing.T) {
	var queue MaxPriorityQueue
	prio := []int{1, 2, 3, 7, 8, 9, 4, 5, 6, 10, 11, 12}
	for _, p := range prio {
		queue.Insert(SimpleQItem{p, "dontCareString"})
	}
	var prev int = 999
	for item := queue.Pop(); item != nil; item = queue.Pop() {
		if item.Priority() > prev {
			t.Errorf("Wrong sequence because %d is larger than %d",
				item.Priority(), prev)
			break
		}
		prev = item.Priority()
	}
}

// This test ensures that insertions and pop operation behave as expected, when
// they are interleaved.
func TestInsertionsInterleavedWithPops(t *testing.T) {
	var queue MaxPriorityQueue

	queue.Insert(SimpleQItem{42, "foo"})
	queue.Insert(SimpleQItem{44, "bar"})
	queue.Insert(SimpleQItem{41, "baz"})

	queue.Pop()
	queue.Insert(SimpleQItem{43, "baz"})
	maxPriorityItem := queue.Pop().(SimpleQItem)
	if maxPriorityItem.Priority() != 43 {
		t.Errorf("Expected 43, got <%d>", maxPriorityItem.Priority())
	}
}

// This test exercises the Pop method immediately after the MaxPriortyQueue is
// instantiated with its zero-value. It is expected not to crash.
func TestResilienceToImmediatePopWhenEmpty(t *testing.T) {
	var queue MaxPriorityQueue
	queue.Pop()
}

// This test ensures that any singularities present when just a single item has
// been inserted and is then popped do not break the expected behaviour.
func TestSingularityOfJustOneInsertionPlusPop(t *testing.T) {
	var queue MaxPriorityQueue
	queue.Insert(SimpleQItem{42, "foo"})
	queue.Pop()
}

// This type exists only to provide a minimum viable implementation of the
// HasPriority interface, but with a token extra field to illustrate generality.
type SimpleQItem struct {
	priority       int
	someOtherField string
}

func (t SimpleQItem) Priority() int {
	return t.priority
}
