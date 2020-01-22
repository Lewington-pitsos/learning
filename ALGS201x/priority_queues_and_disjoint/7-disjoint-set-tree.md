# Tree Implementation of a Disjoint Set

The nice thing about representing elements as array indices and their set ids as the values corresponding to those indices was (if you recall) that the `Find` operation was very quick. The issue with the array implementation was that shuffling around sets was an O(n) operation. This setup was still vastly better than using linked lists though, so let's see if we can at least stick close to it.

## Tree Disjoint sets

Leaving implementation aside for a moment, we will notice that trees can also be used to represent sets. Elements in a set are simply nodes of a tree and the set's id is simply the value of it's root node. But we can also see that trees have the capacity to have a much shorter distance between a given element and it's id than linked lists, which could improve things a lot.

If a linked list has n elements, finding the id of one of it's elements will take n actions in the worst case *even* if you start with a reference to the element itself (since if the element is the first segment of the list you will have to traverse the whole list to find the list's id). In the worst case, a tree with n elements has exactly the same properties (since a tree with only one branch is just a linked list), but trees can be made much more efficient. For a complete, binary tree the worst case for finding the ID of an element in a tree with n elements is log(n) if you start with the reference of that element.

Trees are also just as easy as linked lists to merge into a single set. To unionize sets x and y simply make the root of the tree representing set x and point it to the root of set y (or vise versa), thereby making the root of x a child of y. To keep trees short, you'd always want to point the tree with the least levels to the root of the tree with the most levels.

Ok, so that's nice, but how do we implement it?

## Array Tree Disjoint sets

Recall that it is possible to implement a binary tree as an array. Well it turns out that it is possible to implement arbitrary trees (with unique node values) as arrays in the following way:

1. The value of each tree node will be represented as an index on the array A
2. The parent of each tree node will be represented as the value at that node's index, for instance

   1   2   3   4
 -----------------
 | 3 |   | 3 | 4 |
 -----------------

 Here we see a two trees. The first tree  contains two elements, one with value 1, and the other with value 3. The parent of 1 is 3, as represented by the value at index 1. 3 is the root node (because it's parent is 3 i.e. itself). The second tree contains only 1 element 4.

 This way we can represent multiple trees with arbitrary numbers of children per node *so long as each element in any tree is unique across all trees*.

 This system can easily be used to implement disjoint sets if we treat each tree as its own set and the id of each tree as the value (index) of its root node. Let's see how our 3 methods look under this implementation:

 - `Find(x)`:
    1. if x exists in a set it's index in the array will have a value, if not we're done
    2. If so, that value is the parent of x. If the array value is the same as the index, we are dealing with a singleton set, so the id will just be x
    3. Otherwise, we treat that value as an index to visit X's parent directly and so on until we visit an index whose value is that index, in which case we've found the root node and we return that value.

    This operation is O(l) where l is the levels between x and the root of x's tree. In the worst case this is still O(n), but later we'll show how we can ensure that the worst case doesn't occur.

- `CreateNew(x)`:
    1. First we have to use `Find` to work out if x is already present in a set
    2. if it is, we do nothing
    3. if it isn't, we create a new, singleton tree by setting the x'th index of the array to x

- `Union(x, y)`:
    1. First work out the ids of x and y. While you're at this keep track of which tree is longer.
    2. make the shorter tree a child of the longer one, simply by changing the value of the shorter tree's root to the longer tree's index

`CreateNew` and `Union` are both O(1) operations except that they contain `Find`, so in terms of efficiency, their worst case is just the same as the worst case of `Find`. Currently the worst case of `Find` is O(n), but in reality it's O(the length of the longest branch in any tree), so if we could somehow guarantee short branches, we could ensure that all 3 operations have low costs.

## Merging trees efficiently

One thing we kind of skipped over above was how to determine the height of different trees when calling `Union`. One good (but admittedly space inefficient) way to do this is to keep a second array R that keeps track of the height of all the trees in this disjoint set. The way this will work is: if t is the root root of a tree in the main array then R[t] = the height of t. Under this implementation R will have to be the same size (in terms of allocation) as the main array, since any element in the main array could be the root of a tree.

We achieve this by

1. Updating the `CreateNew` method so that it whenever a new set is created, the rank of that set (0) is added to R under the same index.
2. Updating `Union(x, y)` so that we compare the ranks of x and y before choosing which tree to subordinate to the other, and then finally increasing the rank of the superior tree by 1 IFF the subordinate tree was the same height as the superior tree.

With the last section of 2., notice that if you have 2 trees either they are equal in height, or one is greater. If the trees are equal, subordinating one to the other always results in a new tree of height p + 1, where p was the height of the original trees. If one tree is greater, subordinating the lesser tree to it can never increase the height of the greater tree.

Great, so now we have a time (but arguably not space) efficient tree-merging implementation, that ensures short trees.

## What is the efficiency exactly?

So to be clear we are not interested so much in the efficiency of the actual act of merging two trees, this is clearly O(1) the way we've implemented it, we're more interested in how efficient subsequent `Find` operations performed on those trees are. So let's try to prove something:

1. The maximum height/rank r of a tree t is log{2}(n), where n is the number of nodes in t.
2. put another way n >= 2^r

We're going to prove this by induction for r:

- The base case:
    1. r = 0
    2. where r = 0, according to our implementation, n = 1
    3. 1 = 2^0
    4. 1 = 1

Ok proved, now let's do the step case:

- Step case:
    1. Induction lets us assume that for every tree n >= 2^r
    2. The only way, under our algorithm, we can get a tree of rank R where R > 0, is when two trees are merged and both have the same level. By 1. both of the subtrees must have had at least n, i.e. 2^(R-1) nodes, which means that the resulting tree must have at least 2 * 2^(R-1) nodes, or, more plainly 2^R nodes. 

