package dns_records

var TTLs = []int{
	120, 180, 300, 600, 900, 1800, 3600, 7200, 18000, 43200, 86400, 172800, 432000,
}

var UpstreamHttps = []string{
	"default", "auto", "http", "https",
}

var IPFilterCount = []string{
	"single", "multi",
}

var IPFilterOrder = []string{
	"none", "weighted", "rr",
}

var IPFilterGeoFilter = []string{
	"none", "location", "country",
}

type ARecord struct {
	Type          string       `json:"type"`
	Name          string       `json:"name"`
	TTL           int          `json:"ttl"`
	Cloud         bool         `json:"cloud"`
	UpstreamHttps string       `json:"upstream_https"`
	Value         []Value      `json:"value"`
	IpFilterMode  IpFilterMode `json:"ip_filter_mode"`
}

type AAAARecord struct {
	IP      string
	Country string
	Port    int
	Weight  int
}

type IpFilterMode struct {
	Count     string `json:"count"`
	Order     string `json:"order"`
	GeoFilter string `json:"geo_filter"`
}

type Value struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Port    int    `json:"port"`
	Weight  int    `json:"weight"`
}

type MXRecord struct {
	Host     string
	Priority int
}

type NSRecord struct {
	Host string
}

type SRVRecord struct {
	TargetHost string
	Port       int
	Weight     int
	Priority   int
}

type TXTRecord struct {
	Text string
}

type SPFRecord struct {
	Text string
}

type DKIMRecord struct {
	Text string
}

type ANAMERecord struct {
	Location   string
	HostHeader string
}

type CNAMERecord struct {
	Location   string
	HostHeader string
}

type PTRRecord struct {
	Domain string
}
