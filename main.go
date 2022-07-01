package main

import (
	"WallE/config"
	"WallE/handlers/rest"
	_ "crypto/tls"
	_ "crypto/x509"
	_ "io/ioutil"
	_ "log"

	"github.com/labstack/echo/v4"
)

const localCertFile = "/config/origin_ca_rsa_root.pem"

func main() {

	config := config.InitConfig()
	e := echo.New()
	rest.RegisterMainAPI(e, config)
	e.Logger.Fatal(e.Start(":8080"))
	// rootCAs, _ := x509.SystemCertPool()

	// if rootCAs == nil {
	// 	rootCAs = x509.NewCertPool()
	// }

	// certs, _ := ioutil.ReadFile(localCertFile)

	// if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
	// 	log.Println("No certs appended, using system certs only")
	// }

	// s := http.Server{
	// 	Addr:    ":443",
	// 	Handler: e,
	// 	TLSConfig: &tls.Config{
	// 		RootCAs: rootCAs,
	// 	},
	// }
	// if err := s.ListenAndServeTLS("/config/certificate.crt", "/config/private.key"); err != http.ErrServerClosed {
	// 	e.Logger.Fatal(err)
	// }
}
