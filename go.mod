module github.com/lightninglabs/lightning-terminal

require (
	github.com/btcsuite/btcd v0.22.0-beta.0.20211005184431-e3449998be39
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcutil v1.0.3-0.20210527170813-e2ba6805a890
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/go-errors/errors v1.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/improbable-eng/grpc-web v0.12.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/lightninglabs/faraday v0.2.7-alpha
	github.com/lightninglabs/lightning-node-connect v0.1.7-alpha.0.20220215190639-abe533aa98b8
	github.com/lightninglabs/lndclient v0.14.0-7
	github.com/lightninglabs/loop v0.15.1-beta
	github.com/lightninglabs/pool v0.5.4-alpha.0.20220114202858-525fe156d240
	github.com/lightninglabs/pool/auctioneerrpc v1.0.5
	github.com/lightninglabs/protobuf-hex-display v1.4.3-hex-display
	github.com/lightningnetwork/lnd v0.14.2-beta
	github.com/lightningnetwork/lnd/cert v1.1.0
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/mwitkow/grpc-proxy v0.0.0-20181017164139-0f1106ef9c76
	github.com/rs/cors v1.7.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.4
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/macaroon-bakery.v2 v2.1.0
	gopkg.in/macaroon.v2 v2.1.0
)

replace github.com/lightninglabs/lightning-node-connect => github.com/ellemouton/lightning-node-connect v0.1.5-alpha.0.20220321121554-214be2e1be66

go 1.16
