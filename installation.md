---
layout: page
title: Installation
---

## Package installation

To get started you need to ensure that your [GOPATH environment variable](https://golang.org/doc/code.html#GOPATH) is
pointing to an appropriate directory where you want to install gophercloud:

```bash
mkdir $HOME/go
export GOPATH=$HOME/go
```

Once your environment is set up, you can install the gophercloud package like so:

```bash
go get github.com/rackspace/gophercloud
```

This will install all the source files you need into a `pkg` directory, which is
referenceable from your own source files.

## Credentials

Because you'll be hitting an API, you will need to retrieve your OpenStack
credentials and either store them as environment variables or in your local Go
files. The first method is recommended because it decouples credential
information from source code, allowing you to push the latter to your version
control system without any security risk.

You will need to retrieve the following:

* username
* password
* tenant name or tenant ID
* a valid Keystone identity URL

For users that have the OpenStack dashboard installed, there's a quick shortcut.
If you visit the `project/access_and_security` path in your Horizon dashboard
and click on the "Download OpenStack RC File" button at the top right hand
corner, you will get a bash file that exports all of your access details to
environment variables. To execute the file, run `source admin-openrc.sh` and you
will be prompted for your password.
