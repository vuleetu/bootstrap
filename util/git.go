// Package util provides ...
package util

import (
    "log"
)

type GitPkg struct{}

func (git *GitPkg) Install(){
    log.Print("checking git")
    _, e := sexec("which", []string{"git"})
    if e != nil {
        log.Print("git not found, installing git...")
        config("sudo", []string{"apt-get", "install", "git"})
    } else {
        log.Print("git already exists")
    }
    log.Print("backup original ~/.gitconfig file")
    config("cp", []string{"~/.gitconfig", "~/.gitconfig.orig"})
    log.Print("config git")
    config("git",  []string{"config", "--global", "github.user", "vuleetu"})
    config("git",  []string{"config", "--global", "user.name", "fisher"})
    config("git",  []string{"config", "--global", "user.email", "fisher@yun.io"})
    config("git",  []string{"config", "--global", "alias.unstage", "\"reset HEAD\""})
    config("git",  []string{"config", "--global", "core.editor", "vim"})
    config("git",  []string{"config", "--global", "merge.tool", "vimdiff"})
    config("git",  []string{"config", "--global", "mergetool.vimdiff.keepbackup", "false"})
    config("git",  []string{"config", "--global", "mergetool.vimdiff.cmd", "vimdiff", `vim --noplugin "$PWD/$MERGED"         +":split $PWD/$REMOTE" +":set buftype=nowrite"         +":vertical diffsplit $PWD/$LOCAL" +":set buftype=nowrite"         +":vertical diffsplit $PWD/$BASE" +":set buftype=nowrite"         +":wincmd l"         +":nmap 1 :diffget BASE:diffupdate"         +":nmap 2 :diffget LOCAL:diffupdate"         +":nmap 3 :diffget REMOTE:diffupdate" github.user=vuleetu`})
}

func config(cmd string, args []string){
    lexec(cmd, args)
}
