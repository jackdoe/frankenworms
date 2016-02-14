### just drawing which c-elegans neurons light up in different circumstances

In attempts to understand c-elegans robot better, I wanted to get more hands-on experience in simulating how the connectome behaves under certain pokes and pings

Data comes from from https://github.com/openworm/data-viz/tree/master/HivePlots

```
$ go run *.go
$ while :; do curl http://localhost:8080/ping?DD2; done
$ curl http://localhost:8080/ | dot -x -Tpng > out.png
```

![screenshot](https://raw.githubusercontent.com/jackdoe/worm/master/out.png)

### warning: in no way this is the correct way to do this, but I have to start somewhere :)
