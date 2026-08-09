package main

import (
	"fmt"
	"net"
	"net/http"
)

// Init server & attaching everything
func main() {
	server := NewServer()

	// creating dedicated router with mux
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("web")))

	// Attaching route with mux
	mux.HandleFunc("/send", server.handleSend)

	// Attaching msgs with mux
	mux.HandleFunc("/messages", server.handleMessages)

	ip := getLocalIP()

	fmt.Printf("Server running on LAN: http://%s:8080\n", ip)
	http.ListenAndServe(":8080", mux)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}
