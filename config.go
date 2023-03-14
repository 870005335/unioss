package unioss

type Config struct {
	// (必填，对应阿里云access_key_id、腾讯云secret_id，七牛云accessKey)
	KeyId string
	// (必填，对应阿里云access_key_secret、腾讯云secret_key，七牛云secretKey)
	KeySecret string
	Bucket    string
	// 阿里云必填
	Endpoint string
	// 腾讯云必填
	AppId string
	// 腾讯云必填
	Region string
	// 域名
	Domain string
	// 是否私有 (仅七牛云)
	Private bool
}
