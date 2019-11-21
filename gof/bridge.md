# Bridge

## tldr

Basically: you have lots of different implementations and lots of similar interfaces that need to be fulfilled by them. In response you create a lower level abstraction (probably an interface) for the general functionality that your implementations have (*entirely* without regard to the higher level interfaces you *actually* want fulfilled), call this a `concreteInterface`. You then create a `Implementor` class that takes one `concreteInterface`, which could be any of the base implementations, and use subclasses of this to fulfill all the high level interfaces.

This way you can increase the number of implementations or high level interfaces without causing an explosion of subclasses. Each new high level interface is met with a single `Implementor` class, new implementations don't cause any issues. 


## longer

You have an interface with two methods:

```
public class interface Bird {
    Fly()
    Walk()
}
```

Great. You know you'll have lots of different birds, but that they can all walk and fly. Now you can perform actions against any of them without worrying about the individual bird implementations.

Let's also assume that there are three major implementations you are using: `EfficientBird` which is super computationally efficient and `StatisticsBird` which is slower but keeps track of what methods are called on it, and `DatabaseBird` that saves certain things to the database. All three are needed in different situations. 

This is still fine, we can swap `EfficientBird`, `StatisticsBird` and `DatabaseBird` around as needed.

But now let's say a new requirement comes up: you need to involve land-based birds in your system which can also run, so now you add a new interface:
 

```
public class interface LandBird : Bird {
    Run()
}
```

Let's imagine that in this the intention of `Run` is basically just to do walk 3 times, i.e. it can be expressed entirely in terms of functionality that already exists in `Bird`. 


Ok, so `LandBird` also needs efficient and statistical implementations. The most straightforward thing to do might be to add the `Run` method to each bird implementation, or possible create a `Land` subclass for each implementation, and use these to implement `LandBird`. 

Both of these ideas quickly become unsustainable though, imagine we also want an `AquaticBird` and a `WaddlingBird`, etc. That's 3 new subclasses or class extensions to write each time.

The alternative is the `Bridge` pattern. Here you stop having your bird implementations directly fulfill your bird interfaces. 

Instead you create a new interface `ConcreteBird` that only contains the most basic bird functionality (`Walk` and `Fly`).

You then create a new `Implementor` class that has access to a `concreteBird`, and uses the `concreteBird`'s public methods to fulfill more advanced functions like `Run` or `Waddle`.

This way, for each new variant of the `Bird` interface you require, you only need to create one additional Implementor class, rather than a new concrete bird class for each of your 3 implementations (or a new method on each bird).
