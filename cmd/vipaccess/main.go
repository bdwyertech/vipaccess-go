package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bdwyertech/vipaccess-go/vipaccess"
)

func main() {
	png := flag.String("png", "", "png path to write QR code")
	flag.Parse()

	p := vipaccess.GenerateRandomParameters()

	c, err := vipaccess.GenerateCredential(p)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Validate(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("OTP credential: %s\nExpires: %s (%s)\n",
		c.URI(), c.Expires, -time.Since(c.Expires))

	if *png != "" {
		f, err := os.Create(*png)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(c.QRCodePNG()); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("wrote QR code to %s\n", *png)
	}
}
