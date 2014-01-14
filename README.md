glogbridge
==========
[![Build Status](https://travis-ci.org/kevinwallace/glogbridge.png?branch=master)](https://travis-ci.org/kevinwallace/glogbridge)

Package `glogbridge` redirects output from the built-in [log](http://golang.org/pkg/log/) package to the [glog](https://github.com/golang/glog) logging library.

Usage
-----

Just import it:

    import _ "github.com/kevinwallace/glogbridge"

Now all those log lines from third-party libraries will show up in your glog logs at `WARNING` level.
