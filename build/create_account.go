package build

import (
	"errors"

	"github.com/openbankit/go-base/xdr"
)

// CreateAccount groups the creation of a new CreateAccountBuilder with a call
// to Mutate.
func CreateAccount(muts ...interface{}) (result CreateAccountBuilder) {
	result.Mutate(muts...)
	return
}

// CreateAccountMutator is a interface that wraps the
// MutateCreateAccount operation.  types may implement this interface to
// specify how they modify an xdr.PaymentOp object
type CreateAccountMutator interface {
	MutateCreateAccount(*xdr.CreateAccountOp) error
}

// CreateAccountBuilder helps to build CreateAccountOp structs.
type CreateAccountBuilder struct {
	O   xdr.Operation
	CA  xdr.CreateAccountOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *CreateAccountBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case CreateAccountMutator:
			err = mut.MutateCreateAccount(&b.CA)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutateCreateAccount for Destination sets the CreateAccountOp's Destination
// field
func (m Destination) MutateCreateAccount(o *xdr.CreateAccountOp) error {
	return setAccountId(m.AddressOrSeed, &o.Destination)
}

// MutateCreateAccount for AccountType sets the CreateAccountOp's
// AccountType field
// TODO: FIX IT
// func (m CreateAccountWithScratch) MutateCreateAccount(o *xdr.CreateAccountOp) (err error) {
// 	o.Body.AccountType = xdr.Uint32(m.AccountType)
// 	switch m.AccountType {
// 	case xdr.AccountTypeAccountScratchCard:
// 		o.Body.ScratchCard.Asset = m.Asset
// 		o.Body.ScratchCard.Amount = m.Amount
// 	}
// 	return
// }
