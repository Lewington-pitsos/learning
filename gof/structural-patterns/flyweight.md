# Flyweight

A flyweight is an object that is used to represent a very tiny and/or fine-grained concept, one so tiny that it could argued that perhaps it *should* be represented byt a primitive like an int or a char instead. Flyweights should be treated like you would treat a char or an integer: i.e. for each type there is only ONE flyweight instance that gets referenced wherever that type is needed. 

For instance, imagine you are building a chessboard system. One could argue that it is not worth having any code to represent each square on the chessboard. It might be easier to simply store a width and a height somewhere and then have pieces track their grid positions.

If you did decide to track each board square though, the objects you use to do this should probably be super lightweight though, simply for performance reasons. Objects like this are called flyweights.

## Qualities of flyweights

One way you can make an object lightweight is by having it store as little internal data as possible within itself. Of course it is not always easy to see which data is necessarily kept in an object and which can be stored elsewhere, but a good framework for working this out is to think of the *intrinsic* verses *intrinsic* state of the object. 

### Intrinsic and Extrinsic data:
Intrinsic and Extrinsic are negations of one another. Intrinsic data is data that remains true of the object, no matter what context it appears in. If data is intrinsic to a class of objects AC, then you should be able to swap the usage contexts of two instances of that class (say A1 and A2), and that data should remain correct. A common example of data that is extrinsic is location information. Location information is only true for a given object relative to the context in which it is being used, ergo context dependent. An object that represents a chess board square might naturally store the grid reference of that square, but this would be an example of extrinsic state, since if you used that object in a different context, or swapped it with another object representing a square, that information would become false.

