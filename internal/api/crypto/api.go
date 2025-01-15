package api

import (
	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

type Service interface {
	Remove()
}

var _ desc.ServerInterface = (*Implementation)(nil)

type Implementation struct {
	cryptoService Service
}

func NewImplementation(cryptoService Service) *Implementation {
	return &Implementation{
		cryptoService: cryptoService,
	}
}
