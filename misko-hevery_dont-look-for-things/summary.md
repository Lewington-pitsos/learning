# Don't Look For Things - Misko Hevery

This is a talk about how to do dependency injection. Said another way, it is about how to allow your business logic access to the objects it needs in order to function. In practice this often manifests as "what exactly do I pass in to the constructors/methods of my classes?"

## 1. Try to make constructors as slim as possible.

Testing is important. Testing a method should be easy. All you should need to do is pass some inputs to the method and assert some things about the output/object state, possible a few times. In practice though testing a single method is hard because in order to test the method you first need to initialize the class that method sits on, and initialization is often expensive.

This is bad. 

The solution is: make sure the initialization process of your classes is as minimal as possible, in particular, if a complex process required for initialization CAN be placed outside of a class constructor (or method involved in the setting-up of a class instance), do that, and pass in the (simpler) results of that process instead.

Consider a class C that requires a HTTP request to be made as part of its construction, because C needs the response for some reason. This a HTTP an expensive and complicated process. In this case C only actually needs the RESPONSE. This is a good example of what we just talked about, we can alter the constructor so that it requires the response as a parameter, and has no knowledge of the HTTP request code. The complex process can then be placed OUTSIDE the class in a separate Factory class, which just passes in the response. 

This provides some advantages:
1. Your tests become cheaper. You can mock the HTTP process, or, even simpler, create a dummy response to pass in.
2. Your tests become less complex. You no longer need to worry about errors arising from some issue in the complex HTTP process.
3. In practice your class probably becomes more reusable. In reality you might someday want to reuse the code of C in a context that does not involve a HTTP request. If this ever occurs you will be ahead of the game, because you can simple define a new factory method for that context that does not involve the HTTP stuff, rather than having to rewrite C's constructor.

A good way to tell if your constructor is a bit too extensive is if it contains references to things the class otherwise probably has no business knowing about. A class constructor should only know about things in the class. Extra things related to the class should be in a Factory method. 

An analogy: There is a car factory. During the construction process a huge, programmed metal arm needs to screw in four bolts. In this system the concept of the car includes the four bolts, but it does not include the complex robotic arm. Put another way, as the car goes about its business in the world, almost all that business involves the four bolts. If they were absent, the car would fall apart, so they are involved in the system that is the car. The business of the car does NOT involve the robotic arm. Once the arm has placed the bolts, anything can happen to it and the car's business remains unchanged. 

If we were to model this system sing code, and the car was its own class, the car's constructor should include a reference to four bolt instances, but should NOT include any reference to the robotic arm. 

## 2. The Law of Demeter: Avoid giving your business logic elements it does not need.

