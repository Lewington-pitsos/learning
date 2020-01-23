# Don't Look For Things - Misko Hevery

This is a talk about how to do dependency injection. Said another way, it is about how to allow your business logic access to the objects it needs in order to function. In practice this often manifests as "what exactly do I pass in to the constructors/methods of my classes?"

## 1. Try to make constructors as slim as possible.

Testing is important. Testing a method should be easy. All you should need to do is pass some inputs to the method and assert some things about the outputs/object state, possibly a few times. In practice though testing a single method is often hard because in order to test the method you first need to initialize the class that method sits on, and initialization is often expensive.

This is bad. 

The solution is: make sure the initialization process of your classes is as minimal as possible, in particular, if a complex process required for initialization CAN be placed outside of a class constructor (or any method involved in the setting-up of a class instance), do that, and pass in the (simpler) results of that process instead.

Consider a class C that requires a HTTP request to be made as part of its construction, because C needs the response for some reason. This HTTP request is an expensive and complicated process. In this case C only actually needs the RESPONSE. This is a good example of what we just talked about: we can alter the constructor so that it requires the response as a parameter, and has no knowledge of the HTTP request code. The complex process can then be placed OUTSIDE the class in a separate Factory class*, which just passes in the response. 

This provides some advantages:
1. Your tests become cheaper. You can mock the HTTP process, or, even simpler, create a dummy response to pass in.
2. Your tests become less complex. You no longer need to worry about errors arising from some issue in the complex HTTP process.
3. In practice your class probably becomes more reusable. In reality you might someday want to reuse the code of C in a context that does not involve a HTTP request. If this ever occurs you will be ahead of the game, because you can simply define a new factory method for that context that does not involve the HTTP stuff, rather than having to rewrite C's constructor.

A good way to tell if your constructor is a bit too extensive is if it contains references to things the class otherwise probably has no business knowing about. A class constructor should only know about things in the class. Extra things related to the class should be in a Factory method. 

### *

Factory classes. This is actually possibly the biggest upshot of this whole spiel: extract complex and out-of-place setup behaviour to a factory class method. This way you don't pollute your class OR the code that initiates the setup procedure.
 
## 2. The Law of Demeter: Avoid giving your business logic elements it does not need.

The law of demeter is that classic thing where

    "try not to let anything in your codebase know about anything that isn't related to it"

Naturally the main reason for this is "coupling". 

Another reason (brought up in the talk) is that violations of the law of demeter directly make testing much harder. As we said earlier, instantiating objects is often expensive. It can also cascade, as objects require other yet objects to be instantiated and so on. The functionality you want to test on a given method, or even a given class is often actually quite cheap to test (in terms of mental and computational load) so every object instantiated is a relatively significant burden increase.

Furthermore, if one of your methods includes a reference to a complex system, every time you want to test that method you are going to have to refresh your memory of how that complex system works *and* how it is used in the method under test, since NEITHER is likely to be clear at a glance. Simple systems don't cause the same problem to the same degree.

All of this being the case, having your methods (constructor or otherwise) refer to systems they could get away with not referring to introduces a big testing cost. The more numerous and complex the systems the greater the cost.

In particular, you should ensure that your methods only reference the simplest and least objects possible.

### Getter objects

In particular, some codebases include a "getter" class that just holds references to a bunch of things that might be needed around the codebase. These are kind of the perfect storm of what we just spoke of because they are at best opaque and at worst complex and opaque and always have an expensive construction process that references large quantities of other objects. This makes testing crazy annoying, since for every class you want to test you need to instantiate a hoard of classes, regardless of how simple that class is. PLUS you need a refreshed on the internal workings of the getter class and, possibly even worse, the expectations the class your testing have about the state of the getter class every time you want to revisit the test.

Don't use these pls.

### Nested Objects and Superfluous Parameters

One thing that sometimes happens is: you have a class A. A has a child B and B has a child C. C requires a further classes D and E in order to set itself up, but A and B do not. This is annoying, because lots of possible setups kind of suck.

