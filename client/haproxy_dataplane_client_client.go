// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/practical-coder/hdc/client/acl"
	"github.com/practical-coder/hdc/client/acl_runtime"
	"github.com/practical-coder/hdc/client/backend"
	"github.com/practical-coder/hdc/client/backend_switching_rule"
	"github.com/practical-coder/hdc/client/bind"
	"github.com/practical-coder/hdc/client/cache"
	"github.com/practical-coder/hdc/client/cluster"
	"github.com/practical-coder/hdc/client/configuration"
	"github.com/practical-coder/hdc/client/declare_capture"
	"github.com/practical-coder/hdc/client/defaults"
	"github.com/practical-coder/hdc/client/dgram_bind"
	"github.com/practical-coder/hdc/client/discovery"
	"github.com/practical-coder/hdc/client/filter"
	"github.com/practical-coder/hdc/client/frontend"
	"github.com/practical-coder/hdc/client/global"
	"github.com/practical-coder/hdc/client/group"
	"github.com/practical-coder/hdc/client/health"
	"github.com/practical-coder/hdc/client/http_after_response_rule"
	"github.com/practical-coder/hdc/client/http_check"
	"github.com/practical-coder/hdc/client/http_error_rule"
	"github.com/practical-coder/hdc/client/http_errors"
	"github.com/practical-coder/hdc/client/http_request_rule"
	"github.com/practical-coder/hdc/client/http_response_rule"
	"github.com/practical-coder/hdc/client/information"
	"github.com/practical-coder/hdc/client/log_forward"
	"github.com/practical-coder/hdc/client/log_target"
	"github.com/practical-coder/hdc/client/mailer_entry"
	"github.com/practical-coder/hdc/client/mailers"
	"github.com/practical-coder/hdc/client/maps"
	"github.com/practical-coder/hdc/client/nameserver"
	"github.com/practical-coder/hdc/client/peer"
	"github.com/practical-coder/hdc/client/peer_entry"
	"github.com/practical-coder/hdc/client/process_manager"
	"github.com/practical-coder/hdc/client/reloads"
	"github.com/practical-coder/hdc/client/resolver"
	"github.com/practical-coder/hdc/client/ring"
	serverops "github.com/practical-coder/hdc/client/server"
	"github.com/practical-coder/hdc/client/server_switching_rule"
	"github.com/practical-coder/hdc/client/server_template"
	"github.com/practical-coder/hdc/client/service_discovery"
	"github.com/practical-coder/hdc/client/sites"
	"github.com/practical-coder/hdc/client/specification"
	"github.com/practical-coder/hdc/client/specification_openapiv3"
	"github.com/practical-coder/hdc/client/spoe"
	"github.com/practical-coder/hdc/client/spoe_transactions"
	"github.com/practical-coder/hdc/client/stats"
	"github.com/practical-coder/hdc/client/stick_rule"
	"github.com/practical-coder/hdc/client/stick_table"
	"github.com/practical-coder/hdc/client/storage"
	"github.com/practical-coder/hdc/client/tcp_check"
	"github.com/practical-coder/hdc/client/tcp_request_rule"
	"github.com/practical-coder/hdc/client/tcp_response_rule"
	"github.com/practical-coder/hdc/client/transactions"
	"github.com/practical-coder/hdc/client/user"
	"github.com/practical-coder/hdc/client/userlist"
)

