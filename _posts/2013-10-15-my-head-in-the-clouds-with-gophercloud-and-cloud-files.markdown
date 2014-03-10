---
layout: post
title: "My Head In the Clouds with Gophercloud and Cloud Files"
date: 2013-10-15 12:00
comments: true
author: Samuel A. Falvo II
published: true
---

What ever happened to Gophercloud's Cloud Files support?
And, when will Gophercloud finally become multi-vendor?
I plan on completing these two items in that exact order
in the coming months, but
it will take some time.

<!-- more -->

First, let me give the status update with respect to Cloud Files support.
In the working world, you're often pulled in a plurality of directions at once.
As you can imagine, Gophercloud is but one of many projects I'm involved with.
While this doesn't pull me off of Gophercloud full-time,
it does cut into the time I can spend working on it.

For the past month or so, whenever I could sit down and work on Gophercloud,
I spent the time studying OpenStack, Microsoft, Amazon, and Google solutions for cloud file storage,
looking for any commonality between them that I could capitalize on when defining the Gophercloud cloud files API.
Support for OpenStack cloud files comes relatively easy.
Laying the plans to support the equivalent services from Amazon, Microsoft's Azure, and Google's AppEngine all using the same interfaces, however,
proved substantially more daunting than I anticipated.
It's not that these services' particular APIs were all that hard to use.
Indeed, each of these services offers, superficially, the same class of service.
Their respective details differ, however, in such a way as to impede easy abstraction.

To restore the parity of feet and terra firma,
and actually start pushing code again instead of daydream,
I decided to punt on supporting everything from everyone,
and just focus once again on OpenStack.
It's what I know, and,
based on exchanges on the development mailing list and in Github tickets,
it's what those actually using Gophercloud seem to expect the most.
So, expect more pull requests soon.

When, however, will the multi-vendor support become a reality?
This has been a contentious issue in the past, and
Rackspace's SDK support priorities have always been Openstack, Rackspace, and then other providers.
However, my gut tells me that Gophercloud lacks more contributors largely because it hasn't addressed support for more providers yet.
Thus, I feel motivated to address multi-vendor support as soon as I'm happy with its cloud files support.

Once I define the precedent for an API's basic design,
filling in missing functionality typically involves a handful of lines of code for each end-point.
Indeed, if Go supported macros,
they could help trivially implement many of the compute-related functions.
The real excitement lies with making Gophercloud multi-vendor capable,
for in there lies the more interesting problems to solve.

Unless I receive strong support for the staying the course,
I think the time has come that I should invest effort in widening support for multiple vendors.
After all, this has always been part of Gophercloud's promise from day one,
and I'd feel disingenuous for going too long without addressing this feature.
Besides, having existing infrastructure for more platforms might provide the way by which additional contributions,
especially for non-OpenStack platforms, come.
Taken at face value, "Build it and they will come," may not always hold true.
Sometimes, I think, folks just want a paved road to the construction site.

