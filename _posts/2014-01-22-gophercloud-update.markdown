---
layout: post
title: "Gophercloud Update"
date: 2014-01-22 13:00
comments: true
author: Samuel A. Falvo II
published: true
---

I'm happy to report that Gophercloud has received some contributor interest recently.
In particular, Gophercloud now supports Rackspace API Keys for authentication.
The publicly-visible interface for using API keys follows the conventions set previously by the `AuthOptions` structure,
which means the *interface* should work with any provider that supports API keys.
In terms of *implementation*, it only works with Rackspace for now.
We welcome contributions, even if only as a feature-request, from others who'd like to see this work with other providers.

I also received several contributions for its Cloud Files support over the past several days.
While the Cloud Files API remains in flux, for it's still very new, we broke ground and we're hammering the details out now.
We won't merge to master until we have an API we're happy with;
however, folks interested in exercising the new API can do so by changing to the `cloud-files` topic branch.
Interested in helping out?  Join us today and leave your mark!

What's coming down the pike?
Version 0.1.0 will introduce support for arbitrary fixed and floating IP pool names, as well as support for security groups.
As I write this, we're closing in at 50% complete for v0.1.0.
Looking forward, it looks like version 0.1.1 will support a more complete images API.

Perhaps most significantly, we're looking at refreshing the API completely for v0.2.0.
We're finding the current API design just isn't scalable along the axes we'd like.
First, it's taking longer and longer to run tests, and we'd like to reduce the time taken.
Next, it's nearly impossible for every user to run every test against their provider of choice,
as few providers support all available extensions to OpenStack.
Next, running the current set of tests requires you have an account through a real OpenStack provider, which costs you money.
I feel it shouldn't have to cost the user real money to test a package.
Finally, adding support for OpenStack extensions yields messy code.
If you remember the good ol' days with Windows 3.1 and its heavily overloaded `WINDOWS` directory,
that's not entirely dissimilar to what's happening with Gophercloud's package layout, data structures, and interfaces in versions 0.1.x.
The v0.2.0 API will address these issues thanks to
a new package hierarchy layout,
a complete refactoring of the request/response phases,
and extensive use of the virtually forgotten Carrier-Rider design pattern, which happens to apply nicely for our needs.

As usual, if anyone wishes to contribute to the project, we'd love to see you on our [developers group](https://groups.google.com/forum/#!forum/gophercloud-dev).
For those who have already contributed, I thank you deeply!

