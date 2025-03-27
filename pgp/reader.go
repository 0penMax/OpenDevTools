package pgp

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"io"
	"openDevTools/models"
	"time"
)

func ReadPublicKeyData(data io.Reader) ([]models.ResultItem, error) {
	// Parse the armored key ring.
	entities, err := openpgp.ReadArmoredKeyRing(data)
	if err != nil {
		return nil, errors.Join(errors.New("Error reading armored key ring:"), err)
	}

	result := make([]models.ResultItem, 0)
	// Iterate over each parsed entity (public key).
	for _, entity := range entities {
		result = append(result, models.ResultItem{
			Name:  "Primary Keys:",
			Value: "",
		})
		result = append(result, models.ResultItem{
			Name:  "Fingerprint",
			Value: fmt.Sprintf("%X", entity.PrimaryKey.Fingerprint),
		})
		result = append(result, models.ResultItem{
			Name:  "Key ID",
			Value: fmt.Sprintf("%X", entity.PrimaryKey.KeyId),
		})

		result = append(result, models.ResultItem{
			Name:  "Creation Time",
			Value: fmt.Sprintf("%v", entity.PrimaryKey.CreationTime),
		})
		result = append(result, models.ResultItem{
			Name:  "Public Key Algorithm",
			Value: fmt.Sprintf("%v", entity.PrimaryKey.PubKeyAlgo),
		})

		if bitLength, err := entity.PrimaryKey.BitLength(); err != nil {
			result = append(result, models.ResultItem{
				Name:  "Bit Length:",
				Value: fmt.Sprintf("  error - %v\n", err),
			})
		} else {
			result = append(result, models.ResultItem{
				Name:  "Bit Length",
				Value: fmt.Sprint(bitLength),
			})
		}

		result = append(result, models.ResultItem{
			Name:  "User Identities:",
			Value: "",
		})

		// Each identity includes the user ID string and its self-signature.
		for identityName, identity := range entity.Identities {
			result = append(result, models.ResultItem{
				Name:  "Identity",
				Value: identityName,
			})
			result = append(result, models.ResultItem{
				Name:  "UserId",
				Value: fmt.Sprintf("%v", identity.UserId),
			})
			result = append(result, models.ResultItem{
				Name:  "Self-Signature Creation Time",
				Value: fmt.Sprintf("%v", identity.SelfSignature.CreationTime),
			})

			// If KeyLifetimeSecs is set, calculate the expiration time.
			if identity.SelfSignature.KeyLifetimeSecs != nil {
				lifetime := time.Duration(*identity.SelfSignature.KeyLifetimeSecs) * time.Second
				expiration := entity.PrimaryKey.CreationTime.Add(lifetime)
				result = append(result, models.ResultItem{
					Name:  "Key Expiration Time",
					Value: fmt.Sprintf("%v", expiration),
				})
			} else {
				result = append(result, models.ResultItem{
					Name:  "Key Expiration Time",
					Value: "No expiration set (key does not expire)",
				})
			}
		}

		result = append(result, models.ResultItem{
			Name:  "Subkeys:",
			Value: "",
		})
		// Loop through subkeys, if any.
		for _, subkey := range entity.Subkeys {
			result = append(result, models.ResultItem{
				Name:  "Subkey Fingerprint",
				Value: fmt.Sprintf("%X", subkey.PublicKey.Fingerprint),
			})
			result = append(result, models.ResultItem{
				Name:  "Subkey Key ID",
				Value: fmt.Sprintf("%X", subkey.PublicKey.KeyId),
			})
			result = append(result, models.ResultItem{
				Name:  "Creation Time",
				Value: fmt.Sprintf("%v", subkey.PublicKey.CreationTime),
			})
			result = append(result, models.ResultItem{
				Name:  "Public Key Algorithm",
				Value: fmt.Sprintf("%v", subkey.PublicKey.PubKeyAlgo),
			})
			if bitLength, err := subkey.PublicKey.BitLength(); err != nil {
				result = append(result, models.ResultItem{
					Name:  "Bit Length",
					Value: fmt.Sprintf("error - %v", err),
				})
			} else {
				result = append(result, models.ResultItem{
					Name:  "Bit Length",
					Value: fmt.Sprintf("%d", bitLength),
				})
			}

			if subkey.Sig != nil && subkey.Sig.KeyLifetimeSecs != nil {
				lifetime := time.Duration(*subkey.Sig.KeyLifetimeSecs) * time.Second
				expiration := subkey.PublicKey.CreationTime.Add(lifetime)
				result = append(result, models.ResultItem{
					Name:  "Subkey Expiration Time",
					Value: fmt.Sprintf("%v", expiration),
				})
				fmt.Println("  Subkey Expiration Time:", expiration)
			} else {
				result = append(result, models.ResultItem{
					Name:  "Subkey Expiration Time",
					Value: fmt.Sprintf("%s", "No expiration set for this subkey"),
				})
			}
		}
	}

	return result, nil

}
