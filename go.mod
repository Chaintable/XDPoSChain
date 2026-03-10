module github.com/XinFinOrg/XDPoSChain

go 1.24.0

require (
	github.com/Chaintable/pipeline v0.0.62-xdc-v2.6.8-debank-2
	github.com/VictoriaMetrics/fastcache v1.13.2
	github.com/cespare/cp v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/docker v1.4.2-0.20180625184442-8e610b2b55bf
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fatih/color v1.13.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/snappy v1.0.0
	github.com/gorilla/websocket v1.5.0
	github.com/holiman/uint256 v1.3.2
	github.com/huin/goupnp v1.3.0
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mattn/go-colorable v0.1.13
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.5
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/pkg/errors v0.9.1
	github.com/prometheus/prometheus v1.7.2-0.20170814170113-3101606756c5
	github.com/rs/cors v1.7.0
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/stretchr/testify v1.10.0
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	golang.org/x/crypto v0.36.0
	golang.org/x/sync v0.12.0
	golang.org/x/sys v0.36.0
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/consensys/gnark-crypto v0.10.0
	github.com/crate-crypto/go-kzg-4844 v1.0.0
	github.com/deckarep/golang-set/v2 v2.7.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/dop251/goja v0.0.0-20200721192441-a695b0cdd498
	github.com/ethereum/c-kzg-4844 v1.0.0
	github.com/fjl/gencodec v0.0.0-20230517082657-f9840df7b83e
	github.com/fsnotify/fsnotify v1.8.0
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/google/gofuzz v1.2.0
	github.com/google/uuid v1.6.0
	github.com/influxdata/influxdb-client-go/v2 v2.4.0
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c
	github.com/karalabe/hid v1.0.1-0.20240306101548-573246063e52
	github.com/kevinburke/go-bindata v3.23.0+incompatible
	github.com/kylelemons/godebug v1.1.0
	github.com/mattn/go-isatty v0.0.17
	github.com/protolambda/bls12-381-util v0.0.0-20220416220906-d8552aa452c7
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible
	github.com/status-im/keycard-go v0.3.3
	github.com/urfave/cli/v2 v2.27.5
	golang.org/x/text v0.23.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2 v1.32.5 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.6 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.28.5 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.46 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.20 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.4.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.66.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.1 // indirect
	github.com/aws/smithy-go v1.22.1 // indirect
	github.com/bits-and-blooms/bitset v1.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/deepmap/oapi-codegen v1.6.0 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/garslo/gogen v0.0.0-20170306192744-1d203ffc1f61 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/kilic/bls12-381 v0.1.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/v3 v3.5.10 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/grpc v1.58.3 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools v2.2.0+incompatible // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
