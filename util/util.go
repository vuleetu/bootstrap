// Package util provides ...
package util

import (
    "os"
    "os/exec"
    "bytes"
    "log"
)


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

func sexec(cd string, args []string) (string, error){
    var out bytes.Buffer
    in := os.Stdin
    cmd := exec.Command(cd, args...)
    cmd.Stdin = in
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    return out.String(),err
}

func Install(p pkg){
    p.Install()
}

func config(cmd string, args []string){
    lexec(cmd, args)
}
