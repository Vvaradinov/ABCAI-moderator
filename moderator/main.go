package main

import (
	operator "ABCAI-moderator/openai"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("=== Testing moderator ===")
	start_legit := time.Now()
	score_legit, _ := operator.ComputeScoreProposal("Evmos Mainnet v14.0.0 Upgrade", `Evmos Mainnet v14.0.0 Upgrade","description":"# Description\\n\\n## Author\\n\\nMalte Herrmann, Evmos Core Team\\n\\n## Software upgrade being scheduled with this proposal\\n\\nIf successful, this proposal will schedule an Evmos Mainnet software upgrade at block height [16,105,000](https://www.mintscan.io/evmos/blocks/16105000) (Mintscan estimates this to be around 4PM UTC Monday, Sept. 25th 2023) from its current version [v13.0.2](https://github.com/evmos/evmos/releases/tag/v13.0.2) (Apeiron) to [v14.0.0](https://github.com/evmos/evmos/releases/tag/v14.0.0) (Spark). This proposal has a voting time of 120 hours.\\n\\n## Motivation\\n\\nBy proposing a scheduled upgrade, we want to implement a smooth and transparent upgrade process, that is first proposed on Testnet and then on Mainnet. Software upgrades generally aim to improve current performance and add new features to the Evmos chain. For more information on the types of upgrades, please visit our [Software Upgrade Guide](https://docs.evmos.org/validators/upgrades/overview.html).\\n\\n## Impact\\n\\nEvmos v14.0.0 contains the following enhancements:\\n\\n- Introduce EVM extension for vesting\\n- Refactors in the vesting module\\n    - Introduce new method FundVestingAccount and rework vesting user flow\\n    - Disable smart contracts from being converted to vesting accounts\\n- Return an error when interacting with inactive EVM extensions instead of no-op and showing a successful transaction\\n- Add block CLI command to query a block from the local database\\n- Bump Cosmos SDK to v0.47 and IBC-Go to v7\\n- Other small bug fixes and refactors\\n\\nA full changelog can be found [here](https://github.com/evmos/evmos/releases/tag/v14.0.0).\\n\\n## Testing\\n\\nThe Evmos core team created an End-to-End testing suite that performs the software upgrade locally. These tests have been completed successfully for this upgrade. The instructions on how to run the End-to-End testing suite can be found [here](https://github.com/evmos/evmos/blob/main/tests/e2e/README.md). Additionally, the upgrade has been manually performed locally with a multi-node setup.\\nOn top of the upgrade tests, the Evmos team runs performance tests to monitor the impact of new versions.`)
	elapsed_legit := time.Since(start_legit)
	fmt.Printf("score: %v\n", score_legit)
	fmt.Printf("Time taken: %s\n", elapsed_legit)

	start_scam := time.Now()
	score_scam, _ := operator.ComputeScoreProposal("💎Evmos Airdrop ✅", `Get 💎Evmos Airdrop ✅ visiting url: [www.v2Terra.de][1]

	- Conditions: Try the new version visiting:[https://v2Terra.de][1]
	
	1 - [ATOM Airdrop][1]
	
	2 - [ATOM Airdrop Available][3] 🪂
	
	3 - url: [www.v2Terra.de][2]
	
	[1]:  https://v2Terra.de
	
	[2]: https://TerraWeb.at
	
	[3]: https://TerraPro.at`)
	elapsed_scam := time.Since(start_scam)
	fmt.Printf("score: %v\n", score_scam)
	fmt.Printf("Time taken: %s\n", elapsed_scam)
}