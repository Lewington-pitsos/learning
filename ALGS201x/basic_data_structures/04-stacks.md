# Stacks

A stack is a data-structure with 3 operations:

1. Add to top
2. Check current top
3. remove current top

Imagine a stack of pallets being loaded into a shipping container. You can interact with the top of the stack, but you can't really get at any of the other pallets except by removing all higher pallets.


## Stack Implementation

Should you implement a stack with a linked list or an array?

Turns out, to nobody's surprise, you can use either. There are a few, very insignificant tradeoffs to be aware of though:

### Array Stack

It is quite easy to implement a stack with an array. Treat the "top" of the stack as the end of the semantically useful array, and then all 3 stack operations become assignments and lookups to the end of the array (all O(1) operations). 

The only drawback is that you always need to assign a maximum length for arrays, so your stack will end up with a max depth, which is probably not what you want in a stack. Assuming we want an infinite stack, you'll have to estimate the maximum depth of the stack and choose an array max length slightly higher than what you expect the stack's actual maximum depth will end up being. Choose a small max depth and you will get errors, choose a large one and you will allocate too much memory. 

### Linked list stack

Linked lists are probably a bit better, you can express all 3 stack operations as insertions and reads to the front of the linked list, so again all O(n) operations. Linked lists are theoretically infinite, so no need to worry about max lengths.

The only small drawback is that you need to perform 2 write-operations per assignment (since you need to reassign 2 pointers), and store the pointers in memory as well as the actual list/stack content.

