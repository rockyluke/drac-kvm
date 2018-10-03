# drac-kvm

[![License][license-img]][license-url]
[![Build][build-img]][build-url]

## Overview

The integrated Dell Remote Access Controller or DRAC (iDRAC) is an out-of-band
management platform  on certain Dell  servers.  It provides  functionality that
helps you deploy,  update, monitor and maintain Dell PowerEdge  servers with or
without a systems management software agent.

[dell.com](https://www.dell.com/)

A preliminary  implementation of iLO  (Integrated Lights Out) KVM  is available
for version iLO 3 and iLO 4.

[hp.com](https://www.hpe.com/)

Support for Supermicro KVM implementation was added in version 2.0.0.

[supermicro.com](https://www.supermicro.com/)

## Description

A simple CLI launcher for Dell DRAC and HP iLO KVM sessions

This has been tested on the following Dell servers:

* 11th Generation (eg: Dell R710 / iDRAC6)
* 12th Generation (eg: Dell R720 / iDRAC7)
* 13th Generation (eg: Dell R730 / iDRAC8)

This has been tested on the following HP servers:

* 7th Generation (eg: HP DL120 G7)
* 8th Generation (eg: HP DL160 G8)

This has been tested on the following SuperMicro servers:

* PCS S2420Q-M5

## Setup

It requires  that you  have java  installed on  your machine  (specifically the
`javaws` binary).

### Go

If you  already have Go  configured on  your system then  you can just  run the
following to quickly install it:

```bash
go get github.com/rockyluke/drac-kvm
```

### Homebrew

If you already  have Homebrew configured on  your system then you  can just run
the following to quickly install it:

```bash
brew tap rockyluke/devops
brew install drac-kvm
```

## Usage

Fed up of logging into the DRAC web interface just to launch a KVM session?
This simple Go program should help ease the pain.

```bash
drac-kvm --help
Usage of ./drac-kvm:
  -d, --delay int
    	Number of seconds to delay for javaws to start up & read jnlp before deleting it (default 10)
  -h, --host string
    	The DRAC host (or IP)
  -j, --javaws string
    	The path to javaws binary (default "/usr/bin/javaws")
  -k, --keep-jnlp
    	Keep JNLP files and do not clean them after failed start
  -p, --password
    	Prompt for password (optional, will use default vendor if not present)
  -u, --username string
    	The KVM username
  -V, --vendor string
    	The KVM Vendor
  -v, --version int
    	KVM vendor specific version for idrac: (6, 7 or 8) (default -1)
  -w, --wait
    	Wait for java console process end
```

### Example using default dell credentials (root/calvin)

```bash
drac-kvm -h 10.25.1.100
2014/06/26 16:01:11 Detecting iDRAC version...
2014/06/26 16:01:11 Found iDRAC version 7
2014/06/26 16:01:11 Launching DRAC KVM session to 10.25.1.100
```

### Example using custom credentials

```bash
drac-kvm -h 10.25.1.100 -u bob -p
Password: **********
2014/06/26 16:01:11 Detecting iRAC version...
2014/06/26 16:01:11 Found iDRAC version 7
2014/06/26 16:01:11 Launching DRAC KVM session to 10.25.1.100
```

### Configuration file

You can create a configuration file

```bash
cat ~/.drackvmrc
# Override the hardcoded defaults for username and password.
# Useful if your environment has consistent usernames and
# passwords for the KVMs.
[defaults]
javaws_path = /Library/Internet Plug-Ins/JavaAppletPlugin.plugin/Contents/Home/bin/javaws
username = foo
password = bar

[192.168.0.42]
vendor = dell
username = foo
password = bar

[web-1]
vendor = hp
host = 10.33.0.1
username = root
password = password4root

[web-2]
vendor = supermicro
host = 10.33.0.2
username = root
```

## Java Security

Some versions of java are not compatible with older versions of ilo providers. If you receive an error like this

[java-error](https://user-images.githubusercontent.com/67790/46364353-db72b080-c675-11e8-9d7e-3ab0540c51d6.png)

You have to edit `$JAVA_HOME/lib/security/java.security` file and remove following entry `3DES_EDE_CBC` from line starting with `jdk.tls.disabledAlgorithms`. This is global setting and it can affect your system security.

```diff

- jdk.tls.disabledAlgorithms=SSLv3, RC4, MD5withRSA, DH keySize < 1024, EC keySize < 224, DES40_CBC, RC4_40, 3DES_EDE_CBC
+ jdk.tls.disabledAlgorithms=SSLv3, RC4, MD5withRSA, DH keySize < 1024, EC keySize < 224, DES40_CBC, RC4_40
```

## Credits

@jamesdotcuff [blog post](http://blog.jcuff.net/2013/10/fun-with-idrac.html)
@PaulMaddox [initial release](https://github.com/PaulMaddox/drac-kvm)

## Development

Feel free to contribute on GitHub.

## Miscellaneous

```
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

[license-img]: https://img.shields.io/badge/license-ISC-blue.svg
[license-url]: LICENSE
[build-img]: https://travis-ci.org/rockyluke/drac-kvm.svg?branch=master
[build-url]: https://travis-ci.org/rockyluke/drac-kvm
