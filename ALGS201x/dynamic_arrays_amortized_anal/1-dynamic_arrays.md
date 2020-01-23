# Dynamic Arrays

Ok, so as we know, memory gets allocated to arrays at compile time (in C presumably). This can be bad because sometimes you don't know how big you want your arrays to be at compile time.

## Semi-Dynamic Arrays: Dynamically Allocated arrays

Some languages let you allocate array memory at runtime, when you declare an array variable. This is better but sometimes tou don't even know how large your array needs to be when you declare it at runtime.

What we really want is to be able to shrink and grow arrays that we have already created at runtime. But this goes against the whole idea of arrays. The whole way arrays work is because they consist of a bunch of contiguous memory. Growing an array at runtime can only occur if we extent that contiguous section of memory, but we have no guarantee that the memory section we want to extend into isn't being used by something else. If we ensure that the memory section IS free for our array to grow into, that's the same as allocating it to the array in the first place. We can't extend the array in some other, free section of memory because that would make the array's allocated memory no longer contiguous (meaning the usual array implementation no longer works).


## The Solution: Resizeable arrays/Dynamic arrays

It's kind of what you'd expect. You stop dealing with arrays *directly* and instead create a wrapper system. The wrapper allows you to perform all the usual array operations on the internal array, with one difference: if you ask the wrapper to add an element while its internal array is full, the wrapper will create a new, larger array, copy the contents of the original array over to the new array, add the new element to the new array, and then deallocate the old array. This way we can treat the wrapper class as if it actually *is* a magical, growing array without breaking the laws that govern arrays. Systems like this are called dynamic arrays.

In general dynamic arrays are not much more expensive than normal arrays, so long as you are sensible about choosing how large to make the new, larger arrays. 