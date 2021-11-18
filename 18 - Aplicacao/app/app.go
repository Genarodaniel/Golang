package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a aplicacao de linha de comando
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicacao de CLI"
	app.Usage = "Buscar Ip e Name Servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "amazon.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},

		{
			Name:   "ns",
			Usage:  "Busca servidores NS de um endereço",
			Flags:  flags,
			Action: buscarNS,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {

	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarNS(c *cli.Context) {
	host := c.String("host")

	nameServers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range nameServers {
		fmt.Println(server)
	}
}
