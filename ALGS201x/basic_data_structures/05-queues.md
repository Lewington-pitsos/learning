# Queues

The exact same as a stack, you keep adding elements to the top, but you can't read from the top, you can only take elements from the bottom. These are like real world bread lines. You can add new queuers to the queue, but you can only serve queuers at the start of the queue, those that have been in the queue the longest.

## Queue implementations

Should you use a linked list or an array?

Turns out you can do either again, but both are a bit more complex.

## Linked list queue

You could just use a regular linked list, add new elements to the start, take from the end whenever you need, but in the vanilla linked list, reading from the end is O(n). A better idea is a linked list with a tail pointer. In this case you enqueue new elements to the end (O(1) with a tail pointer) and dequeue from the start (also O(n)). You can't do this the other way around (enqueue at the start, dequeue from the end) because the dequeueing operation involves setting the pointer of the new queue end element to nil. The tail pointer only points to the literal last queue element, so reassigning the pointer of the second last element would require traversing the entire list from the start.

## Array queue

If you tried to implement a queue with a normal array there would be a lot of shifting involved, and that would be terrible. 

There is a solution though: an array with separate read and write indices. The read and write indices start pointing to the first array element. An Enqueue operation adds the element to the current write index and increments that index. This way you can store a bunch of elements in order. The Dequeue operation reads and returns the element at the read index, and then increments the read index. This way all the elements are dequeued in order. If the read index catches up to the write index, the queue is empty. If either index reaches the end of the allocated array, that index wraps around and begins from the start again.

Pretty nifty actually.