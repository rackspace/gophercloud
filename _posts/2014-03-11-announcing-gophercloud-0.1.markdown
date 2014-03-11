---
layout: post
title: "Announcing Gophercloud 0.1"
date: 2014-03-11 12:00
author: Samuel A. Falvo II
published: true
---

I'm pleased to announce that the final PR for Gophercloud 0.1 has successfully merged to master!  Gophercloud 0.1 is out!

Gophercloud 0.1 is an incremental and bug-fix release that also adds support for the following Nova extensions:

* Floating IPs
* Security Groups
* Security Group Default Rules

In addition, 0.1 supports OpenStack deployments that do not expose `Public` and `Private` IP pools.  Now, pools may possess any name configured by the OpenStack administrator.

As an added bonus to software developers, Gophercloud's `osutil` sub-package provides a more convenient method of building an `AuthOptions` structure from a standard set of environment variables.  That's less code you have to write for every application.

Developers can grab Gophercloud on our [Github page.](https://github.com/rackspace/gophercloud)  If you're interested in contributing to our upcoming 0.2.0 API, feel free to get in contact with us!  You can learn more by [visiting the our community page.](http://gophercloud.io/community.html)

