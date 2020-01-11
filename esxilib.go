package main

import (
        "log"
        "strings"
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
    m := make(map[string]string)
    result := cmd(hostip,"vim-cmd vmsvc/getallvms")
    log.Printf(result)
    lines := strings.Split(result,"\n")
    lines = lines[1:]
    for _, theline := range lines {
        log.Printf("Line = %s",theline)
        v := strings.Fields(theline)
        if (len(v) > 0){
           m[v[1]] = v[0]
           }
        }
    log.Printf("%v",m)
 


}

func main(){

    getvms("192.168.1.150")

}

