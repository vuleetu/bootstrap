// Package main provides ...
package main

import (
    "log"
    "os"
    "os/exec"
    "bytes"
)

func  main() {
   defer (func(){
        recover()
        return
   })()
   log.SetPrefix("[config]") 
   zsh()
}

func zsh() {
    log.Print("check zsh")
    s, e := sexec("which", []string{"zsh"})
    if e != nil {
        log.Print("get zsh")
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
    }
}

func lexec(cd string, args []string){
    in := os.Stdin
    out := os.Stdout
    e := os.Stderr
    cmd := exec.Command(cd, args...)
    cmd.Stdin = in
    cmd.Stdout = out 
    cmd.Stderr = e
    err := cmd.Run()
    if err != nil {
        log.Print(err.Error())
        panic("something wrong happend") 
    }   
} 

func sexec(cd string, args []string) string, error{
    var out bytes.Buffer
    in := os.Stdin
    cmd := exec.Command(cd, args...)
    cmd.Stdin = in
    cmd.Stdout = out 
    cmd.Stderr = out
    err := cmd.Run()
    return out.String(),err
} 
