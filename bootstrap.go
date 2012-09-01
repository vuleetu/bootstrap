// Package main provides ...
package main

import (
    "./util"
)

func  main() {
    var git util.GitPkg
    util.Install(&git)
    var zsh util.ZshPkg
    util.Install(&zsh)
    var vim util.VimPkg
    util.Install(&vim)
    var chrome util.ChromePkg
    util.Install(&chrome)
}


