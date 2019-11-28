# Decorator pattern

# Simply Put: 

You can change the behaviour of an object at runtime by creating a wrapper that (A) implements the same interface and (B) reroutes all method calls to the inner object, but also secretly adds the desired behaviour at the same time. This is the *Decorator Pattern* and the wrapper is known as the *decorator*. You can nest several Decorators on top of each other, and add and remove them at runtime without causing issues.

## Longer

You have a class `Bird`. You sometimes want to be able to change the functionality of `Bird` instances at runtime. Say sometimes birds in your system get sick. While sick these birds act differently, but maintain the same interface. 

How do you implement this?

Simple approach is just add some more code to `Bird`. For each method in `Bird` you might have an if clause that checks if the bird instance is sick and acts accordingly. There is actually nothing inherently wrong with this on a small scale, but it starts becoming complicated when you have to cater for other conditions too like, say broodiness, old age, etc. 

Another approach might be to make a bunch of different `Bird` subclasses, `SickBird`, `OldBird` etc. You then have an `OmniBird` class which swaps these smaller classes in and out depending on conditions. This way you get a few implementations, but they are (mostly) kept separate and so as a whole the system is probably manageably complex. You might end up with some code duplication, but hopefully this can be avoided with external (or possibly static on `Bird`) helper methods/classes. However, if conditions overlap then this system cannot cope except by creating a class for every possible conjunction of conditions. This is unacceptable mainly because it will be very difficult to implement all these slightly different classes without either lot of code duplication or a lot of complexity in the code reuse system that runs between the classes, or both.

A better approach is the *decorator pattern*: you create a bunch of "decorator" objects that implement the same interface as `Bird` but expect a reference to a `Bird` that actually does most of the heavy lifting. But these decorators also add additional behaviour. For instance the `SickBirdDecorator` might re-implement `Bird`'s fly method, and make extra `Bird.ReduceStamina()` calls to the inner `Bird` before rerouting `Walk()` method calls. They can also be added and removed at runtime, which is nice. Generally all the Decorators will subclass from the same `BirdDecorator` class, since their shared reliance on bird means there's likely some common code we can put in a superclass. 

Decorators can be nested, and code reuse isn't so difficult since most of the work is still being done by `Bird`


## Issues

Identity is a bit fucked. One nice things about decorators is that you can treat a decorated object exactly the same as the original object. BUT just watch out when comparing for object equality, because a decorated object does not have the same identity as its decorated version unless you mess with the identity function.


