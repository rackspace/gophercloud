---
layout: post
title:  "Programming Without Lambdas"
date:   2013-12-08 19:00:00
---

<blockquote>
 <p>
  That's another aspect of your application where you can save 100% of the code.
  I remain adament that local variables are not only useless, they're harmful.
  And, if you're writing code that use them, you're writing . . . non-optimal code?
 </p>
 <p>
  . . .
 </p>
 <p>
  It is necessary to have variables.
  ColorForth has got a whole slew of, well, system variables for some things . . . that are necessary.
  And, it is very useful when you're editing something to have the cursor position stored in a variable,
  so that when you come back to it the cursor is still there, and you can pick up where you left off.
 </p>
 <p>
  Variables are essential.
  I don't see any use for a small number of variables.
 </p>
 <small>Charles H. Moore, 1999, on the topic of local variables in Forth software.  <a href="https://www.youtube.com/watch?v=NK0NwqF8F0k">From a video recording made at iTVc, Inc.</a>, around 52 minutes, 55 seconds into the video.</small>
</blockquote>

Forth embodies a very libertarian philosophy of program construction.
You're free to make software as shoddy or as elegant as you want,
as carelessly or as meticulously as you desire,
as aloof or as structured as you care,
and as informal or as formally proven as you wish.
Because Forth lacks in-built conventions on how a program should "look", Forth earned itself a dubious "write-only" reputation.
Paradoxically, while some languages like Haskell have a reputation for requiring excessive discipline to use well,
I've yet to see a language as dependent on discipline and personal responsibility as I have with Forth.
I argue, the stronger the type-checker, the less discipline you need to produce reliable, readable software.

A set of programming conventions or patterns helps produce easily maintained software that also satisfies high reliability requirements, while minimizing the cognitive load on the programmer(s).
About three years ago, I published a "programming pattern" for Forth titled [Declarative, Imperative, then Inquisitive](http://sam-falvo.github.io/2010/02/27/declarative-imperative-then-inquisitive/).
It served me very well in the years since, and continues to serve me well today.
However, the DItI pattern assumes a minimum level of capability in the Forth run-time environment which the Kestrel-2's S16X4 family of processors cannot meet in hardware alone.
Therefore, I needed to arrive at a new set of programming techniques specifically for more limited environments, such as the S16X4 family of MISC processors.
In writing out what I've learned, I now have a better understanding of what Chuck Moore tried to teach the greater Forth community back around the turn of the century.

This article contains a collection of independent yet related techniques which I've developed on my own over the last three years of Kestrel-2 development and enjoyment.
First, I review the DItI pattern and discuss the slight changes necessary to adapt it to the Kestrel-2.
If you haven't read the DItI pattern yet, [you might benefit from doing so now](http://sam-falvo.github.io/2010/02/27/declarative-imperative-then-inquisitive/).
I then discuss treating software as units of virtual hardware, a modularization technique suitable for Forth in general, but the S16X4 Machine Forth in particular.
I explain why a preference for global variables exists over locals and record fields, and provide two strategies for overcoming their limitations when working with a plurality of records.
Finally, I briefly talk about tasks, in the multitasking sense of the word.

I mentioned enjoyment previously.
Do not underestimate the value of the enjoyment factor:
I typically have only a few hours per week for actual Kestrel development, so
it remains in my best interest to balance carefully my desire to document all of my software against minimizing the time taken to see something that works.
Therefore, any coding pattern must meet two important criteria:
1) It must apply generally to all classes of software, from application development to low-level systems software development, and,
2) It must directly support my writing of comment-free code in a manner that allows me to re-acquaint myself with the code no less than a year or two later, after not reading it until then.
I believe the techniques presented in this article fulfill both criteria at least as well as the DItI pattern.

### Adapting the Declarative, Imperative, then Inquisitive Pattern

Forth words may take one of several different forms.
A word may *enforce state*.
For instance, when writing a video game, you might find a word `-collision` which *guarantees* all subsequent code that a collision between the player and some other object has not yet happened.
If a collision occurred, then the burden of handling this exceptional case falls, directly or indirectly, on `-collision`.
A word may *effect state*.
For instance, a word `revealed` might draw a sprite on the video game display, thus revealing the image if it isn't already visible.
A word may *unconditionally instruct* the computer what to do next.
For instance, a word `cls` might clear the screen by zeroing out the video frame buffer, without bothering to check if it's already zeroed first.
Finally, a word may *ask a question*.
For example, as part of its implementation, `-collision` might use a word `collided?` to determine if the player has intersected any other object on the screen.

