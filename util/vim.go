// Package util provides ...
package util

import (
    "log"
)


type VimPkg struct{}


func (vim *VimPkg) Install(){
    log.Print("checking vim")
    _, e := sexec("which", []string{"zsh"})
    if e != nil {
        log.Print("vim not found, installing vim...")
        config("sudo", []string{"apt-get", "install", "vim"})
    } else {
        log.Print("vim already exists")
    }
    log.Print("backup original .vimrc file")
    config("mv", []string{"~/.vimrc", "~/.vimrc.orig"})
    log.Print("downloading .vimrc")
    config("git",  []string{"clone", "https://github.com/vuleetu/vimrc.git", "~/vimrc/"})
    log.Print("link new .vimrc file")
    config("ln", []string{"-s", "~/vimrc/.vimrc", "~/.vimrc"})
    log.Print("install vim plugin")
    config("vim", []string{"+BundleInstall", "+qa"})
}
