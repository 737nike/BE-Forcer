package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Is Fortnite in lobby?")
	fmt.Println("(y/N): ")
	var answer string
	fmt.Scan(&answer)
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)
	if answer == "y" || answer == "yes" {
		fmt.Println("Starting BE Service.")
		startbesvc := exec.Command("cmd", "/C", "net start BEService")
		startbesvc.Run()
		time.Sleep(1 * time.Second)
		fmt.Printf("Killing EAC Service.")
		killeacsvc := exec.Command("cmd", "/C", "net stop EasyAntiCheat")
		killeacsvc.Run()
		time.Sleep(1 * time.Second)
		fmt.Println("Killing EAC Launcher.")
		killeac := exec.Command("cmd", "/C", "taskkill /IM \"EasyAntiCheat Launcher\" /F")
		killeac.Run()
		time.Sleep(3 * time.Second)
		fmt.Println("Starting BE Launcher")
		//Using syscall to spawn a parrent process
		var sI syscall.StartupInfo
		var pI syscall.ProcessInformation
		argv := syscall.StringToUTF16Ptr("C:\\Program Files\\Epic Games\\Fortnite\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping_BE.exe")
		err := syscall.CreateProcess(
			nil,
			argv,
			nil,
			nil,
			true,
			0,
			nil,
			nil,
			&sI,
			&pI)
		if err != nil {
		}
		time.Sleep(3 * time.Second)
		fmt.Println("BE Forced!")
		os.Exit(0)
	} else {
		main()
	}
}
