package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetBTBPrivKeySignature() string {
	// TODO: get btb key signature from btb-mdm
	key := "P0l1t3kN1k_c1L4C4p"

	return key
}

const urlIntegration = "integrations.externals.http.%s.endpoints.%s"
const host = "integrations.externals.http.%s.host"

func GetIntegURL(integration string, name string) string {
	getHost := fmt.Sprintf(host, integration)
	getString := fmt.Sprintf(urlIntegration, integration, name)
	return viper.GetString(getHost) + viper.GetString(getString)
}

// event driven
// api driven
// db transaction

// signature = token = api driven
// api gateway
// bagusnya secret key di db
// api contract = aturan contract

// efisiensi code
// belajar consum api
