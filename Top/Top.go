package Top

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func Run(name string) {
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	colorBlue := "\033[34m"
	for {
		fmt.Println("|----------------------------------------------------|")
		fmt.Print(string(colorBlue),">> Nombre: ", name, "\n")
		execute()
		fmt.Println(string(colorReset),"")
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
				continue
			} else if com.Text() == "2" {
				break
			}
		}
	}

}

func execute() {
	colorBlue := "\033[34m"
	cpu := "/proc/cpu_top"
	bytesLeidos, err2 := ioutil.ReadFile(cpu)
	if err2 != nil {
		log.Println(err2)
		return
	}
	leido := string(bytesLeidos[:])
	result := leido[:len(leido)-3] + "]"
	fmt.Println(string(colorBlue),string(result))

}
