#image2asiic
[![Gobuild Download](http://gobuild.io/badge/github.com/Alienero/image2asiic/downloads.svg)](http://gobuild.io/github.com/Alienero/image2asiic)                    
A image to asiic tool written in golang

## Install
You can use gobuild to download。

##Usage
```sh
./image2asiic -path=[image path] -q=[quality] -to=[out file path]
```
if `-to` is not set, image2asiic will print to stdout.        
In `-q` 1 is the highest level quality.

##Example
`-q=3`           
![image2asiic](https://raw.githubusercontent.com/Alienero/image2asiic/master/example/ex.png "image2asiic")
`-q=1`
![image2asiic2](https://raw.githubusercontent.com/Alienero/image2asiic/master/example/ex2.png "image2asiic2")
