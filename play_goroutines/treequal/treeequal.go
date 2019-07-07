/*
This program is  about using two goroutines to walk two binary trees at the
same time, and assessing if the trees emit the same sequence of node values
(despite potentially being different trees). The algorithm fails-fast.

It also demonstrates the use of closing a channel to signal a quit mandate,
because a read from it will then succeed for the first time. An idiomatic
signallying method that can be consumed by multiple goroutines.
*/

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
walkR is a recursive walker of the tree rooted at *t*. It emits each node's
values on the *elements* channel. The walk sequence is deterministic, recursing
the left child first, then emitting the parent value, then recursing the
right child.
*/
func walkR(t *tree.Tree, elements chan int, quitC chan int) {
	if t == nil {
		return // Not an error, it's just the recursion cannot carry on on this branch.
	}
	walkR(t.Left, elements, quitC)
	// In the normal run we just send the parent node value out on the
	// elements channel. But, if the (external) element comparison between the two
	// trees concludes early because the traversal of the other tree finished
	// before this one had exhausted the tree, then that would leave this
	// goroutine trying to continue, but blocked by the main goroutine no
	// longer trying to consume its output values. In that case the main goroutine
	// wlll have closed the quit channel - which make the attempted read from
	// it below succeed (receiving value zero by definition).
	select {
	case elements <- t.Value:

	case <-quitC:
		return // Will happen <D> times as the recursive call-stack unwinds.
	}
	walkR(t.Right, elements, quitC)
}

/*
walkThenCloseChannel orchestrates the recursive walk, and carries the extra
responsibility of closing the elements channel when the recursive walk
is complete - to allow the caller to detect that the process has finished.
*/
func walkThenCloseChannel(t *tree.Tree, elements chan int, quitC chan int) {
	walkR(t, elements, quitC)
	close(elements)
}

/*
areSame determines if two binary trees produce the same sequence when
traversed. It calls *walkThenCloseChannel* in two separate goroutines, one
for each tree.
*/
func areSame(t1, t2 *tree.Tree) bool {
	t1Elements := make(chan int)
	t2Elements := make(chan int)
	quitC := make(chan int)
	// Arranging to close the quit channel when this funtion is ready to
	// return, will signal a potentially still-running traversal Go routine to stop and
	// unwind its return stack.
	defer close(quitC)
	go walkThenCloseChannel(t1, t1Elements, quitC)
	go walkThenCloseChannel(t2, t2Elements, quitC)
	for {
		// Consume the traversed values from both trees until one of them has closed
		// the channel to signal  that the traversal of that tree is complete. If it
		// is signalled on both channels at the same time, it means the trees
		// are equal. Otherwise they are not.
		t1e, ok1 := <-t1Elements
		t2e, ok2 := <-t2Elements
		if ok1 == false && ok2 == false {
			// both traversals finished at the same time, and both goroutines
			// will have terminated consequently.
			return true
		}
		if ok1 == false || ok2 == false {
			return false
		}
		if t1e != t2e {
			return false
		}
	}
	fmt.Printf("XXXXXX allegedly unreachable code just got reached\n")
	return true
}

/*
This program creates two binary trees and then uses the same() function to
determine if they yield the same sequence of values when traversed.
*/
func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	same := areSame(t1, t2)
	fmt.Printf("Tree equality is: %v\n", same)
}

// todos
// when close if run out of input?
// how/why signal goroutines to stop
// what defer close?
