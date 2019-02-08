package module

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/linkai-io/am/am"
	"github.com/linkai-io/am/pkg/convert"
	"github.com/linkai-io/am/pkg/retrier"
	service "github.com/linkai-io/am/protocservices/module"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type Config struct {
	ModuleType am.ModuleType
	Timeout    int
}

type Client struct {
	client         service.ModuleClient
	conn           *grpc.ClientConn
	defaultTimeout time.Duration
	config         *Config
	key            string
}

func New() *Client {
	return &Client{defaultTimeout: (time.Second * 60)}
}

func (c *Client) Init(data []byte) error {
	var err error
	c.config, err = c.parseConfig(data)
	if err != nil {
		return err
	}

	if c.config.Timeout != 0 {
		c.defaultTimeout = (time.Second * time.Duration(c.config.Timeout))
	}

	c.key = am.KeyFromModuleType(c.config.ModuleType)
	if c.key == "" {
		return errors.New("unknown module type passed to init")
	}

	conn, err := grpc.DialContext(context.Background(), "srv://consul/"+c.key, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = service.NewModuleClient(conn)
	return nil
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.defaultTimeout = timeout
}

func (c *Client) parseConfig(data []byte) (*Config, error) {
	config := &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Client) Analyze(ctx context.Context, userContext am.UserContext, address *am.ScanGroupAddress) (*am.ScanGroupAddress, map[string]*am.ScanGroupAddress, error) {
	var err error
	var resp *service.AnalyzedResponse
	in := &service.AnalyzeRequest{
		UserContext: convert.DomainToUserContext(userContext),
		Address:     convert.DomainToAddress(address),
	}

	ctxDeadline, cancel := context.WithTimeout(context.Background(), c.defaultTimeout)
	defer cancel()

	err = retrier.RetryIfNot(func() error {
		var retryErr error

		resp, retryErr = c.client.Analyze(ctxDeadline, in)
		if retryErr != nil {
			log.Warn().Str("client", c.key).Err(retryErr).Msg("module analyze returned error")
		}
		return retryErr
	}, "rpc error: code = Unavailable desc")

	if err != nil {
		return nil, nil, err
	}

	addrs := make(map[string]*am.ScanGroupAddress, len(resp.Addresses))
	for key, val := range resp.Addresses {
		addrs[key] = convert.AddressToDomain(val)
	}

	return convert.AddressToDomain(resp.Original), addrs, nil
}
