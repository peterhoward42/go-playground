/*
Package maxpriorityqueue and the primary type therein: MaxPriorityQueue,
implements the well known container/storage idiom, for items having a known
(integer) priority value. Items are popped in maximum priority order, regardless
of the order in which they are added.  It works for any type that satisfies the
HasPriority interface, and thus stores only interface-values (which are small)
for the items inserted.  The zero-value of the type is immediately usable.

IMPLEMENTATION: The implementation is optimised for large data sets by using a
heap under the hood, and thus performs insertion and pop operations in O(log n)
time.  The heap implementation is the well known array-backed storage mechanism
in which index doubling/halving can be used to navigate parent-child
relationships, and maintenance of the heap property does not require fresh memory
allocation. The underlying array storage is wrapped in a slice, that is
initialised to zero size.  Go's built-in append method is used to add things to
the slice, and is entrusted to re-allocate and copy the backing array fairly
optimally whenever there is insufficient capacity to add new items. This cost
could be mitigated by providing a method to allocate the initial storage size,
but would be at the expense of the simplicity gained by having the zero-value of
the queue useful as is.

Added this line to explore git.
*/
package maxpriorityqueue

// The MaxPriorityQueue type is the queue provided by this package.
type MaxPriorityQueue struct {
	items []HasPriority
}

// The Insert method injects the given itemToInsert into the queue, in accordance
// with its priority relative to the items already present - as discovered by
// calling the Priority method on the itemToInsert.
func (q *MaxPriorityQueue) Insert(itemToInsert HasPriority) {
	q.items = append(q.items, itemToInsert)
	q.bubbleUp(len(q.items) - 1)
}

// The Pop method removes the highest priority item that is present from the
// queue and returns it to the caller.
func (q *MaxPriorityQueue) Pop() HasPriority {
	if len(q.items) == 0 {
		return nil
	}
	toReturn := q.items[0]
	q.items[0] = q.items[len(q.items)-1]
	q.items = q.items[0 : len(q.items)-1]
	q.bubbleDown(0)
	return toReturn
}

// HasPriority is an interface that declares "I have an integer priority, and you
// can find out what it is via my accessor method."

type HasPriority interface {
	Priority() int
}

//-----------------------------------------------------------------------------
// Non exported code below
//-----------------------------------------------------------------------------

func (q *MaxPriorityQueue) bubbleUp(index int) {
	if parentIdx := index / 2; parentIdx >= 0 {
		if q.items[index].Priority() > q.items[parentIdx].Priority() {
			q.swap(index, parentIdx)
			q.bubbleUp(parentIdx)
		}
	}
}

func (q *MaxPriorityQueue) bubbleDown(index int) {
	highestChild, noChildren := q.maxPrioChild(index)
	if noChildren {
		return
	}
	if q.items[highestChild].Priority() > q.items[index].Priority() {
		q.swap(index, highestChild)
		q.bubbleDown(highestChild)
	}
}

func (q *MaxPriorityQueue) maxPrioChild(parentIdx int) (childIdx int,
	noChildren bool) {
	left := parentIdx * 2
	right := left + 1
	if left > len(q.items)-1 {
		noChildren = true
		return
	}
	noChildren = false
	if (right > len(q.items)-1) ||
		(q.items[left].Priority() > q.items[right].Priority()) {
		childIdx = left
	} else {
		childIdx = right
	}
	return
}

func (q *MaxPriorityQueue) swap(i, j int) {
	tmp := q.items[i]
	q.items[i] = q.items[j]
	q.items[j] = tmp
}
