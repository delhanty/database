package am

import (
	"context"
)

const (
	RNWebData             = "lrn:service:webdata:feature:"
	RNWebDataResponses    = "lrn:service:webdata:feature:responses"
	RNWebDataCertificates = "lrn:service:webdata:feature:certificates"
	RNWebDataSnapshots    = "lrn:service:webdata:feature:snapshots"
	WebDataServiceKey     = "webdataservice"
)

// HTTPResponse represents a captured network response
type HTTPResponse struct {
	Scheme            string                 `json:"scheme"`
	HostAddress       string                 `json:"host_address"`
	IPAddress         string                 `json:"ip_address"`
	ResponsePort      string                 `json:"response_port"`
	RequestedPort     string                 `json:"requested_port"`
	RequestID         string                 `json:"request_id,omitempty"`
	Status            int                    `json:"status"`
	StatusText        string                 `json:"status_text"`
	URL               string                 `json:"url"`
	Headers           map[string]interface{} `json:"headers"`
	MimeType          string                 `json:"mime_type"`
	RawBody           string                 `json:"raw_body,omitempty"`
	RawBodyLink       string                 `json:"raw_body_link"`
	RawBodyHash       string                 `json:"raw_body_hash"`
	ResponseTimestamp int64                  `json:"response_timestamp"`
	IsDocument        bool                   `json:"is_document"`
	WebCertificate    *WebCertificate        `json:"web_certificate,omitempty"`
}

// WebCertificate contains extracted certificate data from an HTTPResponse
type WebCertificate struct {
	Protocol                          string   `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string   `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  string   `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                            string   `json:"cipher"`                            // Cipher name.
	Mac                               string   `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId                     int      `json:"certificateId"`                     // Certificate ID value.
	SubjectName                       string   `json:"subjectName"`                       // Certificate subject name.
	SanList                           []string `json:"sanList"`                           // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string   `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                         int64    `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                           int64    `json:"validTo"`                           // Certificate valid to (expiration) date
	CertificateTransparencyCompliance string   `json:"certificateTransparencyCompliance"` // Whether the request complied with Certificate Transparency policy enum values: unknown, not-compliant, compliant
}

// WebData is the primary container of a scangroup address's http response data.
type WebData struct {
	Address           *ScanGroupAddress `json:"address"`
	Responses         []*HTTPResponse   `json:"responses"`
	Snapshot          string            `json:"snapshot,omitempty"`
	SnapshotLink      string            `json:"snapshot_link"`
	SerializedDOM     string            `json:"serialized_dom,omitempty"`
	SerializedDOMHash string            `json:"serialized_dom_hash"`
	SerializedDOMLink string            `json:"serialized_dom_link"`
	ResponseTimestamp int64             `json:"response_timeestamp"`
}

// WebSnapshot for returning serialized dom and image snapshot links
type WebSnapshot struct {
	OrgID             int    `json:"org_id"`
	GroupID           int    `json:"group_id"`
	SnapshotLink      string `json:"snapshot_link"`
	SerializedDOMLink string `json:"serialized_dom_link"`
	ResponseTimestamp int64  `json:"response_timeestamp"`
}

// WebResponseFilter used to filter results when searching web data.
type WebResponseFilter struct {
	OrgID             int   `json:"org_id"`
	GroupID           int   `json:"group_id"`
	WithResponseTime  bool  `json:"with_response_time"`
	SinceResponseTime int64 `json:"since_response_time"`
	Start             int64 `json:"start"`
	Limit             int   `json:"limit"`
}

// WebCertificateFilter used to filter results when searching web data.
type WebCertificateFilter struct {
	OrgID             int   `json:"org_id"`
	GroupID           int   `json:"group_id"`
	WithResponseTime  bool  `json:"with_response_time"`
	SinceResponseTime int64 `json:"since_response_time"`
	WithValidTo       bool  `json:"with_valid_to"`
	ValidToTime       int64 `json:"valid_to_time"`
	Start             int64 `json:"start"`
	Limit             int   `json:"limit"`
}

type WebSnapshotFilter struct {
	OrgID             int   `json:"org_id"`
	GroupID           int   `json:"group_id"`
	WithResponseTime  bool  `json:"with_response_time"`
	SinceResponseTime int64 `json:"since_response_time"`
	Start             int64 `json:"start"`
	Limit             int   `json:"limit"`
}

type WebDataService interface {
	Init(config []byte) error
	IsAuthorized(ctx context.Context, userContext UserContext, resource, action string) bool
	Update(ctx context.Context, userContext UserContext, webData *WebData) (int, error)
	GetResponses(ctx context.Context, userContext UserContext, filter *WebResponseFilter) (int, []*HTTPResponse, error)
	GetCertificates(ctx context.Context, userContext UserContext, filter *WebCertificateFilter) (int, []*WebCertificate, error)
	GetSnapshots(ctx context.Context, userContext UserContext, filter *WebSnapshotFilter) (int, []*WebSnapshot, error)
}
