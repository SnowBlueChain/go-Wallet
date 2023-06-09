package xrpgrp

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hiromaily/go-crypto-wallet/pkg/config"
	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/api/xrpgrp/xrp"
	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/coin"
	"github.com/hiromaily/go-crypto-wallet/pkg/ws"
)

// NewWSClient try to connect Ripple Server by web socket
func NewWSClient(conf *config.Ripple) (*ws.WS, *ws.WS, error) {
	publicURL := conf.WebsocketPublicURL
	if publicURL == "" {
		if publicURL = xrp.GetPublicWSServer(conf.NetworkType).String(); publicURL == "" {
			return nil, nil, errors.New("websocket URL is not found")
		}
	}
	public, err := ws.New(context.Background(), publicURL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "fail to call ws.New() for public API: %s", publicURL)
	}

	// acceptable without adminClient
	adminURL := conf.WebsocketAdminURL
	if adminURL == "" {
		return public, nil, nil
	}
	admin, err := ws.New(context.Background(), adminURL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "fail to call ws.New() for admin API: %s", adminURL)
	}

	return public, admin, nil
}

// NewGRPCClient try to connect gRPC Server
func NewGRPCClient(conf *config.RippleAPI) (*grpc.ClientConn, error) {
	if conf.URL == "" {
		return nil, errors.New("url for grpc server is not defined in config")
	}
	var opts []grpc.DialOption
	if !conf.IsSecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial(conf.URL, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to call grpc.Dial: %s", conf.URL)
	}
	return conn, nil
}

// NewRPCClient RPCClient, maybe not used
//func NewRPCClient(conf *config.Ripple) *jsonrpc.RPCClient {
//	if conf.JSONRpcURL == "" {
//		return nil
//	}
//	rpcClient := jsonrpc.NewClient(conf.JSONRpcURL)
//	return &rpcClient
//}

// NewRipple creates Ripple instance according to coinType
func NewRipple(wsPublic *ws.WS, wsAdmin *ws.WS, api *xrp.RippleAPI, conf *config.Ripple, logger *zap.Logger, coinTypeCode coin.CoinTypeCode) (Rippler, error) {
	switch coinTypeCode {
	case coin.XRP:
		ripple, err := xrp.NewRipple(context.Background(), wsPublic, wsAdmin, api, coinTypeCode, conf, logger)
		if err != nil {
			return nil, errors.Wrap(err, "fail to call xrp.NewRipple()")
		}
		return ripple, err
	}
	return nil, errors.Errorf("coinType %s is not defined", coinTypeCode.String())
}
