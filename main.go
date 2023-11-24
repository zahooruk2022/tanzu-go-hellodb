package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"tanzu-accelerator-go-hellodb-main/visits"

	"github.com/nebhale/client-go/bindings"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {

	b := bindings.FromServiceBindingRoot()
	b = bindings.Filter(b, "postgresql")
	if len(b) != 1 {
		return fmt.Errorf("expected 1 binding of type 'postgresql': got %d", len(b))
	}

	u, ok := bindings.Get(b[0], "username")
	if !ok {
		return fmt.Errorf("no username in binding")
	}

	p, ok := bindings.Get(b[0], "password")
	if !ok {
		return fmt.Errorf("no password in binding")
	}

	d, ok := bindings.Get(b[0], "database")
	if !ok {
		return fmt.Errorf("no database in binding")
	}

	host, ok := bindings.Get(b[0], "host")
	if !ok || host == "" {
		host, ok = bindings.Get(b[0], "privateIP")
		if !ok || host == "" {
			host, ok = bindings.Get(b[0], "publicIP")
			if !ok {
				return fmt.Errorf("no host, publicIP or privateIP in binding")
			}
		}
	}

	provider := "unknown"
	provider, _ = bindings.GetProvider(b[0])

	bURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", u, p, host, d)
	fmt.Println("Binding: ", bURL)
	pgURL, err := url.Parse(bURL)
	if err != nil {
		return fmt.Errorf("unable to parse URL: %w", err)
	}

	v, err := visits.NewLog(pgURL)
	if err != nil {
		return fmt.Errorf("unable to create VisitorLog: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		id, err := v.NewVisit(req.Context())
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(res, "%s", err.Error())
		}

		fmt.Fprintf(res, "Tanzu does Go too!\n\nvisit: %s\nprovider: %s", id, provider)
	})

	err = http.ListenAndServe(":8080", mux)
	return err
}
