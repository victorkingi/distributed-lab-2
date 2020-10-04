# RPC in GO

## 1: Deploying the Chat System

Following from the end of last week's lab, you should now have a distributed
system -- a chat client and server. Hopefully you've already tested these out 
locally, but now you need to put them to the test by running them in a properly
distributed fashion.

For this, we will make use of AWS. Follow the
[guides](https://github.com/UoB-CSA/setup-guides/aws) for setting
up instances, accessing them, and opening required ports.

In particular, you should:

+ create two t2.micro instances
+ make sure Go is installed on your instances
+ load your client/server code onto the instances (e.g., via git)
+ start your server running on one instance
+ connect to your server with your client on another instance
+ connect to your server from your local machine
+ communicate between clients via your (genuinely) distributed system!


## 2: Using RPC - Secret Strings

Follow the video from this week to create a simple RPC system that allows
clients to call the functions of a server via a defined interface. As with last
week, this is best attempted in stages:

+ Stage 1: Write server code to enable access to the "secret" string
manipulation function. Test it by writing a client that sends a string to be
reversed.
+ Stage 2: Enable Premium Tier service by implementing the `FastReverse`
function in the server.
+ Stage 3: Update your client to read words from the `wordlist` file and reverse them all.
+ (Optional) run multiple server instances, and speed up the work by
having your client split the load between servers.

## 3: 99 Bottles of Beer

So far we've been focusing on client-server systems. However, there are times
when we don't really want a distinction in roles between components -- when we
might want them to act as peers, all running the same code.

For this task, you are going to solve the ["99 Bottles of
Beer"](https://en.wikipedia.org/wiki/99_Bottles_of_Beer) problem in a
distributed fashion. Rather than a single process singing the lyrics all by itself,
you're going to run at least three instances of your code, on different
machines, and have these buddies share the work of singing the song
verse-by-verse, in order. 

Here's how it should go:

**Buddy 1**: "99 bottles of beer on the wall, 99 bottles of beer. Take one down, pass it around..."
**Buddy 2**: "98 bottles of beer on the wall, 98 bottles of beer. Take one down, pass it around..."
**Buddy 3**: "97 bottles of beer on the wall, 97 bottles of beer. Take one down, pass it around..."
**Buddy 1**: "96 bottles of beer on the wall, 96 bottles of beer. Take one down, pass it around..."

All the way down until there are no more bottles of beer on the wall, which the
final singer should note appropriately with your preferred ending to the song.

Use 3 for testing, but your solution should be able to handle any number of
buddies singing along. There should be no difference in the code, just different
flags on the commandline.

###Hints: 

1. Each process needs to accept an `ip:port` string for the 'next' buddy who will
follow on from them in the song. You'll have to configure them in a loop. 
2. You don't want clients to try connect to each other straight away, or you
won't have time to set the final process running so that the first can connect.
3. When you set up the processes, you'll also need some way to indicate which of them 
should start the song (I suggest allowing any `n` bottles of beer, for testing purposes). 
4. You may need to look at `client.Go` rather than `client.Call`. 
