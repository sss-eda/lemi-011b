package main

// https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html
// To run:
// go run main.go
// Command-line options:
//   -production : enables HTTPS on port 443
//   -redirect-to-https : redirect HTTP to HTTTPS

// func makeHTTPToHTTPSRedirectServer() *http.Server {
// 	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
// 		newURI := "https://" + r.Host + r.URL.String()
// 		http.Redirect(w, r, newURI, http.StatusFound)
// 	}
// 	mux := &http.ServeMux{}
// 	mux.HandleFunc("/", handleRedirect)
// 	return &http.Server{
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 5 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 		Handler:      mux,
// 	}
// }

// func parseFlags() {
// 	flag.BoolVar(&flgProduction, "production", false, "if true, we start HTTPS server")
// 	flag.BoolVar(&flgRedirectHTTPToHTTPS, "redirect-to-https", false, "if true, we redirect HTTP to HTTPS")
// 	flag.Parse()
// }

// func main() {
// 	ctx := context.Background()

// 	configurator, err := configuration.NewService(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	config := &Config{}
// 	err = configurator.ParseENV(ctx, config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var m *autocert.Manager

// 	var httpsSrv *http.Server
// 	if config.Environment == Production {
// 		hostPolicy := autocert.HostWhitelist(config.TLS.Hosts)
// 		func(ctx context.Context, host string) error {
// 			// Note: change to your real host
// 			allowedHost := "www.mydomain.com"
// 			if host == allowedHost {
// 				return nil
// 			}
// 			return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
// 		}

// 		dataDir := "."
// 		m = &autocert.Manager{
// 			Prompt:     autocert.AcceptTOS,
// 			HostPolicy: hostPolicy,
// 			Cache:      autocert.DirCache(dataDir),
// 		}

// 		httpsSrv = &http.Server{
// 			ReadTimeout:  5 * time.Second,
// 			WriteTimeout: 5 * time.Second,
// 			IdleTimeout:  120 * time.Second,
// 			Handler:      &http.ServeMux{},
// 		}
// 		httpsSrv.Addr = ":443"
// 		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

// 		go func() {
// 			fmt.Printf("Starting HTTPS server on %s\n", httpsSrv.Addr)
// 			err := httpsSrv.ListenAndServeTLS("", "")
// 			if err != nil {
// 				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
// 			}
// 		}()
// 	}

// 	var httpSrv *http.Server
// 	if flgRedirectHTTPToHTTPS {
// 		httpSrv = makeHTTPToHTTPSRedirectServer()
// 	} else {
// 		httpSrv = &http.Server{
// 			ReadTimeout:  5 * time.Second,
// 			WriteTimeout: 5 * time.Second,
// 			IdleTimeout:  120 * time.Second,
// 			Handler:      &http.ServeMux{},
// 		}
// 	}
// 	// allow autocert handle Let's Encrypt callbacks over http
// 	if m != nil {
// 		httpSrv.Handler = m.HTTPHandler(httpSrv.Handler)
// 	}

// 	httpSrv.Addr = httpPort
// 	fmt.Printf("Starting HTTP server on %s\n", httpPort)
// 	err := httpSrv.ListenAndServe()
// 	if err != nil {
// 		log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
// 	}
// }
