package services

import (
	"WallE/domains"
	"WallE/models"
)

type serviceTransaksi struct {
	repo domains.TransaksiDomain
}

func (s *serviceTransaksi) TransaksiBaru(transaksi models.Transaksi) error {
	return s.repo.TransaksiBaru(transaksi)
}
