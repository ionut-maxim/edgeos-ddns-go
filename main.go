package main

import (
	"log"
	"log/syslog"
	"os"

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

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "ddns-go")
	if e == nil {
		log.SetOutput(logwriter)
	}

	api := cfClient(cfApiKey, cfEmail)

	app, recipient := poClient(poToken, poUser)

	zoneID := getZoneID(cfZoneName, api)
	id, cfIp := getDNSRecord(cfRecordName, zoneID, api)

	ifaceIp := getInterfaceIP(ifaceName)

	log.Print("Cloudflare IP: '" + cfIp + "'; pppoe0 IP: '" + ifaceIp.String() + "'")

	if cfIp != ifaceIp.String() {
		if ipInRange(ifaceIp, cgNatRange) {
			restartInterface(ifaceName)

			notify("Restarted pppoe connection\nNew IP is"+getInterfaceIP(ifaceName).String(), *recipient, *app)
		}

		updateDNSRecord(zoneID, id, ifaceIp.String(), api)

		notify("Updated Cloudflare with IP "+ifaceIp.String(), *recipient, *app)
	}

}
