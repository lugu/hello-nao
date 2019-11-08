// Package main contains a generated proxy
// .

package main

import (
	"bytes"
	"fmt"
	bus "github.com/lugu/qiloop/bus"
	basic "github.com/lugu/qiloop/type/basic"
	object "github.com/lugu/qiloop/type/object"
)

// Constructor gives access to remote services
type Constructor struct {
	session bus.Session
}

// Services gives access to the services constructor
func Services(s bus.Session) Constructor {
	return Constructor{session: s}
}

// ALBehaviorManager is the abstract interface of the service
type ALBehaviorManager interface {
	// StartBehavior calls the remote procedure
	StartBehavior(behavior string) error
	// StopAllBehaviors calls the remote procedure
	StopAllBehaviors() error
	// GetBehaviorNames calls the remote procedure
	GetBehaviorNames() ([]string, error)
}

// ALBehaviorManagerProxy represents a proxy object to the service
type ALBehaviorManagerProxy interface {
	object.Object
	bus.Proxy
	ALBehaviorManager
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
func (c Constructor) ALBehaviorManager() (ALBehaviorManagerProxy, error) {
	proxy, err := c.session.Proxy("ALBehaviorManager", 1)
	if err != nil {
		return nil, fmt.Errorf("contact service: %s", err)
	}
	return MakeALBehaviorManager(c.session, proxy), nil
}

// StartBehavior calls the remote procedure
func (p *proxyALBehaviorManager) StartBehavior(behavior string) error {
	var err error
	var buf bytes.Buffer
	if err = basic.WriteString(behavior, &buf); err != nil {
		return fmt.Errorf("serialize behavior: %s", err)
	}
	_, err = p.Call("startBehavior", buf.Bytes())
	if err != nil {
		return fmt.Errorf("call startBehavior failed: %s", err)
	}
	return nil
}

// StopAllBehaviors calls the remote procedure
func (p *proxyALBehaviorManager) StopAllBehaviors() error {
	var err error
	var buf bytes.Buffer
	_, err = p.Call("stopAllBehaviors", buf.Bytes())
	if err != nil {
		return fmt.Errorf("call stopAllBehaviors failed: %s", err)
	}
	return nil
}

// GetBehaviorNames calls the remote procedure
func (p *proxyALBehaviorManager) GetBehaviorNames() ([]string, error) {
	var err error
	var ret []string
	var buf bytes.Buffer
	response, err := p.Call("getBehaviorNames", buf.Bytes())
	if err != nil {
		return ret, fmt.Errorf("call getBehaviorNames failed: %s", err)
	}
	resp := bytes.NewBuffer(response)
	ret, err = func() (b []string, err error) {
		size, err := basic.ReadUint32(resp)
		if err != nil {
			return b, fmt.Errorf("read slice size: %s", err)
		}
		b = make([]string, size)
		for i := 0; i < int(size); i++ {
			b[i], err = basic.ReadString(resp)
			if err != nil {
				return b, fmt.Errorf("read slice value: %s", err)
			}
		}
		return b, nil
	}()
	if err != nil {
		return ret, fmt.Errorf("parse getBehaviorNames response: %s", err)
	}
	return ret, nil
}