// Default haproxy dataplane client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new haproxy dataplane client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *HaproxyDataplaneClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new haproxy dataplane client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *HaproxyDataplaneClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new haproxy dataplane client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *HaproxyDataplaneClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(HaproxyDataplaneClient)
	cli.Transport = transport
	cli.ACL = acl.New(transport, formats)
	cli.ACLRuntime = acl_runtime.New(transport, formats)
	cli.Backend = backend.New(transport, formats)
	cli.BackendSwitchingRule = backend_switching_rule.New(transport, formats)
	cli.Bind = bind.New(transport, formats)
	cli.Cache = cache.New(transport, formats)
	cli.Cluster = cluster.New(transport, formats)
	cli.Configuration = configuration.New(transport, formats)
	cli.DeclareCapture = declare_capture.New(transport, formats)
	cli.Defaults = defaults.New(transport, formats)
	cli.DgramBind = dgram_bind.New(transport, formats)
	cli.Discovery = discovery.New(transport, formats)
	cli.Filter = filter.New(transport, formats)
	cli.Frontend = frontend.New(transport, formats)
	cli.Global = global.New(transport, formats)
	cli.Group = group.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.HTTPAfterResponseRule = http_after_response_rule.New(transport, formats)
	cli.HTTPCheck = http_check.New(transport, formats)
	cli.HTTPErrorRule = http_error_rule.New(transport, formats)
	cli.HTTPErrors = http_errors.New(transport, formats)
	cli.HTTPRequestRule = http_request_rule.New(transport, formats)
	cli.HTTPResponseRule = http_response_rule.New(transport, formats)
	cli.Information = information.New(transport, formats)
	cli.LogForward = log_forward.New(transport, formats)
	cli.LogTarget = log_target.New(transport, formats)
	cli.MailerEntry = mailer_entry.New(transport, formats)
	cli.Mailers = mailers.New(transport, formats)
	cli.Maps = maps.New(transport, formats)
	cli.Nameserver = nameserver.New(transport, formats)
	cli.Peer = peer.New(transport, formats)
	cli.PeerEntry = peer_entry.New(transport, formats)
	cli.ProcessManager = process_manager.New(transport, formats)
	cli.Reloads = reloads.New(transport, formats)
	cli.Resolver = resolver.New(transport, formats)
	cli.Ring = ring.New(transport, formats)
	cli.Server = serverops.New(transport, formats)
	cli.ServerSwitchingRule = server_switching_rule.New(transport, formats)
	cli.ServerTemplate = server_template.New(transport, formats)
	cli.ServiceDiscovery = service_discovery.New(transport, formats)
	cli.Sites = sites.New(transport, formats)
	cli.Specification = specification.New(transport, formats)
	cli.SpecificationOpenapiv3 = specification_openapiv3.New(transport, formats)
	cli.Spoe = spoe.New(transport, formats)
	cli.SpoeTransactions = spoe_transactions.New(transport, formats)
	cli.Stats = stats.New(transport, formats)
	cli.StickRule = stick_rule.New(transport, formats)
	cli.StickTable = stick_table.New(transport, formats)
	cli.Storage = storage.New(transport, formats)
	cli.TCPCheck = tcp_check.New(transport, formats)
	cli.TCPRequestRule = tcp_request_rule.New(transport, formats)
	cli.TCPResponseRule = tcp_response_rule.New(transport, formats)
	cli.Transactions = transactions.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.Userlist = userlist.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// HaproxyDataplaneClient is a client for haproxy dataplane client
