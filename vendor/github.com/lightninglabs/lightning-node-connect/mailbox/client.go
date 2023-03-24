package mailbox

import (
	"bytes"
	"context"
	"net"
	"sync"
)

// Client manages the mailboxConn it holds and refreshes it on connection
// retries.
type Client struct {
	mailboxConn *ClientConn

	connData *ConnData

	status   ClientStatus
	statusMu sync.Mutex

	sid [64]byte

	ctx context.Context
}

// NewClient creates a new Client object which will handle the mailbox
// connection.
func NewClient(ctx context.Context, connData *ConnData) (*Client, error) {
	sid, err := connData.SID()
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:      ctx,
		connData: connData,
		status:   ClientStatusNotConnected,
		sid:      sid,
	}, nil
}

// Dial returns a net.Conn abstraction over the mailbox connection. Dial is
// called everytime grpc retries the connection. If this is the first
// connection, a new ClientConn will be created. Otherwise, the existing
// connection will just be refreshed.
func (c *Client) Dial(_ context.Context, serverHost string) (net.Conn, error) {
	// If there is currently an active connection, block here until the
	// previous connection as been closed.
	if c.mailboxConn != nil {
		log.Debugf("Dial: have existing mailbox connection, waiting")
		<-c.mailboxConn.Done()
		log.Debugf("Dial: done with existing conn")
	}

	log.Debugf("Client: Dialing...")

	sid, err := c.connData.SID()
	if err != nil {
		return nil, err
	}

	// If the SID has changed from what it was previously, then we close any
	// previous connection we had.
	if !bytes.Equal(c.sid[:], sid[:]) && c.mailboxConn != nil {
		err := c.mailboxConn.Close()
		if err != nil {
			log.Errorf("could not close mailbox conn: %v", err)
		}

		c.mailboxConn = nil
	}

	c.sid = sid

	if c.mailboxConn == nil {
		mailboxConn, err := NewClientConn(
			c.ctx, c.sid, serverHost, func(status ClientStatus) {
				c.statusMu.Lock()
				c.status = status
				c.statusMu.Unlock()
			},
		)
		if err != nil {
			return nil, &temporaryError{err}
		}
		c.mailboxConn = mailboxConn
	} else {
		mailboxConn, err := RefreshClientConn(c.ctx, c.mailboxConn)
		if err != nil {
			return nil, &temporaryError{err}
		}
		c.mailboxConn = mailboxConn
	}

	return c.mailboxConn, nil
}

// ConnStatus returns last determined client connection status.
func (c *Client) ConnStatus() ClientStatus {
	c.statusMu.Lock()
	defer c.statusMu.Unlock()

	return c.status
}
