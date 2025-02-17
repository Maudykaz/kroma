package immutables_test

import (
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/immutables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBuildKroma(t *testing.T) {
	results, err := immutables.BuildOptimism(immutables.ImmutableConfig{
		"L2StandardBridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2CrossDomainMessenger": {
			"otherMessenger": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2ERC721Bridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
			"messenger":   common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"KromaMintableERC721Factory": {
			"remoteChainId": big.NewInt(1),
			"bridge":        common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"ValidatorRewardVault": {
			"validatorPoolAddress": common.HexToAddress("0x1234567890123456789012345678901234567890"),
			"rewardDivider":        new(big.Int).SetUint64(24 * 7),
		},
		"L1FeeVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"ProtocolVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
	}, false)
	require.Nil(t, err)
	require.NotNil(t, results)

	contracts := map[string]bool{
		"GasPriceOracle":             true,
		"L1Block":                    true,
		"L2CrossDomainMessenger":     true,
		"L2StandardBridge":           true,
		"L2ToL1MessagePasser":        true,
		"ValidatorRewardVault":       true,
		"ProtocolVault":              true,
		"L1FeeVault":                 true,
		"KromaMintableERC20Factory":  true,
		"L2ERC721Bridge":             true,
		"KromaMintableERC721Factory": true,
	}

	// Only the exact contracts that we care about are being
	// modified
	require.Equal(t, len(results), len(contracts))

	for name, bytecode := range results {
		// There is bytecode there
		require.Greater(t, len(bytecode), 0)
		// It is in the set of contracts that we care about
		require.True(t, contracts[name])
	}
}
