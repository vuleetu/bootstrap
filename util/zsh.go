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
        config("sudo", []string{"apt-get", "install zsh"})
        log.Print("zsh install success")
    } else {
        log.Print("zsh already exists")
    }
    log.Print("download oh-my-zsh")
    config("git",  []string{"clone", "https://github.com/robbyrussell/oh-my-zsh.git", "~/.oh-my-zsh/"})
    log.Print("backup original .zshrc file")
    config("cp", []string{"~/.zshrc", "~/.zshrc.orig"})
    log.Print("create new .zshrc file")
    config("cp", []string{"~/.oh-my-zsh/templates/zshrc.zsh-template", "~/.zshrc" })
    log.Print("change default shell to zsh")
    config("chsh", []string{"-s", "/bin/zsh" })
}
