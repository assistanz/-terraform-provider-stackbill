package api

var (
	computeOfferingList = "compute/computeOfferingList"
)

// Compute Offering List Api
func GetComputeOfferingListApi(zoneUuid string) string {
	return EndPoint + "/" + computeOfferingList + "?zoneUuid=" + zoneUuid
}
