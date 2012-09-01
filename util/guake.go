// Package util provides ...
package util

import (
    "log"
)

type GuakePkg struct{}

func (guake *GuakePkg) Install(){
    log.Print("checking guake")
    _, e := sexec("which", []string{"guake"})
    if e != nil {
        log.Print("guake not found, installing guake...")
        config("sudo", []string{"apt-get", "install", "guake"})
    } else {
        log.Print("guake already exists")
    }
}


