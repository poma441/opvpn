package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//ToDO: add parameter for CA name
func CreateCA() string {
	log_string := ""
	_, err := os.Stat("certs")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("certs", 0777)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization:  []string{"ORGANIZATION_NAME"},
			Country:       []string{"COUNTRY_CODE"},
			Province:      []string{"PROVINCE"},
			Locality:      []string{"CITY"},
			StreetAddress: []string{"ADDRESS"},
			PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPriv, _ := rsa.GenerateKey(rand.Reader, 2048)
	caPub := &caPriv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, caCert, caCert, caPub, caPriv)
	if err != nil {
		log.Println("create ca failed", err)
		return "create ca failed"
	}

	// Public key
	caCertOut, err := os.Create("certs/ca.crt")
	if err != nil {
		log.Println("Cannot create ca.crt! ", err)
		return "Cannot create file ca.crt!"
		// panic(err)
	}
	pem.Encode(caCertOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	caCertOut.Close()
	log.Print("written CA cert.pem\n")
	log_string += "written CA cert.pem; "

	// Private key
	caKeyOut, err := os.OpenFile("certs/ca.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println("Cannot create ca.key! ", err)
		return "Cannot create ca.key! "
		// panic(err)
	}
	pem.Encode(caKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPriv)})
	caKeyOut.Close()
	log.Print("written CA ca.key\n")
	log_string += "written CA ca.key; "
	return log_string
}

//ToDO: add parameter for Server name
func CreateServer(caCert_path string, caPriv_path string) {
	catls, err := tls.LoadX509KeyPair(caCert_path, caPriv_path)
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}
	//Creating server certificate and priv key
	serverCert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"ORGANIZATION_NAME"},
			Country:       []string{"COUNTRY_CODE"},
			Province:      []string{"PROVINCE"},
			Locality:      []string{"CITY"},
			StreetAddress: []string{"ADDRESS"},
			PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	serverPriv, _ := rsa.GenerateKey(rand.Reader, 2048)
	serverPub := &serverPriv.PublicKey

	// Sign the certificate
	serverCert_b, err := x509.CreateCertificate(rand.Reader, serverCert, ca, serverPub, catls.PrivateKey)
	if err != nil {
		log.Println("Cannot create Server cert", err)
	}
	// Public key
	certOut, err := os.Create("certs/server.crt")
	if err != nil {
		log.Println("Cannot create Server cert file", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: serverCert_b})
	certOut.Close()
	log.Print("written Server cert.pem\n")

	// Private key
	keyOut, err := os.OpenFile("certs/server.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("Cannot create Server key file", err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(serverPriv)})
	keyOut.Close()
	log.Print("written Server key.pem\n")
}

//ToDO: add parameter for client name
func CreateClient(clientNum int, caCert_path string, caPriv_path string) {
	// clientNum = strconv.Itoa(clientNum)
	catls, err := tls.LoadX509KeyPair(caCert_path, caPriv_path)
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}
	// Creating client certificate and priv key
	clientCert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"ORGANIZATION_NAME"},
			Country:       []string{"COUNTRY_CODE"},
			Province:      []string{"PROVINCE"},
			Locality:      []string{"CITY"},
			StreetAddress: []string{"ADDRESS"},
			PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	clientPriv, _ := rsa.GenerateKey(rand.Reader, 2048)
	clientPub := &clientPriv.PublicKey

	// Sign the certificate
	clientCert_b, err := x509.CreateCertificate(rand.Reader, clientCert, ca, clientPub, catls.PrivateKey)
	if err != nil {
		log.Println("Cannot create Client cert", err)
	}
	// Public key
	certOut, err := os.Create("certs/client" + strconv.Itoa(clientNum) + ".crt")
	if err != nil {
		log.Println("Cannot create Client cert file", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: clientCert_b})
	certOut.Close()
	log.Print("written Client" + strconv.Itoa(clientNum) + " cert.pem\n")

	// Private key
	keyOut, err := os.OpenFile("certs/client"+strconv.Itoa(clientNum)+".key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("Cannot create Client key file", err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientPriv)})
	keyOut.Close()
	log.Print("written Client" + strconv.Itoa(clientNum) + " key.pem\n")
}

func CreateTA() string {
	cmd := exec.Command("openvpn", "--genkey", "secret", "certs/ta.key")
	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
		return "Cannot execute command to create TA.key; "
	}
	log.Print("written ta.key\n")
	return "created ta.key; "
}

func CreateDH() string {
	cmd := exec.Command("openssl", "dhparam", "-out", "/etc/openvpn/certs/dh2048.pem 2048")
	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
		return "Cannot execute command to create dh2048.pem; "
	}
	log.Print("written Client dh2048.pem\n")
	return "created dh2048.pem; "
}
