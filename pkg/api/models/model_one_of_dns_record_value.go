/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

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
