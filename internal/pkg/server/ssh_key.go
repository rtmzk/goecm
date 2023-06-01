package server

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

func GenSshKeys(info *v1.KeyInfo) error {
	privKey, pubKey, err := MakeSSHKeyPair()
	if err != nil {
		return err
	}
	info.PublicKey = pubKey
	info.PrivateKey = privKey

	var count int64
	count, _ = store.Client().SSHKey().Get(context.Background(), info, metav1.GetOptions{})

	if count == 1 {
		return nil
	}
	err = store.Client().SSHKey().Create(context.Background(), info, metav1.CreateOptions{})
	if err != nil {
		return nil
	}

	if err = SavePrivKey(info.PrivateKey); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("../../internal/goecmserver/template/add_ssh_key.sh.tpl")
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, pubKey)
	data, _ := ioutil.ReadAll(&buff)
	log.Info(string(data))
	if err != nil {
		return err
	}
	err = saveFile("add_host_key.sh", string(data))
	if err != nil {
		return err
	}

	return nil
}

func GetSshKey() (string, string, error) {
	var key v1.KeyInfo
	var err error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", err
	}
	return key.PublicKey, key.PrivateKey, nil
}

//nolint: not used.
func GenKeyScript() ([]byte, error) {
	pub_key, _, err := GetSshKey()
	if err != nil {
	}

	tmpl, err := template.ParseFiles("../../internal/goecmserver/template/add_ssh_key.sh.tpl")
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, pub_key)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func SavePrivKey(privateKey string) error {
	if utils.IsNotExist(utils.UserHome() + "/.ssh") {
		os.MkdirAll(utils.UserHome()+"/.ssh", 0600)
	}
	file, err := os.OpenFile(utils.UserHome()+"/.ssh/id_rsa", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(privateKey))
	if err != nil {
		return err
	}
	return err
}

func saveFile(name string, data string) error {
	file, err := os.OpenFile(utils.UserHome()+"/.ssh/"+name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return nil
}

func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

func EncodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}

//EncodeSSHKey
func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func MakeSSHKeyPair() (string, string, error) {
	pkey, pubkey, err := GenerateKey(2048)
	if err != nil {
		return "", "", err
	}

	pub, err := EncodeSSHKey(pubkey)
	if err != nil {
		return "", "", err
	}

	return string(EncodePrivateKey(pkey)), string(pub), nil
}
