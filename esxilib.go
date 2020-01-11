package main

import (
        "log"
        //"strings"
        "os/exec"
)


func cmd(hostip string,thecmd string) string {
    cmdline := "ssh root@" + hostip + " \"" + thecmd + "\"" + ";exit"
    log.Printf("Cmd = %s",cmdline)
    cmd := exec.Command("bash")
    cmdWriter, _ := cmd.StdinPipe()
    cmdWriter.Write([]byte(cmdline + "\n"))
    cmdWriter.Write([]byte("exit"    + "\n"))
    cmd.Wait()
//    if err != nil {
//        log.Fatalf("cmd.Run() failed with %s\n", err)
//    }
    out, _ := cmd.CombinedOutput()
    return(string(out))


}
func getvms(hostip string) {
    result := cmd(hostip,"vim-cmd vmsvc/getallvms")
    lines := result.Split("\n")
    lines = lines[1:]

    log.Printf(result)
 


}

func main(){

    getvms("192.168.1.150")

}

