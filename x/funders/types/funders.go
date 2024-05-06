package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (f *Funding) GetScore(whitelist map[string]WhitelistCoinEntry) (score uint64) {
	for _, coin := range f.Amounts {
		if entry, found := whitelist[coin.Denom]; found {
			score += uint64(entry.CoinWeight.MulInt64(coin.Amount.Int64()).TruncateInt64())
		}
	}

	return
}

// CleanAmountsPerBundle removes every coin in amounts per bundle
// which is not present in the amounts coins list
func (f *Funding) CleanAmountsPerBundle() {
	amountsPerBundle := sdk.NewCoins()

	for _, coin := range f.AmountsPerBundle {
		if found, _ := f.Amounts.Find(coin.Denom); found {
			amountsPerBundle = amountsPerBundle.Add(coin)
		}
	}

	f.AmountsPerBundle = amountsPerBundle
}

func (f *Funding) ChargeOneBundle(whitelist map[string]WhitelistCoinEntry) (payouts sdk.Coins) {
	chargable := f.Amounts.Min(f.AmountsPerBundle)

	// only charge coins which are whitelisted
	for _, coin := range chargable {
		if _, found := whitelist[coin.Denom]; found {
			payouts = payouts.Add(coin)
		}
	}

	f.TotalFunded = f.TotalFunded.Add(payouts...)
	f.Amounts = f.Amounts.Sub(payouts...)
	f.CleanAmountsPerBundle()
	return
}

func (f *Funding) IsActive() (isActive bool) {
	return !f.Amounts.IsZero()
}

func (f *Funding) IsInactive() (isInactive bool) {
	return !f.IsActive()
}

// SetInactive removes a funding from active fundings
func (fs *FundingState) SetInactive(funding *Funding) {
	for i, funderAddress := range fs.ActiveFunderAddresses {
		if funderAddress == funding.FunderAddress {
			fs.ActiveFunderAddresses[i] = fs.ActiveFunderAddresses[len(fs.ActiveFunderAddresses)-1]
			fs.ActiveFunderAddresses = fs.ActiveFunderAddresses[:len(fs.ActiveFunderAddresses)-1]
			break
		}
	}
}

// SetActive adds a funding to active fundings
func (fs *FundingState) SetActive(funding *Funding) {
	for _, funderAddress := range fs.ActiveFunderAddresses {
		if funderAddress == funding.FunderAddress {
			return
		}
	}
	fs.ActiveFunderAddresses = append(fs.ActiveFunderAddresses, funding.FunderAddress)
}
