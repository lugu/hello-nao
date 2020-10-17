// Package main contains a generated proxy
// .

package main

import (
	"context"
	"fmt"
	bus "github.com/lugu/qiloop/bus"
)

// ALBehaviorManagerProxy represents a proxy object to the service
type ALBehaviorManagerProxy interface {
	StartBehavior(behavior string) error
	StopAllBehaviors() error
	GetBehaviorNames() ([]string, error)
	GetBehaviorsByTag(tag string) ([]string, error)
	GetTagList() ([]string, error)
	// Generic methods shared by all objectsProxy
	bus.ObjectProxy
	// WithContext can be used cancellation and timeout
	WithContext(ctx context.Context) ALBehaviorManagerProxy
}

// proxyALBehaviorManager implements ALBehaviorManagerProxy
type proxyALBehaviorManager struct {
	bus.ObjectProxy
	session bus.Session
}

// MakeALBehaviorManager returns a specialized proxy.
func MakeALBehaviorManager(sess bus.Session, proxy bus.Proxy) ALBehaviorManagerProxy {
	return &proxyALBehaviorManager{bus.MakeObject(proxy), sess}
}

// ALBehaviorManager returns a proxy to a remote service
func ALBehaviorManager(session bus.Session) (ALBehaviorManagerProxy, error) {
	proxy, err := session.Proxy("ALBehaviorManager", 1)
	if err != nil {
		return nil, fmt.Errorf("contact service: %s", err)
	}
	return MakeALBehaviorManager(session, proxy), nil
}

// WithContext bound future calls to the context deadline and cancellation
func (p *proxyALBehaviorManager) WithContext(ctx context.Context) ALBehaviorManagerProxy {
	return MakeALBehaviorManager(p.session, p.Proxy().WithContext(ctx))
}

// StartBehavior calls the remote procedure
func (p *proxyALBehaviorManager) StartBehavior(behavior string) error {
	var ret struct{}
	args := bus.NewParams("(s)", behavior)
	resp := bus.NewResponse("v", &ret)
	err := p.Proxy().Call2("startBehavior", args, resp)
	if err != nil {
		return fmt.Errorf("call startBehavior failed: %s", err)
	}
	return nil
}

// StopAllBehaviors calls the remote procedure
func (p *proxyALBehaviorManager) StopAllBehaviors() error {
	var ret struct{}
	args := bus.NewParams("()")
	resp := bus.NewResponse("v", &ret)
	err := p.Proxy().Call2("stopAllBehaviors", args, resp)
	if err != nil {
		return fmt.Errorf("call stopAllBehaviors failed: %s", err)
	}
	return nil
}

// GetBehaviorNames calls the remote procedure
func (p *proxyALBehaviorManager) GetBehaviorNames() ([]string, error) {
	var ret []string
	args := bus.NewParams("()")
	resp := bus.NewResponse("[s]", &ret)
	err := p.Proxy().Call2("getBehaviorNames", args, resp)
	if err != nil {
		return ret, fmt.Errorf("call getBehaviorNames failed: %s", err)
	}
	return ret, nil
}

// GetBehaviorsByTag calls the remote procedure
func (p *proxyALBehaviorManager) GetBehaviorsByTag(tag string) ([]string, error) {
	var ret []string
	args := bus.NewParams("(s)", tag)
	resp := bus.NewResponse("[s]", &ret)
	err := p.Proxy().Call2("getBehaviorsByTag", args, resp)
	if err != nil {
		return ret, fmt.Errorf("call getBehaviorsByTag failed: %s", err)
	}
	return ret, nil
}

// GetTagList calls the remote procedure
func (p *proxyALBehaviorManager) GetTagList() ([]string, error) {
	var ret []string
	args := bus.NewParams("()")
	resp := bus.NewResponse("[s]", &ret)
	err := p.Proxy().Call2("getTagList", args, resp)
	if err != nil {
		return ret, fmt.Errorf("call getTagList failed: %s", err)
	}
	return ret, nil
}
