### just drawing which c-elegans neurons light up in different circumstances

In attempts to understand c-elegans robot better, I wanted to get more hands-on experience in simulating how the connectome behaves under certain pokes and pings

Data comes from from https://github.com/openworm/data-viz/tree/master/HivePlots

```
$ go run *.go
$ while :; do curl http://localhost:8080/ping?DD2; done
$ curl http://localhost:8080/ | dot -x -Tpng > out.png
```

![screenshot](https://raw.githubusercontent.com/jackdoe/worm/master/out.png)

when you start it it spins up and starts doing things forever, due to the nature of the connections

```
2016/02/14 16:58:33 AVER activated, nReceived: 110
2016/02/14 16:58:33 DA3 activated, nReceived: 28
2016/02/14 16:58:33 DD5 activated, nReceived: 47
2016/02/14 16:58:33 AIML activated, nReceived: 6
2016/02/14 16:58:33 AS10 activated, nReceived: 6
2016/02/14 16:58:33 AVJL activated, nReceived: 33
2016/02/14 16:58:33 SIADL activated, nReceived: 5
2016/02/14 16:58:33 URAVL activated, nReceived: 5
2016/02/14 16:58:33 PVNL activated, nReceived: 12
2016/02/14 16:58:33 VD4 activated, nReceived: 57
2016/02/14 16:58:33 ADLL activated, nReceived: 2
2016/02/14 16:58:33 SIADR activated, nReceived: 3
...
```

use `$ curl http://localhost:8080/debug` to toggle debug print

### warning: this is not the correct way to do this, but I have to start somewhere :)

## links 

* [C. Elegans Connectome Research](http://www.connectomeengine.com/Home/CElegans)
* [PyOpenWorm](https://github.com/openworm/PyOpenWorm)
* [OpenWorm](https://github.com/openworm/)
* [OpenWorm.org](http://www.openworm.org/)
* [GENERAL NATURE OF THE GENETIC CODE FOR PROTEINS](https://profiles.nlm.nih.gov/ps/access/SCBCBJ.pdf)
* [C. Elegans Hermaphrodite sensory receptors](http://wormatlas.org/hermaphrodite/nervous/Images/neurotable1leg.htm)
* [wormbook mechanosensation](http://www.wormbook.org/chapters/www_mechanosensation/mechanosensation.html)
