package resources

type KeyPair struct {
	Alg string
	D   string
	Dp  string
	Dq  string
	E   string
	Kid string
	Kty string
	N   string
	P   string
	Use string
	Q   string
	Qi  string
}

func GetKeyPair() KeyPair {
	return KeyPair{
		P:   "4Tsx1IzEo8BuYSCwTS0yKLcGP99D0h4Y-FW66uJawZWM16mZFPptCxSPqttCUrp5rYoiCbNCS3_kojWgkDKS_c3pKe4npuOtbI4HnSNhB9X0WItcvVhRKYUoZCUtT3FoNJfceYGE1jog5mYs4dJxuGFHIXs-kErIK8foRtZnBH0",
		Kty: "RSA",
		Q:   "zVbqLt79Sxr70itecj4wInMpYiuFZziQQdsw52yow2whXutU77_jmZ4kFpELuaeOg6H0zLRI23aEll11bLSuEPIr1ni9zTpR12Rgx0X5vEhI1jE5tY71wf6XEc9_ao4SCNyjWI59ufUk8NhPKx2b0PcynNqbcy7gTSRMF0QGD6U",
		D:   "BHoKy4GVG4ddZEF6bf5V_DEPMUb8zHfWOjOobRytbIizCIw0028csWbtixKQ6AFUEXWQGauKwggsIcTuBQU0Nw_0XifOtVdOD2xNKNugwLpH4G_h_p14u1iaCQEDkcPHt7tWMgDfQQFdpgxdRa1ZHjqQGcAz7nNhbVWTTI0c8SZ33i1fhFziuB-OYnW1fqZsEW_Htc9cDNQCuBKFJjkk3kmDadWksErh2OFBSyrtzLAP2bOUufNhCcxIaZYryVuK4agvk1HpXQ4a2v8V3yfO6bwRU4g2MZDhQv7n4sa_ME6kw6Pmm41cPjqz2G9LRmzo1o35_ag_xYOeRiJrDM_OgQ",
		E:   "AQAB",
		Use: "sig",
		Kid: "stETwk2s1n0-jK8XZ7xDawkfExuSS9zkc_ih-TAXl7A",
		Qi:  "usRjzkde4QmtpxqcAPsHdYJKi3d9HGwWECSPltdFC5qu3sd-yINbB15Xgb1pqcjp3CbSDCWv2hd2fB1YU8uuQIvGbc-fkUxslsfz-cZ6pwBArUGpSl9fjxue3t9nqmcZBD9yes-47a0jE1_F3OayfVUMJyzgC3v7e6SUjW8sCu4",
		Dp:  "d7UOdvmaSa4s_FJzYZFz50_fcnMniLWWb-agwdYshlcF1Fm_kvbGnez_rr96MJ4LrcCM9rrfsWD4E36NKBa5KUwXcMGzMXAw7FFQmqQDHjcI90aas8gM2xR6sz9PJboQuGM-OxYrk_CxK8OCd0NyTZzmQVVPm0EjwqEtQvkRJ8E",
		Alg: "RS256",
		Dq:  "fcdPi4ZsASV2Ozb1k2UkdvWz07506NVtB-oL-rZEafBCQLJGRNV1xRNpZ6lRVVedpGQAX6PPPR9L-A8nbn7-YML2t1keaiffGeESpv3mC777XriocN4xC7O3NeG0WsDLq0H7jLM7sDzJc_4MAI0SdN320r3VmqiMHLSvJ1EzpUE",
		N:   "tKje00qALNfKWBkDszHPN1ZJgt2lvyeEvDLbjMjpfdACt_-WAySXaT2geRTN9kN6w7mgCWbYhiMFWbRNO8dyK4i83VkK5ncT-hEd0XbLuJivuBSPhtoQheKVD50YZguVXIpkd6fsHl0bpsfGT9x_7uzSBReYYP64Ox0a0Wek7aC9LMmt8N2IPOCBzVVceyJoJb-TylL5rvJMfWl97ucAgMt8eBdmxFYGZQEGOAFGVRWEroWQp-AdsKEOn3kzDj3xTc6TPYPTy_FhIBzdMsMhrnjvnpBqWyaT2Ycqjk-3U6DfO9r5EeWGqW6FxKeDHr343fatQMQrejFEO8MOiJc3kQ",
	}
}
