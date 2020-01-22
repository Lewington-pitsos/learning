# Building a heap Efficiency

So, earlier we stated that it cost n * log(n) to build a binary heap out of an array, where n is the number of elements in that array. This makes sense, since to order a binary heap you must call `siftDown` on each element in the relevant tree (i.e. n times) and each `siftDown` call costs log(n).

... or does it. Recall that log(n) is the *maximum* levels in a tree with n nodes. But `siftDown` only costs 1 per element in the SUB-TREE whose root it is called upon. Most of the `siftDown` calls will be made on the roots of sub-trees that have less than log(n) levels. Is it possible the heap sort is even more efficient than we suspected?

## How efficient exactly?

Imagine a full, complete binary heap. Calling siftDown on the leaves of this tree will cost 0. To be generous though, let's imagine that it actually costs 1. If the tree really is full there will be n/2 leaves, so the first cost we incur is

n/2 * 1, or just n/2

The parents of those leaves will cost 2 each, and there are exactly n/4 of these, so the next level of cost will be

n/4 * 2 or 2n / 4

The Grandparents of those leaves will cost 3 each and there are n/8 of these:

n/8 * 3 or 3n/8

Following this pattern, the total cost will look something like:


n/2 + 2n/4 + 3n/8 + 4n/16 + .... 

With the final, root costing of course log(n), and there being only a single root node that costs this much. This huge sum isn't super useful though. Is there any way we can wrangle it to be a bit more useful?

Well for starters, every term has its numerator divided by n, so we can extract that so that the sum is:

n * (1/2 + 2/4 + 3/8 + 4/16)

And we can simplify this further by pointing out that each fraction takes the form of

i / 2 ^i

E.g. 2 / 2^2 =  2/4 and 3 / 2^3 = 3/8

Next we note, and we can do this geometrically by drawing lines, that a more common huge sum, of the form

1/2 + 1/4 + 1/8 + 1/16 

Is asymptotically 1.

Next, imagine we subtracted this new sum from our original one (leaving n out for the moment), we'd get:

1/4 + 2/8 + 3/16 + 4/32 + 5/64....

Which we can also prove geometrically, is asymptotically 1.

So the sum itself turned out to be 2, meaning that the total cost in the end is n * 2, or 2n.

## So what?

So ultimately we now know that the cost of constructing a binary heap is 2n. This is (in most cases) slightly more efficient than what we thought it was before, which was n * log(n), but both amount are in the same order, namely O(n). 

We didn't change any of our conclusions, but now we have a more accurate understanding of the cost.

## Implication for sorting

The efficiency of the heap sort algorithm overall isn't improved either, since the actual sorting bit (where you keep dequeueing he max element from the binary tree and shoving it at the end of the array) costs n * log(n) (you call `extractMax` n times, and the cost of `extractMax` is still log(n)). O(n * log(n)) is a higher order than simply O(n) so the overall sorting cost is still O(n * log(n)). Nothing gained.

However, we know that the build part of the heap sort algorithm is O(n), so if we, for some reason didn't need to run the whole of the sorting part of the algorithm, we might have an O(n) algorithm. 

For instance, if we only wanted the single highest value from the array, the sorting part of the algorithm would only cost log(n), and 2n + log(n) is in the order of n. Of course, in reality if we just wanted the single highest value we could get at it very simply just by iterating once, so this isn't very useful.

Can we be a little more useful though? We know that if we limit the sorting part to just 1 we have an O(n) algorithm, same if we limit it to just 2. What is the highest number such that, if we limited the sorting portion of the heap sort algorithm to just that number, we still end up with an O(n) algorithm?

Well recall that the cost for the whole list is n * log(n)... what is the highest number such that, if you multiply it by log(n), you don't get higher than n? The answer is n / log(n). If K is the number of elements you want sorted, and k <= n / log(n), then the max cost of the sorting portion of the algorithm is n / log(n) * log(n), or just n. This makes the cost for the ENTIRE algorithm 3n, which means that the whole heap sort algorithm is an order O(n) operation (as long as K <= n/log(n)).

