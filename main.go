package main

import (
	"log"
	"log/syslog"
	"os"

	"edgeos-ddns/dns/cloudflare"
	"edgeos-ddns/notification/pushover"

	_ "github.com/joho/godotenv/autoload" // Autoload .env file containing our secrets
)

func main() {
	cfApiKey := os.Getenv("CF_API_KEY")
	cfEmail := os.Getenv("CF_API_EMAIL")
	cfZoneName := os.Getenv("CF_ZONE")
	cfRecordName := os.Getenv("CF_RECORD_NAME")
	ifaceName := os.Getenv("IFACE_NAME")
	cgNatRange := os.Getenv("CGNAT_RANGE")
	poToken := os.Getenv("PUSHOVER_TOKEN")
	poUser := os.Getenv("PUSHOVER_USER")

	logwriter, err := syslog.New(syslog.LOG_NOTICE, "ddns-go")
	if err == nil {
		log.SetOutput(logwriter)
	}

	api, err := cloudflare.New(cfApiKey, cfEmail)
	if err != nil {
		log.Fatalf("unable to create cloudflare client: %v", err)
		return
	}

	id, cfIp := cloudflare.GetRecord(cfRecordName, zoneID, api)

	ifaceIp := getInterfaceIP(ifaceName)

	log.Print("Cloudflare IP: '" + cfIp + "'; pppoe0 IP: '" + ifaceIp.String() + "'")

	if cfIp != ifaceIp.String() {
		if ipInRange(ifaceIp, cgNatRange) {
			restartInterface(ifaceName)

			message := "Restarted pppoe connection\nNew IP is" + getInterfaceIP(ifaceName).String()
			pushover.Notify(message, poToken, poUser)
			if err != nil {
				log.Fatalf("unable to ")
			}
		}

		cloudflare.UpdateRecord(zoneID, id, ifaceIp.String(), api)

		pushover.Notify("Updated Cloudflare with IP "+ifaceIp.String(), poToken, poUser)
	}

}
