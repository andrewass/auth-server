package resources

type JWK struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	N   string `json:"n"`
}

func GetJwks() []JWK {
	return []JWK{
		{
			Kty: "RSA",
			E:   "AQAB",
			Use: "sig",
			Kid: "M7QzWoOvRz0Tv111qCnpRThoh0QvEyRZyRlzp8mNgFc",
			Alg: "RS256",
			N:   "ivRDN04Z_wCqSxaJdHCr8Gnu-TIEypUWQKT68gSiIbMSGwrF91ia4Ec5bMeHtqn3b40yF2uW2Zl_jDYQrZRGxiTXZnRdkOYmeyF8VpPg21S1ajS_lf9H8OBsbugV1HYs8bO1slvuWob6KT6pEPrl7IFtd0WSG2_tBxbRtb4wO7DDVSuY3MUMiK078DgM77jW_m71_JzMBPJUlEwEV3SsEyOMZcBDHfrt1OmaMw26CRMBbEncZaEHIrn1U363iMOdYWrVVUP5n3Mlk4GWGM9ZseGdAXpdSGGccXQodwya1q6Kt_0vmnrlOql9u5KQWifel9-j3SODMsXJyfJinfYdQw",
		},
	}
}
