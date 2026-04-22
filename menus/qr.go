package menus

import (
	"bufio"
	"fmt"
	"openDevTools/QR"
	"openDevTools/models"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func showQRMenu() {
	ClearScreen()
	var m Menu

	m.title = "QR Code"
	m.desc = "Encode and decode QR Code."

	m.navItems = append(m.navItems, navItem{
		name: "Decode QR Code",
		do:   decodeQRCodeMenu,
	})
	m.navItems = append(m.navItems, navItem{
		name: "Encode QR Code",
		do:   showSelectTypeQRMenu,
	})

	m.show()

}

func decodeQRCodeMenu() {
	ClearScreen()
	var m Menu

	m.title = "Decode QR Code"
	m.desc = ""

	m.navItems = append(m.navItems, navItem{
		name: "From clipboard",
		do:   showScanQrCodeFromClipboard,
	})

	m.navItems = append(m.navItems, navItem{
		name: "From file",
		do:   showSelectFile4ScanQr,
	})

	m.show()

}

func showSelectFile4ScanQr() {
	ClearScreen()
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	filepath, ok := OpenFileDialog(nil)

	if !ok {
		pterm.Warning.Println("openFileDialog cancelled")
		return
	}

	pterm.Info.Println("filepath:", filepath)

	data, err := os.ReadFile(filepath)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	result, err := QR.Scan(data)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	showQrCodeResultTable(result)
}

func showScanQrCodeFromClipboard() {
	ClearScreen()

	pterm.Info.Println("Reading image from clipboard...")

	data, err := readImgFromClipboard()
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	result, err := QR.Scan(data)
	if err != nil {
		pterm.Warning.Println(err)
		return
	}

	showQrCodeResultTable(result)

}

func showQrCodeResultTable(results []models.ResultItem) {
	tableHeader := pterm.TableData{
		{"index", "value"},
	}

	showTable(tableHeader, results)
}

func showSelectTypeQRMenu() {
	reader := bufio.NewReader(os.Stdin)
	options := []QR.QRType{QR.TypeURL, QR.TypeVCard, QR.TypeWiFi, QR.TypeSMS, QR.TypeTel, QR.TypeEmail, QR.TypeGeo, QR.TypeEvent, QR.TypePay}

	ClearScreen()
	var m Menu

	m.title = "QR Code Generator"
	m.desc = "Select a specialized QR payload type to build and encode."

	var selected QR.QRType

	for _, option := range options {
		m.navItems = append(m.navItems, navItem{
			name: string(option),
			do: func() {
				selected = option
			},
		})
	}

	m.show()

	payload := make(map[string]string)

	switch selected {
	case QR.TypeURL:
		fmt.Print("Enter URL (e.g., https://example.com): ")
		payload["url"], _ = reader.ReadString('\n')
		payload["url"] = strings.TrimSpace(payload["url"])
	case QR.TypeVCard:
		fmt.Print("Full name: ")
		payload["fn"], _ = reader.ReadString('\n')
		payload["fn"] = strings.TrimSpace(payload["fn"])
		fmt.Print("Organization: ")
		payload["org"], _ = reader.ReadString('\n')
		payload["org"] = strings.TrimSpace(payload["org"])
		fmt.Print("Phone: ")
		payload["tel"], _ = reader.ReadString('\n')
		payload["tel"] = strings.TrimSpace(payload["tel"])
		fmt.Print("Email: ")
		payload["email"], _ = reader.ReadString('\n')
		payload["email"] = strings.TrimSpace(payload["email"])
	case QR.TypeWiFi:
		fmt.Print("SSID: ")
		payload["ssid"], _ = reader.ReadString('\n')
		payload["ssid"] = strings.TrimSpace(payload["ssid"])
		fmt.Print("Authentication (WPA/WEP/nopass): ")
		payload["auth"], _ = reader.ReadString('\n')
		payload["auth"] = strings.TrimSpace(payload["auth"])
		fmt.Print("Password (leave blank for nopass): ")
		payload["password"], _ = reader.ReadString('\n')
		payload["password"] = strings.TrimSpace(payload["password"])
	case QR.TypeSMS:
		fmt.Print("Phone number: ")
		payload["phone"], _ = reader.ReadString('\n')
		payload["phone"] = strings.TrimSpace(payload["phone"])
		fmt.Print("Message: ")
		payload["message"], _ = reader.ReadString('\n')
		payload["message"] = strings.TrimSpace(payload["message"])
	case QR.TypeTel:
		fmt.Print("Phone number: ")
		payload["phone"], _ = reader.ReadString('\n')
		payload["phone"] = strings.TrimSpace(payload["phone"])
	case QR.TypeEmail:
		fmt.Print("Email address: ")
		payload["email"], _ = reader.ReadString('\n')
		payload["email"] = strings.TrimSpace(payload["email"])
		fmt.Print("Subject (optional): ")
		payload["subject"], _ = reader.ReadString('\n')
		payload["subject"] = strings.TrimSpace(payload["subject"])
		fmt.Print("Body (optional): ")
		payload["body"], _ = reader.ReadString('\n')
		payload["body"] = strings.TrimSpace(payload["body"])
	case QR.TypeGeo:
		fmt.Print("Latitude: ")
		payload["lat"], _ = reader.ReadString('\n')
		payload["lat"] = strings.TrimSpace(payload["lat"])
		fmt.Print("Longitude: ")
		payload["lon"], _ = reader.ReadString('\n')
		payload["lon"] = strings.TrimSpace(payload["lon"])
	case QR.TypeEvent:
		fmt.Print("Summary: ")
		payload["summary"], _ = reader.ReadString('\n')
		payload["summary"] = strings.TrimSpace(payload["summary"])
		fmt.Print("Start (YYYYMMDDTHHMMSSZ): ")
		payload["start"], _ = reader.ReadString('\n')
		payload["start"] = strings.TrimSpace(payload["start"])
		fmt.Print("End (YYYYMMDDTHHMMSSZ): ")
		payload["end"], _ = reader.ReadString('\n')
		payload["end"] = strings.TrimSpace(payload["end"])
		fmt.Print("Location (optional): ")
		payload["location"], _ = reader.ReadString('\n')
		payload["location"] = strings.TrimSpace(payload["location"])
	case QR.TypePay:
		fmt.Print("Payment payload (e.g., EMV QR string or UPI/VPA): ")
		payload["payment"], _ = reader.ReadString('\n')
		payload["payment"] = strings.TrimSpace(payload["payment"])
	}

	qrData := QR.QRSelection{Type: selected, Payload: payload}

	data, err := QR.BuildQRString(qrData)
	if err != nil {
		pterm.Error.Println(err)
		return
	}

	pterm.Info.Println("qrCode image filetype - PNG")
	processAndSaveImg(QR.Generate, []byte(data))

}
