### just drawing which c-elegans neurons light up in different circumstances

In attempts to understand c-elegans robot better, I wanted to get more hands-on experience in simulating how the connectome behaves under certain pokes and pings

Data comes from from https://github.com/openworm/data-viz/tree/master/HivePlots

![screenshot](https://raw.githubusercontent.com/jackdoe/worm/master/out.gif)

the object moves without any programming about how it should move, it just has hooks to muscle neurons as if it owns left/right wheels, but the network lights up on its own

### start it

```
$ go run *.go
$ curl http://localhost:8080/ping?ASGR # give it food, and watch it go
$ curl http://localhost:8080/ | dot -x -Tsvg > out.svg # will take a while
```

![screenshot](https://raw.githubusercontent.com/jackdoe/worm/master/out.png)

when you start it it spins up and starts doing things forever, due to the nature of the connections

```
2016/02/14 23:12:05 move left: -4
2016/02/14 23:12:05 move right: -24
2016/02/14 23:12:05 move left: -15
2016/02/14 23:12:05 move right: -1
2016/02/14 23:12:06 move left: -16
2016/02/14 23:12:06 move right: 3
2016/02/14 23:12:06 move left: 31
2016/02/14 23:12:06 move right: -13
2016/02/14 23:12:06 move left: -12
```

use `$ curl http://localhost:8080/debug` to toggle debug print

```
2016/02/14 21:20:21 RIBR activated, nReceived: 31
2016/02/14 21:20:21 VD12 activated, nReceived: 9
2016/02/14 21:20:21 VA5 activated, nReceived: 20
2016/02/14 21:20:21 AVL activated, nReceived: 26
2016/02/14 21:20:21 DD2 activated, nReceived: 95
2016/02/14 21:20:21 RMFL activated, nReceived: 11
2016/02/14 21:20:21 AS6 activated, nReceived: 10
2016/02/14 21:20:21 DA5 activated, nReceived: 25
2016/02/14 21:20:21 VD2 activated, nReceived: 76
2016/02/14 21:20:21 AS9 activated, nReceived: 20
2016/02/14 21:20:21 VB8 activated, nReceived: 28
2016/02/14 21:20:21 AVAR activated, nReceived: 233
```

### world

for now just spawn the c.elegans as a player in simple terminal 2d world, started by `go run *.go -game` and then show it food with `curl http://localhost:8080/ping?ASGR`

### muscle

i used map from neuron to fake body parts from https://github.com/Connectome/GoPiGo/blob/master/connectome.py

```
def AVFL():
        ...
        postSynaptic['HSNL'][nextState] += 1
        postSynaptic['MVL11'][nextState] += 1
        postSynaptic['MVL12'][nextState] += 1
```

i just map to:

```
AVFL;LEFT;FAKE;1;FAKE
AVFL;LEFT;FAKE;1;FAKE
```

and on i link neuron `LEFT` to `body.left` muscle, so `body.left.delta()` shows how much it wants to move left

### warning: this is not the correct way to do this, but I have to start somewhere :)

## links 

* [C. Elegans Connectome Research](http://www.connectomeengine.com/Home/CElegans)
* [PyOpenWorm](https://github.com/openworm/PyOpenWorm)
* [OpenWorm](https://github.com/openworm/)
* [OpenWorm.org](http://www.openworm.org/)
* [GENERAL NATURE OF THE GENETIC CODE FOR PROTEINS](https://profiles.nlm.nih.gov/ps/access/SCBCBJ.pdf)
* [C. Elegans Hermaphrodite sensory receptors](http://wormatlas.org/hermaphrodite/nervous/Images/neurotable1leg.htm)
* [wormbook mechanosensation](http://www.wormbook.org/chapters/www_mechanosensation/mechanosensation.html)
* [C. Elegans GoPiGo](https://github.com/Connectome/GoPiGo)