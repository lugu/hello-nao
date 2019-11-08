// Package main contains a generated proxy
// .

package main

import (
	"bytes"
	"fmt"
	bus "github.com/lugu/qiloop/bus"
	basic "github.com/lugu/qiloop/type/basic"
	object "github.com/lugu/qiloop/type/object"
	"io"
	"log"
)

// Constructor gives access to remote services
type Constructor struct {
	session bus.Session
}

// Services gives access to the services constructor
func Services(s bus.Session) Constructor {
	return Constructor{session: s}
}

// BehaviorFailed is serializable
type BehaviorFailed struct {
	P0 string
	P1 string
	P2 string
}

// readBehaviorFailed unmarshalls BehaviorFailed
func readBehaviorFailed(r io.Reader) (s BehaviorFailed, err error) {
	if s.P0, err = basic.ReadString(r); err != nil {
		return s, fmt.Errorf("read P0 field: %s", err)
	}
	if s.P1, err = basic.ReadString(r); err != nil {
		return s, fmt.Errorf("read P1 field: %s", err)
	}
	if s.P2, err = basic.ReadString(r); err != nil {
		return s, fmt.Errorf("read P2 field: %s", err)
	}
	return s, nil
}

// writeBehaviorFailed marshalls BehaviorFailed
func writeBehaviorFailed(s BehaviorFailed, w io.Writer) (err error) {
	if err := basic.WriteString(s.P0, w); err != nil {
		return fmt.Errorf("write P0 field: %s", err)
	}
	if err := basic.WriteString(s.P1, w); err != nil {
		return fmt.Errorf("write P1 field: %s", err)
	}
	if err := basic.WriteString(s.P2, w); err != nil {
		return fmt.Errorf("write P2 field: %s", err)
	}
	return nil
}

// ALFrameManager is the abstract interface of the service
type ALFrameManager interface {
	// GetMethodList calls the remote procedure
	GetMethodList() ([]string, error)
	// NewBehaviorFromFile calls the remote procedure
	NewBehaviorFromFile(xmlFilePath string, name string) (string, error)
	// PlayBehavior calls the remote procedure
	PlayBehavior(id string) error
	// ExitBehavior calls the remote procedure
	ExitBehavior(id string) error
	// Behaviors calls the remote procedure
	Behaviors() ([]string, error)
	// SubscribeBehaviorPlayed subscribe to a remote signal
	SubscribeBehaviorPlayed() (unsubscribe func(), updates chan string, err error)
	// SubscribeBehaviorStopped subscribe to a remote signal
	SubscribeBehaviorStopped() (unsubscribe func(), updates chan string, err error)
	// SubscribeBehaviorFailed subscribe to a remote signal
	SubscribeBehaviorFailed() (unsubscribe func(), updates chan BehaviorFailed, err error)
}

// ALFrameManagerProxy represents a proxy object to the service
type ALFrameManagerProxy interface {
	object.Object
	bus.Proxy
	ALFrameManager
}

// proxyALFrameManager implements ALFrameManagerProxy
type proxyALFrameManager struct {
	bus.ObjectProxy
	session bus.Session
}

// MakeALFrameManager returns a specialized proxy.
func MakeALFrameManager(sess bus.Session, proxy bus.Proxy) ALFrameManagerProxy {
	return &proxyALFrameManager{bus.MakeObject(proxy), sess}
}

// ALFrameManager returns a proxy to a remote service
func (c Constructor) ALFrameManager() (ALFrameManagerProxy, error) {
	proxy, err := c.session.Proxy("ALFrameManager", 1)
	if err != nil {
		return nil, fmt.Errorf("contact service: %s", err)
	}
	return MakeALFrameManager(c.session, proxy), nil
}

// GetMethodList calls the remote procedure
func (p *proxyALFrameManager) GetMethodList() ([]string, error) {
	var err error
	var ret []string
	var buf bytes.Buffer
	response, err := p.Call("getMethodList", buf.Bytes())
	if err != nil {
		return ret, fmt.Errorf("call getMethodList failed: %s", err)
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
		return ret, fmt.Errorf("parse getMethodList response: %s", err)
	}
	return ret, nil
}

// NewBehaviorFromFile calls the remote procedure
func (p *proxyALFrameManager) NewBehaviorFromFile(xmlFilePath string, name string) (string, error) {
	var err error
	var ret string
	var buf bytes.Buffer
	if err = basic.WriteString(xmlFilePath, &buf); err != nil {
		return ret, fmt.Errorf("serialize xmlFilePath: %s", err)
	}
	if err = basic.WriteString(name, &buf); err != nil {
		return ret, fmt.Errorf("serialize name: %s", err)
	}
	response, err := p.Call("newBehaviorFromFile", buf.Bytes())
	if err != nil {
		return ret, fmt.Errorf("call newBehaviorFromFile failed: %s", err)
	}
	resp := bytes.NewBuffer(response)
	ret, err = basic.ReadString(resp)
	if err != nil {
		return ret, fmt.Errorf("parse newBehaviorFromFile response: %s", err)
	}
	return ret, nil
}

