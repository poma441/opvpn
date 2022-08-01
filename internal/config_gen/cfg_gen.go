package ovpn_config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type ConfJson struct {
	ServerIP    string `json:"serverIP"`
	Lvl         string `json:"tunnel_lvl"`
	Port        string `json:"port"`
	Proto       string `json:"proto"`
	AdapterName string `json:"dev"`
	Cipher      string `json:"cipher"`
	AddrPool    string `json:"ifconfig-pool"`
	Netmask     string `json:"netmask"`
	Route       string `json:"push"`
}

func (config_data *ConfJson) CreateServerConfigAndCcd(config_directives []string, path_to_srv_dir string) {
	_, err := os.Stat(path_to_srv_dir + "/ccd")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path_to_srv_dir+"/ccd", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	text := [15]string{
		config_directives[0] + config_data.Port,
		config_directives[1] + config_data.Proto,
		config_directives[2] + config_data.AdapterName,
		config_directives[3] + config_data.Cipher,
		config_directives[4],
		config_directives[5],
		config_directives[6] + path_to_srv_dir + "/ccd",
		config_directives[7] + path_to_srv_dir + "/ca.crt",
		config_directives[8] + path_to_srv_dir + "/server.crt",
		config_directives[9] + path_to_srv_dir + "/server.key",
		config_directives[10] + path_to_srv_dir + "/ta.key 0",
		config_directives[11],
		config_directives[12],
	}

	file, err := os.Create(path_to_srv_dir + "/server.conf")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for i := 0; i < 13; i++ {
		file.WriteString(text[i] + "\n")
	}
	//Generating ccd content
	IPool := strings.Split(config_data.AddrPool, ",")
	for i := 0; i < len(IPool); i++ {
		file, err := os.Create(path_to_srv_dir + "/ccd/client" + strconv.Itoa(i+1))

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString("ifconfig-push " + IPool[i] + " " + config_data.Netmask + "\n")
		file.WriteString("push \"" + config_data.Route + "\"" + "\n")
	}

}

func ReadKeyFile(pathToKey string) []byte {
	data, err := ioutil.ReadFile(pathToKey)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (config_data *ConfJson) CreateClientConf(config_directives []string, path_to_client_dir string, path_to_keys string) {
	//Creating clients dir if not exists
	_, err := os.Stat(path_to_client_dir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path_to_client_dir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	//Creating configs
	IPool := strings.Split(config_data.AddrPool, ",")
	for i := 0; i < len(IPool); i++ {
		file, err := os.Create(path_to_client_dir + "/client" + strconv.Itoa(i+1) + ".ovpn")

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(config_directives[13] + "\n")
		file.WriteString(config_directives[1] + config_data.Proto + "\n")
		file.WriteString(config_directives[2] + config_data.AdapterName + "\n")
		file.WriteString(config_directives[14] + config_data.ServerIP + "\n")
		file.WriteString(config_directives[15] + "\n")
		file.WriteString("<ca>\n")
		file.WriteString(string(ReadKeyFile(path_to_keys + "/ca.crt")))
		file.WriteString("</ca>\n")
		file.WriteString("<cert>\n")
		file.WriteString(string(ReadKeyFile(path_to_keys + "/client" + strconv.Itoa(i+1) + ".crt")))
		file.WriteString("</cert>\n")
		file.WriteString("<key>\n")
		file.WriteString(string(ReadKeyFile(path_to_keys + "/client" + strconv.Itoa(i+1) + ".key")))
		file.WriteString("</key>\n")
		file.WriteString("<tls-auth>\n")
		file.WriteString(string(ReadKeyFile(path_to_keys + "/ta.key")))
		file.WriteString("</tls-auth>\n")
		file.WriteString(config_directives[3] + config_data.Cipher + "\n")
	}
}

// func main() {
// 	config_data := conf_json{"srv_ip", "tap", "1194", "udp", "tap0", "aes", "10.1.244.249,10.1.244.250", "255.255.255.0", "route 10.1.0.0 255.255.0.0 10.1.254.1"}
// 	config_directives := []string{
// 		"port ",
// 		"proto ",
// 		"dev ",
// 		"cipher ",
// 		"mode server",
// 		"tls-server",
// 		"client-config-dir ",
// 		"ca ",
// 		"cert ",
// 		"key ",
// 		"tls-auth ",
// 		"client-to-client",
// 		"max-routes-per-client 2048",
// 		//Client conf directives
// 		"client", //14
// 		"remote ",
// 		"tls-client",
// 		"pull",
// 	}
// 	// CreateServerConfigAndCcd(config_directives, config_data, "confs/server")
// 	// CreateServerConfigAndCcd(config_directives, config_data, "/etc/openvpn/server")
// 	CreateClientConf(config_directives, config_data, "clients", "certs")
// }
