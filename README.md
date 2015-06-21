# Hubstatus

A small nifty tool to check the status of github.com and all its related 
services. If you are experiencing problems connecting to GitHub, run Hubstatus 
to check to the current system status.

# Installation

Installation requires the [go toolchain](https://golang.org/dl/).

> $ go get github.com/umahmood/hubstatus <br>
> $ cd $GOPATH/src/github.com/umahmood/hubstatus <br>
> $ go install

# Usage

List the current system status:

> $ hubstatus <br>
> Status: good <br>
> Last updated: 2015-06-20 23:15:07 +0000 UTC

Get the last human communication:

> $ hubstatus -last-human <br>
> Status: good <br>
> Message: Everything operating normally. <br>
> Created on: 2015-05-06 12:20:14 +0000 UTC

Get the most recent human communications:

> $ hubstatus -recent-human <br>
> Status: good <br>
> Message: Battlestation fully operational <br>
> Created_on: 2012-12-07T18:11:55Z
> 
> Status: minor <br>
> Message: Almost done reticulating splines <br>
> Created on: 2012-12-05T12:08:33Z

Output responses in JSON format:

> $ hubstatus -j <br>
> {"status":"good","last_updated":"2015-06-20T23:47:22Z"}

For help run:

> $ hubstatus -h <br>
> Usage of hubstatus: <br>
> -j=false: Output responses as JSON. <br>
> -last-human=false: Returns the last human communication, status, and timestamp. <br>
> -recent-human=false: Returns the most recent human communications with status and timestamp.

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
