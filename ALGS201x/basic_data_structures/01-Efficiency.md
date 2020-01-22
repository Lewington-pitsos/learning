# Efficiency

Efficiency is often important in computing, but not usually tiny amount of efficiency. Usually the thing at issue is the *order* of efficiency of an operation. Let's use a real world example here.

## Ball in Bucket

Let's use a real world example of an operation: Imagine a huge parking lot with 1000 numbered buckets, each with different things in them. Imagine you have a special teleportation device which takes you to whichever bucket you like, if you just say its number. The thing you're going to be doing (i.e. the *type of operation* you'll be performing) is inspecting the contents of buckets. Ok cool.

### Orders of Efficiency. 

Now consider: which of the following tasks can you perform quicker:

1. Inspect the contents of bucket 99.
2. Inspect the contents of buckets until you find a bucket that contains the puppy (there is only one puppy).

Clearly in this case `1.` is much quicker to perform, you just say "take me to 99" and check what's in that bucket. Moreover, this would be the case no matter which bucket you were told to inspect or how many buckets there were: as long as each bucket is numbered it only ever takes you one try. 

Conversely `2.` will probably take you a lot longer. What you'd probably end up doing is teleporting to each bucket in sequence starting from bucket 1, checking it, and moving on if the bucket does not contain a puppy. This *could* take 1000 tries, until you find the puppy, and on average it will take 500. Importantly, the amount of effort of operation `2.` depends on the number of buckets. If there are only a few buckets, it's not that much harder than `1.`, but if there are 1 billion buckets it becomes immensely harder.

In computational terms, `1.` and `2.` fall into different categories of efficiency. `1.` is considered a *constant-time* operation, since regardless of the context it is used in (i.e. which bucket you are assigned and the total number of buckets) the operation will take the same amount of time, namely one try. We can represent this fact in notation as:

    O(1)

Which says (roughly) that the "Operation" completes in 1 action. An O(3) operation might look like this:


3. Inspect the contents of bucket 99, then inspect the contents of bucket 44, then inspect the contents of bucket 876.

In general, constant time operations are considered very efficient. You want as many as possible.

`2.` is considered a *linear-time* operation, since as the number of buckets in total grows, the more actions `2.` will take (on average) to complete. However, because each bucket has a random chance of containing the puppy, on average we will find the puppy about halfway through the forest of buckets. In notation we would write this as 

    O(n/2)

Which means that this operation will take `n/2` actions to complete, where (in this case) `n` is the total number of buckets. Importantly when we write `O(n/2)` the `n` could refer to a lot of things, not just numbers of buckets, but the key meaning to take out of `O(n/2)` is that there is *some* aspect of the operation which will take more or less time depending on the context. An example of a pure `O(n)` operation might be:

4. Inspect the contents of every bucket.

Since this will require one action for every bucket.

Finally, let's have an example of an *exponential-time* operation, one that matches the notation `O(n^2)`:

5. Each bucket contains `n` balls (where `n` is the total number of buckets) each inscribed with the bucket's number. Make it so that each bucket contains one ball from every bucket. You cannot hold more than one type of ball at once.

To complete this operation, you will need to visit the first bucket, take out all but one of it's balls, and then visit every other bucket and drop one ball in, and then repeat for every other bucket in the field. In other words, if there are 8 buckets, you are going to have to visit each bucket 8 times, and generally if there are `n` buckets, you will need to visit each bucket `n` times. This is very inefficient. These kinds of operations are to be avoided where possible.


## Why orders of efficiency?

You might be thinking that in certain cases operation `2.` might take *less* time than operation `3.`. Consider, for instance the case when there are only 3 buckets. In this case `2.` will take only 1.5 actions to complete, while `3.` will always take 3. Why is it that computer programmers put `3.` in a more efficient category than `2.`? Try to answer this for yourself.

The answer is, roughly speaking, that in modern times, individual computer actions are incredibly swift. Performing 1 operations or performing 100 won't cause any difference a human might notice, and because of this it makes sense to clump constant-time operations like `1.` and `3.` together, since even though `1.` is more efficient than `3.`, the difference isn't noticeable or important. 

The real issue is *scale*. Computers are very fast, but they often have to deal with large scales. Performing a lookup in an array of 1000 or more elements is a very common task in computing. Because of this operations like `1.` and `3.` which are *scale-independent* are very useful, since even if you have a big program that deals with huge scales of data, the program won't require too many actions. 

Consider, for instance, a program which is centered around a scale of 1,000 (i.e. something like a field of 1,000 buckets, a very normal scale in computing). If this program has to call operation `1.` 100 times, and then operation `3.` 100 times we still have only 400 actions in total for the whole program. If we have to call operation `2` (which as we recall has an efficiency of `O(n/2)`, or `O(500)` in this case), only 10 times, we already have 5,000 actions. And if we increase the number of times we have to use operation `2.`, or the scale of the program, the number of actions quickly skyrockets. Finally, if we had to perform operation `5.` only once, we instantly end up with 1,000,000 actions.

So, because computer actions are so cheap, and because computers often perform operations at great scale, computer programmers are very interested in the *scaleability* of operations when discussing their efficiency, and not particularly concerned with anything else. If an operation is scaleable, it's considered efficient.

## Orders only

At the end of the day, because we only care about scale, we don't really differentiate between O(1) and O(50), both are basically treated as O(1), as the critical point is simply that we have something better than O(n). 