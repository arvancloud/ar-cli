package dns_records

import (
	"encoding/json"
	"github.com/ebrahimahmadi/ar-cli/pkg/helpers"
	"github.com/ebrahimahmadi/ar-cli/pkg/validator"
)

func (record *ARecord) Build() ([]byte, error) {
	record.
		setType().
		validateIPFilterCount().
		validateIPFilterOrder().
		validateIPGeoFilter().
		validateTTL().
		validateUpstreamHttps()

	return json.MarshalIndent(&record, "", " ")
}

func (record *ARecord) setType() *ARecord {
	record.Type = "a"
	return record
}

func (record *ARecord) validateTTL() *ARecord {
	if ok, tlsValidation := validator.HasInt(record.TTL, TTLs); !ok {
		err := helpers.ToBeColored{Expression: tlsValidation.Error()}
		err.StdoutError().StopExecution()
	}

	return record
}

func (record *ARecord) validateUpstreamHttps() *ARecord {
	if ok, tlsValidation := validator.HasString(record.UpstreamHttps, UpstreamHttps); !ok {
		err := helpers.ToBeColored{Expression: tlsValidation.Error()}
		err.StdoutError().StopExecution()
	}
	return record
}

func (record *ARecord) validateIPFilterCount() *ARecord {
	if ok, tlsValidation := validator.HasString(record.IpFilterMode.Count, IPFilterCount); !ok {
		err := helpers.ToBeColored{Expression: tlsValidation.Error()}
		err.StdoutError().StopExecution()
	}
	return record
}

func (record *ARecord) validateIPFilterOrder() *ARecord {
	if ok, tlsValidation := validator.HasString(record.IpFilterMode.Order, IPFilterOrder); !ok {
		err := helpers.ToBeColored{Expression: tlsValidation.Error()}
		err.StdoutError().StopExecution()
	}
	return record
}

func (record *ARecord) validateIPGeoFilter() *ARecord {
	if ok, tlsValidation := validator.HasString(record.IpFilterMode.GeoFilter, IPFilterGeoFilter); !ok {
		err := helpers.ToBeColored{Expression: tlsValidation.Error()}
		err.StdoutError().StopExecution()
	}
	return record
}