A simple example, framed in the context of a hypothetical video game, will help illustrate the pattern succinctly:

    : -collision    collided? IF explode  R> DROP THEN ;
    : player-stuff  -collision ...stuff here... ;

In this example, `-collision` fulfills the *declarative* word role,
guaranteeing any software it returns to that the player has not collided with anything.
It works by first invoking `collided?` (an *inquisitive* word) to see if a player has collided with another object in a hypothetical video game,
and if so, to `explode` (an *imperative* action) the player.
Observe that `-collision` handles the error directly.
The `R> DROP` phrase permutes the return stack by discarding the (partial) continuation that invoked `-collision`.
This returns not to the caller, but to *its* caller instead.
Thus, whatever software `-collision` returns to is safe in the knowledge that no collision occured; but,
whatever software called `player-stuff` may safely assume the player update succeeded, collision or otherwise.
In effect, this one line of code is equivalent to the following C code:

    void player_stuff() {
        if(player_collided()) {
            explode_player();
            return;
        }
        /* ... */
    }

However, the S16X4 family lacks a return stack all-together.
Words that enforce state, then, no longer enjoy the ability to manipulate a return stack to alter control flow.
Instead, we must use gaurds in front of options to gate whether or not a word executes at run-time.
We rewrite our example as follows:

    :, +collision       collided? if, explode then, ;,
    :, -collision       collided? -1 #, xor, if, ( stuff here ) then, ;,
    :, player-stuff     +collision -collision ;,

Slightly more complex, yet no less declarative in nature.
`+collision` and `-collision` become *options* in a set of decisions that `player-stuff` must perform.

The lack of a return stack coupled with only a very shallow data stack forces strict structured-programming approaches to problem solutions,
where you may enter a block of code at precisely one point,
and you leave it at precisely one other point.
Traditional Forth programming styles, used to their comfortable amounts of both return and data stack resources, simply don't apply to many MISC architectures.

### Software as Logical Hardware

Write software as if you're building hardware.
We must address the computer for what it really is:
a physical, general-purpose machine which multiplexes its time amongst one or more virtual, purpose-built machines.
Just as with digital electronics, it pays to design software in a manner that you can connect on some common substrate.
In the physical domain, these components connect to each other through *buses* (one or more wires intended to send a value to receivers);
with software, through variables in memory.
A variable holds some output of a module, and may serve as an input to another.

