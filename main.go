package main

import (
	"fmt"
	_ "image/png"
	"strings"
	"time"

	ui "github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
	"github.com/hashicorp/mdns"
	"github.com/lugu/qiloop"
	"github.com/lugu/qiloop/bus/session/token"
)

var (
	behaviorManager ALBehaviorManagerProxy
	behaviors       []string

	servers = []*mdns.ServiceEntry{}

	state      ui.UpdateFn
	isScanning = true
)

func readField(infoFields []string, field, ifMissing string) string {
	for _, f := range infoFields {
		keys := strings.SplitN(f, "=", 2)
		if len(keys) > 1 && keys[0] == field {
			return keys[1]
		}
	}
	return ifMissing
}

func connect(serverURL string) error {

	user, token := token.GetUserToken()
	session, err := qiloop.NewSession(serverURL, user, token)
	if err != nil {
		return err
	}

	behaviorManager, err = Services(session).ALBehaviorManager()
	if err != nil {
		return err
	}

	behaviors, err = behaviorManager.GetBehaviorNames()
	if err != nil {
		return err
	}
	return nil
}

func stateScanningScreen(w *ui.Window) {

	for _, s := range servers {
		protocol := readField(s.InfoFields, "Protocol", "unknown")
		robotType := readField(s.InfoFields, "RobotType", "unknown")

		if robotType == "unknown" {
			continue
		}
		if protocol == "unknown" {
			switch s.Port {
			case 9559:
				protocol = "tcp"
			case 9503, 9443:
				protocol = "tcps"
			default:
				continue
			}
		}

		label := fmt.Sprintf("%s: %s (%s)", robotType, s.Host, s.AddrV4)
		if w.ButtonText(label) {
			url := fmt.Sprintf("%s://%s:%d", protocol, s.AddrV4, s.Port)
			err := connect(url)
			if err != nil {
				state = stateErrorScreen(err)
			} else {
				state = stateBehaviorScreen
			}
		}
	}
	w.Label("Using Multicast DNS to discover robots.", "CC")
	if isScanning {
		w.Label("Scanning...", "CC")
	}
}

func stateErrorScreen(err error) ui.UpdateFn {
	return func(w *ui.Window) {
		w.Label("Connection error", "CC")
		w.Row(200).Dynamic(1)
		w.LabelWrap(err.Error())
	}
}

func stateBehaviorScreen(w *ui.Window) {
	for _, b := range behaviors {
		label := strings.Trim(b, "_Behaviors__animations")
		if w.ButtonText(label) {
			err := behaviorManager.StopAllBehaviors()
			if err != nil {
				state = stateErrorScreen(err)
			}
			err = behaviorManager.StartBehavior(b)
			if err != nil {
				state = stateErrorScreen(err)
			}
		}
	}
}

func updatefn(w *ui.Window) {
	w.Row(50).Dynamic(1)
	w.LabelWrap("")
	state(w)
}

func discover(wnd ui.MasterWindow) {
	entries := make(chan *mdns.ServiceEntry, 20)

	go func() {
		for {
			s, ok := <-entries
			if ok {
				servers = append(servers, s)
				wnd.Changed()
			} else {
				isScanning = false
				return
			}
		}
	}()

	params := mdns.DefaultParams("_naoqi._tcp")
	params.Entries = entries
	params.Timeout = time.Minute
	mdns.Query(params)
	close(entries)
}

func main() {
	state = stateScanningScreen
	wnd := ui.NewMasterWindow(0, "Hello NAO", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 2.0))

	go discover(wnd)
	wnd.Main()
}
