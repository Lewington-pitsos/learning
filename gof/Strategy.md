# Strategy Pattern

Actually kind of similar to the decorator pattern. In the decorator pattern you have an object A with behaviour B and you want it's behaviour to be C. Instead of making a subclass of A, you wrap A in a decorator A' that has the same interface as A and re-routs most method calls to A, but also ensures behaviour C. As far as any client is concerned, nothing has changed, and you didn't have to alter A at all to achieve the behaviour change.

With the `strategy pattern` you isolate the code in A (let's call this code F) that would need to change in order to to make A's behaviour C. You rip F out of a and replace it in A with F', an interface whose implementation could either behave like B or C. You then code the C implementation and make that fulfill F'. Hey Presto, A now behaves like C.

Like the Decorator case, none of A's clients are any the wiser.

Unlike the decorator case you DID actually have to make changes to A, but as long as these changes were easy enough to isolate, they shouldn't have caused any real increase in complexity. The only penalty is some extra boilerplate when instantiating A (you have to pass in an F' implementation). 

Strategies can also be used in place of conditionals. If you find that you have a class that is behaving differently in different circumstances that you have to keep checking for, possibly instead implement two separate strategies and swap these between each other. This DOES still involve a conditional, but the logic of the different paths is now hidden.

## When to use which

If the thing that needs to vary in order to get A's behaviour from B to C is hidden way deep down in the inner workings of A and heavily relied on within A any decorator pattern implementation is likely to increase complexity a lot. This is because the decorator pattern cannot alter A at all, so any decorator will need to know a lot about A and do clever method rerouting in order to compensate. The strategy pattern is agnostic to how entangled the code in question is with the inner workings of A because all it does is replace that functionality with some implementation and A can rely on it in exactly the same way as before.

In addition, if A is a big class with lots of exposed methods, any decorator will probably end up with a lot of boilerplate in terms of re-routing calls. Strategies are agnostic to the overall side and number of exposed methods on A because they don't (need to) interact with the exposed interface of A.

The decorator pattern is probably more appropriate in cases where A is lightweight. In such cases the additional complexity and boilerplate involved in managing F' and its implementations internally can be significant proportional to A itself, and so probably not worth it. Decorators however, leave A entirely unchanged. Overall then, a single use of the Decorator pattern has a smaller cost than a single use of a strategy pattern, so if you believe the required change to A's behaviour is only a once-off, probably use a decorator.

Additionally, making A rely on internal strategy classes can impose an extra cost on clients: the cost of choosing a strategy to provision A with on instantiation. This cost is not great, but again, if A is lightweight, it can be significant proportionally. Decorators only add an extra burden to clients that what A to have additional behaviour. The rest of the code system can happily deal with plain old A. 
