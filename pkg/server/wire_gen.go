// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"infra-practice/pkg/presentation/controller"
)

// Injectors from wire.go:

func initialize() (*Server, error) {
	sampleController := controller.NewSampleController()
	server := NewServer(sampleController)
	return server, nil
}

// wire.go:

var ProvideApp = wire.NewSet(
	ProvidePresentation,
	NewServer,
)

var ProvidePresentation = wire.NewSet(controller.NewSampleController)
