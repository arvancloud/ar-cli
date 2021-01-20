/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

import (
	"github.com/masihyeganeh/ar-cli/pkg/utl"
)

type OneOfDnsRecordValue struct {
	ARecord
	AaaaRecord
	MxRecord
	NsRecord
	SrvRecord
	TxtRecord
	SpfRecord
	DkimRecord
	AnameRecord
	CnameRecord
	PtrRecord
}

func (o OneOfDnsRecordValue) MarshalJSON() ([]byte, error) {
	return utl.EncodeInner(o.ARecord, o.AaaaRecord, o.MxRecord, o.NsRecord, o.SrvRecord, o.TxtRecord, o.SpfRecord, o.DkimRecord, o.AnameRecord, o.CnameRecord, o.PtrRecord)
}
