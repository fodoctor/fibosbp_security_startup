package config

/**
 * TODO
 * Your bp private key && public key
 */
const (
	BP_ACCOUNT_NAME = "ChainDoctor1"
	BP_PRIVATE_KEY = "5KTb1FyrXuZ6W866hwpAiMQSKZVrifJgBzSg44jH4jowZieZqBN"
	BP_PUBLIC_KEY  = "FO7nzhUjEJUxVhFdusQXMoPvLpDmt62bPMFzH4u56m6ZYzkTVT6z"
)

/**
 * bpconfig instance object
 */
var BPConfigInstance BPConfig

type BPConfig struct {
	FibOSBinPath string
	P2PSeedArray []string
}

func (self *BPConfig) LoadConfig(configFile string) {

}
