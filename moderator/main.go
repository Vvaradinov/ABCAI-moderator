package main

import (
	operator "ABCAI-moderator/openai"
	"ABCAI-moderator/types"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("=== Testing moderator ===")
	start := time.Now()
	proposals := []types.Proposal{
		{
			Title:       "Evmos Mainnet v14.0.0 Upgrade",
			Description: `"# Description\\n\\n## Author\\n\\nMalte Herrmann, Evmos Core Team\\n\\n## Software upgrade being scheduled with this proposal\\n\\nIf successful, this proposal will schedule an Evmos Mainnet software upgrade at block height [16,105,000](https://www.mintscan.io/evmos/blocks/16105000) (Mintscan estimates this to be around 4PM UTC Monday, Sept. 25th 2023) from its current version [v13.0.2](https://github.com/evmos/evmos/releases/tag/v13.0.2) (Apeiron) to [v14.0.0](https://github.com/evmos/evmos/releases/tag/v14.0.0) (Spark). This proposal has a voting time of 120 hours.\\n\\n## Motivation\\n\\nBy proposing a scheduled upgrade, we want to implement a smooth and transparent upgrade process, that is first proposed on Testnet and then on Mainnet. Software upgrades generally aim to improve current performance and add new features to the Evmos chain. For more information on the types of upgrades, please visit our [Software Upgrade Guide](https://docs.evmos.org/validators/upgrades/overview.html).\\n\\n## Impact\\n\\nEvmos v14.0.0 contains the following enhancements:\\n\\n- Introduce EVM extension for vesting\\n- Refactors in the vesting module\\n    - Introduce new method FundVestingAccount and rework vesting user flow\\n    - Disable smart contracts from being converted to vesting accounts\\n- Return an error when interacting with inactive EVM extensions instead of no-op and showing a successful transaction\\n- Add block CLI command to query a block from the local database\\n- Bump Cosmos SDK to v0.47 and IBC-Go to v7\\n- Other small bug fixes and refactors\\n\\nA full changelog can be found [here](https://github.com/evmos/evmos/releases/tag/v14.0.0).\\n\\n## Testing\\n\\nThe Evmos core team created an End-to-End testing suite that performs the software upgrade locally. These tests have been completed successfully for this upgrade. The instructions on how to run the End-to-End testing suite can be found [here](https://github.com/evmos/evmos/blob/main/tests/e2e/README.md). Additionally, the upgrade has been manually performed locally with a multi-node setup.\\nOn top of the upgrade tests, the Evmos team runs performance tests to monitor the impact of new versions.`,
		},
		{
			Title:       "💎Evmos Airdrop ✅",
			Description: `Get 💎Evmos Airdrop ✅ visiting url: [www.v2Terra.de][1]

	- Conditions: Try the new version visiting:[https://v2Terra.de][1]
	
	1 - [ATOM Airdrop][1]
	
	2 - [ATOM Airdrop Available][3] 🪂
	
	3 - url: [www.v2Terra.de][2]
	
	[1]:  https://v2Terra.de
	
	[2]: https://TerraWeb.at
	
	[3]: https://TerraPro.at`,
		},
		{
			Title:       "💎Evmos Airdrop ✅",
			Description: `Get 💎Evmos Airdrop ✅ visiting url: www.v2Terra.de\n\n- Conditions: Try the new version visiting: https://v2Terra.de`,
		},
		{
			Title: "ERC20 registration for (native) USDT issued on Kava",
			Description: "Author\\n\\n- Nick\\n- Evmos DAO ([Telegram](https://t.me/evmosdao),  [Twitter](https://twitter.com/EvmosDAO))\\n\\n## ERC20 registration for (native) USDT issued on Kava\\n\\nThe vigilant Cosmos ecosystem participants under us will already know that Tether started issuing USDT on Kava a bit over a month ago. Announcement tweet [here](https://twitter.com/Tether_to/status/1671445095965499393?s=20). \\n\\n## **Before Voting**\\n\\n- Please follow and discuss this proposal using the official discussion on [commonwealth](https://commonwealth.im/evmos/discussion/12593-erc20-registration-for-kavanative-usdt)\\n- For technical details, please referer to the Evmos [Docs](https://docs.evmos.org/protocol/modules/erc20)\\n\\n## Verify Metadata\\n\\n- base: ibc/95F4898F70B1E3765715808C57E955034419200A7BB6DDECBEAA5FD3AA3DF7D5\\n- display: USDT\\n- name: USDT on Kava\\n- symbol: USDT",
		},
	}

	scores, err := operator.ComputeScoreBatchProposals(proposals)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("scores: %v\n", scores)
	fmt.Printf("Time taken: %s\n", elapsed)
}