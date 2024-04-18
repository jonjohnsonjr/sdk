/*
Copyright 2021 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package v1

import (
	"context"
	"fmt"
	"net/url"
	"time"

	delegate "chainguard.dev/go-grpc-kit/pkg/options"
	"github.com/chainguard-dev/clog"
	"google.golang.org/grpc"

	"chainguard.dev/sdk/auth"
)

type Clients interface {
	Records() RecordsClient
	RecordContexts() RecordContextsClient
	Sboms() SbomsClient
	Signatures() SignaturesClient
	PolicyResults() PolicyResultsClient
	VulnReports() VulnReportsClient

	Close() error
}

func NewClients(ctx context.Context, addr string, token string) (Clients, error) {
	uri, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tenant service address, must be a url: %w", err)
	}

	target, opts := delegate.GRPCOptions(*uri)

	// TODO: we may want to require transport security at some future point.
	if cred := auth.NewFromToken(ctx, token, false); cred != nil {
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	} else {
		clog.FromContext(ctx).Warn("No authentication provided, this may end badly.")
	}

	var cancel context.CancelFunc
	if _, timeoutSet := ctx.Deadline(); !timeoutSet {
		ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
		defer cancel()
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the iam server: %w", err)
	}

	return &clients{
		records:        NewRecordsClient(conn),
		recordContexts: NewRecordContextsClient(conn),
		sboms:          NewSbomsClient(conn),
		vulnReports:    NewVulnReportsClient(conn),
		signatures:     NewSignaturesClient(conn),
		policyResults:  NewPolicyResultsClient(conn),

		conn: conn,
	}, nil
}

func NewClientsFromConnection(conn *grpc.ClientConn) Clients {
	return &clients{
		records:        NewRecordsClient(conn),
		recordContexts: NewRecordContextsClient(conn),
		sboms:          NewSbomsClient(conn),
		vulnReports:    NewVulnReportsClient(conn),
		signatures:     NewSignaturesClient(conn),
		policyResults:  NewPolicyResultsClient(conn),
		// conn is not set, this client struct does not own closing it.
	}
}

type clients struct {
	records RecordsClient

	recordContexts RecordContextsClient
	sboms          SbomsClient
	signatures     SignaturesClient
	policyResults  PolicyResultsClient
	vulnReports    VulnReportsClient

	conn *grpc.ClientConn
}

func (c *clients) Records() RecordsClient {
	return c.records
}

func (c *clients) RecordContexts() RecordContextsClient {
	return c.recordContexts
}

func (c *clients) Sboms() SbomsClient {
	return c.sboms
}

func (c *clients) Signatures() SignaturesClient {
	return c.signatures
}

func (c *clients) PolicyResults() PolicyResultsClient {
	return c.policyResults
}

func (c *clients) VulnReports() VulnReportsClient {
	return c.vulnReports
}

func (c *clients) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
