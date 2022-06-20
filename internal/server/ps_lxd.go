package server

import (
	"fmt"
	"io/ioutil"

	lxd "github.com/lxc/lxd/client"
)

func (s *GeneralServer) genlxd() {
	cert, _ := ioutil.ReadFile("/home/twitchgg/.config/lxc/client.crt")
	key, _ := ioutil.ReadFile("/home/twitchgg/.config/lxc/client.key")
	c, err := lxd.ConnectLXD("https://10.10.10.197:8443/", &lxd.ConnectionArgs{
		InsecureSkipVerify: true,
		TLSClientCert:      string(cert),
		TLSClientKey:       string(key),
	})
	if err != nil {
		panic(err)
	}
	cs, err := c.GetContainers()
	if err != nil {
		panic(err)
	}
	for _, c := range cs {
		fmt.Println("container:", c.Name, "", c.Status, c.IsActive())
	}
}
