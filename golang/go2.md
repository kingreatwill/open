[TOC]

# [Toward Go 2](https://blog.golang.org/toward-go2)

Russ Cox  
13 July 2017

#### Introduction

[This is the text of [my talk today](https://www.youtube.com/watch?v=0Zbh_vmAKvk) at Gophercon 2017, asking for the entire Go community's help as we discuss and plan Go 2.]

On September 25, 2007, after Rob Pike, Robert Griesemer, and Ken Thompson had been discussing a new programming language for a few days, Rob suggested the name “Go.”

![](https://blog.golang.org/toward-go2/mail.png)

The next year, Ian Lance Taylor and I joined the team, and together the five of us built two compilers and a standard library, leading up to the [open-source release](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html) on November 10, 2009.

![](https://blog.golang.org/toward-go2/tweet.png)

For the next two years, with the help of the new Go open source community, we experimented with changes large and small, refining Go and leading to the [plan for Go 1](https://blog.golang.org/preview-of-go-version-1), proposed on October 5, 2011.

![](toward-go2/go1-preview.png)

With more help from the Go community, we revised and implemented that plan, eventually [releasing Go 1](https://blog.golang.org/go-version-1-is-released) on March 28, 2012.

![](toward-go2/go1-release.png)

The release of Go 1 marked the culmination of nearly five years of creative, frenetic effort that took us from a name and a list of ideas to a stable, production language. It also marked an explicit shift from change and churn to stability.

In the years leading to Go 1, we changed Go and broke everyone's Go programs nearly every week. We understood that this was keeping Go from use in production settings, where programs could not be rewritten weekly to keep up with language changes. As the [blog post announcing Go 1](https://blog.golang.org/go-version-1-is-released) says, the driving motivation was to provide a stable foundation for creating reliable products, projects, and publications (blogs, tutorials, conference talks, and books), to make users confident that their programs would continue to compile and run without change for years to come.

After Go 1 was released, we knew that we needed to spend time using Go in the production environments it was designed for. We shifted explicitly away from making language changes toward using Go in our own projects and improving the implementation: we ported Go to many new systems, we rewrote nearly every performance-critical piece to make Go run more efficiently, and we added key tools like the [race detector](https://blog.golang.org/race-detector).

Now we have five years of experience using Go to build large, production-quality systems. We have developed a sense of what works and what does not. Now it is time to begin the next step in Go's evolution and growth, to plan the future of Go. I'm here today to ask all of you in the Go community, whether you're in the audience at GopherCon or watching on video or reading the Go blog later today, to work with us as we plan and implement Go 2.

In the rest of this talk, I'm going to explain our goals for Go 2; our constraints and limitations; the overall process; the importance of writing about our experiences using Go, especially as they relate to problems we might try to solve; the possible kinds of solutions; how we will deliver Go 2; and how all of you can help.

#### Goals

The goals we have for Go today are the same as in 2007. We want to make programmers more effective at managing two kinds of scale: production scale, especially concurrent systems interacting with many other servers, exemplified today by cloud software; and development scale, especially large codebases worked on by many engineers coordinating only loosely, exemplified today by modern open-source development.

These kinds of scale show up at companies of all sizes. Even a five-person startup may use large cloud-based API services provided by other companies and use more open-source software than software they write themselves. Production scale and development scale are just as relevant at that startup as they are at Google.

Our goal for Go 2 is to fix the most significant ways Go fails to scale.

(For more about these goals, see Rob Pike's 2012 article “[Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)” and my GopherCon 2015 talk “[Go, Open Source, Community](https://blog.golang.org/open-source).”)

#### Constraints

The goals for Go have not changed since the beginning, but the constraints on Go certainly have. The most important constraint is existing Go usage. We estimate that there are at least [half a million Go developers worldwide](https://research.swtch.com/gophercount), which means there are millions of Go source files and at least a billion of lines of Go code. Those programmers and that source code represent Go's success, but they are also the main constraint on Go 2.

Go 2 must bring along all those developers. We must ask them to unlearn old habits and learn new ones only when the reward is great. For example, before Go 1, the method implemented by error types was named `String`. In Go 1, we renamed it `Error`, to distinguish error types from other types that can format themselves. The other day I was implementing an error type, and without thinking I named its method `String` instead of `Error`, which of course did not compile. After five years I still have not completely unlearned the old way. That kind of clarifying renaming was an important change to make in Go 1 but would be too disruptive for Go 2 without a very good reason.

Go 2 must also bring along all the existing Go 1 source code. We must not split the Go ecosystem. Mixed programs, in which packages written in Go 2 import packages written in Go 1 and vice versa, must work effortlessly during a transition period of multiple years. We'll have to figure out exactly how to do that; automated tooling like go fix will certainly play a part.

To minimize disruption, each change will require careful thought, planning, and tooling, which in turn limits the number of changes we can make. Maybe we can do two or three, certainly not more than five.

I'm not counting minor housekeeping changes like maybe allowing identifiers in more spoken languages or adding binary integer literals. Minor changes like these are also important, but they are easier to get right. I'm focusing today on possible major changes, such as additional support for error handling, or introducing immutable or read-only values, or adding some form of generics, or other important topics not yet suggested. We can do only a few of those major changes. We will have to choose carefully.

#### Process

That raises an important question. What is the process for developing Go?

In the early days of Go, when there were just five of us, we worked in a pair of adjacent shared offices separated by a glass wall. It was easy to pull everyone into one office to discuss some problem and then go back to our desks to implement a solution. When some wrinkle arose during the implementation, it was easy to gather everyone again. Rob and Robert's office had a small couch and a whiteboard, so typically one of us went in and started writing an example on the board. Usually by the time the example was up, everyone else had reached a good stopping point in their own work and was ready to sit down and discuss it. That informality obviously doesn't scale to the global Go community of today.

Part of the work since the open-source release of Go has been porting our informal process into the more formal world of mailing lists and issue trackers and half a million users, but I don't think we've ever explicitly described our overall process. It's possible we never consciously thought about it. Looking back, though, I think this is the basic outline of our work on Go, the process we've been following since the first prototype was running.

![](toward-go2/process.png)

Step 1 is to use Go, to accumulate experience with it.

Step 2 is to identify a problem with Go that might need solving and to articulate it, to explain it to others, to write it down.

Step 3 is to propose a solution to that problem, discuss it with others, and revise the solution based on that discussion.

Step 4 is to implement the solution, evaluate it, and refine it based on that evaluation.

Finally, step 5 is to ship the solution, adding it to the language, or the library, or the set of tools that people use from day to day.

The same person does not have to do all these steps for a particular change. In fact, usually many people collaborate on any given step, and many solutions may be proposed for a single problem. Also, at any point we may realize we don’t want to go further with a particular idea and circle back to an earlier step.

Although I don't believe we've ever talked about this process as a whole, we have explained parts of it. In 2012, when we released Go 1 and said that it was time now to use Go and stop changing it, we were explaining step 1. In 2015, when we introduced the Go change proposal process, we were explaining steps 3, 4, and 5. But we've never explained step 2 in detail, so I'd like to do that now.

(For more about the development of Go 1 and the shift away from language changes, see Rob Pike and Andrew Gerrand's OSCON 2012 talk “[The Path to Go 1](https://blog.golang.org/the-path-to-go-1).” For more about the proposal process, see Andrew Gerrand's GopherCon 2015 talk “[How Go was Made](https://www.youtube.com/watch?v=0ht89TxZZnk)” and the [proposal process documentation](https://golang.org/s/proposal).)

#### Explaining Problems

![](toward-go2/process2.png)

There are two parts to explaining a problem. The first part—the easier part—is stating exactly what the problem is. We developers are decently good at this. After all, every test we write is a statement of a problem to be solved, in language so precise that even a computer can understand it. The second part—the harder part—is describing the significance of the problem well enough that everyone can understand why we should spend time solving it and maintaining a solution. In contrast to stating a problem precisely, we don't need to describe a problem's significance very often, and we're not nearly as good at it. Computers never ask us “why is this test case important? Are you sure this is the problem you need to solve? Is solving this problem the most important thing you can be doing?” Maybe they will someday, but not today.

Let's look at an old example from 2011. Here is what I wrote about renaming os.Error to error.Value while we were planning Go 1.

![](toward-go2/error.png)

It begins with a precise, one-line statement of the problem: in very low-level libraries everything imports "os" for os.Error. Then there are five lines, which I've underlined here, devoted to describing the significance of the problem: the packages that "os" uses cannot themselves present errors in their APIs, and other packages depend on "os" for reasons having nothing to do with operating system services.

Do these five lines convince _you_ that this problem is significant? It depends on how well you can fill in the context I've left out: being understood requires anticipating what others need to know. For my audience at the time—the ten other people on the Go team at Google who were reading that document—those fifty words were enough. To present the same problem to the audience at GothamGo last fall—an audience with much more varied backgrounds and areas of expertise—I needed to provide more context, and I used about two hundred words, along with real code examples and a diagram. It is a fact of today's worldwide Go community that describing the significance of any problem requires adding context, especially illustrated by concrete examples, that you would leave out when talking to coworkers.

Convincing others that a problem is significant is an essential step. When a problem appears insignificant, almost every solution will seem too expensive. But for a significant problem, there are usually many solutions of reasonable cost. When we disagree about whether to adopt a particular solution, we're often actually disagreeing about the significance of the problem being solved. This is so important that I want to look at two recent examples that show this clearly, at least in hindsight.

#### Example: Leap seconds

My first example is about time.

Suppose you want to time how long an event takes. You write down the start time, run the event, write down the end time, and then subtract the start time from the end time. If the event took ten milliseconds, the subtraction gives a result of ten milliseconds, perhaps plus or minus a small measurement error.

start := time.Now()       // 3:04:05.000
event()
end := time.Now()         // 3:04:05.010

elapsed := end.Sub(start) // 10 ms

This obvious procedure can fail during a [leap second](https://en.wikipedia.org/wiki/Leap_second). When our clocks are not quite in sync with the daily rotation of the Earth, a leap second—officially 11:59pm and 60 seconds—is inserted just before midnight. Unlike leap years, leap seconds follow no predictable pattern, which makes them hard to fit into programs and APIs. Instead of trying to represent the occasional 61-second minute, operating systems typically implement a leap second by turning the clock back one second just before what would have been midnight, so that 11:59pm and 59 seconds happens twice. This clock reset makes time appear to move backward, so that our ten-millisecond event might be timed as taking negative 990 milliseconds.

start := time.Now()       // 11:59:59.995
event()
end := time.Now()         // 11:59:59.005 (really 11:59:60.005)

elapsed := end.Sub(start) // –990 ms

Because the time-of-day clock is inaccurate for timing events across clock resets like this, operating systems now provide a second clock, the monotonic clock, which has no absolute meaning but counts seconds and is never reset.

Except during the odd clock reset, the monotonic clock is no better than the time-of-day clock, and the time-of-day clock has the added benefit of being useful for telling time, so for simplicity Go 1’s time APIs expose only the time-of-day clock.

In October 2015, a [bug report](https://golang.org/issue/12914) noted that Go programs could not time events correctly across clock resets, especially a typical leap second. The suggested fix was also the original issue title: “add a new API to access a monotonic clock source.” I argued that this problem was not significant enough to justify new API. A few months earlier, for the mid-2015 leap second, Akamai, Amazon, and Google had slowed their clocks a tiny amount for the entire day, absorbing the extra second without turning their clocks backward. It seemed like eventual widespread adoption of this “[leap smear](https://developers.google.com/time/smear)” approach would eliminate leap-second clock resets as a problem on production systems. In contrast, adding new API to Go would add new problems: we would have to explain the two kinds of clocks, educate users about when to use each, and convert many lines of existing code, all for an issue that rarely occurred and might plausibly go away on its own.

We did what we always do when there's a problem without a clear solution: we waited. Waiting gives us more time to add experience and understanding of the problem and also more time to find a good solution. In this case, waiting added to our understanding of the significance of the problem, in the form of a thankfully [minor outage at Cloudflare](https://www.theregister.co.uk/2017/01/04/cloudflare_trips_over_leap_second/). Their Go code timed DNS requests during the end-of-2016 leap second as taking around negative 990 milliseconds, which caused simultaneous panics across their servers, breaking 0.2% of DNS queries at peak.

Cloudflare is exactly the kind of cloud system Go was intended for, and they had a production outage based on Go not being able to time events correctly. Then, and this is the key point, Cloudflare reported their experience in a blog post by John Graham-Cumming titled “[How and why the leap second affected Cloudflare DNS](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/).” By sharing concrete details of their experience with Go in production, John and Cloudflare helped us understand that the problem of accurate timing across leap second clock resets was too significant to leave unfixed. Two months after that article was published, we had designed and implemented a solution that will [ship in Go 1.9](https://beta.golang.org/doc/go1.9#monotonic-time) (and in fact we did it with [no new API](https://golang.org/design/12914-monotonic)).

#### Example: Alias declarations

My second example is support for alias declarations in Go.

Over the past few years, Google has established a team focused on large-scale code changes, meaning API migration and bug fixes applied across our [codebase of millions of source files and billions of lines of code](http://cacm.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/pdf) written in C++, Go, Java, Python, and other languages. One thing I've learned from that team's work is the importance, when changing an API from using one name to another, of being able to update client code in multiple steps, not all at once. To do this, it must be possible to write a declaration forwarding uses of the old name to the new name. C++ has #define, typedef, and using declarations to enable this forwarding, but Go has nothing. Of course, one of Go's goals is to scale well to large codebases, and as the amount of Go code at Google grew, it became clear both that we needed some kind of forwarding mechanism and also that other projects and companies would run into this problem as their Go codebases grew.

In March 2016, I started talking with Robert Griesemer and Rob Pike about how Go might handle gradual codebase updates, and we arrived at alias declarations, which are exactly the needed forwarding mechanism. At this point, I felt very good about the way Go was evolving. We'd talked about aliases since the early days of Go—in fact, the first spec draft has [an example using alias declarations](https://go.googlesource.com/go/+/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec#1182)—but each time we'd discussed aliases, and later type aliases, we had no clear use case for them, so we left them out. Now we were proposing to add aliases not because they were an elegant concept but because they solved a significant practical problem with Go meeting its goal of scalable software development. I hoped this would serve as a model for future changes to Go.

Later in the spring, Robert and Rob wrote [a proposal](https://golang.org/design/16339-alias-decls), and Robert presented it in a [Gophercon 2016 lightning talk](https://www.youtube.com/watch?v=t-w6MyI2qlU). The next few months did not go smoothly, and they were definitely not a model for future changes to Go. One of the many lessons we learned was the importance of describing the significance of a problem.

A minute ago, I explained the problem to you, giving some background about how it can arise and why, but with no concrete examples that might help you evaluate whether the problem might affect you at some point. Last summer’s proposal and the lightning talk gave an abstract example, involving packages C, L, L1, and C1 through Cn, but no concrete examples that developers could relate to. As a result, most of the feedback from the community was based on the idea that aliases only solved a problem for Google, not for everyone else.

Just as we at Google did not at first understand the significance of handling leap second time resets correctly, we did not effectively convey to the broader Go community the significance of handling gradual code migration and repair during large-scale changes.

In the fall we started over. I gave a [talk](https://www.youtube.com/watch?v=h6Cw9iCDVcU) and wrote [an article presenting the problem](https://talks.golang.org/2016/refactor.article) using multiple concrete examples drawn from open source codebases, showing how this problem arises everywhere, not just inside Google. Now that more people understood the problem and could see its significance, we had a [productive discussion](https://golang.org/issue/18130) about what kind of solution would be best. The outcome is that [type aliases](https://golang.org/design/18130-type-alias) will be [included in Go 1.9](https://beta.golang.org/doc/go1.9#language) and will help Go scale to ever-larger codebases.

#### Experience reports

The lesson here is that it is difficult but essential to describe the significance of a problem in a way that someone working in a different environment can understand. To discuss major changes to Go as a community, we will need to pay particular attention to describing the significance of any problem we want to solve. The clearest way to do that is by showing how the problem affects real programs and real production systems, like in [Cloudflare's blog post](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/) and in [my refactoring article](https://talks.golang.org/2016/refactor.article).

Experience reports like these turn an abstract problem into a concrete one and help us understand its significance. They also serve as test cases: any proposed solution can be evaluated by examining its effect on the actual, real-world problems the reports describe.

For example, I've been examining generics recently, but I don't have in my mind a clear picture of the detailed, concrete problems that Go users need generics to solve. As a result, I can't answer a design question like whether to support generic methods, which is to say methods that are parameterized separately from the receiver. If we had a large set of real-world use cases, we could begin to answer a question like this by examining the significant ones.

As another example, I’ve seen proposals to extend the error interface in various ways, but I haven't seen any experience reports showing how large Go programs attempt to understand and handle errors at all, much less showing how the current error interface hinders those attempts. These reports would help us all better understand the details and significance of the problem, which we must do before solving it.

I could go on. Every major potential change to Go should be motivated by one or more experience reports documenting how people use Go today and why that's not working well enough. For the obvious major changes we might consider for Go, I'm not aware of many such reports, especially not reports illustrated with real-world examples.

These reports are the raw material for the Go 2 proposal process, and we need all of you to write them, to help us understand your experiences with Go. There are half a million of you, working in a broad range of environments, and not that many of us. Write a post on your own blog, or write a [Medium](https://www.medium.com/) post, or write a [Github Gist](https://gist.github.com/) (add a `.md` file extension for Markdown), or write a [Google doc](https://docs.google.com/), or use any other publishing mechanism you like. After you've posted, please add the post to our new wiki page, [golang.org/wiki/ExperienceReports](https://golang.org/wiki/ExperienceReports).

#### Solutions

![](toward-go2/process34.png)

Now that we know how we're going to identify and explain problems that need to be solved, I want to note briefly that not all problems are best solved by language changes, and that's fine.

One problem we might want to solve is that computers can often compute additional results during basic arithmetic operations, but Go does not provide direct access to those results. In 2013, Robert proposed that we might extend the idea of two-result (“comma-ok”) expressions to basic arithmetic. For example, if x and y are, say, uint32 values, `lo, hi = x * y` would return not only the usual low 32 bits but also the high 32 bits of the product. This problem didn't seem particularly significant, so we [recorded the potential solution](https://golang.org/issue/6815) but didn't implement it. We waited.

More recently, we designed for Go 1.9 a [math/bits package](https://beta.golang.org/doc/go1.9#math-bits) that contains various bit manipulation functions:

package bits // import "math/bits"

func LeadingZeros32(x uint32) int
func Len32(x uint32) int
func OnesCount32(x uint32) int
func Reverse32(x uint32) uint32
func ReverseBytes32(x uint32) uint32
func RotateLeft32(x uint32, k int) uint32
func TrailingZeros32(x uint32) int
...

The package has good Go implementations of each function, but the compilers also substitute special hardware instructions when available. Based on this experience with math/bits, both Robert and I now believe that making the additional arithmetic results available by changing the language is unwise, and that instead we should define appropriate functions in a package like math/bits. Here the best solution is a library change, not a language change.

A different problem we might have wanted to solve, after Go 1.0, was the fact that goroutines and shared memory make it too easy to introduce races into Go programs, causing crashes and other misbehavior in production. The language-based solution would have been to find some way to disallow data races, to make it impossible to write or at least to compile a program with a data race. How to fit that into a language like Go is still an open question in the programming language world. Instead we added a tool to the main distribution and made it trivial to use: that tool, the [race detector](https://blog.golang.org/race-detector), has become an indispensible part of the Go experience. Here the best solution was a runtime and tooling change, not a language change.

There will be language changes as well, of course, but not all problems are best solved in the language.

#### Shipping Go 2

![](toward-go2/process5.png)

Finally, how will we ship and deliver Go 2?

I think the best plan would be to ship the [backwards-compatible parts](https://golang.org/doc/go1compat) of Go 2 incrementally, feature by feature, as part of the Go 1 release sequence. This has a few important properties. First, it keeps the Go 1 releases on the [usual schedule](https://golang.org/wiki/Go-Release-Cycle), to continue the timely bug fixes and improvements that users now depend on. Second, it avoids splitting development effort between Go 1 and Go 2. Third, it avoids divergence between Go 1 and Go 2, to ease everyone's eventual migration. Fourth, it allows us to focus on and deliver one change at a time, which should help maintain quality. Fifth, it will encourage us to design features to be backwards-compatible.

We will need time to discuss and plan before any changes start landing in Go 1 releases, but it seems plausible to me that we might start seeing minor changes about a year from now, for Go 1.12 or so. That also gives us time to land package management support first.

Once all the backwards-compatible work is done, say in Go 1.20, then we can make the backwards-incompatible changes in Go 2.0. If there turn out to be no backwards-incompatible changes, maybe we just declare that Go 1.20 _is_ Go 2.0. Either way, at that point we will transition from working on the Go 1.X release sequence to working on the Go 2.X sequence, perhaps with an extended support window for the final Go 1.X release.

This is all a bit speculative, and the specific release numbers I just mentioned are placeholders for ballpark estimates, but I want to make clear that we're not abandoning Go 1, and that in fact we will bring Go 1 along to the greatest extent possible.

#### Help Wanted

**We need your help.**

The conversation for Go 2 starts today, and it's one that will happen in the open, in public forums like the mailing list and the issue tracker. Please help us at every step along the way.

Today, what we need most is experience reports. Please tell us how Go is working for you, and more importantly not working for you. Write a blog post, include real examples, concrete detail, and real experience. And link it on our [wiki page](https://golang.org/wiki/ExperienceReports). That's how we'll start talking about what we, the Go community, might want to change about Go.

Thank you.


# [Go 2, here we come!](https://blog.golang.org/go2-here-we-come)

Robert Griesemer  
29 November 2018

#### Background

At GopherCon 2017, Russ Cox officially started the thought process on the next big version of Go with his talk [The Future of Go](https://www.youtube.com/watch?v=0Zbh_vmAKvk) ([blog post](https://blog.golang.org/toward-go2)). We have called this future language informally Go 2, even though we understand now that it will arrive in incremental steps rather than with a big bang and a single major release. Still, Go 2 is a useful moniker, if only to have a way to talk about that future language, so let’s keep using it for now.

A major difference between Go 1 and Go 2 is who is going to influence the design and how decisions are made. Go 1 was a small team effort with modest outside influence; Go 2 will be much more community-driven. After almost 10 years of exposure, we have learned a lot about the language and libraries that we didn’t know in the beginning, and that was only possible through feedback from the Go community.

In 2015 we introduced the [proposal process](https://golang.org/s/proposal) to gather a specific kind of feedback: proposals for language and library changes. A committee composed of senior Go team members has been reviewing, categorizing, and deciding on incoming proposals on a regular basis. That has worked pretty well, but as part of that process we have ignored all proposals that are not backward-compatible, simply labeling them Go 2 instead. In 2017 we also stopped making any kind of incremental backward-compatible language changes, however small, in favor of a more comprehensive plan that takes the bigger picture of Go 2 into account.

It is now time to act on the Go 2 proposals, but to do this we first need a plan.

#### Status

At the time of writing, there are around 120 [open issues labeled Go 2 proposal](https://github.com/golang/go/issues?page=1&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2&utf8=%E2%9C%93). Each of them proposes a significant library or language change, often one that does not satisfy the existing [Go 1 compatibility guarantee](https://golang.org/doc/go1compat). Ian Lance Taylor and I have been working through these proposals and categorized them ([Go2Cleanup](https://github.com/golang/go/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2+label%3AGo2Cleanup), [NeedsDecision](https://github.com/golang/go/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3Aproposal+label%3AGo2+label%3ANeedsDecision), etc.) to get an idea of what’s there and to make it easier to proceed with them. We also merged related proposals and closed the ones which seemed clearly out of the scope of Go, or were otherwise unactionable.

Ideas from the remaining proposals will likely influence Go 2’s libraries and languages. Two major themes have emerged early on: support for better error handling, and generics. [Draft designs](https://blog.golang.org/go2draft) for these two areas have been published at this year’s GopherCon, and more exploration is needed.

But what about the rest? We are [constrained](https://blog.golang.org/toward-go2) by the fact that we now have millions of Go programmers and a large body of Go code, and we need to bring it all along, lest we risk a split ecosystem. That means we cannot make many changes, and the changes we are going to make need to be chosen carefully. To make progress, we are implementing a new proposal evaluation process for these significant potential changes.

#### Proposal evaluation process

The purpose of the proposal evaluation process is to collect feedback on a small number of select proposals such that a final decision can be made. The process runs more or less in parallel to a release cycle and consists of the following steps:

1. _Proposal selection_. The Go team selects a small number of [Go 2 proposals](https://github.com/golang/go/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3AGo2+label%3AProposal) that seem worth considering for acceptance, without making a final decision. See below for more on the selection criteria.

2. _Proposal feedback_. The Go team sends out an announcement listing the selected proposals. The announcement explains to the community the tentative intent to move forward with the selected proposals and to collect feedback for each of them. This gives the community a chance to make suggestions and express concerns.

3. _Implementation_. Based on that feedback, the proposals are implemented. The target for these significant language and library changes is to have them ready to submit on day 1 of an upcoming release cycle.

4. _Implementation feedback_. During the development cycle, the Go team and community have a chance to experiment with the new features and collect further feedback.

5. _Launch decision_. At the end of the three month [development cycle](https://github.com/golang/go/wiki/Go-Release-Cycle) (just when starting the three month repo freeze before a release), and based on the experience and feedback gathered during the release cycle, the Go team makes the final decision about whether to ship each change. This provides an opportunity to consider whether the change has delivered the expected benefits or created any unexpected costs. Once shipped, the changes become part of the language and libraries. Excluded proposals may go back to the drawing board or may be declined for good.

With two rounds of feedback, this process is slanted towards declining proposals, which will hopefully prevent feature creep and help with keeping the language small and clean.

We can’t go through this process for each of the open Go 2 proposals, there are simply too many of them. That’s where the selection criteria come into play.

#### Proposal selection criteria

A proposal must at the very least:

1. _address an important issue for many people_,

2. _have minimal impact on everybody else_, and

3. _come with a clear and well-understood solution_.

Requirement 1 ensures that any changes we make help as many Go developers as possible (make their code more robust, easier to write, more likely to be correct, and so on), while requirement 2 ensures we are careful to hurt as few developers as possible, whether by breaking their programs or causing other churn. As a rule of thumb, we should aim to help at least ten times as many developers as we hurt with a given change. Changes that don't affect real Go usage are a net zero benefit put up against a significant implementation cost and should be avoided.

Without requirement 3 we don’t have an implementation of the proposal. For instance, we believe that some form of genericity might solve an important issue for a lot of people, but we don’t yet have a clear and well-understood solution. That’s fine, it just means that the proposal needs to go back to the drawing board before it can be considered.

#### Proposals

We feel that this is a good plan that should serve us well but it is important to understand that this is only a starting point. As the process is used we will discover the ways in which it fails to work well and we will refine it as needed. The critical part is that until we use it in practice we won't know how to improve it.

A safe place to start is with a small number of backward-compatible language proposals. We haven’t done language changes for a long time, so this gets us back into that mode. Also, the changes won’t require us worrying about breaking existing code, and thus they serve as a perfect trial balloon.

With all that said, we propose the following selection of Go 2 proposals for the Go 1.13 release (step 1 in the proposal evaluation process):

1. [_#20706_](https://github.com/golang/go/issues/20706) _General Unicode identifiers based on_ [_Unicode TR31_](http://unicode.org/reports/tr31/): This addresses an important issue for Go programmers using non-Western alphabets and should have little if any impact on anyone else. There are normalization questions which we need to answer and where community feedback will be important, but after that the implementation path is well understood. Note that identifier export rules will not be affected by this.

2. [_#19308_](https://github.com/golang/go/issues/19308), [_#28493_](https://github.com/golang/go/issues/28493) _Binary integer literals and support for_ in number literals_: These are relatively minor changes that seem hugely popular among many programmers. They may not quite reach the threshold of solving an “important issue” (hexadecimal numbers have worked well so far) but they bring Go up to par with most other languages in this respect and relieve a pain point for some programmers. They have minimal impact on others who don’t care about binary integer literals or number formatting, and the implementation is well understood.

3. [_#19113_](https://github.com/golang/go/issues/19113) _Permit signed integers as shift counts_: An estimated 38% of all non-constant shifts require an (artificial) uint conversion (see the issue for a more detailed break-down). This proposal will clean up a lot of code, get shift expressions better in sync with index expressions and the built-in functions cap and len. It will mostly have a positive impact on code. The implementation is well understood.

#### Next steps

With this blog post we have executed the first step and started the second step of the proposal evaluation process. It’s now up to you, the Go community, to provide feedback on the issues listed above.

For each proposal for which we have clear and approving feedback, we will move forward with the implementation (step 3 in the process). Because we want the changes implemented on the first day of the next release cycle (tentatively Feb. 1, 2019) we may start the implementation a bit early this time to leave time for two full months of feedback (Dec. 2018, Jan. 2019).

For the 3-month development cycle (Feb. to May 2019) the chosen features are implemented and available at tip and everybody will have a chance to gather experience with them. This provides another opportunity for feedback (step 4 in the process).

Finally, shortly after the repo freeze (May 1, 2019), the Go team makes the final decision whether to keep the new features for good (and include them in the Go 1 compatibility guarantee), or whether to abandon them (final step in the process).

(Since there is a real chance that a feature may need to be removed just when we freeze the repo, the implementation will need to be such that the feature can be disabled without destabilizing the rest of the system. For language changes that may mean that all feature-related code is guarded by an internal flag.)

This will be the first time that we have followed this process, hence the repo freeze will also be a good moment to reflect on the process and to adjust it if necessary. Let’s see how it goes.

Happy evaluating!












# [Next steps toward Go 2](https://blog.golang.org/go2-next-steps)

Robert Griesemer, for the Go team  
26 June 2019

#### Status

We’re well on the way towards the release of Go 1.13, hopefully in early August of this year. This is the first release that will include concrete changes to the language (rather than just minor adjustments to the spec), after a longer moratorium on any such changes.

To arrive at these language changes, we started out with a small set of viable proposals, selected from the much larger list of [Go 2 proposals](https://github.com/golang/go/issues?utf8=%E2%9C%93&q=is%3Aissue+is%3Aopen+label%3AGo2+label%3AProposal), per the new proposal evaluation process outlined in the “[Go 2, here we come!](https://blog.golang.org/go2-here-we-come)” blog post. We wanted our initial selection of proposals to be relatively minor and mostly uncontroversial, to have a reasonably high chance of having them make it through the process. The proposed changes had to be backward-compatible to be minimally disruptive since [modules](https://blog.golang.org/using-go-modules), which eventually will allow module-specific language version selection, are not the default build mode quite yet. In short, this initial round of changes was more about getting the ball rolling again and gaining experience with the new process, rather than tackling big issues.

Our [original list of proposals](https://blog.golang.org/go2-here-we-come) – [general Unicode identifiers](https://golang.org/issue/20706), [binary integer literals](https://golang.org/issue/19308), [separators for number literals](https://golang.org/issue/28493), [signed integer shift counts](https://golang.org/issue/19113) – got both trimmed and expanded. The general Unicode identifiers didn’t make the cut as we didn’t have a concrete design document in place in time. The proposal for binary integer literals was expanded significantly and led to a comprehensive overhaul and modernization of [Go’s number literal syntax](https://golang.org/design/19308-number-literals). And we added the Go 2 draft design proposal on [error inspection](https://golang.org/design/go2draft-error-inspection), which has been [partially accepted](https://golang.org/issue/29934#issuecomment-489682919).

With these initial changes in place for Go 1.13, it’s now time to look forward to Go 1.14 and determine what we want to tackle next.

#### Proposals for Go 1.14

The goals we have for Go today are the same as in 2007: to [make software development scale](https://blog.golang.org/toward-go2). The three biggest hurdles on this path to improved scalability for Go are package and version management, better error handling support, and generics.

With Go module support getting increasingly stronger, support for package and version management is being addressed. This leaves better error handling support and generics. We have been working on both of these and presented [draft designs](https://golang.org/design/go2draft) at last year’s GopherCon in Denver. Since then we have been iterating those designs. For error handling, we have published a concrete, significantly revised and simplified proposal (see below). For generics, we are making progress, with a talk (“Generics in Go” by Ian Lance Taylor) [coming up](https://www.gophercon.com/agenda/session/49028) at this year’s GopherCon in San Diego, but we have not reached the concrete proposal stage yet.

We also want to continue with smaller improvements to the language. For Go 1.14, we have selected the following proposals:

[#32437](https://golang.org/issue/32437). A built-in Go error check function, “try” ([design doc](https://golang.org/design/32437-try-builtin)).

This is our concrete proposal for improved error handling. While the proposed, fully backwards-compatible language extension is minimal, we expect an outsize impact on error handling code. This proposal has already attracted an enormous amount of comments, and it’s not easy to follow up. We recommend starting with the [initial comment](https://golang.org/issue/32437#issue-452239211) for a quick outline and then to read the detailed design doc. The initial comment contains a couple of links leading to summaries of the feedback so far. Please follow the feedback recommendations (see the “Next steps” section below) before posting.

[#6977](https://golang.org/issue/6977). Allow embedding overlapping interfaces ([design doc](https://golang.org/design/6977-overlapping-interfaces)).

This is an old, backwards-compatible proposal for making interface embedding more tolerant.

[#32479](https://golang.org/issue/32479) Diagnose `string(int)` conversion in `go vet`.

The `string(int)` conversion was introduced early in Go for convenience, but it is confusing to newcomers (`string(10)` is `"\n"` not `"10"`) and not justified anymore now that the conversion is available in the `unicode/utf8` package. Since removing this conversion is not a backwards-compatible change, we propose to start with a `vet` error instead.

[#32466](https://golang.org/issue/32466) Adopt crypto principles ([design doc](https://golang.org/design/cryptography-principles)).

This is a request for feedback on a set of design principles for cryptographic libraries that we would like to adopt. See also the related [proposal to remove SSLv3 support](https://golang.org/issue/32716) from `crypto/tls`.

#### Next steps

We are actively soliciting feedback on all these proposals. We are especially interested in fact-based evidence illustrating why a proposal might not work well in practice, or problematic aspects we might have missed in the design. Convincing examples in support of a proposal are also very helpful. On the other hand, comments containing only personal opinions are less actionable: we can acknowledge them but we can’t address them in any constructive way. Before posting, please take the time to read the detailed design docs and prior feedback or feedback summaries. Especially in long discussions, your concern may have already been raised and discussed in earlier comments.

Unless there are strong reasons to not even proceed into the experimental phase with a given proposal, we are planning to have all these implemented at the start of the [Go 1.14 cycle](https://golang.org/wiki/Go-Release-Cycle) (beginning of August, 2019) so that they can be evaluated in practice. Per the [proposal evaluation process](https://blog.golang.org/go2-here-we-come), the final decision will be made at the end of the development cycle (beginning of November, 2019).

Thank you for helping making Go a better language!