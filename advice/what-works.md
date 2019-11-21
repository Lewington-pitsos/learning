# What works

1. Actually have a plan of some sort.

When starting work on a new system ensure that you take at the very least 5 minutes to make a rough plan. This might include:

- listing the major system requirements
- working out what the essential components will (probably) be
- drawing a flow chart that shows how these interact

If you don't do this you will probably end up writing code that has major design implications without realizing it, which is the same as leaving major design elements up to chance. In addition, if you start coding a system without any picture of how it should look in the end you will end up semi-unconcsiously building one as you go. This picture can actually make it more difficult to see a good design later, because you have unconsciously integrated it into your thinking too much.


2. Come up for air.

Whenever you feel flustered, frustrated or unsure of how to proceed, take 10 minutes or so to have another high level look at what you're coding. Try asking the following questions:

- does your current work reflect the original rough system plan?
- have you learned anything that might allow you to improve this plan?  

Personally I have a tendency to make a plan at the beginning and then never look at it again until I am finished, at which point the system I have built looks very different from what I originally envisioned. The implication of this is that I actually ended up doing a lot of the work without reference to any plan at all, which, as we said in 1., is probably bad.


3. Focus on big issues rather than little issues.

When building a system you will have some requirements that it needs to fulfil. These are non-negotiable. If the system does not meet them it is not complete.

Requirements aren't everything though when the system is part of a large project. A well built system will have a number of good qualities in addition to fulfilling its requirements. The larger the project the system is part of the more important these are. 

- Flexible: it should be easy to change the system in response to new demands. The more likely the demand, the more easy the change should be to implement.
- Parsable: it should be easy to understand, both for someone else, and for yourself 2 months form now.
- Reusable: it should be able to perform its function independently to the specific context it was designed for.

Let's introduce a concept called "code quality" that basically means: "the extent to which your code achieves the above". 

## Big and Little issues

Little issues are instances of bad code that do not reduce our code quality by any significant degree. Examples of little issues:

- code duplication
- big unweildly methods
- methods being private or public
- inefficient methods
- a single large object whose internal workings could be expressed as a bunch of smaller objects (internal being the operative word)
- objects having access to data they don't need

Big issues are instances of bad code that directly and significantly reduce our code quality. Examples are:

- Large chains of nested objects 

    These make changes and parsing more difficult as you have to dig your way through many method calls. They also make changes more erorr prone as your eyes glaze over while applying trivial changes to lots of different method signitures.

- Tightly coupled objects 
    
    These make changes more difficult, and can lead to unseen effects when changes are made.

- Abstractions that don't match the concepts they stand for 
    
    These and can lead to the coder forming misleading mental models of the entire system, making good design impossible. Mismatched abstractions of this kind also tend to be hard to reuse, since generally the problem you are trying to solve though reuse actually requires a single and whole concept.

- Objects with unclear responsibilities

    It is much eaiser to parse a system where every element has a single responsibility and the relations between each element are straightforward. Systems where different objects do similar things without any particular reason are going to be difficult to understand and change in the future, as future coders will have to understand, line by line, the code that makes up the elements, rather than being able to rely on abstract concepts.

- Objects requiring references to elements that, conceptually speaking, they should have nothing to do with

    This makes it harder than it should be to change the system, since now you have an arbetrary dependency injection to worry about. 