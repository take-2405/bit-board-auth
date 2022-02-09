package disutil

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func CreateFireBaseConfig() error {
	fp, err := os.Create("./firebase-auth.json")
	if err != nil {
		err = errors.Wrap(err, "os create firebase-auth.json")
		return errors.WithStack(err)
	}
	defer fp.Close()

	file := fmt.Sprintf(` {
    "type": "%s",
    "project_id": "%s",
    "private_key_id": "%s",
    "private_key": "%s",
    "client_email": "%s",
    "client_id": "%s",
    "auth_uri": "%s",
    "token_uri": "%s",
    "auth_provider_x509_cert_url": "%s",
    "client_x509_cert_url": "%s"
}`,
		os.Getenv("FS_TYPE"),
		os.Getenv("FS_PROJECT_ID"),
		os.Getenv("FS_PRIVATE_KEY_ID"),
		os.Getenv("FS_PRIVATE_KEY"),
		os.Getenv("FS_CLIENT_EMAIL"),
		os.Getenv("FS_CLIENT_ID"),
		os.Getenv("FS_AUTH_URI"),
		os.Getenv("FS_TOKEN_URI"),
		os.Getenv("FS_AUTH_PROVIDER_X509_CERT_URL"),
		os.Getenv("FS_CLIENT_X509_CERT_URL"))

	_, err = fp.Write(([]byte)(file))
	if err != nil {
		err = errors.Wrap(err, "write content on firebase-auth.json")
		return errors.WithStack(err)
	}
	return nil
}
