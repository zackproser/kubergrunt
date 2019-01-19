package helm

import (
	"fmt"

	"github.com/gruntwork-io/kubergrunt/kubectl"
	"github.com/gruntwork-io/kubergrunt/tls"
)

func StoreCertificateKeyPairAsKubernetesSecret(
	kubectlOptions *kubectl.KubectlOptions,
	secretName string,
	secretNamespace string,
	labels map[string]string,
	annotations map[string]string,
	nameBase string,
	certificateKeyPairPath tls.CertificateKeyPairPath,
) error {
	secret := kubectl.PrepareSecret(secretNamespace, secretName, labels, annotations)
	err := kubectl.AddToSecretFromFile(secret, fmt.Sprintf("%s.crt", nameBase), certificateKeyPairPath.CertificatePath)
	if err != nil {
		return err
	}
	err = kubectl.AddToSecretFromFile(secret, fmt.Sprintf("%s.pem", nameBase), certificateKeyPairPath.PrivateKeyPath)
	if err != nil {
		return err
	}
	err = kubectl.AddToSecretFromFile(secret, fmt.Sprintf("%s.pub", nameBase), certificateKeyPairPath.PublicKeyPath)
	if err != nil {
		return err
	}
	return kubectl.CreateSecret(kubectlOptions, secret)
}
