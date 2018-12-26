package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

//const (
//	// Version message
//	AppMajor uint = 0
//	AppMinor uint = 0
//	AppPatch uint = 6
//	AppName       = "Copernicus"
//
//	defaultConfigFile  = "bitcioncash.conf"
//	defaultDataDirname = "bitcoincash"
//	defaultProjectDir  = "github.com/coeprnet/copernicus"
//)

// All command args
var (
	showVersion        = flag.Bool("version", false, "View the Copernicus version information")
	configFile         = flag.String("configfile", "bitcoincash.conf", "Path to configuration file")
	datadir            = flag.String("datadir", "data", "Path to configuration file")
	reindex            = flag.Bool("reindex", false, "reindex")
	excessiveblocksize = flag.Uint64("excessiveblocksize", 0, "")

	// RPC
	rpcListeners = flag.String("rpclistens", "127.0.0.1:18334",
		"Add an interface/port to listen for RPC connections (default port: 8334, testnet: 18334)")
	rpcUser       = flag.String("rpcuser", "", "Username for RPC connections")
	rpcPass       = flag.String("rpcpass", "", "Password for RPC connections")
	rpcMaxClients = flag.Int("rpcmaxclients", 100, "Max number of RPC clients for standard connections")

	// Log
	level = flag.String("debugerror", "info", "Define level of log, include trace, debug, info, warn, error")

	// Mempool
	maxMempool         = flag.Int64("maxmempool", 300000000, "")
	limitAncestorcount = flag.Int("limitancestorcount", 50000, "")

	// P2PNet
	testNet      = flag.Bool("testnet", false, "")
	regTest      = flag.Bool("regtest", false, "")
	whiteList    = flag.String("whitelist", "", "whitelist")
	banThreshold = flag.Uint("banscore", 100, "")

	// Chain
	assumeVlid          = flag.String("assumevalid", "", "")
	utxoHashStartHeight = flag.Int("utxohashstartheight", -1, "")
	utxoHashEndHeight   = flag.Int("utxohashendheight", -1, "")

	// Wallet
	walletEnable = flag.Bool("walletenable", false, "Enable wallet")

	// Other
	replayProtectionActivationTime = flag.Int64("replayprotectionactivationtime", -1, "")
	magneticAnomalyTime            = flag.Int64("magneticanomalyactivationtime", -1, "")
	stopAtHeight                   = flag.Int("stopatheight", -1, "")
	promiscuousMempoolFlags        = flag.String("promiscuousmempoolflags", "", "")
	blockVersion                   = flag.Int("blockversion", -1, "regtest block version")
	spendZeroConfChange            = flag.Uint("spendzeroconfchange", 1, "")
	maxTimeAdjustment              = flag.Uint("maxtimeadjustment", 4200, "Maximum allowed median peer time offset adjustment. Local perspective of time may be influenced by peers forward or backward by this amount.")
	minimumChainWork               = flag.String("minimumchainwork", "", "")
)

// rpc is an rpc config type, it used to config rpc.
type rpc struct {
	RPCListeners         []string // Add an interface/port to listen for RPC connections (default port: 8334, testnet: 18334)
	RPCUser              string   // Username for RPC connections
	RPCPass              string   // Password for RPC connections
	RPCLimitUser         string   //Username for limited RPC connections
	RPCLimitPass         string   //Password for limited RPC connections
	RPCCert              string   `default:""` //File containing the certificate file
	RPCKey               string   //File containing the certificate key
	RPCMaxClients        int      //Max number of RPC clients for standard connections
	RPCMaxWebsockets     int      //Max number of RPC websocket connections
	RPCMaxConcurrentReqs int      //Max number of concurrent RPC requests that may be processed concurrently
	RPCQuirks            bool     //Mirror some JSON-RPC quirks of Bitcoin Core -- NOTE: Discouraged unless interoperability issues need to be worked around
}

func initRpc() *rpc {
	return &rpc{
		RPCListeners:  str2strs(*rpcListeners),
		RPCUser:       *rpcUser,
		RPCPass:       *rpcPass,
		RPCMaxClients: *rpcMaxClients,
	}
}

type mempool struct {
	MinFeeRate           int64  //
	LimitAncestorCount   int    // Default for -limitancestorcount, max number of in-mempool ancestors
	LimitAncestorSize    int    // Default for -limitancestorsize, maximum kilobytes of tx + all in-mempool ancestors
	LimitDescendantCount int    // Default for -limitdescendantcount, max number of in-mempool descendants
	LimitDescendantSize  int    // Default for -limitdescendantsize, maximum kilobytes of in-mempool descendants
	MaxPoolSize          int64  `default:"300000000"` // Default for MaxPoolSize, maximum megabytes of mempool memory usage
	MaxPoolExpiry        int    `default:"336"`       // Default for -mempoolexpiry, expiration time for mempool transactions in hours
	CheckFrequency       uint64 `default:"4294967296"`
}

func initMempool() *mempool {
	return &mempool{
		MaxPoolSize:        *maxMempool,
		LimitAncestorCount: *limitAncestorcount,
	}
}

