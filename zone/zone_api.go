package zone

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Zone Api
func NewZoneApi() ZoneApiI {
	return &zoneApi{}
}

// Zone api interface
type ZoneApiI interface {
	ListZones(string, interface{}) (string, error)
}

// Zone api object
type zoneApi struct {
}

// List Zones
// TODO - Documentation
func (z *zoneApi) ListZones(uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetZoneListApi()
	if uuid != "" {
		endPoint += "?uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
