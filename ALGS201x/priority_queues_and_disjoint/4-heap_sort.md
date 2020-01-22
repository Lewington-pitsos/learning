# Heap Sort

Binary heaps are good for sorting. A cheap way to sort an array for instance, is to implement a priority queue using a binary heap, insert all the array elements into it, and then use ExtractMax (i.e. remove and return the root) on the binary heap to return all the enqueued elements in sorted order.

## Sorting an array using heap-sort in place

Let's imagine you're concerned about memory, or just a huge nerd and you want to:

1. sort an array A using heap-sort
2. not allocate ANY additional memory, you want to just use the array.

In this case we can't satisfy 2. if we implement a separate priority queue and feed the results back in to our original array. Instead we need to turn the array into a binary heap (using the procedure for doing this we discussed earlier), and then sort the binary-heap-array IN PLACE, whilst still using it to do the sorting. 

As far as I'm concerned, this is way to complicated to be worthwhile, but hey.

### Converting an unsorted array into a binary heap

Recall that you can use an array to represent a binary heap. You just need to ensure that the correct nodes are matched up to the correct array indices, so that the binary heap index traversal rules apply.

What you actually do is:

1. Set the `size` of the desired binary tree. Perhaps your array has 15 segments allocated, but you only care about the first 12 elements. You need your new binary tree to keep track of this fact, you should set its `size` to 12. If you don't keep track of the size you have no way of knowing whether elements you are accessing using the binary tree traversal rules are real array elements or our-of-bounds cruft.
2. Great, now we have a binary tree stored as an array, but it isn't a max heap yet (unless you are very lucky) since the values in array (and hence in the tree nodes) are not in any kind of order. The array will not be useful for heap sorting until each node's children have values lower than it.
3. So the next step is to order the tree. The quickest way of ordering a randomly ordered tree is by calling `siftDown` on every sub-tree that has a chance of not being correctly ordered. Note that leaf nodes do not fall into this category. While each is it's own subtree, it is a subtree of a single node, so it has no "order". So, we don't need to call `siftDown` on every node. Instead we just need to call it on the first `size`/2 (rounded down) indices of the array (these are indices that correspond to non-leaf nodes). We do this in DESCENDING index order, because the sift-action performed on a given sub-tree (containing a single parent and it's children) only really makes sense if the trees branching off all the children are already well-ordered. If this is not the case, subsequent re-ordering of the descendent trees might de-order the original sub-tree (as high-priority nodes lower in the tree are sifted up). Recall that `siftDown` is a recursive function that will continue to sift a low priority node down until that node is no longer breaking the binary heap rules, this may entail several node "swaps", possibly more the higher the closer the node you begin sifting is to the root.
4. Having performed `size`/2 `siftDown` operations in the correct order, we will have ourselves a binary max heap, implemented as an array with a `size`.

Ok, great. The cost of that whole procedure, by the way, is O(n/2 * log(n)). This is because we call `siftDown` n/2 times, and each `siftDown` operation costs, at worst log(n). What we don't have yet though, is a sorted array. A is now being purposed as a binary heap, so it's nodes are in a wacky order. Let's actually do the sort.

### Converting a binary heap stored in an array into a sorted array

We have to ensure that as we start sorting the array we don't violate the constraints of the binary heap, or else future attempts to get the max value from the heap will fail. This is actually quite doable though:

Recall how the `takeMax` method on a binary heap works: 

1. You swap the value of the last/youngest/highest-index element in the tree with the value of the root element.
2. You expunge that element from the tree entirely
3. you `siftDown` the root of the tree (which now has some random, probably low value)

So, simply calling `takeMax` on the binary heap array will:

1. Send the highest value in the binary tree to the last (real/non-cruft) index of the array
2. Shrink the binary tree so it no longer contains that index
3. re-order the shrunken tree so it conforms to binary heap rules

So: all we need to do is call `takeMax` n times on the binary heap and it will automatically sort the array in (ascending) order in place. Weirdly painless.

### Heap sort and quick sort

Heap sort turns out to be quite efficient at order O(n * log(n)). It can also be performed in place, so it is memory efficient too. Because of this, heap sort is a viable alternative to quicksort, which is the currently most-used algorithm (in practice quicksort tends to be a bit faster most of the time because the arrays we actually sort are often kind of small).

In fact, for real nerds the optimal sort is a hybrid of quicksort and heapsort where you start using quicksort, and if quicksort's recursion runs too deep you cancel it and switch over to heapsort. 

Point is: heapsort is actually legit useful. 