package gmtls

type Certificate struct{}

func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
	return Certificate{}, nil
}
