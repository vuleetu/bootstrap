package util

import (
    "log"
)

type ZshPkg struct{}

func (zsh *ZshPkg) Install() {
    log.Print("check zsh")
    _, e := sexec("which", []string{"zsh"})
    if e != nil {
        log.Print("zsh not found, installing zsh...")
        cmd, args := "sudo", []string{"apt-get", "install zsh"}
        lexec(cmd, args)
        log.Print("download oh-my-zsh")
        cmd1, args1 := "git",  []string{"clone", "https://github.com/robbyrussell/oh-my-zsh.git", "~/.oh-my-zsh/"}
        lexec(cmd1, args1)
        log.Print("backuo original .zshrc file")
        cmd2, args2 := "cp", []string{"~/.zshrc", "~/.zshrc.orig"}
        lexec(cmd2, args2)
        log.Print("create new .zshrc file")
        cmd3, args3 := "cp", []string{"~/.oh-my-zsh/templates/zshrc.zsh-template", "~/.zshrc" }
        lexec(cmd3, args3)
        log.Print("change default shell to zsh")
        cmd4, args4 := "chsh", []string{"-s", "/bin/zsh" }
        lexec(cmd4, args4)
        log.Print("zsh install success")
    } else {
        log.Print("zsh already exists")
    }
}
