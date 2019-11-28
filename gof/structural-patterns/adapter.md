# Adapter

An adapter is supposed to allow system A to interact with system B, without system A having to know anything about the interface of system B. 

So, for example, system B might implement interface B', and you might want to use system B as part of system A. But B' might not make sense in the context of system A. To prevent B' from polluting system A, you might create an *adapter* object C with interface C' makes sense in the context of system A. C simply wraps around an implementation of B' and converts C' calls to the appropriate calls or sequences of calls on B'.

## Examples

You are designing a system that involves roads and cars travelling along them. One of the elements needed in this system is `Roundabout`. Roundabout needs to update the positions of all the cars currently on it each tick. It's interface only consists of `Update()` `GetCarPositions()` and `AddCar(Car)`. 

Working out a car's position on a roundabout requires a some complex circle-based mathematics though. 

From another project you have a `Circle` class that represents a circular space, and tracks the positions of objects inside it. This class actually contains most of what Roundabout needs to do, but it's interface has nothing to do with cars and it doesn't update itself. 

This is a perfect place to use an **Adapter**. You can implement it using inheritance or composition. 

### Inheritance

You create a RoundAbout class that inherits from Circle. You make RoundAbout implement the required interface mostly by calling it's parent's methods, but also perhaps by defining a few other methods to fill the gaps. 

#### Advantages

You can overwrite Circle's methods very easily if need be, just by adding method with the same signature to RoundAbout.

#### Disadvantages

Since RoundAbout is inheriting *directly* from Circle, if Circle had a subclass (say `Donut`) that you wanted to use instead for certain round about elements in your system, you'd need to define a whole new RoundAbout class that inherits from Donut instead.

### Composition

You create a RoundAbout class that is initialized with a reference to a Circle instance. You make RoundAbout implement the required interface by calling the Circle instance's methods, and fill the gaps as above. 



