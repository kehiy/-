# Abstract

In this blog post i'm going to talk about some of ideas, which is in my mind about one week, this is simply just about predicting the future, nut I need to say **it's not going to be 100% practical** and it's better to consider this just as a fun idea.

# How?

I'm going to explain this idea like how I found it, so the seed of this idea was the block hashes and timestamps in blockchain. if you are not familiar with this the below image is good one, you can see blockchain data structure clearly.

![blockchain-data-structure](./Blockchain-data-structure.ppm)

As you can see each block contains the previous one [hash](https://en.wikipedia.org/wiki/Hash_function), and by the way the first block have 0 as previous block hash. 

Now, let have similar data structure like this, we call blockchain, time chain (the word most worse idea) call blocks nodes and keep data field like this:

- timestamp
- sequence number
- previous node hash
- event list (a list of events that we define them soon)

When we start to making this node for example each 10 seconds, we have a lot of hashes for each node. each hash is a number, which we can represent as base 10 number. we need to find a general formula fo n*th* block hash, something like a relation between sequence number of node and the block hash. till now, I wasn't find any way for that to be honest, but [Chaos theory](https://en.wikipedia.org/wiki/Chaos_theory) was one of thing I found, which is talk about patterns in Chaotic systems in simple term. as you know our sequence of hashes is also a Chaotic system, for example one of adjectives of a Chaotic system is `it must be sensitive to initial conditions.` like our system. but as I said, I don't have any good idea for this part (which is the main part, and that's the reason I called it a fun idea) till now.

![Chaos theory](./Figure-1The-Mandelbrot-set-1.jpg)

But we consider that we found the general formula to continue. by having the formula, how can we do our job?

At first consider that we have a time chain that makes a new node each 24 hours at the end of day. and we have about 100 nodes and also we have the general formula, we are going to predict a basketball match between team A and B in next week. first we need to define a source code which going to get the result of game from a web API for example and add it as a event in 107th node (it will happen next week BTW). now we cna calculate the 107th node hash with general formula (which is it really going to be if we have correct formula), then we can calculate 7 node with next day till next week timestamps, and with 101, 102, 103, ..., 107 nodes with  empty event and we keep node 107 hash with empty event, next we do this again and for second time we are going to consider 107 node as and node with a event which our bot will add to it if team A wins for example. now we have 3 hashes, one which is really going to be, two which is a 107th node with empty event and three is a 107th node hash which contain an win event. now we can simply compare hash 1 with 2 and 3!
If we got 1 = 2, team going to lose, and if we got 1 = 3, team going to win.

this is also and better approach to show this instead of nodes and blocks, if the hash number line crosses over the event point. event will happen. (we locate the event point in the hash number we find in the chart and the time will be next week also the cuts on lines represents each node or day)

![time-chain](./tC.png)

