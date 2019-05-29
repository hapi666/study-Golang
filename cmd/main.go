package main

import (
	"log"
	"os/exec"
)

func main(){
	//tick:= time.Tick(1*time.Second)
	//for {
	//	select {
	//	case <-tick:
	//		command := `./operation.sh .`
	//		cmd := exec.Command("/bin/bash", "-c", command)
	//		output, err := cmd.Output()
	//		if err != nil {
	//			log.Fatalf("Execute Shell:%s failed with error:%s", command, err.Error())
	//		}
	//		log.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	//	}
	//
	//}
	command := `./operation.sh .`
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Execute Shell:%s failed with error:%s", command, err.Error())
	}
	log.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

}