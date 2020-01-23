# Amortized Analysis

Basically this is working out the average efficiency. Consider, for instance the "add new element" operation on a dynamic array. This operation is usually O(1), but occasionally (i.e. whenever we happen to have reached the capacity of the inner array), it's an O(n) operation. Clearly there is a big difference between these two cases, and neither seems to capture the true cost of that operation. Amortized analysis, i.e. finding the average cost for a number of operations, is what we need in this case.

There are a few ways to do this:

## Aggregate method

Basically what this means is doing a bunch of math. Work out what your algorithm costs in all it's cases and then, based on that, work out the total cost if the algorithm is run `n` times. The only result you are interested in is what *order* of expense the operation falls into. If the total cost increases at a constant rate as we increase `n` you've got yourself some nice O(1). Alternatively, if the incremental cost gets larger as N gets larger, you have an O(n) cost.

## Bankers method

This one is a bit weird. So you work out all the costs of all the different cases of your algorithm as above. Then you pretend as if you had to "pay" the cost of every case in tokens in order to perform the operation under that case. So in the case of the dynamic array's add-new-element operation:

1. case 1 (no resize) has a cost of 1
2. case 2 (resize) has a cost of 1 + n, where n is the number of elements currently in the array

Now, with infinite tokens, clearly I can perform as many operations as I like, but the whole idea of the bankers method is to work out "how many tokens do I need to acquire every time I perform an operation, in order to ensure that I can pay for the expensive cases when they arise?".

For instance, if I am only given a single token every operation, I can pay for the first time I add a new element (or the first few, depending on how large my initial array is), but as soon as I reach case 2. I will need more than 1 token, and since I have none saved up, I cannot perform the operation. This means that a single token per operation is not enough to pay for the add-new-element operation.

We might go ahead and trial different numbers now, but to skip forward a bit, it turns out that 3 tokens per operation is enough to pay for this operations. Every time we hit a 1. operation, we pay one token and save the other two in a bank, and it turns out that whenever we hit a 2. case, we always have more than enough tokens to pay for it.

The upshot of this is that the amortized cost of the add-new-element operation is a little less than O(3), i.e. in the order of O(1). The thing to look out for is when the tokens required to proceed keep increasing more and more as you cycle through the cases. This is a sign that we probably have an amortized cost to this operation in the order of O(n).

## Physicist's Method

You kind of do the opposite of the bankers method. Conceptually speaking the amortized cost of an operation using the physicists method is the actual cost of that operation, plus the *potential* for additional future cost incurred by that operation. I.e. in the dynamic array case, the cost of case 1. is 1, but each time you do case 1 the potential for a case 2., and hence a large cost, increases.

Concretely, what you do is a lot of maths. The physicist's method amortization cost formula looks like this:

where `t` is the current operation step
where `ct` is the actual cost of `t`
where `h` is the state of the program
where `P(h)` is the Amortization function, a function that takes a program state and returns the potential cost arising from that state.

The formula goes:

    Amortization Cost per Operation = ct + P(ht) - P(h{t - 1})

In other words, the Amortization cost is the actual cost, plus the difference between the potential cost BEFORE the operation was performed on this step and the potential cost after the operation was performed on this step.

Naturally `h`: what state to take into account, and `P(h)`: how to convert the state into a cost score, are important considerations.

In this case it turns out that the state to take into account is just the capacity `cap` and the size `s` of the currently used internal array, and the cost function is:

    P(cap, size) = 2s - cap

Using these definitions, and a bunch of math we eventually can determine that the Amortized cost of both cases is 3, which is the same as what we got with the banker's method.


## Conclusion

So what the fuck is all of this good for? Well, what we learned from that is that one particular method for implementing dynamic arrays, namely the method where you double the size of the internal array each time, has an efficiency in the order of O(1), which is good. There are other implementations that have much worse efficiency, for instance where you increase the size of the internal array by just 1 each time you increase it's size. You can confirm this using any of the 3 above amortization methods.

Amortization can show you which algorithms are efficient and which aren't, when the algorithm has different costs depending on its state.