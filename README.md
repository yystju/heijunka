# heijunka
Algorithm of Heijunka in golang...

## Code structure

1. The code structure is aimed for developing under the `Theia-go` docker environment.

    So if you wanna use it as a libray, you'd better change the structure. 
    
    For now, I will develop it on [Theia IDE](https://theia-ide.org/) to see if it's OK to be used in real business development.

2. The code is assumed to be directly as $GOPATH. 

    The main.go is the entry with `package main`.

    The heijunka package is placed under src/heijunka folder.

3. dependencies

    go get github.com/BurntSushi/toml

