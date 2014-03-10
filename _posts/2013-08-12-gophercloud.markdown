---
layout: post
title: "Gophercloud: A multi-cloud software development kit for Go"
date: 2013-08-13 08:00
comments: true
author: Samuel A. Falvo II
published: true
---

<img src="http://developer.rackspace.com/images/2013-08-12-gophercloud/gophercloud.png" alt="Gophercloud Logo" />

[Go][1]. It’s more than the opposite of stop and unless you’ve been living under a rock you’ve probably heard about this exciting modern language created out of Google by Rob Pike and Ken Thompson, both important names in the Unix community.

Go is a bit of a bridge between the performance of static-typed languages and the usability and ease of programming associated with dynamic languages. If you’re unfamiliar with what this means, it’s kind of like getting the good stuff from C and the good stuff from Python.<!--More-->

Go, as with Python, favors the "[There should be one-- and preferably only one --obvious way to do it.][2]" philosophy, which naturally leads to a well-factored standard library, testing, and code formatting tools. While Go follows in the footsteps of C/C++ syntactically, it implements some sugar which makes writing Go code almost as easy and compact as writing Python. This helps maintain code legibility without sacrificing the desirable properties of static typing. Everybody’s happy.

Go also shares some attributes with Erlang; in particular, its “[goroutine][3]” construct makes writing concurrent software much easier than managing threads by hand, even when compared to languages which offer built-in support for threading, such as Java. It’s those attributes that make it attractive for running massive processes in parallel.

Organizations of all sizes look to Go to improve their infrastructure in some way. Google, for example, successfully applied Go to both their central download server (dl.google.com) and to various subsystems in YouTube. [Docker.io][4] uses it for implementing their virtual machine container software. Finally, Rackspace's own [Exceptional.io][5] has been using it for about a year now to help implement [Airbrake][6].

For some users, Go offers the enhanced reliability that static typing promises. It’s like wrapping your code in Kevlar. For others, it offers cost-savings in the data center, as Go's statically compiled binaries eliminate the overhead incurred through the interpretation of more dynamic languages.  Small, highly efficient, single binary deployments mean more efficient utilization of existing infrastructure.

Released in 2009, Go has proven its worth and is here to stay.  Its
popularity grows daily. That's why Rackspace is investing in it by
creating a multi-cloud (supporting [OpenStack][8] first, [Rackspace Cloud][9] second) software development kit (SDK) for the Go programming community
called [Gophercloud][7] (Mascot: gopher, check! Cloud? Check!)!

Rackspace firmly believes in a tool and API ecosystem where developers in any language aren’t locked-in on the API, or code level. This is why we actively contribute to packages such as [Fog][10] (Ruby), [jclouds][11] (Java), [pkgcloud][12] (node.js) and [libcloud][13] adding OpenStack support first, and then layering on Rackspace Cloud specific extensions.

[OpenStack][8] is the open source, community driven “cloud infrastructure” project that Rackspace helped found, and many other companies and individuals contribute to on a daily basis. Our investment in it is a firm belief that lock-in to proprietary systems and APIs on any level is bad for developers and actively harmful. However, we realize that it is not the only “cloud” out there – so when we look at tooling such as [Gophercloud][7], and other SDKs, devops tools, etc – we take a “what is good for the community” first approach. 

This means that supporting and making tools and SDKs – such as [Gophercloud][7] – that are designed to support many cloud hosts and help you, as an application developer not get locked into specific APIs on a code level is paramount to our ideals and goals.

Hence, Gophercloud – this SDK aims to package OpenStack, Rackspace Cloud and other Cloud provider APIs in a form more easily consumable by Go developers. Gophercloud uses Go's unique support for structurally-typed interfaces to help isolate applications from any one specific OpenStack provider, and works hard to adhere to established Go-community standards.


[Gophercloud][7] is still very new – it’s 100 percent open source, community focused. It’s under a liberal Apache 2 license (just like OpenStack) and we welcome contribution, feedback, thoughts, ideas – everything from the developer community. We want to make an SDK that fits the needs of Go users everywhere. Your patches, feedback, bug reports, everything: we want to hear from you!

If you're interested in getting started with or contributing to [Gophercloud][7], check out our GitHub repository at [https://github.com/rackspace/gophercloud][7]. You can also email the [gophercloud][14] mailing list.

We'd love to hear from you and welcome your contributions!

[1]: http://golang.org/
[2]: http://www.python.org/dev/peps/pep-0020/
[3]: https://gobyexample.com/goroutines
[4]: http://www.docker.io/
[5]: http://www.exceptional.io/
[6]: https://airbrake.io/pages/home
[7]: https://github.com/rackspace/gophercloud
[8]: http://www.openstack.org/
[9]: http://www.rackspace.com/cloud/
[10]: http://fog.io/
[11]: http://jclouds.incubator.apache.org/
[12]: https://github.com/nodejitsu/pkgcloud
[13]: http://libcloud.apache.org/
[14]: https://groups.google.com/forum/#!forum/gophercloud-dev
