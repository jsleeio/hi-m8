# hi-m8

## what is this?

Just a tiny webserver, intended for testing web plumbing. Always listens on
TCP/3000.

## how?

```
$ ./hi-m8 -message="g'day github" &
[1] 22846

$ curl -s 0:3000 | od -c
0000000    g   '   d   a   y       g   i   t   h   u   b
0000014

$ kill %1
[1]+  Terminated: 15          ./hi-m8 -message="g'day github"
```

## more?

There actually is a bit more, but not much.

```
$ ./hi-m8 --help
Usage of ./hi-m8:
  -listen string
    	[address]:port to bind to (default ":3000")
  -message string
    	specify a message to be returned to clients (default "hi m8")
  -path string
    	specify a path to respond to (default "/")
```
