// Package util provides ...
package util

import (
    "log"
)

type ChromePkg struct{}

func (chrome *ChromePkg) Install(){
    log.Print("checking chromium-browser")
    _, e := sexec("which", []string{"chromium-browser"})
    if e != nil {
        log.Print("chromium-browser not found, installing chromium-browser...")
        config("sudo", []string{"apt-get", "install", "chromium-browser"})
    } else {
        log.Print("chromium-browser already exists")
    }
}

