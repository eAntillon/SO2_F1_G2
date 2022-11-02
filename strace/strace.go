package strace

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	bit "github.com/eAntillon/SO2_F1_G2/Reportes"
)

func Run(name string) {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	colorGreen := "\033[32m"
	for {
		fmt.Println("|----------------------------------------------------|")
		fmt.Print(string(colorGreen), ">> Nombre: ", name, "\n")
		execute()
		fmt.Println(string(colorReset), "")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|                Ejecutar de nuevo?                  |")
		fmt.Println("|  1)  Si                                            |")
		fmt.Println("|  2)  No, regresar                                  |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  Seleccione una opcion:                            |")
		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "1" {
				bit.AgregarActividad("SYSCALL")
				continue
			} else if com.Text() == "2" {
				break
			}
		}
	}

}

func execute() {
	fmt.Println("")
	fmt.Println("################### STRACE SYSTEM ###################")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("|    Ingrese el comando para ejecutar             :  |")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("")
	fmt.Print(">> ")
	com := bufio.NewScanner(os.Stdin)
	if com.Scan() {
		if com.Text() == "exit" {
			return
		} else {
			strace(strings.Fields(com.Text()))
		}
	}
}

func strace(command []string) {
	var regs syscall.PtraceRegs
	var ss syscallCounter

	ss = ss.init()

	cmd := exec.Command(command[0], command[1])
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			ss.inc(regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}

	ss.print()
}
