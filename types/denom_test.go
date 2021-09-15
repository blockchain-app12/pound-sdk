package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	bps  = "bps"  // 1 (base denom unit)
	mbps = "mbps" // 10^-3 (milli)
	ubps = "ubps" // 10^-6 (micro)
	nbps = "nbps" // 10^-9 (nano)
)

func TestRegisterDenom(t *testing.T) {
	bpsUnit := OneDec() // 1 (base denom unit)

	require.NoError(t, RegisterDenom(bps, bpsUnit))
	require.Error(t, RegisterDenom(bps, bpsUnit))

	res, ok := GetDenomUnit(bps)
	require.True(t, ok)
	require.Equal(t, bpsUnit, res)

	res, ok = GetDenomUnit(mbps)
	require.False(t, ok)
	require.Equal(t, ZeroDec(), res)

	// reset registration
	denomUnits = map[string]Dec{}
}

func TestConvertCoins(t *testing.T) {
	bpsUnit := OneDec() // 1 (base denom unit)
	require.NoError(t, RegisterDenom(bps, bpsUnit))

	mbpsUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	require.NoError(t, RegisterDenom(mbps, mbpsUnit))

	ubpsUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	require.NoError(t, RegisterDenom(ubps, ubpsUnit))

	nbpsUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	require.NoError(t, RegisterDenom(nbps, nbpsUnit))

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), bps, Coin{}, true},
		{NewCoin(bps, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(bps, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(bps, NewInt(5)), mbps, NewCoin(mbps, NewInt(5000)), false},       // bps => mbps
		{NewCoin(bps, NewInt(5)), ubps, NewCoin(ubps, NewInt(5000000)), false},    // bps => ubps
		{NewCoin(bps, NewInt(5)), nbps, NewCoin(nbps, NewInt(5000000000)), false}, // bps => nbps

		{NewCoin(ubps, NewInt(5000000)), mbps, NewCoin(mbps, NewInt(5000)), false},       // ubps => mbps
		{NewCoin(ubps, NewInt(5000000)), nbps, NewCoin(nbps, NewInt(5000000000)), false}, // ubps => nbps
		{NewCoin(ubps, NewInt(5000000)), bps, NewCoin(bps, NewInt(5)), false},            // ubps => bps

		{NewCoin(mbps, NewInt(5000)), nbps, NewCoin(nbps, NewInt(5000000000)), false}, // mbps => nbps
		{NewCoin(mbps, NewInt(5000)), ubps, NewCoin(ubps, NewInt(5000000)), false},    // mbps => ubps
	}

	for i, tc := range testCases {
		res, err := ConvertCoin(tc.input, tc.denom)
		require.Equal(
			t, tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		require.Equal(
			t, tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}

	// reset registration
	denomUnits = map[string]Dec{}
}
