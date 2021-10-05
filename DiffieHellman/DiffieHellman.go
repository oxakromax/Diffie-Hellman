package DiffieHellman

import (
	"encoding/json"
	"errors"
	"github.com/oxakromax/Diffie-Hellman/utils"
	"math/big"
)

type Diffie struct {
	privateKey, comunnicatorKey, commonKey *big.Int
	PublicKey, G, P                        *big.Int
}

func (d *Diffie) setComunnicatorKey(comunnicatorKey *big.Int) {
	d.comunnicatorKey = comunnicatorKey
	d.commonKey = new(big.Int).Exp(d.comunnicatorKey, d.privateKey, d.P)
}

func (d *Diffie) PartnerConfig(p *Diffie) {
	// For offline Porpoises you can configure
	if d.P == nil || d.G == nil {
		d.P, d.G = p.P, p.G
		d.privateKey = big.NewInt(utils.GetRandomN64(1, d.P.Int64()-int64(1)))
		d.PublicKey = new(big.Int).Exp(d.G, d.privateKey, d.P)
	}
	if p.PublicKey != nil {
		d.setComunnicatorKey(p.PublicKey)
	}
}
func (d *Diffie) JsonConfig(data []byte) error {
	// For easy configuration from a Json received by Fibber, Echo, Gin... etc
	n := new(Diffie)
	err := json.Unmarshal(data, n)
	if n.P == nil || n.G == nil {
		return errors.New("Empty Partner, You may want to use FirstConfig() instead")
	}
	d.PartnerConfig(n)
	return err
}
func (d *Diffie) FirstConfig() {
	// Automatic Configuration if it's the first time creating the struct
	d.P = big.NewInt(utils.PickRandomGorP(utils.PrimeNumbers()))
	d.G = big.NewInt(utils.PickRandomGorP(utils.PrimitiveRoots(d.P.Int64())))
	d.privateKey = big.NewInt(utils.GetRandomN64(1, d.P.Int64()-int64(1)))
	d.PublicKey = new(big.Int).Exp(d.G, d.privateKey, d.P)
}

func (d Diffie) GetKey() int64 {
	return d.commonKey.Int64()
}
