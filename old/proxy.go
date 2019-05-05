package mjpeg

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

// Flags for the source MJPEG stream and the server port
var url = flag.String("url", "", "URL for the MJPEG stream")
var addr = flag.String("addr", ":8080", "Port to listen on")

// RelayStream reads a MJPEG stream and creates a relay for n number of clients to access
func RelayStream() error {

	// Parse the input flags for port and MJPEG source url
	flag.Parse()
	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Println("Relay stream started")

	// Serve the MJPEG video in the default route of /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Read the mjpeg stream
		resp, err := http.Get(*url)
		if err != nil {
			log.Println("Error fetching response")
			return
		}

		// Copy headers from the response to our relay
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))

		// Copy the body
		io.Copy(w, resp.Body)
		resp.Body.Close()

		log.Println("Relay stream stopped")
	})

	// Listen on the specifed port and serve the MJPEG
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Println("Error listening on the port")
		return err
	}

	return nil
}