type p2pNet struct {
	ListenAddrs         []string `validate:"require" default:"1234"`
	MaxPeers            int      `default:"128"`
	TargetOutbound      int      `default:"64"`
	DisableBanning      bool
	BanThreshold        uint32 `default:"100"`
	TestNet             bool
	RegTest             bool
	SimNet              bool
	DisableListen       bool     `default:"true"`
	BlocksOnly          bool     //Do not accept transactions from remote peers.
	BanDuration         int64    `default:"86400"` // How long to ban misbehaving peers
	Proxy               string   // Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)
	UserAgentComments   []string // Comment to add to the user agent -- See BIP 14 for more information.
	DisableDNSSeed      bool     //Disable DNS seeding for peers
	DisableRPC          bool
	DisableTLS          bool
	Whitelists          []*net.IPNet
	NoOnion             bool     `default:"true"`  // Disable connecting to tor hidden services
	Upnp                bool     `default:"false"` // Use UPnP to map our listening port outside of NAT
	ExternalIPs         []string // Add an ip to the list of local addresses we claim to listen on to peers
	MaxTimeAdjustment   uint64   `default:"4200"`
	ConnectPeersOnStart []string
	//AddCheckpoints      []model.Checkpoint
}

func initP2PNet() *p2pNet {
	return &p2pNet{
		TestNet:      *testNet,
		RegTest:      *regTest,
		Whitelists:   getNetFromString(*whiteList),
		BanThreshold: uint32(*banThreshold),
	}
}

func getNetFromString(netStr string) []*net.IPNet {
	if netStr == "" {
		return nil
	}

	netlists := str2strs(netStr)
	nets := make([]*net.IPNet, len(netlists))

	for _, addr := range netlists {
		_, ipnet, err := net.ParseCIDR(addr)
		if err != nil {
			ip := net.ParseIP(addr)
			if ip == nil {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("[Error]The whitelist value of '%s' is invalid", addr))
				continue
			}

			var bits int
			if ip.To4() == nil {
				// IPv6
				bits = 128
			} else {
				bits = 32
			}

			ipnet = &net.IPNet{
				IP:   ip,
				Mask: net.CIDRMask(bits, bits),
			}
		}
		nets = append(nets, ipnet)
	}

	return nets
}

type addrMgr struct {
	SimNet       bool
	ConnectPeers []string
}

func initAddrMgr() *addrMgr {
	return &addrMgr{}
}

type protocol struct {
	NoPeerBloomFilters bool `default:"true"`
	DisableCheckpoints bool `default:"true"`
}

func initProtocol() *protocol {
	return &protocol{}
}

type script struct {
	AcceptDataCarrier   bool `default:"true"`
	MaxDatacarrierBytes uint `default:"223"`
	IsBareMultiSigStd   bool `default:"true"`
	//use promiscuousMempoolFlags to make more or less check of script, the type of value is uint
	PromiscuousMempoolFlags string
	Par                     int `default:"32"`
}

func initScript() *script {
	return &script{}
}

type txOut struct {
	DustRelayFee int64 `default:"83"`
}

func initTxOut() *txOut {
	return &txOut{}
}

type chain struct {
	AssumeValid         string
	UtxoHashStartHeight int32 `default:"-1"`
	UtxoHashEndHeight   int32 `default:"-1"`
}

func initChain() *chain {
	return &chain{
		AssumeValid:         *assumeVlid,
		UtxoHashStartHeight: int32(*utxoHashStartHeight),
		UtxoHashEndHeight:   int32(*utxoHashEndHeight),
	}
}

type mining struct {
	BlockMinTxFee int64  // default DefaultBlockMinTxFee
	BlockMaxSize  uint64 // default DefaultMaxGeneratedBlockSize
	Strategy      string `default:"ancestorfeerate"` // option:ancestorfee/ancestorfeerate
}

func initMining() *mining {
	return &mining{}
}

type pProf struct {
	IP   string `default:"localhost"`
	Port string `default:"6060"`
}

func initPprof() *pProf {
	return &pProf{}
}

type blockIndex struct {
	CheckBlockIndex bool
}

func initBlockIndex() *blockIndex {
	return &blockIndex{}
}

type wallet struct {
	Enable              bool `default:"false"`
	Broadcast           bool `default:"false"`
	SpendZeroConfChange bool `default:"true"`
}

func initWallet() *wallet {
	return &wallet{
		Enable: *walletEnable,
	}
}

type log struct {
	Level    string   //description:"Define level of log,include trace, debug, info, warn, error"
	Module   []string // only output the specified module's log when using log.Print(...)
	FileName string   // the name of log file
}

func initLog() *log {
	return &log{
		Level: *level,
	}
}

type config struct {
	DataDir            string `default:"data"`
	Reindex            bool
	Excessiveblocksize uint64
	RPC                rpc
	Log                log
	Mempool            mempool
	P2PNet             p2pNet
	AddrMgr            addrMgr
	Protocol           protocol
	Script             script
	TxOut              txOut
	Chain              chain
	Mining             mining
	PProf              pProf
	BlockIndex         blockIndex
	Wallet             wallet
}

// New is the way to get a instance of config.

func New() *config {
	flag.Parse()
	return initConfig()
}

func initConfig() *config {
	return &config{
		DataDir:            *datadir,
		Reindex:            *reindex,
		Excessiveblocksize: *excessiveblocksize,
		RPC:                *initRpc(),
		Log:                *initLog(),
		Mempool:            *initMempool(),
		P2PNet:             *initP2PNet(),
		AddrMgr:            *initAddrMgr(),
		Protocol:           *initProtocol(),
		Script:             *initScript(),
		TxOut:              *initTxOut(),
		Chain:              *initChain(),
		Mining:             *initMining(),
		PProf:              *initPprof(),
		BlockIndex:         *initBlockIndex(),
		Wallet:             *initWallet(),
	}
}

func str2strs(str string) []string {
	return strings.Split(str, ",")
}

func version() {
	fmt.Sprintf("%d.%d.%d", AppMajor, AppMinor, AppPatch)
	os.Exit(0)
}

func main() {
	conf := New()
	fmt.Printf("%+v\n", *conf)
}
