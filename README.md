# SparkleBase Daemon

**Sparkle** is the data collection daemon for [sparklebase.com](https://sparklebase.com).

It will collect the following data points:
* Architecture and OS version (if available)
* Memory information such as total installed memory, memory in use, swap, etc…
* Uptime information

*Besides mentioned metrics no personal data will be collected!* 

## Installation

### Daemon installation
Requirements: `golang` distribution of your system vendor (Linux, MacOS, *BSD).

```
$ go get github.com/sparklebase/sparkle
```
This will install the daemon in your `$GOPATH`. If you did not change the location it will install the daemon in `~/go/bin/sparkle`


### Receiving Metrics
**Sparkle** will take a host-update token which you can obtain in the SparkleBase Dashboard upon adding a new Host (Add Host button).

*Example:*
```
$ sparkle b7b8bc83-bedc-4e58-88d6-cf484a5d1798
```
Will transmit the data once and exit.

If you wish to send data on a regular basis you can make use of cronjobs:

```
$ crontab -e

*/15 * * * * /home/myself/go/bin/sparkle b7b8bc83-bedc-4e58-88d6-cf484a5d1798
```
Will transmit data every 15 minutes.

### Options
```
$ sparkle --help
usage: sparkle [<flags>] <token>

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --target="https://sparklebase.com/api/update-host"
             Server URL
  --version  Show application version.

Args:
  <token>  Host specific update token
```

