package utils

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func getMacAndIpAddress() map[string]string {
	deviceNetAddrs := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("error", fmt.Sprintf("oops! %v", err))
		return deviceNetAddrs
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println("error", fmt.Sprintf("oops! %v", err))
			return deviceNetAddrs
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				deviceNetAddrs["ip"] = v.IP.String()
			case *net.IPAddr:
				deviceNetAddrs["ip"] = v.IP.String()
			}
			mac := i.HardwareAddr.String()
			if mac != "" {
				deviceNetAddrs["mac"] = mac
			}
		}
	}

	return deviceNetAddrs
}

func getOs() map[string]string {
	osDetails := make(map[string]string)
	osDetails["os"] = runtime.GOOS
	osDetails["arch"] = runtime.GOARCH

	return osDetails
}

func getDeviceName() *string {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Println("unable to get device hostname")
		return nil
	}

	hostname := string(out)[:len(out)-1]
	return &hostname
}

func generateFingerprint() string {
	var fp bytes.Buffer
	deviceNetDetails := getMacAndIpAddress()
	osDetails := getOs()
	deviceName := getDeviceName()
	_, ipOk := deviceNetDetails["ip"]
	if ipOk {
		fp.WriteString(deviceNetDetails["ip"])
		fp.WriteString(":")
	}
	_, macAddrOk := deviceNetDetails["mac"]
	if macAddrOk {
		fp.WriteString(deviceNetDetails["mac"])
		fp.WriteString(":")
	}
	_, osOk := osDetails["os"]
	if osOk {
		fp.WriteString(osDetails["os"])
		fp.WriteString(":")
	}
	_, archOk := osDetails["arch"]
	if archOk {
		fp.WriteString(osDetails["arch"])
		fp.WriteString(":")
	}
	if deviceName != nil {
		fp.WriteString(*deviceName)
		fp.WriteString(":")
	}
	fpString := fp.String()

	hash := sha256.Sum256([]byte(fpString))
	hex := fmt.Sprintf("%x", hash)

	return hex
}

func GetDeviceMetadata() map[string]string {
	metadata := make(map[string]string)
	metadata["fingerprint"] = generateFingerprint()
	deviceName := getDeviceName()
	if deviceName != nil {
		metadata["device_name"] = *deviceName
	}
	osDetails := getOs()
	netDetails := getMacAndIpAddress()
	for k, v := range osDetails {
		metadata[k] = v
	}

	for k, v := range netDetails {
		metadata[k] = v
	}

	return metadata
}
