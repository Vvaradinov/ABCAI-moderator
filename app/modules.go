package app

//
//import (
//	"cosmossdk.io/x/upgrade"
//	upgradetypes "cosmossdk.io/x/upgrade/types"
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//	"github.com/cosmos/cosmos-sdk/x/consensus"
//	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
//	distr "github.com/cosmos/cosmos-sdk/x/distribution"
//	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
//	"github.com/cosmos/cosmos-sdk/x/genutil"
//	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
//	"github.com/cosmos/cosmos-sdk/x/gov"
//	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
//	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
//	sdkparams "github.com/cosmos/cosmos-sdk/x/params"
//	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
//	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
//	"github.com/cosmos/cosmos-sdk/x/staking"
//	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
//	// IBC Imports
//	//"github.com/cosmos/ibc-go/modules/capability"
//	//capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
//	//"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
//	//ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
//	//ibc "github.com/cosmos/ibc-go/v7/modules/core"
//	//ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
//	//ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
//	//ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
//)
//
//var (
//	ModuleBasics = module.NewBasicManager(
//		auth.AppModuleBasic{},
//		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
//		bank.AppModuleBasic{},
//		//capability.AppModuleBasic{},
//		staking.AppModuleBasic{},
//		distr.AppModuleBasic{},
//		gov.NewAppModuleBasic(
//			[]govclient.ProposalHandler{
//				paramsclient.ProposalHandler,
//				//ibcclientclient.UpdateClientProposalHandler,
//				//ibcclientclient.UpgradeProposalHandler,
//			},
//		),
//		consensus.AppModuleBasic{},
//		//ibc.AppModuleBasic{},
//		//ibctm.AppModuleBasic{},
//		upgrade.AppModuleBasic{},
//		//transfer.AppModuleBasic{},
//		// TODO: Add module AppModuleBasic
//	)
//
//	maccPerms = map[string][]string{
//		authtypes.FeeCollectorName:     nil,
//		distrtypes.ModuleName:          nil,
//		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
//		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
//		govtypes.ModuleName:            {authtypes.Burner},
//		//ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
//	}
//)
//
//func appModules(
//	app *App,
//	appCodec codec.cdc,
//	txConfig client.txConfig,
//) []module.AppModule {
//
//	return []module.AppModule{
//		genutil.NewAppModule(
//			app.AccountKeeper,
//			app.StakingKeeper,
//			app,
//			txConfig,
//		),
//		auth.NewAppModule(appCodec, app.AccountKeeper, nil, app.GetSubspace(authtypes.ModuleName)),
//		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
//		//capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
//		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
//		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
//		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
//		upgrade.NewAppModule(app.UpgradeKeeper, app.AccountKeeper.AddressCodec()),
//		sdkparams.NewAppModule(app.ParamsKeeper),
//		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
//
//		// IBC Modules
//		//ibc.NewAppModule(app.IBCKeeper),
//		//transfer.NewAppModule(app.TransferKeeper),
//
//		// TODO: Add new app module constructor
//	}
//}
//
//func orderBeginBlockers() []string {
//	return []string{
//		upgradetypes.ModuleName,
//		//capabilitytypes.ModuleName,
//		distrtypes.ModuleName,
//		stakingtypes.ModuleName,
//		//authtypes.ModuleName,
//		//banktypes.ModuleName,
//		//govtypes.ModuleName,
//		//ibcexported.ModuleName,
//		//ibctransfertypes.ModuleName,
//		genutiltypes.ModuleName,
//		//paramstypes.ModuleName,
//		//consensusparamtypes.ModuleName,
//		// TODO: Add module name
//	}
//}
//
//func orderEndBlockers() []string {
//	return []string{
//		govtypes.ModuleName,
//		stakingtypes.ModuleName,
//		//ibcexported.ModuleName,
//		//ibctransfertypes.ModuleName,
//		//capabilitytypes.ModuleName,
//		//authtypes.ModuleName,
//		//banktypes.ModuleName,
//		//distrtypes.ModuleName,
//		genutiltypes.ModuleName,
//		//paramstypes.ModuleName,
//		//upgradetypes.ModuleName,
//		//consensusparamtypes.ModuleName,
//		// TODO: Add module name
//	}
//}
//
//func orderGenesisBlockers() []string {
//	return []string{
//		//capabilitytypes.ModuleName,
//		authtypes.ModuleName,
//		banktypes.ModuleName,
//		distrtypes.ModuleName,
//		stakingtypes.ModuleName,
//		govtypes.ModuleName,
//		genutiltypes.ModuleName,
//		//ibctransfertypes.ModuleName,
//		//ibcexported.ModuleName,
//		paramstypes.ModuleName,
//		upgradetypes.ModuleName,
//		consensusparamtypes.ModuleName,
//		// TODO: Add module name
//	}
//}
