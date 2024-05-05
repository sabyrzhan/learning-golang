package facade_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacade(t *testing.T) {
	p2pService := NewP2PPaymentFacade()
	err := p2pService.SendPayment("+1111111", "+1222222", 200, EUR, "1234")
	assert.Nil(t, err)
}
