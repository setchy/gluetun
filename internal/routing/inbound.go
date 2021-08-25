package routing

import (
	"errors"
	"fmt"
	"net"
)

const (
	inboundTable    = 200
	inboundPriority = 100
)

var (
	errDefaultIP   = errors.New("cannot get default IP address")
	errRuleAdd     = errors.New("cannot add rule")
	errRouteAdd    = errors.New("cannot add route")
	errRuleDelete  = errors.New("cannot delete rule")
	errRouteDelete = errors.New("cannot delete route")
)

func (r *Routing) routeInboundFromDefault(defaultGateway net.IP,
	defaultInterface string) (err error) {
	if err := r.addRuleInboundFromDefault(inboundTable); err != nil {
		return fmt.Errorf("%w: %s", errRuleAdd, err)
	}

	defaultDestination := net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(0, 0, 0, 0)}
	if err := r.addRouteVia(defaultDestination, defaultGateway, defaultInterface, inboundTable); err != nil {
		return fmt.Errorf("%w: %s", errRouteAdd, err)
	}

	return nil
}

func (r *Routing) unrouteInboundFromDefault(defaultGateway net.IP,
	defaultInterface string) (err error) {
	defaultDestination := net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(0, 0, 0, 0)}
	if err := r.deleteRouteVia(defaultDestination, defaultGateway, defaultInterface, inboundTable); err != nil {
		return fmt.Errorf("%w: %s", errRouteDelete, err)
	}

	if err := r.delRuleInboundFromDefault(inboundTable); err != nil {
		return fmt.Errorf("%w: %s", errRuleDelete, err)
	}

	return nil
}

func (r *Routing) addRuleInboundFromDefault(table int) (err error) {
	defaultIP, err := r.DefaultIP()
	if err != nil {
		return fmt.Errorf("%w: %s", errDefaultIP, err)
	}

	if err := r.addIPRule(defaultIP, table, inboundPriority); err != nil {
		return fmt.Errorf("%w: %s", errIPRuleAdd, err)
	}

	return nil
}

func (r *Routing) delRuleInboundFromDefault(table int) (err error) {
	defaultIP, err := r.DefaultIP()
	if err != nil {
		return fmt.Errorf("%w: %s", errDefaultIP, err)
	}

	if err := r.deleteIPRule(defaultIP, table, inboundPriority); err != nil {
		return fmt.Errorf("%w: %s", errIPRuleAdd, err)
	}

	return nil
}