package models

type QueryReq struct {
	Domain       string
	TlsVersions  []string
	CipherSuites []string
}