A very basic solution might be: pass B and D into A, then have A pass E and D to B, and finally have B pass E and D to C. This sucks though, because you have un-needed, possibly misleading parameters, plus un-needed references to E and D in B and A's constructors. This is worse if there's more parameters like E and D that C needs. Changes are going to be really boring and annoying.

    NewA(... param1 D, param2 E){ ... }

    NewB(... param1 D, param2 E){ ... }

    NewC(param1 D, param2 E){ ... }


One possible improvement might be to bundle ALL the parameters, those needed by A, B or C, into a single object called like, Params. Then simply pass Params through all 3. Any future changes are made to params and only the class that the changes relate to and all is well. The issue here is that this kind of violates the law of demeter, because A, B and C have access to data that they don't need. 

This is not so large a violation though, and one could argue that it is still preferable to passing each parameter down the tree, especially as the number of parameters increase. There's a 3rd better option though: Factory methods.

See, by having A know about C's parameters in the first place we're already kind of violating the law of demeter. A good constructor for A should really look like:

    NewA(child B){ ... }

Is this bad though? Now the construction logic, which should really belong inside A, because nothing else cares about B or C, is now the responsibility of whatever is instantiating A's. That sucks.

Au contraire: you just put all that logic into a factory method, one that instantiates C, passes that to B and finally passes *that* to A. Now when you want to instantiate an A, you pass all the original parameters in to the factory method, D and E included, and it will spit out an A. As long as you have a factory method, instantiating A remains just as cheap as before. 

This includes testing. Don't think you have to manually instantiate a B whenever you need to test A. Worst case, you can just rely on the factory method again and we've lost nothing. Best case you can now create a dummy of B and pass it in easily, or create a fancy, dummy-creating alternative factory methods. Before both of these would have been impossible.

And Note that the factory method DOES clear up our original issue: A and B no longer have superfluous references to D and E, and if C acquires more dependencies only the factory method will need to worry about them. 

You can almost have this rule, in fact: **never call "new" or the language equivalent, EXCEPT inside factory methods, or possibly tests**. This applies more the more complex the object being instantiated.


### Lifetime factory methods

You might argue that there is one downside with this strategy: you will need to create a factory method for every class, thereby doubling the number of classes.

In reality this is not the case though. Usually you only need one factory method for each "lifecycle" of classes i.e. those classes that get instantiated and destroyed at the same time, and usually there aren't that many lifecycles. For instance above we had a single factory class for A, B and C. At most you might need to add some new methods to that factory if you need to instantiate B and C independently elsewhere in your code (unlikely), but 1 factory per class is never going to happen.

## Validation

Smaller point, but validation might not be worthwhile. Firstly, because tests are better than validation, and can do the same job. I.e. the idea behind validating is to prevent mysterious silent failures later on by catching the problem when it is obvious. Good testing will do the same job because it will either break when the problem becomes possible, or allow the coder to pinpoint the issue by giving the green light to everything else. 

Both testing and validation are protection against future mistakes, but testing is better because it can be systematized. Testing occurs OUTSIDE your actual code, so you can develop testing paradigms and test infrastructure whose effectiveness is only bounded by how clever you are. Validation, on the other hand occurs IN your code, so any paradigms you come up with must take your code into account. This basically means that creating good paradigms is very very hard, which basically means that you can't systematize validation. As projects grow, things that are systematic quickly become more and more valuable, and things that aren't quickly become hazardous. 

So, given that both testing and validation can do the same job, and only testing can be systematized, you should always test and never validate. When you add tests you're building the momentum of your test system, when you validate you're adding cruft to your code.

Also validation can fuck with you when you're trying to test if you do it wrong.

## Segregation

A good paradigm for your entire application is a split between instantiation related code and business logic related code. Naturally you are bound to instantiate some things in your business logic, but if you are careful you should be able to limit this to data classes and things like arrays. Any class that has its own semantics in your application, or requires a complex setup procedure should only be instantiated by separate factory classes.

If you can do this you will save yourself a lot of headache.

In addition, if you do it properly it is possible to auto-generate most of your factory methods (in like C with its decades of people making tools for it).