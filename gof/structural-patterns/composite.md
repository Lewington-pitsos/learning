# Composite

The composite pattern occurs when a single instance of an object and a group of instances can be treated in exactly the same way. For e.g. the code is responsible for individuals and the code responsible for groups conform to the same interface.

For instance, lets say we're simulating an army. You might have a `soldier` class representing an individual entity, and maybe a bunch of different `soldier` subclasses for soldiers with more health or less damage etc, and a `unit`, class representing a group of `soldier`'s. You want to be able to treat a `unit` in the same was as a `soldier`, so you define an interface `combatant` and make them both implement it. This way you can also have `unit`'s comprised of `unit`s rather than `soldier`s.

In composite pattern terms the `soldier` is a "leaf", the `unit` is a "composite" and anything that implements `combatant` (i.e. both) is considered a "component".

## Different implementations

All well and good, but clearly `unit` has some functionality that `soldier` doesn't. For instance, `unit` will need methods like `removeRandomCombatant`, which clearly wouldn't make sense on a single `soldier`.  

Basically the response to this is: just give `soldier` a dummy version of `removeRandomCombatant`. If `soldier.removeRandomCombatant()` gets called just do a noop and nothing bad should happen *in general*. If calling composite methods on leaves might lead to anomalous shit happening, maybe don't use the pattern. Alternatively you can add a runtime check of whether a component is a composite or a leaf. for instance `combatant.isSoldier()`, and `soldier` and `unit` implementations return different values. This does reduce the main advantage of the composite pattern though, which is that clients don't have to worry about what they're dealing with, so if you find yourself considering doing this maybe, again don't use the pattern.

Also, note that if, conceptually, you treat `soldier` not as an individual but as *a group always containing one individual*, methods like `removeRandomCombatant` or `combatantCount` actually do make conceptual sense. You might even implement a method like `soldier.removeRandomCombatant` and have it convert the internals of `soldier` to something that works like an empty `unit`.

## Benefits

Main thing is it allows the client to do a lot less legwork.

## Downsides

It can make your code too general. Let's say you want to conceptually distinguish a `unit` that only contains other `units` rather than individual soldiers. It is not straightforward how you would do this, and you probably would have to rely on checking at run-time. Basically don't use the composite pattern unless your system can be expressed in very general terms.

