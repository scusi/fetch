fetch
=====

a simple fetch cmd tool to download files from the web, implemented in go

Basically i missed the 'fetch' tool i knew from FreeBSD on my Mac OX, so i wrote
a very simple fetch replacement in go for my self.

Requirements
============

You need to have a working go environment on your system in order to compile and install this program.
To learn more about go vistit http://golang.org/

Download and Install
====================

<pre>
go get github.com/scusi/fetch
cd fetch
go install
</pre>

Usage
=====

After install you can call it like this:

```fetch http://somehost.tld/somepath/somefile.ext```

The above command will download to _somefile.ext_

In case there is no filename within the url, like in the following example a generic filename is used.

```fetch http://somehost.tld/somepath/```

The above command will download, whatever is coming back as _outfile_, since there is no filename which could be extracted from the URL.

Proxy support
-------------

If you want to use a proxy you should export the environment variable HTTP_PROXY before you call fetch. This is needed only once. After your HTTP_PROXY environment variable was set _fetch_ will regorcnize it and use it.

```
export HTTP_PROXY="http://myproxy.tld:8080/"
fetch http://somehost.tld/somepath/file.ext
```
