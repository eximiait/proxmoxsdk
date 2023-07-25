package proxmoxsdk

import (
	"errors"
	"net/url"

	proxmox_api_go "github.com/Telmate/proxmox-api-go/proxmox"
)

// LogFunc represents a flexiable and injectable logger function which fits to most of logger libraries
type LogFunc func(format string, v ...interface{})

// Connection represents an HTTP connection to the engine server.
// It is intended as the entry point for the SDK, and it provides access to the `system` service and, from there,
// to the rest of the services provided by the API.
type Connection struct {
	url      *url.URL
	username string
	password string
	// Debug options
	logFunc LogFunc
	timeout int
	// proxmox_api_go.Client client
	client  *proxmox_api_go.Client
	headers map[string]string
}

// URL returns the base URL of this connection.
func (c *Connection) URL() string {
	return c.url.String()
}

// Test tests the connectivity with the server using the system service.
func (c *Connection) Test() error {
	/*_, err := c.SystemService().Get().Send()
	if err != nil {
		return fmt.Errorf("failed to validate the connection (%w)", err)
	}*/
	return nil
}

// authenticate uses OAuth to do authentication
func (c *Connection) authenticate() (string, error) {
	/*if c.ssoToken == "" {
		token, err := c.getAccessToken()
		if err != nil {
			return "", err
		}
		c.ssoToken = token
	}*/
	//return c.ssoToken, nil
	return "", nil
}

// Close releases the resources used by this connection.
func (c *Connection) Close() error {
	//return c.CloseIfRevokeSSOToken(true)
	return nil
}

// NewConnectionBuilder creates the `ConnectionBuilder struct instance
func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{
		conn: &Connection{},
		err:  nil}
}

// ConnectionBuilder represents a builder for the `Connection` struct
type ConnectionBuilder struct {
	conn *Connection
	err  error
}

// URL sets the url field for `Connection` instance
func (connBuilder *ConnectionBuilder) URL(urlStr string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	// Save the URL:
	useURL, err := url.Parse(urlStr)
	if err != nil {
		connBuilder.err = err
		return connBuilder
	}
	connBuilder.conn.url = useURL
	return connBuilder
}

// Username sets the username field for `Connection` instance
func (connBuilder *ConnectionBuilder) Username(username string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.username = username
	return connBuilder
}

// Password sets the password field for `Connection` instance
func (connBuilder *ConnectionBuilder) Password(password string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	connBuilder.conn.password = password
	return connBuilder
}

// LogFunc sets the logging function field for `Connection` instance
func (connBuilder *ConnectionBuilder) LogFunc(logFunc LogFunc) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.logFunc = logFunc
	return connBuilder
}

// Timeout sets the timeout field for `Connection` instance
func (connBuilder *ConnectionBuilder) Timeout(timeout int) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}
	connBuilder.conn.timeout = timeout
	return connBuilder
}

// Headers sets a map of custom HTTP headers to be added to each HTTP request
func (connBuilder *ConnectionBuilder) Headers(headers map[string]string) *ConnectionBuilder {
	// If already has errors, just return
	if connBuilder.err != nil {
		return connBuilder
	}

	if connBuilder.conn.headers == nil {
		connBuilder.conn.headers = map[string]string{}
	}

	for hk, hv := range headers {
		connBuilder.conn.headers[hk] = hv
	}
	return connBuilder
}

// Build constructs the `Connection` instance
func (connBuilder *ConnectionBuilder) Build() (*Connection, error) {
	// If already has errors, just return
	if connBuilder.err != nil {
		return nil, connBuilder.err
	}

	// Check parameters
	if connBuilder.conn.url == nil {
		return nil, errors.New("the URL must not be empty")
	}
	if len(connBuilder.conn.username) == 0 {
		return nil, errors.New("the username must not be empty")
	}
	if len(connBuilder.conn.password) == 0 {
		return nil, errors.New("the password must not be empty")
	}

	connBuilder.conn.client = &proxmox_api_go.Client{
		Username:    connBuilder.conn.username,
		Password:    connBuilder.conn.password,
		ApiUrl:      connBuilder.conn.url.Opaque,
		TaskTimeout: connBuilder.conn.client.TaskTimeout,
	}
	return connBuilder.conn, nil
}
