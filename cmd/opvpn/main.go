package main

import (
	"net/http"
	ovpn_config "opvpn/internal/config_gen"
	server "opvpn/internal/keys"
	manage_server "opvpn/internal/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type pki struct {
	CA          bool `json:"ca"`
	Server      bool `json:"server"`
	ClientCount int  `json:"clients"`
}

type server_management struct {
	Command string `json:"cmd"`
}

var keys_struct pki
var config_data ovpn_config.ConfJson
var server_mng server_management

func addPKI(c *gin.Context) {
	if err := c.BindJSON(&keys_struct); err != nil {
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	if keys_struct.CA {
		server.CreateCA()
		server.CreateTA()
	}
	if keys_struct.Server {
		server.CreateServer("certs/ca.crt", "certs/ca.key")
	}
	for i := 0; i < keys_struct.ClientCount; i++ {
		server.CreateClient(i+1, "certs/ca.crt", "certs/ca.key")
	}

	c.JSON(http.StatusCreated, gin.H{
		"CA":          "created",
		"Server_cert": "created",
		"Client_keys": "created",
	})
}

func CreateConf(c *gin.Context) {
	if err := c.BindJSON(&config_data); err != nil {
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	config_directives := []string{
		"port ",
		"proto ",
		"dev ",
		"cipher ",
		"mode server",
		"tls-server",
		"client-config-dir ",
		"ca ",
		"cert ",
		"key ",
		"tls-auth ",
		"client-to-client",
		"max-routes-per-client 2048",
		//Client conf directives
		"client",
		"remote ",
		"tls-client",
		"pull",
	}
	config_data.CreateServerConfigAndCcd(config_directives, "confs/server")
	config_data.CreateClientConf(config_directives, "clients", "certs")
	c.JSON(http.StatusCreated, gin.H{"msg": "created"})
}

func ManageServer(c *gin.Context) {
	if err := c.BindJSON(&server_mng); err != nil {
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")

	args := []string{"openvpn", server_mng.Command}

	c.JSON(http.StatusOK, gin.H{"msg": manage_server.ExecCommand("service", args)})
}

func main() {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// Create API route group
	router.POST("/keys", addPKI)
	router.POST("/management", ManageServer)
	router.POST("/conf", CreateConf)

	// Start listening and serving requests
	router.Use(cors.New(config))
	router.Run(":8080")

}
