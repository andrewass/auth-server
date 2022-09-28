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
			Kid: "stETwk2s1n0-jK8XZ7xDawkfExuSS9zkc_ih-TAXl7A",
			Alg: "RS256",
			N:   "tKje00qALNfKWBkDszHPN1ZJgt2lvyeEvDLbjMjpfdACt_-WAySXaT2geRTN9kN6w7mgCWbYhiMFWbRNO8dyK4i83VkK5ncT-hEd0XbLuJivuBSPhtoQheKVD50YZguVXIpkd6fsHl0bpsfGT9x_7uzSBReYYP64Ox0a0Wek7aC9LMmt8N2IPOCBzVVceyJoJb-TylL5rvJMfWl97ucAgMt8eBdmxFYGZQEGOAFGVRWEroWQp-AdsKEOn3kzDj3xTc6TPYPTy_FhIBzdMsMhrnjvnpBqWyaT2Ycqjk-3U6DfO9r5EeWGqW6FxKeDHr343fatQMQrejFEO8MOiJc3kQ",
		},
	}
}