type HaproxyDataplaneClient struct {
	ACL acl.ClientService

	ACLRuntime acl_runtime.ClientService

	Backend backend.ClientService

	BackendSwitchingRule backend_switching_rule.ClientService

	Bind bind.ClientService

	Cache cache.ClientService

	Cluster cluster.ClientService

	Configuration configuration.ClientService

	DeclareCapture declare_capture.ClientService

	Defaults defaults.ClientService

	DgramBind dgram_bind.ClientService

	Discovery discovery.ClientService

	Filter filter.ClientService

	Frontend frontend.ClientService

	Global global.ClientService

	Group group.ClientService

	Health health.ClientService

	HTTPAfterResponseRule http_after_response_rule.ClientService

	HTTPCheck http_check.ClientService

	HTTPErrorRule http_error_rule.ClientService

	HTTPErrors http_errors.ClientService

	HTTPRequestRule http_request_rule.ClientService

	HTTPResponseRule http_response_rule.ClientService

	Information information.ClientService

	LogForward log_forward.ClientService

	LogTarget log_target.ClientService

	MailerEntry mailer_entry.ClientService

	Mailers mailers.ClientService

	Maps maps.ClientService

	Nameserver nameserver.ClientService

	Peer peer.ClientService

	PeerEntry peer_entry.ClientService

	ProcessManager process_manager.ClientService

	Reloads reloads.ClientService

	Resolver resolver.ClientService

	Ring ring.ClientService

	Server serverops.ClientService

	ServerSwitchingRule server_switching_rule.ClientService

	ServerTemplate server_template.ClientService

	ServiceDiscovery service_discovery.ClientService

	Sites sites.ClientService

	Specification specification.ClientService

	SpecificationOpenapiv3 specification_openapiv3.ClientService

	Spoe spoe.ClientService

	SpoeTransactions spoe_transactions.ClientService

	Stats stats.ClientService

	StickRule stick_rule.ClientService

	StickTable stick_table.ClientService

	Storage storage.ClientService

	TCPCheck tcp_check.ClientService

	TCPRequestRule tcp_request_rule.ClientService

	TCPResponseRule tcp_response_rule.ClientService

	Transactions transactions.ClientService

	User user.ClientService

	Userlist userlist.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *HaproxyDataplaneClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ACL.SetTransport(transport)
	c.ACLRuntime.SetTransport(transport)
	c.Backend.SetTransport(transport)
	c.BackendSwitchingRule.SetTransport(transport)
	c.Bind.SetTransport(transport)
	c.Cache.SetTransport(transport)
	c.Cluster.SetTransport(transport)
	c.Configuration.SetTransport(transport)
	c.DeclareCapture.SetTransport(transport)
	c.Defaults.SetTransport(transport)
	c.DgramBind.SetTransport(transport)
	c.Discovery.SetTransport(transport)
	c.Filter.SetTransport(transport)
	c.Frontend.SetTransport(transport)
	c.Global.SetTransport(transport)
	c.Group.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.HTTPAfterResponseRule.SetTransport(transport)
	c.HTTPCheck.SetTransport(transport)
	c.HTTPErrorRule.SetTransport(transport)
	c.HTTPErrors.SetTransport(transport)
	c.HTTPRequestRule.SetTransport(transport)
	c.HTTPResponseRule.SetTransport(transport)
	c.Information.SetTransport(transport)
	c.LogForward.SetTransport(transport)
	c.LogTarget.SetTransport(transport)
	c.MailerEntry.SetTransport(transport)
	c.Mailers.SetTransport(transport)
	c.Maps.SetTransport(transport)
	c.Nameserver.SetTransport(transport)
	c.Peer.SetTransport(transport)
	c.PeerEntry.SetTransport(transport)
	c.ProcessManager.SetTransport(transport)
	c.Reloads.SetTransport(transport)
	c.Resolver.SetTransport(transport)
	c.Ring.SetTransport(transport)
	c.Server.SetTransport(transport)
	c.ServerSwitchingRule.SetTransport(transport)
	c.ServerTemplate.SetTransport(transport)
	c.ServiceDiscovery.SetTransport(transport)
	c.Sites.SetTransport(transport)
	c.Specification.SetTransport(transport)
	c.SpecificationOpenapiv3.SetTransport(transport)
	c.Spoe.SetTransport(transport)
	c.SpoeTransactions.SetTransport(transport)
	c.Stats.SetTransport(transport)
	c.StickRule.SetTransport(transport)
	c.StickTable.SetTransport(transport)
	c.Storage.SetTransport(transport)
	c.TCPCheck.SetTransport(transport)
	c.TCPRequestRule.SetTransport(transport)
	c.TCPResponseRule.SetTransport(transport)
	c.Transactions.SetTransport(transport)
	c.User.SetTransport(transport)
	c.Userlist.SetTransport(transport)
}
