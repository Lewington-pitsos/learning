# Do plan B first.

When designing a new system there is a school of thought planning the entire system out before writing a single line of code. This is probably not a good idea though. 

## Why?

In my experience it is easy to plan incorrectly. By definition when you plan you are working at a higher level of abstraction than when you are implementing the plan, often much higher. I find that when doing this it is quite easy to accidentally "gloss over" or miss important elements of the system. If this occurs it can cause significant time losses, since aspects your plan are now based on something impossible. These aspects, along with any code you have written to implement them will need to be scrapped and re-written to at least some degree, which is costly, and becomes more costly the later you discover your error. 

In addition it tempts human error. When you are 1/3rd of the way through implementing your carefully orchestrated plan for a given system and discover that your plan is not feasible and will need to be changed, you will feel (if you are anything like me) a strong impulse to preserve as much of the existing plan and codebase as possible. This impulse is not rational, and can easily lead the developer to hastily "shoe-horn" in poorly documented design components and code structures where they don't really belong to tragic effect.

The more detailed or complex your plan is, the more costly (in terms both of time and your own reserves of discipline) it will be to rework if you discover that it contains errors. 

Lastly, unless you understand the nature of the system you are trying to code *very well* you are more or less guaranteed to make plan-breaking errors of this kind. Perhaps this is a small point, but I think it bears mentioning since I find that most of the time any code system I want to build relies on something external, perhaps a new library or a website's API that I do not fully understand. As a result, a developer determined to create a comprehensive plan in cases like this probably needs to take the time to learn the new subsystems *well* before even beginning to plan, which takes discipline.

## What is the alternative?

Essentially: if creating a coherent plan for your system is proving difficult, you are likely to make planning errors, which means a lot of extra work for yourself. Instead it is probably better to plan the bits of it that *are* easy to plan and do the rest by the seat of your pants. This may not result in a good system, but in my experience it doesn't take too much time and either (A) leads to a passable system that looks kind of gross, but actually never seems to cause you any issues, or (B) allows you to see more or less exactly how the system *should* be constructed, allowing you to re-write it into a good system without running into planning errors.





## What is the alternative?

The alternative is "Plan B".

## What does that mean?

When I begin planning a new system I generally see, in very vague terms, a bunch of different candidate plans. Some more complex, some more elegant, but generally at least one will occur to me that is very basic and procedural, the kind of thing a very new programmer might write, probably involving nested for loops and very few separate objects. This is plan B, and as I begin coding one of the more elegant solutions, I tend to take some comfort form the fact that, if this elegant solution blows up in my face, I can always, in the worst case, just revert to plan B and at least my job will get (in the very loosest sense) done.

## Plan B sounds bad. Why is spaghetti code an alternative to planning?

In a nutshell: because it is much quicker and less error prone, and easy to refactor later if it becomes a problem. 