// PlayBehavior calls the remote procedure
func (p *proxyALFrameManager) PlayBehavior(id string) error {
	var err error
	var buf bytes.Buffer
	if err = basic.WriteString(id, &buf); err != nil {
		return fmt.Errorf("serialize id: %s", err)
	}
	_, err = p.Call("playBehavior", buf.Bytes())
	if err != nil {
		return fmt.Errorf("call playBehavior failed: %s", err)
	}
	return nil
}

// ExitBehavior calls the remote procedure
func (p *proxyALFrameManager) ExitBehavior(id string) error {
	var err error
	var buf bytes.Buffer
	if err = basic.WriteString(id, &buf); err != nil {
		return fmt.Errorf("serialize id: %s", err)
	}
	_, err = p.Call("exitBehavior", buf.Bytes())
	if err != nil {
		return fmt.Errorf("call exitBehavior failed: %s", err)
	}
	return nil
}

// Behaviors calls the remote procedure
func (p *proxyALFrameManager) Behaviors() ([]string, error) {
	var err error
	var ret []string
	var buf bytes.Buffer
	response, err := p.Call("behaviors", buf.Bytes())
	if err != nil {
		return ret, fmt.Errorf("call behaviors failed: %s", err)
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
		return ret, fmt.Errorf("parse behaviors response: %s", err)
	}
	return ret, nil
}

// SubscribeBehaviorPlayed subscribe to a remote property
func (p *proxyALFrameManager) SubscribeBehaviorPlayed() (func(), chan string, error) {
	propertyID, err := p.SignalID("behaviorPlayed")
	if err != nil {
		return nil, nil, fmt.Errorf("property %s not available: %s", "behaviorPlayed", err)
	}

	handlerID, err := p.RegisterEvent(p.ObjectID(), propertyID, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("register event for %s: %s", "behaviorPlayed", err)
	}
	ch := make(chan string)
	cancel, chPay, err := p.SubscribeID(propertyID)
	if err != nil {
		return nil, nil, fmt.Errorf("request property: %s", err)
	}
	go func() {
		for {
			payload, ok := <-chPay
			if !ok {
				// connection lost or cancellation.
				close(ch)
				return
			}
			buf := bytes.NewBuffer(payload)
			_ = buf // discard unused variable error
			e, err := basic.ReadString(buf)
			if err != nil {
				log.Printf("unmarshall tuple: %s", err)
				continue
			}
			ch <- e
		}
	}()

	return func() {
		p.UnregisterEvent(p.ObjectID(), propertyID, handlerID)
		cancel()
	}, ch, nil
}

// SubscribeBehaviorStopped subscribe to a remote property
func (p *proxyALFrameManager) SubscribeBehaviorStopped() (func(), chan string, error) {
	propertyID, err := p.SignalID("behaviorStopped")
	if err != nil {
		return nil, nil, fmt.Errorf("property %s not available: %s", "behaviorStopped", err)
	}

	handlerID, err := p.RegisterEvent(p.ObjectID(), propertyID, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("register event for %s: %s", "behaviorStopped", err)
	}
	ch := make(chan string)
	cancel, chPay, err := p.SubscribeID(propertyID)
	if err != nil {
		return nil, nil, fmt.Errorf("request property: %s", err)
	}
	go func() {
		for {
			payload, ok := <-chPay
			if !ok {
				// connection lost or cancellation.
				close(ch)
				return
			}
			buf := bytes.NewBuffer(payload)
			_ = buf // discard unused variable error
			e, err := basic.ReadString(buf)
			if err != nil {
				log.Printf("unmarshall tuple: %s", err)
				continue
			}
			ch <- e
		}
	}()

	return func() {
		p.UnregisterEvent(p.ObjectID(), propertyID, handlerID)
		cancel()
	}, ch, nil
}

// SubscribeBehaviorFailed subscribe to a remote property
func (p *proxyALFrameManager) SubscribeBehaviorFailed() (func(), chan BehaviorFailed, error) {
	propertyID, err := p.SignalID("behaviorFailed")
	if err != nil {
		return nil, nil, fmt.Errorf("property %s not available: %s", "behaviorFailed", err)
	}

	handlerID, err := p.RegisterEvent(p.ObjectID(), propertyID, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("register event for %s: %s", "behaviorFailed", err)
	}
	ch := make(chan BehaviorFailed)
	cancel, chPay, err := p.SubscribeID(propertyID)
	if err != nil {
		return nil, nil, fmt.Errorf("request property: %s", err)
	}
	go func() {
		for {
			payload, ok := <-chPay
			if !ok {
				// connection lost or cancellation.
				close(ch)
				return
			}
			buf := bytes.NewBuffer(payload)
			_ = buf // discard unused variable error
			e, err := readBehaviorFailed(buf)
			if err != nil {
				log.Printf("unmarshall tuple: %s", err)
				continue
			}
			ch <- e
		}
	}()

	return func() {
		p.UnregisterEvent(p.ObjectID(), propertyID, handlerID)
		cancel()
	}, ch, nil
}
