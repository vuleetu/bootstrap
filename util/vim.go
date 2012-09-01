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
        cmd, args := "sudo", []string{"apt-get", "install", "vim"}
        lexec(cmd, args)
    } else {
        log.Print("vim already exists")
    }
    log.Print("backup original .vimrc file")
    cmd2, args2 := "mv", []string{"~/.vimrc", "~/.vimrc.orig"}
    lexec(cmd2, args2)
    log.Print("downloading .vimrc")
    cmd1, args1 := "git",  []string{"clone", "https://github.com/vuleetu/vimrc.git", "~/vimrc/"}
    lexec(cmd1, args1)
    log.Print("link new .vimrc file")
    cmd3, args3 := "ln", []string{"-s", "~/vimrc/.vimrc", "~/.vimrc"}
    lexec(cmd3, args3)
    log.Print("install vim plugin")
    cmd4, args4 := "vim", []string{"+BundleInstall", "+qa"}
    lexec(cmd4, args4)
}
