# drac-kvm ![License badge][license-img]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Frockyluke%2Fdrac-kvm.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Frockyluke%2Fdrac-kvm?ref=badge_shield)

## Overview

The integrated Dell Remote Access Controller  or DRAC (iDRAC) is an out-of-band
management platform  on certain Dell  servers.  It provides  functionality that
helps you deploy,  update, monitor and maintain Dell PowerEdge  servers with or
without a systems management software agent.

[dell.com](https://www.dell.com/)

A preliminary  implementation of iLO  (Integrated Lights Out) KVM  is available
for version iLO 3 and iLO 4.

[hp.com](https://www.hpe.com/)

## Description

A simple CLI launcher for Dell DRAC and HP iLO KVM sessions

This has been tested on the following Dell servers:

 * 11th Generation (eg: Dell R710 / iDRAC6)

 * 12th Generation (eg: Dell R720 / iDRAC7)

 * 13th Generation (eg: Dell R730 / iDRAC8)

This has been tested on the following HP servers:

 * 7th Generation (eg: HP DL120 G7)

 * 8th Generation (eg: HP DL160 G8)

## Setup

It requires  that you  have java  installed on  your machine  (specifically the
`javaws` binary).

### Docker

#### Linux

```bash
docker run \
       -e DISPLAY=${DISPLAY} \
	   -e XAUTHORITY=/tmp/.Xauthority \
	   -v ${HOME}/.Xauthority:/tmp/.Xauthority \
	   -v /tmp/.X11-unix:/tmp/.X11-unix \
	   rockyluke/drac-kvm
```

#### macOS

```bash
brew install socat

DISPLAY_ADDRESS=$(ifconfig|grep 'inet '|grep -v '127.0.0.1'| head -1|awk '{print $2}')

socat TCP-LISTEN:6001,reuseaddr,fork UNIX-CLIENT:\"$DISPLAY\" &
docker run \
       -e DISPLAY=${DISPLAY_ADDRESS} \
	   -e XAUTHORITY=/tmp/.Xauthority \
	   -v ${HOME}/.Xauthority:/tmp/.Xauthority \
	   -v /tmp/.X11-unix:/tmp/.X11-unix \
	   rockyluke/drac-kvm
```

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
Usage of drac-kvm
  -h, --host="some.hostname.com": The DRAC host (or IP)
  -j, --javaws="/usr/bin/javaws": The path to javaws binary
  -p, --password=false: Prompt for password (optional, will use 'calvin' if not present)
  -u, --username="": The DRAC username
  -v, --version=-1: iDRAC version (6, 7 or 8)
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
username = foo
password = bar

[192.168.0.42]
username = foo
password = bar

[web-1]
host = 10.33.0.1
username = root
password = password4root

[web-2]
host = 10.33.0.2
username = root
```

```bash
# Use IP, username and password from drackvmrc for this ip
drac-kvm -h 192.168.0.42
2014/06/26 16:01:11 Detecting iRAC version...
2014/06/26 16:01:11 Found iDRAC version 7
2014/06/26 16:01:11 Launching DRAC KVM session to 192.168.0.42

# Same for host alias web-1
drac-kvm -h web-1
2014/06/26 16:01:11 Detecting iRAC version...
2014/06/26 16:01:11 Found iDRAC version 7
2014/06/26 16:01:11 Launching DRAC KVM session to 10.33.0.1

# Specify -p option in order to do not use defaut password
drac-kvm -h web-2 -p
Password: **********
2014/06/26 16:01:11 Detecting iRAC version...
2014/06/26 16:01:11 Found iDRAC version 7
2014/06/26 16:01:11 Launching DRAC KVM session to 10.33.0.2
```

## Credits

@jamesdotcuff's helpful blog post:

http://blog.jcuff.net/2013/10/fun-with-idrac.html

## Development

Feel free to contribute on GitHub.

```
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

[license-img]: https://img.shields.io/badge/license-Apache%202.0-blue.svg


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Frockyluke%2Fdrac-kvm.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Frockyluke%2Fdrac-kvm?ref=badge_large)