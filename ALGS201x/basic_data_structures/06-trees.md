# Trees

Trees are kind of like linked lists on steroids, i.e. the requirement that each list have only one pointer is removed, they can have as many pointers as they wish.

More concretely: a tree can either be empty, or it consists of:

1. a Key (this would be some kind of content, i.e. if the tree consists of numbers, a key might be the number 99)
2. a list of links to child trees

The child trees themselves will consist of further keys and lists of child trees. The only restriction is that a tree can never have any of its ancestors as children. The list of links only flows in one direction, away from the start of the tree


## Traversing trees

Often we will be using recursion to traverse and perform actions upon trees. Recursion is common when dealing with trees.


### Depth first and Breadth first

This is a term from graph theory in general. In the context of trees is means traversing one child tree fully before beginning to traverse any others. The alternative would be breadth-first, in which each child tree is traversed one level, before any additional levels are explored.

Depth-first search is very easy to implement: just do a recursive function that runs itself on all branches before dealing with the tree itself. 

Breadth first is a little harder, but an easy idea is have an algorithm that handles the tree itself, then adds all its children to the queue and then re-runs itself on every element it dequeues from the queue.

### Traversal order

When traversing trees something to keep in mind is: do you traverse the tree's children first or the tree (i.e. it's key) itself? This will have a big impact on the order of operations. 