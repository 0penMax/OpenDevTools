package QR

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type QRType string

const (
	TypeURL   QRType = "URL"
	TypeVCard QRType = "vCard"
	TypeWiFi  QRType = "WiFi"
	TypeSMS   QRType = "SMS"
	TypeTel   QRType = "Tel"
	TypeEmail QRType = "Email"
	TypeGeo   QRType = "Geo"
	TypeEvent QRType = "Event"
	TypePay   QRType = "Payment"
	TypeRaw   QRType = "Raw"
)

type QRSelection struct {
	Type    QRType
	Payload map[string]string
}

// BuildQRString takes a QRSelection and returns the payload string for QR encoding.
func BuildQRString(sel QRSelection) (string, error) {
	if sel.Payload == nil {
		sel.Payload = map[string]string{}
	}

	switch sel.Type {
	case TypeURL:
		u := strings.TrimSpace(sel.Payload["url"])
		if u == "" {
			return "", errors.New("url is required")
		}
		// ensure scheme
		if !strings.Contains(u, "://") {
			u = "https://" + u
		}
		return u, nil

	case TypeVCard:
		fn := sel.Payload["fn"]
		org := sel.Payload["org"]
		tel := sel.Payload["tel"]
		email := sel.Payload["email"]
		var b strings.Builder
		b.WriteString("BEGIN:VCARD\r\nVERSION:3.0\r\n")
		if fn != "" {
			b.WriteString("FN:" + escapeVC(fn) + "\r\n")
		}
		if org != "" {
			b.WriteString("ORG:" + escapeVC(org) + "\r\n")
		}
		if tel != "" {
			b.WriteString("TEL:" + escapeVC(tel) + "\r\n")
		}
		if email != "" {
			b.WriteString("EMAIL:" + escapeVC(email) + "\r\n")
		}
		b.WriteString("END:VCARD")
		return b.String(), nil

	case TypeWiFi:
		ssid := sel.Payload["ssid"]
		if ssid == "" {
			return "", errors.New("ssid is required")
		}
		auth := sel.Payload["auth"]
		if auth == "" {
			auth = "WPA"
		}
		pass := sel.Payload["password"]
		hidden := sel.Payload["hidden"]
		parts := []string{fmt.Sprintf("WIFI:T:%s", escapeWiFi(auth)), "S:" + escapeWiFi(ssid)}
		if pass != "" {
			parts = append(parts, "P:"+escapeWiFi(pass))
		}
		if hidden == "true" || hidden == "True" {
			parts = append(parts, "H:true")
		}
		return strings.Join(parts, ";") + ";;", nil

	case TypeSMS:
		phone := sel.Payload["phone"]
		if phone == "" {
			return "", errors.New("phone is required")
		}
		msg := sel.Payload["message"]
		return "SMSTO:" + phone + ":" + msg, nil

	case TypeTel:
		phone := sel.Payload["phone"]
		if phone == "" {
			return "", errors.New("phone is required")
		}
		return "TEL:" + phone, nil

	case TypeEmail:
		emailAddr := sel.Payload["email"]
		if emailAddr == "" {
			return "", errors.New("email is required")
		}
		subj := url.QueryEscape(sel.Payload["subject"])
		body := url.QueryEscape(sel.Payload["body"])
		u := "mailto:" + emailAddr
		q := ""
		if subj != "" {
			q = "subject=" + subj
		}
		if body != "" {
			if q != "" {
				q += "&"
			}
			q += "body=" + body
		}
		if q != "" {
			u += "?" + q
		}
		return u, nil

	case TypeGeo:
		lat := strings.TrimSpace(sel.Payload["lat"])
		lon := strings.TrimSpace(sel.Payload["lon"])
		if lat == "" || lon == "" {
			return "", errors.New("lat and lon are required")
		}
		return "geo:" + lat + "," + lon, nil

	case TypeEvent:
		summary := escapeICal(sel.Payload["summary"])
		start := sel.Payload["start"]
		end := sel.Payload["end"]
		loc := escapeICal(sel.Payload["location"])
		var b strings.Builder
		b.WriteString("BEGIN:VCALENDAR\r\nVERSION:2.0\r\nBEGIN:VEVENT\r\n")
		if summary != "" {
			b.WriteString("SUMMARY:" + summary + "\r\n")
		}
		if start != "" {
			b.WriteString("DTSTART:" + start + "\r\n")
		}
		if end != "" {
			b.WriteString("DTEND:" + end + "\r\n")
		}
		if loc != "" {
			b.WriteString("LOCATION:" + loc + "\r\n")
		}
		b.WriteString("END:VEVENT\r\nEND:VCALENDAR")
		return b.String(), nil

	case TypePay:
		if p := sel.Payload["payment"]; p != "" {
			return p, nil
		}
		return "", errors.New("payment payload is required")

	case TypeRaw:
		if txt := sel.Payload["text"]; txt != "" {
			return txt, nil
		}
		return "", errors.New("text is required")

	default:
		return "", errors.New("unsupported QR type")
	}
}

// helpers

func escapeVC(s string) string {
	replacer := strings.NewReplacer("\\", "\\\\", "\n", "\\n", ";", "\\;", ",", "\\,")
	return replacer.Replace(s)
}

func escapeWiFi(s string) string {
	return strings.NewReplacer("\\", "\\\\", ";", "\\;").Replace(s)
}

func escapeICal(s string) string {
	replacer := strings.NewReplacer("\\", "\\\\", ";", "\\;", ",", "\\,", "\n", "\\n")
	return replacer.Replace(s)
}
