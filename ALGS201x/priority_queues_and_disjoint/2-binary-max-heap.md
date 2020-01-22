# Binary Max Heap

A Binary tree is a tree with the following constraint:

1. A parent cannot have more than 2 children

A binary max heap is a binary tree with the following additional constraint:

2. Each child must have a value equal to or less than its parent.

## Basic operations on binary max heaps

There are two very useful "reordering" operations you can perform on binary max heaps: `siftUp` and `siftDown`. Using these two operations you can convert any binary tree into a binary max heap. Once we know how to use these, the general strategy for performing high level operations on a binary max heap will be:

1. ignore the max heap constraints and just add or remove whichever nodes you like wherever you like, so long as the tree remains a binary tree.
2. re-sift the tree until it is a binary max heap again.

Because all the operations in step 1. are O(1) and all the operations in 2. are O(log(n)), the most expensive any operation ends up being is O(log(n)).

### SiftUp

This swaps the value of a child with its parent. You use this operation on a binary tree node whose value is higher than that of its parent (violating the max heap rule). 

### siftDown

This swaps a parent with the higher of its two children. This can also be used to fix a violation. Note that you must swap with the higher priority child, since if you swapped with the lower priority child, you are introducing a new violation as the lower priority child is now the parent of it's higher priority sibling.

## Higher level operations on binary max heaps

### GetMax

GetMax without extracting the node itself is very easy, since the maximum value is guaranteed to be at the tree's root.

### Insert

To insert a node `n` anywhere, simply ignore it's value and the max heap constraint and attach it to any node where doing so does not violate the binary tree constraint. If the insertion does not violate the max heap constraint, then we're done. Otherwise, sift the newly inserted value upwards and check for violations. If there are none, we're done, otherwise sift upwards once again. Importantly, by switching a higher priority child with its lower priority parent you can never introduce MORE violations of the binary heap constraint, if that child is the only element out of whack. The child-now-parent had a priority as least as high as the old parent, so it will be higher priority than the rest of the parent's children. Similarly, the parent-now-child must be higher priority than the old child's children, because a node is always higher priority than all of its descendants in a well formed binary max heap (which as we said earlier, is the case other than that one child).

The worst-case cost of this is one siftUp operation for every level in the tree (i.e. in the case that the new element is the max of the entire tree), so we're looking at O(log(n)) (since the number of levels is log(n) where n is the number of nodes in the tree).

### Dequeue

Here you firstly swap the value of the root node (which is the max) with any leaf node, and then delete the at leaf node. You have now extracted the max value from the tree, the only issue is that you are probably violating the max heap constraint. To fix this simply siftDown the root node (used to be a leaf node), and keep doing so until the max heap constraint is satisfied. Again, the greatest cost of this is one sift per tree level, sl O(n). And again, at worse a siftDown operation where the parent is lower than at least one of it's children can do is retain the single max heap violation. The new parent will be higher than all its former siblings, because siftDown always chooses the highest priority sibling, so at worst the new child will be lower than both of it's new children, something fixed by single siftDown.