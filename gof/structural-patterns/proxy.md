# Proxy

Sometimes it is expensive, complicated or unwise to allow direct access to an object. In such cases you can create a wrapper object that controls attempts to access the object. This wrapper is called a `Proxy`.  

## Description

Usually you refer to and keep track of an object using a simple pointer. You initialize the object, point to that place in memory, and voila: you can forget the rest. Consider the following though:

- An object which is common throughout your system, only actually used in rare cases and super expensive to initialize.

Ideally in this case you might want to have a "fake" or uninitialized version of the object used most of the time, and only initialize it when/if it is needed.

- An object which is common and super expensive to initialize, and which you often only need to query for a cheap subset of its data.

Here you'd ideally only want to initialize the object partly, unless the expensive data needs to be queried.

- An object whose internals cannot legally be accessed without the proper password.

Here you could run into trouble if someone without the password nevertheless has access to the object. Either you have to pass it around in an encrypted form (probably hard), or pass around a fake version which only gains the sensitive data when given the proper password.

In all these cases what you DON'T want is to add complexity to your system on account of the special treatment you want to give these objects. What you really want in all three cases is an object that implements the same interface as the expensive/sensitive object and acts exactly like the fully fledged version, but silently performs the expensive/sensitive operations when they are needed.

Such an object is called a `Proxy`.

## proxy types

### remote proxy

Stands in for an object that is hard to access.

### Virtual proxy

Stands in for an expensive object.

### protection proxy

The security one.

### Smart Reference

This performs additional actions whenever a reference is made to the "real" object, like perhaps counting references made or validating stuff. This one is a little different since it actually adds functionality to the object. It feels a little bit like a `decorator`, but the difference is that the purpose of a decorator is to add to the *general* behaviour of an object, while the smart reference proxy is only supposed to alter access-based behaviour.

