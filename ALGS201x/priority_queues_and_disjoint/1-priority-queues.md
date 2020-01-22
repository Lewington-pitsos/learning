# Priority Queues

A priority queue is a little like the queue datastructure that we talked about earlier. You add elements to it, and retrieve elements from it. The difference is that the insertion order does not matter to a priority queue. Instead a priority queue can use any algorithm for determining which element is "next in line". In a sense a normal queue might be thought of as a special type of priority queue: one that always prioritizes insertion-order.

You would expect a priority queue to have methods like the following at the very least:

    Enqueue(element)
    DequeueMax() element

A super simple example might be a priority queue that stores integers, and prioritizes magnitude, so always dequeues the highest inserted integer until all have been dequeued. 

## Naive Implementations

But how the heck do you implement one of these? 

### Unsorted Array

That's quite easy actually, you can do it with an array, for instance. Just insert new elements to the end of the array and, when it comes to dequeueing, just scan the list for the highest priority element and return it.

That's a bit annoying though, because, while insertion is O(1), dequeueing is O(n), since you have to scan the ENTIRE array to work out which is the highest priority element. Depending on the implementation, you might also be shifting the elements around to fill the blank space.

### Sorted Array

If we sort the array such that the highest priority element is always at the end, dequeueing becomes a simple `pop` operation, which is O(1). But insertion becomes a pain in the ass. 

Firstly you have to work out *where* to insert a new element. This will require scanning the array for the correct place to insert the new element so as not to fuck up the sorted order. This task in itself is not *terrible*, since because the array is already sorted, you can use binary search to locate the correct insertion location, which has a cost of O(log(n)), which is worse than O(1), but usually acceptable. The issue is that shifting is required in order to carry the insertion out, so instantly we're back to log(n).

### Unsorted Linked List

Ok, arrays are garbage, what about an unsorted linked list? Again insertion is easy: just add to the end and enjoy the O(1). Just like the unsorted array though, you have to scan the entire linked list every dequeue, giving us an O(n) dequeue, which is garbage.

### Sorted Linked List

If we order the linked list such that the highest priority element is always at the end (and have a tail pointer) we can dequeue at O(1) time. Insertion takes O(n) though, because, just like before, we will need to traverse the linked list to find the correct insertion point, and you can't use binary search on linked lists like you can with arrays, so this is an O(n) operation. Insertion is easy though, for all the good it does us.


## Good implementation

Basically: use a binary tree.

### Binary tree insertion

This is a O(log(n)) operation, since the cost increases the larger your tree, but that increase itself decreases sharply, so we're actually quite close to constant time. 