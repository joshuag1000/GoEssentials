package goessentials

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func StartWebServer(NonHttpsPort string, SitePort string) {
	ExecPath, err2 := filepath.Abs(filepath.Dir(os.Args[0]))
	if err2 != nil {
		log.Fatal(err2)
	}
	go func() {
		// check for https certificates
		if _, err := os.Stat(ExecPath + "/HTTPS-key/server.crt"); os.IsNotExist(err) {
			fmt.Printf("server.crt does not exist. HTTPS NOT STARTED\n")
		} else if _, err := os.Stat(ExecPath + "/HTTPS-key/server.key"); os.IsNotExist(err) {
			fmt.Printf("server.key does not exist. HTTPS NOT STARTED\n")
		} else {
			// begin https server
			err_https := http.ListenAndServeTLS(":"+SitePort, ExecPath+"/HTTPS-key/server.crt", ExecPath+"/HTTPS-key/server.key", nil)
			if err_https != nil {
				log.Fatal("Web server (HTTPS): \n", err_https)
			}
		}
	}()

	// begin http server
	err_http := http.ListenAndServe(":"+NonHttpsPort, nil)
	if err_http != nil {
		log.Fatal("Web server (HTTP): ", err_http)
	}
}
