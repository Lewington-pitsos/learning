# Proof By Induction

This is a very good and simple idea. You use it to prove any proposition involving integers. It goes like this:

1. Say I have a proposition, f(x) = y. 
2. If I can prove that f(1) = y
3. And I can prove that, assuming f(k) = y, then f(k + 1) = y
4. Then I have proved that f(x) = y IN GENERAL

Ok, lets go again in plainer english:

Let's say I want to convince you that some forumla holds true of every integer. For example, perhaps I want to prove that x = (2x)/2. What could I possibly do to convince you of that if you didn't believe me straight off? One thing I could do easily is prove the case where x = 1:

    1 = (2 * 1) / 2

But how do I prove it IN GENERAL? I don't have time to prove the same thing for every integer. Well if I can prove the following: that IF the theory works for some integer x, then it must also work for the next integer x + 1, then we have something quite nice. 


Assuming x = 2x /2

1. x + 1 = 2(x + 1) / 2 
2. 2x / 2 + 1 = 2(x + 1) / 2  
3. 


If a theory works for the case where x = 1. And we have proven that IF that theory forks for x, it also works for x + 1, then we know that the theory must work for x = 2. We know this because the theory works for the case where x = 1, and we know that if it works for x it must work for x + 1, and 2 is 1 + 1. 

So now we know that the theory works for x = 1 and x = 2. But we don't stop there, because it follows that the theory works for x = 3 (since it works for x = 2, and 3 is 2 + 1 and we have shown that if it works for x it works for x + 1). And so on, for every integer.

Proof by induction. It's good stuff.