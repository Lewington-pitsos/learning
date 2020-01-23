# Arrays

An array is: a contiguous section of memory of finite length, that is divided into *equally sized* portions/segments. So: on declaring an array you already know how many elements it contains. In order to declare an array you need to know the maximum amount of memory any one of those elements could possibly take to store, as that will need to be the size of each array segment.

## What are Arrays good for?

Random access and assigning.

### Random Access

The main thing Arrays are good for is "constant time access" or "random access", which means that it takes the same amount of time to access any element from an array, regardless of how large the array is. This is `O(1)` from the Efficiency section. This is possible because the *find an element in an array* operation is very simple:

So, we know the first memory address in the array, and we know the maximum size of an element in the array, and we know that each element in the array is given the memory space the maximum sized element would need. Based on this, to find the `nth` item in an array `A` with an initial memory address of `m1` and maximum element size of `em`  we just:

1. Find the memory address at `m1` + (n * `em`). This is the element we are looking for.

This holds true no matter what `n` is and no matter how large `A` is.

### Assigning.

For pretty much the same reasons, assigning array elements (i.e. replacing them with other legal elements) is also `O(1)` efficient. All you need to do is find the element you want to reassign using the above process, and then write the new element to that memory address. There is no chance of it overflowing into the next array segment because each segment is sized so that it can fit the largest possible legal element inside it.


## What are arrays bad for?

Moving elements around. Imagine we have an array of numbers with 7 elements where the last 3 are 0:

    -----------------------------
    | 2 | 7 | 3 | 9 | 0 | 0 | 0 |
    -----------------------------

Let's say we wanted to insert a 3 in the second position and displace the 7 the 3 and the 9 to the right, overwriting one 0. How many operations would this take? The answer is 4, one to move each of the displaced elements, and one final one to assign the `2`. Because memory addresses are fixed (i.e. they are not relative to anything), and because arrays rely on memory addresses to determine the contents of their segments the only way to shift an element from one array segment to another is to re-write the memory corresponding to those two segments. If you want to shift 3 elements, you need to do this 3 times.


    -----------------------------
1.  | 2 | 7 | 3 | 9 | 9 | 0 | 0 |
    -----------------------------


    -----------------------------
2.  | 2 | 7 | 3 | 3 | 9 | 0 | 0 |
    -----------------------------


    -----------------------------
3.  | 2 | 7 | 7 | 3 | 9 | 0 | 0 |
    -----------------------------


    -----------------------------
4.  | 2 | 3 | 7 | 3 | 9 | 0 | 0 |
    -----------------------------

Because of this, inserting is considered an `O(n)` operation, where `n` is the number of elements displaced. For large arrays, this can be very expensive.

## Cheap and Expensive Array Operations

As we saw, lookup and reassignment are cheap array operations, while shifting is expensive. This has consequences for higher-level array operations that involve these lower level operations. For instance, removing elements from arrays (i.e. over-writing them with nils), only involves assignment, so it is a cheap high-level operation. Removing elements and filling the gaps with other elements is an expensive operation though, since it involves shifting as well as assignment.

In general you want to use arrays to perform actions that do not involve shifting. Good examples of these are operations on the "end" of an array. Consider an array that is being used to store 3 important numbers, but can store up to 7:

  -----------------------------
  | 8 | 7 | 9 | 0 | 0 | 0 | 0 |
  -----------------------------

The unused slots are all 0's. In one sense the 9 is the "end" of this array, since it is the last element with any semantic meaning. Adding elements to the "end" (in this sense) of an array is a very cheap process (assuming we know which segment the last meaningful segment is), since it only involves a single assignment. Similarly, removing elements from the "end" of the array (i.e. assigning them to 0) is also cheap. In both cases no shifting is required.

Inserting (without overwriting existing elements) or removing elements anywhere else involves shifting though, so those operations are expensive.

## Arrays as lists

In general, humans generally treat arrays like lists, i.e. ordered sets of things. But real life lists are quite different from arrays, and what look like simple operations on a list are actually quite expensive when performed on an array. For instance, one thing lists can't have is blank spaces. Consider a shopping list for instance: 

- Ham

- Milk

- Tomatoes

Someone reading this list would interpret it as a list containing Ham, Milk and Tomatoes

If I realized, after making that list that I didn't need milk, I could cross that element off my list:

- Ham

- ~~Milk~~

- Tomatoes

But someone reading that list would read it and think it was a list containing Ham and Tomatoes. They would NOT read it and think "ah yes, this is a list containing Ham, a blank space, and tomatoes". This is not the same with arrays stored in computer memory. If I wanted to write an array to represent the initial list I might have something like:

 ---------------------------------
 | Ham     | Milk     | Tomatoes |
 ---------------------------------

If I deleted the second entry I would have 

 ---------------------------------
 | Ham     |          | Tomatoes |
 ---------------------------------

But if I wanted to make the array a faithful representation of the second list, I would need to shift the last element toward the beginning as well:

 ---------------------------------
 | Ham     | Tomatoes |          |
 ---------------------------------

Which, if the list was much longer, could end up being a very expensive operation.

The takeaway is: even though we often treat arrays like real-life lists, they aren't quite the same and sometimes operations that are cheap on real life lists end up being expensive when performed on arrays.


## 2-D arrays

Sometimes we want to represent data as 2D arrays. The main upshot of 2d arrays over arrays is that you can access elements using a multi-index. I.e. if we have an 2D array:


  -----------------------------
  | 8 | 7 | 9 | 1 | 4 | 7 | 7 |
  -----------------------------
  | 6 | 1 | 9 | 2 | 0 | 0 | 0 |
  -----------------------------

Here we can access the 2 by specifying (2, 4), meaning row 2, column 4. 

### Implementation

Behind the scenes however, 2-D arrays are often implemented as normal arrays that also store one additional piece of data: the length of the rows. For instance, the 2D array above would probably be stored as:

  ---------------------------------------------------------
  | 8 | 7 | 9 | 1 | 4 | 7 | 7 | 6 | 1 | 9 | 2 | 0 | 0 | 0 |
  ---------------------------------------------------------

And an additional 7 (the row length) would be stored somewhere. This is still exactly the same 2D array though, and if I asked it for the segment (2, 4) it would give me the segment ith the 2 in it. Only it the computer doesn't work out which slot this is using spatial reasoning like you or me, instead it would:

1. work out which row we are dealing with, in this case row 2
2. in response it would set the starting point of the array to just before slot 8 (containing a 6), since this is where the second row starts
3. it would then count 4 slots from there, arriving at the 2.

### Row-Major and Column-Major

The above is what is called "Row Major" 2D array ordering, i.e. where the elements are stored as rows, placed one after the other. The alternative is column-major, where the elements are stored as columns placed one after another:

  ---------------------------------------------------------
  | 8 | 6 | 7 | 1 | 9 | 9 | 1 | 2 | 4 | 0 | 7 | 0 | 7 | 0 |
  ---------------------------------------------------------

And in this case the column length (2) is written down.