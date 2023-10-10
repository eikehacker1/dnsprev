package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		if len(os.Args) != 2 {
			fmt.Println("Uso: go run main.go [IP ou Domínio]")
			return
		}

		arg := os.Args[1]

		ipAddr := net.ParseIP(arg)
		if ipAddr != nil {
			names, err := net.LookupAddr(arg)
			if err != nil {
				fmt.Println("Erro ao fazer a pesquisa reversa de DNS:", err)
				return
			}
			for _, name := range names {
				fmt.Println(name)
			}
		} else {
			ips, err := net.LookupIP(arg)
			if err != nil {
				fmt.Println("Erro ao fazer a pesquisa de DNS:", err)
				return
			}
			for _, ip := range ips {
				fmt.Println(ip.String())
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			arg := scanner.Text()

			ipAddr := net.ParseIP(arg)
			if ipAddr != nil {
				names, err := net.LookupAddr(arg)
				if err != nil {
					fmt.Println("Erro ao fazer a pesquisa reversa de DNS:", err)
					return
				}
				for _, name := range names {
					fmt.Println(name)
				}
			} else {
				ips, err := net.LookupIP(arg)
				if err != nil {
					fmt.Println("Erro ao fazer a pesquisa de DNS:", err)
					return
				}
				for _, ip := range ips {
					fmt.Println(ip.String())
				}
			}
		}
	}
}
