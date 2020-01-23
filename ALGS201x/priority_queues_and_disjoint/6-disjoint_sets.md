# Disjoint sets

This data structure relies in the general idea of a set, i.e. a bunch of elements that are grouped together.

A disjoint set is basically a bag containing sets. It can merge any of the sets inside it together, and given two elements, will be able to tell you whether they belong to the same set. 

Importantly, all elements are treated as unique, so multiple of the same element per set is not allowed. In addition no element can be a member of more than 1 set. In this way you are kind of treating the contents of a disjoint set as if they were a number of unique physical objects that you are clumping into groups.

## API

A basic disjoint set as 3 methods:

1. CreateSet(x)

This creates a singleton set inside the disjoint set that itself only contains x.

2. Union(x, y)

Assumes that x and y are both ALREADY members of a set in the disjoint set. These two sets are merged into a single set, if they were not already a single set.

3. Find(x) -> ID

ID is the id of the set containing X. ID must uniquely identify that set to the exclusion of all other sets in the disjoint set.

Again, pretend we're working with a bunch of unique items, like a packet of numbered ping pong balls. `CreateSet(pingpong1)` is like setting a the pingpong ball with 1 on it in its own area, separate from the rest. Calling `CreateSet(pingpong1)` doesn't really make much sense in our physical analogy since there is only one pingpong ball with 1 on it (what are you supposed to do, move it to a different area?) so it makes sense to treat this as a noop. `CreateSet(pingpong2)` is like setting the 2 ball down elsewhere in the table, and `Union(pingpong1, pingpong2)` is like moving both balls to sit together. Calling  `CreateSet()` with either ball or `Union(pingpong1, pingpong2)` both don't make sense physically, so these should probably be implemented as noops. etc etc.

## Uses: Maze

In general disjoint sets are decent at modelling "connectedness". You can make a union out of elements that relate to one another in a way that is transitive and reflective.

For example, "reachable" tiles in a maze. Some tiles are reachable from other tiles, others aren't. You can model the reachability relations of a maze using a disjoint set where you first make a singleton set fo each tile, then unionize all the sets of tiles that are reachable for one another. You end up with a bunch of large sets within the disjoint sets, each representing an area in the maze, all of whose tiles are reachable to one another, but out of reach to all other tiles.

## Naive Implementations

Obviously we're going to end up using a tree eventually, but let's try with arrays and linked lists first:

### Array Implementation

Ok, so if the elements themselves are integers we can do something kind of interesting and treat elements as indices on an array which holds lists of set ids. So consider:

   1   2   3   4   5   6   7   8   9
 -------------------------------------
 | 1 | 1 |   | 4 | 1 | 4 | 7 |   | 4 |
 -------------------------------------

Here the elements 1, 2, and 5 are in the same set, as are 4, 6, and 9, and the 7 is in its own singleton set. The sets are represented by the having all the indices that correspond to elements in the same set have same value in the array (i.e. the array values are the set IDs). The ids in this example are just the value of the lowest element (i.e. index) in the set, but they could be whatever really. The elements 3 and 8 are not part of this disjoint set, hence they do not have IDS.

This setup makes `CreateSet` and `Find` operations very easy. 

2. `Find(x)`: you simply return the array value at index x
1. `CreateSet(x)`: First call `Find` to work out whether x is already a member of some set. If so you do nothing, if so you simply assign the value of index x in the array to x. 

Both of these are O(1) operations.

`Union(x, y)` is harder though. Here you need to:

1. find the ids of both x and y --> O(1)
2. do nothing if they are the same id --> O(1)
3. Otherwise, iterate though THE ENTIRE array and set all values that equal the higher of x or y to the lower. --> O(n)

In particular, there is no way to work out which elements are in a given set without iterating over the entire array. This is very gay.

### Linked List implementation

As we already know, linked lists are very easy to mutate. In addition, we can see how a single linked list could easily represent a set with a unique id, consider:


1 -> 2 -> 5 -> 
4 -> 6 -> 9 -> 
7 -> 

This is exactly the same disjoint set we previously represented as an array, represented as a bunch of linked lists. We could easily say that the id of each list is simply that list's final element. So here 1, 2, and 5 are all members of a set with id 5. Let's see how the operations do:

1. `Find(x)`: you have to traverse every single linked list in the worst case to even work out whether x is in the disjoint set. Plus, when you find x, you need to traverse the whole of that list to find the id of that set. Terrible stuff.

2. `CreateSet(x)`: You'd like to simply create a new, singleton linked list, but if you do this without checking if x is in the disjoint already you'll fuck everything up, so `createSet` is also O(n), because it involves a `Find`.

3. `Union(x, y)`: This requires TWO find operations so it's even worse. The ONE nice thing though is that the actual joining of sets itself is easy since you just set the pointer of one list to the head of the other list, thereby ensuring that all the elements in both lists now share the same id.
