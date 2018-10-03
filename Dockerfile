# Copyright (c) 2017, rockyluke
#
# Permission  to use,  copy, modify,  and/or  distribute this  software for  any
# purpose  with  or without  fee  is hereby  granted,  provided  that the  above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS"  AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
# REGARD TO  THIS SOFTWARE INCLUDING  ALL IMPLIED WARRANTIES  OF MERCHANTABILITY
# AND FITNESS.  IN NO EVENT SHALL  THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
# INDIRECT, OR  CONSEQUENTIAL DAMAGES OR  ANY DAMAGES WHATSOEVER  RESULTING FROM
# LOSS OF USE, DATA OR PROFITS,  WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
# OTHER  TORTIOUS ACTION,  ARISING  OUT OF  OR  IN CONNECTION  WITH  THE USE  OR
# PERFORMANCE OF THIS SOFTWARE.

FROM golang:1.10.4-stretch as builder

COPY . /go/src/github.com/rockyluke/drac-kvm
WORKDIR /go/src/github.com/rockyluke/drac-kvm

RUN go build

FROM golang:1.10.4-stretch

ENV DEBIAN_FRONTEND="noninteractive" \
    TZ="Europe/Amsterdam"

RUN apt-get update  -qq && \
    apt-get upgrade -qq -y && \
    apt-get install -qq -y \
      icedtea-netx \
      libx11-6 \
      x11-utils

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=builder /go/src/github.com/rockyluke/drac-kvm/drac-kvm /usr/bin/drac-kvm

ENTRYPOINT ["drac-kvm"]

CMD ["--help"]
# EOF
