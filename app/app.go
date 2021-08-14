package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a applicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de Servidor na Internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIP,
		},
		{
			Name:   "servidor",
			Usage:  "Busca do Servidor de endereço na internet",
			Flags:  flags,
			Action: buscarServidor,
		},
	}
	return app
}

func buscarIP(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("IPs do host", host)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Servidor do host", host)
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
