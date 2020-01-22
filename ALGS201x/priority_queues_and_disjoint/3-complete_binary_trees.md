# Complete Binary tree

A binary tree is "complete" if all it's levels are filled completely, except for its last/lowest level, AND if the lowest level is not itself completely filled, all the nodes in it must be in the left-most positions.

## Advantages: minimize levels

As we noted earlier, operations on binary trees take longer the more levels those trees have, so if we can guarantee a low level count we'll end up with a faster binary tree. 

A binary "complete" tree always has the minimal number of levels possible, since each level must be completely filled before you can start adding another one. In addition, a complete binary tree will have exactly log{2}(n) (rounded up) levels, where n in the total number of nodes. It can be proved:

1. We can see by looking at the damn thing that 

Where n is the maximum number of nodes in the final branch of a full, complete tree
Where l is the number of levels on that tree

that n = 2^l - 1.

You can check that out, it's just evident.

2. n + 1 = 2^l
3. log{2}(n+1) = l
4. finally, therefore l is always in the order of log(n)

Because all operations on a complete full tree are on the order of the number of levels of the tree, all tree operations are also, at worst, O(log(n)).

This is in the case when the tree is full as well as complete. When the tree is not even full the above still holds, because 

## Advantages: can be stored as an array

Because of the left-most constraint, it turns out you can store a complete binary tree in an array. Let the index of each node reflect its insertion order, with the root having index 1, it's leftmost child the index 2 and so on. It turns out that if you do this you can traverse the tree without pointers just by using the index of the current node like so:

where `i` is the node's index:

1. i/2 rounded down will be the index of its parent
2. 2i will be the index of its left child, and
3. 21+1 will be the index of it's right child (duh)

You can also chain these to quickly access portions of the tree, e.g. i/4 rounded down will get you a node's grandparent (or the root if it has no grandparent) and index 1 will always return you to the root.

Of course to actually store the data in an array you will probably have to use index 0, so you'll have to convert between the array indices and the (easy-to-traverse) 1-starting tree indices, but otherwise great.

## Preservation

In order to reap these advantages, we need to preserve the "completeness" of the tree in question, i.e. whenever we add or remove a node from the tree we have to ensure that the tree does not become incomplete.

This is very easy though. The insertion methods we discussed earlier would add a new node to the tree to any random position that did not violate the binary tree constraint. We can easily modify this to always perform insertions at the leftmost open position in the highest non-full level. Sifting up does not effect the shape of the tree.

When extracting the root, recall that in the past we first swapped it's value with that of a random leaf, before removing that leaf, and then performed siftDown operations until the binary heap constraint was satisfied. Again, sifting does not effect the shape of a tree, so this will not effect completeness. The only thing we need to do is ensure that we select the most-recently-added leaf when performing the swap-and-remove (this will always be the rightmost leaf on the lowest level). 

