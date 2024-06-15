package main

type resourceType int

const (
	_ resourceType = iota
	resourceTypeEmployee
	resourceTypeWater
)

type resourceConsumer interface {
	consumeResource(*worker)
}
