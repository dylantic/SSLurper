package models

type QueryRequest struct {
	Domain       string
	TlsVersions  []string
	CipherSuites []string
}
