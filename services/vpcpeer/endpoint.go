package vpcpeer

import "sync"

// EndpointMap Endpoint Data
var once sync.Once
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	once.Do(func() {
		EndpointMap = map[string]string{}
	})
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