<a name="ruleofthumb"></a>
Separate software that implements some important piece of functionality (a *module*) and software that connects them together (a *connection* or *configuration*).
While a module makes free reference to the variables it needs,
you should only declare those variables in connection code.
Avoid, where possible, declaring variables inside components themselves (later, [I'll discuss an exception](#modulelocalvariables) to this rule of thumb).
A *component*, as in the physical world, consists of one or more modules inter-connected through some connection logic.

As long as you maintain your components hierarchically,
you should not fear a proliferation of variables.
They work exactly the same way as buses in digital electronics design, as the following table illustrates.

<table class="table">
    <tr>
        <th>Digital Logic</th>
        <th>Variables</th>
    </tr>
    <tr>
        <td>One, and only one, driver may assert a value.  If more than one driver asserts concurrently, the bus enters a race condition which at best presents an indeterminant value to other receivers, and at worst, causes hardware to overheat due to the dead shorts it causes, and the hardware suffers damage.</td>
        <td>One, and only one, module may update a value.  If more than one module updates the same variable, a race condition occurs.  At best, small amounts of information may disappear into the ether, while at worst, the software will crash hard enough to warrant a restart of the process.</td>
    </tr>
    <tr>
        <td>The threat of bus contention implies that buses offer a one-to-many information distribution model.</td>
        <td>A single module might hold responsibility for setting a variable; however, that variable often contains material of interest to a plurality of other modules.  Thus, variables, particularly in functional environments, similarly offer a one-to-many distribution model for the information it provides.</td>
    </tr>
    <tr>
        <td>An engineer remains fully aware of every bus in the system, <i>at the appropriate level of abstraction.</i>  A block diagram of a digital electronics subsystem typically shows only a subset of the buses comprising the actual hardware.  However, undisclosed buses rarely transcend a single functional unit's boundaries.</td>
        <td>Through the hierarchical composition of components, a software engineer remains fully aware of every variable used at that level of the design hierarchy.  Though a variable might sit in global memory as far as the programming language cares, a programmer stipulates that some variables make sense only at certain levels of abstraction.</td>
    </tr>
    <tr>
        <td>Buses exist <i>outside</i> of the components that rely on them.  A speaker doesn't care if you use lamp cord or 24-karat gold-plated Monster cable to hook up your sound system with.  A CMOS transistor doesn't care if the fabricator uses aluminum or copper metalization layer.  Etc.</td>
        <td>With module/connection separation, the programmer allocates storage for variables <i>outside</i> of the units he configures to use them with.  A module relies on a logical name identifying a storage location.  The programmer determines how that name maps to actual memory for his application.</td>
    </tr>
</table>

Take, for example, two programs which add and multiply two values, respectively.
It turns out we can deploy these programs in a larger system in several different ways.
The following figures illustrates two approaches.

<div class="row">
    <div class="col-md-6 img-responsive" style="text-align: center"><img src="{{site.url}}/images/addmul-arch-ex-1.dot.svg" /></div>
    <div class="col-md-6 img-responsive" style="text-align: center"><img src="{{site.url}}/images/addmul-arch-ex-2.dot.svg" /></div>
</div>

The first figure shows how both modules may draw their inputs from a shared set of variables, and deposits their results similarly into a shared output variable.
The second figure shows complete isolation of the respective modules.

In Forth, the adder and multiplier might well look like the following example listings:

    \ file: adder.fs

    : add               add_in_1 @  add_in_2 @  +  add_out ! ;

&nbsp;

    \ file: muler.fs

    : mul               mul_in_1 @  mul_in_2 @  *  mul_out ! ;

Observe how neither `adder.fs` nor `muler.fs`, in any way, care about the other module.
From their own points of view, `add` and `mul` work exclusively with their own set of variables.
How the adder and multiplier work together depends entirely on how these two modules "connect" together.
We define a set of configuration modules, which exist only to configure how these modules interact with each other.

    \ file: connections-fig-1.fs

    variable math_in_1
    variable math_in_2
    variable math_out

    : add_in_1      math_in_1 ;
    : add_in_2      math_in_2 ;
    : add_out       math_out ;
    include adder.fs

    : mul_in_1      math_in_1 ;
    : mul_in_2      math_in_2 ;
    : mul_out       math_out ;
    include muler.fs

&nbsp;

    \ file: connections-fig-2.fs

    variable add_in_1
    variable add_in_2
    variable add_out

    variable mul_in_1
    variable mul_in_2
    variable mul_out

    include adder.fs
    include muler.fs

For those who've read Leo Brodie's [Thinking Forth](http://thinking-forth.sourceforge.net), these configuration modules correspond closely to "load blocks."
For those who have experience programming FPGAs or [GreenArrays chips](http://www.greenarraychips.com), these configuration modules resemble floorplanning.
The differences lie in dimensionality; memory exposes only a single dimension (address), and you place modules relative to each other.

### Prefer Global Variables over Local Variables or Records

Common wisdom suggests avoiding global variables at all costs.
This wisdom certainly applies to common lambda-based programming techniques,
for they fail to explicitly address the clean separation of components and how to connect them.
Only recently has the greater software engineering community attempted to catalog techniques for the latter,
resulting in patterns such as [Dependency Injection](http://en.wikipedia.org/wiki/Dependency_injection).

However, fundamental assumptions that work for contemporary CPUs don't hold for MISC architectures.
In particular, global variable access often consumes significantly less time than fields relative to a base pointer.
Particularly for MISC CPUs, they also require substantially fewer words of memory to access as well.

While contemporary RISC(-derived) processors pipeline the math necessary to compute the addresses of local variables and record fields so that they run as fast as global variable access,
cache effects notwithstanding,
classical CISC and MISC architectures do not.
An engineer using global variables may witness more drastic performance gains with MISC architectures, as the following table illustrates.

<table class="table table-responsive table-striped">
    <caption>
        Table 1.  Single-beat versus read/modify/write cycles for the S16X4, S16X4A, F18A, and eP64 MISC cores, and the MC68000 and W65C816 CISC processors.
    </caption>
    <thead>
        <tr>
            <th>&nbsp;</th>
            <th colspan="2">S16X4(A)</th>
            <th colspan="2">F18A</th>
            <th colspan="2">eP64</th>
            <th colspan="2">MC68000</th>
            <th colspan="2">WD65C816 (native)</th>
        </tr>
        <tr>
            <th>&nbsp;</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th>Single</th>
            <td>7</td>
            <td>2-3</td>
            <td>11</td>
            <td>4</td>
            <td>8-9</td>
            <td>3-4</td>
            <td>12</td>
            <td>16</td>
            <td>5</td>
            <td>5</td>
        </tr>
        <tr>
            <th>R-M-W</th>
            <td>14</td>
            <td>8</td>
            <td>15</td>
            <td>8</td>
            <td>12</td>
            <td>7</td>
            <td>20</td>
            <td>24</td>
            <td>8<sup>1</sup></td>
            <td>8</td>
        </tr>
    </tbody>
</table>
<p class="small"><sup>1</sup>  If, and only if, the direct page register points to the activation frame on the processor stack.</p>

The following code fragments illustrate why.

<table class="table table-responsive table-striped">
    <caption>
        Table 2.  Code fragments for global vs. local/record accesses.  <b>B</b> and <b>V</b> represent <tt>base</tt> and <tt>var</tt> variables.  <b>O</b> represents a numeric offset.
    </caption>
    <thead>
        <tr>
            <th>&nbsp;</th>
            <th colspan="2">S16X4(A)</th>
            <th colspan="2">F18A</th>
            <th colspan="2">eP64</th>
            <th colspan="2">MC68000</th>
            <th colspan="2">WD65C816 (native)</th>
        </tr>
        <tr>
            <th>&nbsp;</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
            <th>Local</th>
            <th>Global</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <th>Single</th>
            <td><pre><b>B</b> @
<b>O</b> #
+ @</pre></td>
            <td><pre><b>V</b> @</pre></td>
            <td><pre><b>B</b> a! @a
<b>O</b> #
nop +
a! @a</pre></td>
            <td><pre><b>V</b> a! @a</pre></td>
            <td><pre><b>B</b> x! @x
<b>O</b> # +
x! @x</pre></td>
            <td><pre><b>V</b> x! @x</pre></td>
            <td><pre>MOVE.W  123(A4),D5</pre></td>
            <td><pre>MOVE.W  <b>V</b>,D5</pre></td>
            <td><pre>LDA <b>O</b>,S</pre></td>
            <td><pre>LDA <b>V</b></pre></td>
        </tr>
        <tr>
            <th>R-M-W</th>
            <td><pre><b>B</b> @
<b>O</b> #
+ @
1 # xor
<b>B</b> @
<b>O</b> #
+ !</pre></td>
            <td><pre><b>V</b> @
1 # xor
<b>V</b> !</pre></td>
            <td><pre><b>B</b> a! @a
<b>O</b> #
nop +
a! @a
1 # xor
!a</pre></td>
            <td><pre><b>V</b> a! @a
1 # xor
!a</pre></td>
            <td><pre><b>B</b> x! @x
<b>O</b> # +
x! @x
1 # xor
!x</pre></td>
            <td><pre><b>V</b> x! @x
1 # xor
!x</pre></td>
            <td><pre>BSET D5,123(A4)</pre></td>
            <td><pre>BSET D5,<b>V</b></pre></td>
            <td><pre>INC <b>O</b><sup>1</sup></pre></td>
            <td><pre>INC <b>V</b></pre></td>
        </tr>
    </tbody>
</table>
<p class="small"><sup>1</sup>  If, and only if, the direct page register points to the activation frame on the processor stack.</p>

Another advantage from looking at table 2 reveals itself through the instruction counts.
Contemporary processors rely on complex addressing modes or pipelines to hide the cost of effective address calculations.
MISC architectures expose these costs directly.
Thus, the more sophisticated the addressing mode, the more instructions necessary to calculate them.
Considering MISC's average instruction density, this may consume considerable memory resources.

Thus, on MISC architectures, globals hold a distinct advantage over local or structured field variables.
They remain considerably faster to access, and the accessing code consumes considerably fewer memory resources.

### Working with Multiple Objects

With a single set of variables declared in a configuration, a software developer may invoke a module to work with a single set of state.
This surprisingly rarely poses a problem in actual practice.
Nonetheless, the value of some programs derives from its facility with handling multiple instances of some object of interest.
Consider, for example, a text editor application.
Few text editors on the market today, open-source or commercial, present the user with a single-buffer limitation.

Let's start with a simple gap-buffer implementation.

    \ File: gap-buffer.fs

    : head?         bs @ gs @ U< ;
    : tail?         ge @ be @ U< ;
    : +space        gs @ ge @ U< 0= IF R> DROP THEN ;
    : +tail         tail? 0= IF R> DROP THEN ;
    : +head         head? 0= IF R> DROP THEN ;

    : ins           +space  gs @ c!  1 gs +! ;
    : del           +tail  1 ge +! ;
    : ->            +tail  ge @ c@  gs @ c!  1 gs +!  1 ge +! ;
    : <-            +head  -1 ge +!  -1 gs +!  gs @ c@  ge @ c! ;
    : end           BEGIN tail? WHILE -> REPEAT ;
    : home          BEGIN head? WHILE <- REPEAT ;

To keep things as simple as possible, we'll start with a single gap buffer instance.

    \ File: config-gap-buffer.fs

    VARIABLE bs ( buffer start )
    VARIABLE gs ( gap start )
    VARIABLE ge ( gap end )
    VARIABLE be ( buffer end )

    INCLUDE gap-buffer.fs

To support multiple gap buffers, we may take one of two approaches.
The first approach introduces a pointer which points to the real gap buffer state.

    \ File: config-gap-buffer-ptr.fs

    VARIABLE current-buffer

    : bs            current-buffer @ ;
    : gs            current-buffer @ CELL+ ;
    : ge            current-buffer @ [ 2 CELLS ] + ;
    : be            current-buffer @ [ 3 CELLS ] + ;

    INCLUDE gap-buffer.fs

This works quite well in practice; however, notice that we perform effective address calculations on every field reference.
This explains why, in table 1 above, so-called local variable references take so much longer than global references.

If dereferencing through a pointer proves unattractive for your needs, you may rely instead on field *caching*.
We continue to rely on global variables to hold our state; however, introduce additional functionality to switch between different instances.

    \ File: config-gap-buffer-cached.fs

    VARIABLE bs ( buffer start )
    VARIABLE gs ( gap start )
    VARIABLE ge ( gap end )
    VARIABLE be ( buffer end )

    INCLUDE gap-buffer.fs

    VARIABLE old    0 old !
    VARIABLE new    0 new !

    : writeback     old @ IF gs @  old @ CELL+ !   ge @ old @ [ 2 CELLS ] + ! THEN ;
    : fetch         new @ @ bs !  new @ CELL+ @ gs !
                    new @ [ 2 CELLS ] + @ ge !  new @ [ 3 CELLS ] + @ be ! ;
    : switch        new @ old ! ;
    : select-buf    old @ new @ XOR IF writeback fetch switch THEN ;

The observant reader might notice that the write-back operation only affects the gap pointers.
Since a buffer's address and extent will not differ during normal operation, no reason exists to waste time writing those fields back.
However, we must reload all four fields from our new gap-buffer instance.

From a performance consideration, each reference to `bs`, `gs`, `ge`, and `be` includes the amortized cost of the most recent cache fill.
Thus, caching works best with relatively small objects which do not change often.
However, dereferencing through a pointer also incurs a constant overhead with each field reference.
If performance generally doesn't matter, choose the approach which enhances program *maintenance*.
Otherwise, make sure to profile your application to determine which is best for your needs.

<a name="modulelocalvariables"></a>
### Module-Local Variables

[Earlier](#ruleofthumb), I mentioned you should strive to not define variables inside the component logic of your software.
Of course, for pragmatic reasons, I exempt variables essential to the internal operation of a module but for which aren't otherwise relevent to a module's public interface.
For example, a module to fill memory with an arbitrary value will require a starting address `start`, an ending address `end`, and the value `fillValue` to fill with.
However, unless the author designs the filler to support interruption, most external software won't care about the the current pointer, which ranges between the start and final addresses.
The author may safely declare the current pointer variable within the filler module.

    \ file: simple-filler.fs

    int, p
    :, word             fillValue @, p @, !,  p @, 2 #, +, p !, ;,
    :, fill             p @, end @, xor, if, word again, then, ;,
    :, fillMem          start @, p !,  fill ;,

The topic of deciding when it's safe to declare module-local variables raises an interesting point of discussion.
When you make use of an internal variable, you'll often find a reusable sub-module that might benefit from isolation.
Let's reconsider our memory filler example above, but in a contrived, time-sensitive application, such as a video game.
Since the amount of time `fillMem` takes to complete its task depends on how much memory we want to fill,
the author might desire something with more predictable timing characteristics to help ensure timing of, e.g., audio or video effects.
Suppose, just to pick a number, we wish to guarantee interruption after the filler moves 256 words.
We may write our filler like so:

    \ file: interruptable-filler.fs

    :, word             fillValue @, p @, !,  p @, 2 #, +, p !,  ctr @, -1 #, +, ctr !, ;,
    :, fill             p @, end @, xor, if,  ctr @, if,  word again, then, then, ;,

In this case, `fill` becomes our entrypoint, but requires a bit more initialization on the front-end.
We need to set `start`, `end`, `fillValue`, `p`, and `ctr` appropriately.
Once we initialize this state, we can invoke `fill` as many times as necessary provided `p` remains unequal to `end`.
Just make sure to reset the `ctr` variable prior to invoking `fill`, so as to maintain desired timing.
Observe the key point, however: `p`, once an internal variable, now stands with `start` and `end` as a first-class interface member.

These niggley details illustrates that truly re-usable code often costs greater complexity.
To automate the process of using `fill` for simpler use-cases, particularly those where we just don't care how long it takes,
we re-institute `fillMem`, but in terms of our now interruptable filler.

    \ file: simple-filler.fs

    int, p
    int, ctr

    include interruptable-filler.fs

    :, (fillMem)        p @, end @, xor, if,  256 #, ctr !,  fill  again, then, ;,
    :, fillMem          start @, p !, (fillMem) ;,

This version takes care of `p` and `ctr` for us, preserving our original interface to `fillMem` from above, but which relies on the interruptable version of `fill`.
If your software relied on both `fillMem` and `fill`, then we'd need to move the `p` and `ctr` declarations to a connector file:

    \ file: toc.fs

    ( ... )
    int, p
    int, ctr
    ( ... )
    include interruptable-filler.fs
    include simple-filler.fs
    ( ... )

I cannot think of any firm guidelines on when a sub-module and its dependent variables *should* be refactored into connectors and includable modules.
Knowing when to perform this transformation comes with experience and per-project contextual knowledge.

### Tasks and Program Inversion

Tasks represent a special class of imperative functionality:
it potentially may perform a different operation on each activation.
Philosophically, tasks view *returning* as a kind of system call operation,
while invoking the task word implies *resuming* from where the task left off when it last returned.
Put another way, a return instruction relinguishes the task's claim on the CPU, while
a program which invokes it again "schedules" it.
For those familiar with Jackson Structured Programming, you may recognize this as "program inversion."

A task's *variables* defines its entire state.
Put another way, a task regularly updates its variables for later resumption prior to returning.
Unlike more traditional threading,
CPU state plays no role in defining the task's current state.
You call a task to schedule CPU time to the task, and it returns when it volunteers to relinguish it.

### Conclusion

I discussed several techniques which, used collectively, enable me to write reliable software for the Kestrel at its lowest level.
Developed out of necessity, these techniques resonate well with what Chuck Moore advocated throughout the years.
Many in the Forth community dismissed his contributions as unnecessarily controversial.
Hopefully, this documentation will help other interested Forth programmers get a better understanding on the best practices long thought so controversial.

