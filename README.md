
# Election Chaincode

  

An attribute-based access controlled chaincode backend for a voting application. 

# Tech Stack

- Golang 1.22.1
- Hyperleldger Fabric Contract API v1.2.2
- LevelDB
- Docker


# Setup (Non-Live)

## Installing Dependencies

- This is a Golang app. Download Golang here (https://go.dev/)

- The non-live environment is the standard Fabric Test Network. To find out how to set it up, go here (https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html)

## Start the network and deploy the smart contract
### Start the Network

Once the dependencies have been installed or downloaded, we can use the Fabric test network to deploy and interact with the smart contract. To do this, run the following command to go to the test network directory and bring down any running test networks.

```
cd fabric-samples/test-network
./network.sh down
```
Run the following command to deploy the test network using Certificate Authorities:
```
./network.sh up createChannel -ca
```
Then, deploy the chaincode into a channel on the network.
```
./network.sh deployCC -ccn election -ccp /mnt/c/users/others/documents/golang_project/hyperledger_fabric_related/election_code -ccl go
```

### Register the sample identities
First, set the following environment variables in order to use the Fabric CA Client
```
export PATH=${PWD}/../bin:${PWD}:$PATH 
export FABRIC_CFG_PATH=$PWD/../config/
```
As a sample, we will create the identities using the Org1 CA. Set the fabric ca home to the msp of the org1 ca admin.
  
```
export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.example.com/
```

Then, we will register and enroll the user, and copy the required YAML for that user.

### For the user with the attribute 'election.admin'
```
fabric-ca-client register --id.name election.admin3 --id.secret election.admin1 --id.type client --id.affiliation org1 --id.attrs 'election.admin=true:ecert' --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```
```
fabric-ca-client enroll -u https://election.admin3:election.admin1@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/election.admin3@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```

```
cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/election.admin3@org1.example.com/msp/config.yaml"
```
### For the user with the attribute 'election.voter'
```
fabric-ca-client register --id.name election.voter3 --id.secret election.voter1 --id.type client --id.affiliation org1 --id.attrs 'election.voter=true:ecert' --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```
```
fabric-ca-client enroll -u https://election.voter3:election.voter1@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/election.voter3@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```

```
cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/election.voter3@org1.example.com/msp/config.yaml"
```
## Before Invoking
Before the functions are invoked, the required environment variables must be set.

### Common Exports
```
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051
export TARGET_TLS_OPTIONS=(-o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt")
```
### User-specific exports
#### For the user with the attribute 'election.admin'
```
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/election.admin3@org1.example.com/msp
```
#### For the user with the attribute 'election.voter'
```
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/election.voter3@org1.example.com/msp
```

## Additional Info
The Hyperledger Fabric Documentation can be read here (https://hyperledger-fabric.readthedocs.io/)

The Hyperledger Fabric CA Documentation can be read here (https://hyperledger-fabric-ca.readthedocs.io/en/latest/)

### Folder Structure
```
├───smart-contract
│   ├───constants
│   ├───core
│   ├───errors
│   ├───helpers
│   └───structs
│       └───fun
└───vendor
    ├───github.com
    │   ├───go-openapi
    │   │   ├───jsonpointer
    │   │   ├───jsonreference
    │   │   │   └───internal
    │   │   ├───spec
    │   │   └───swag
    │   ├───gobuffalo
    │   │   ├───envy
    │   │   ├───packd
    │   │   │   └───internal
    │   │   │       └───takeon
    │   │   │           └───github.com   
    │   │   │               └───markbates
    │   │   │                   └───errx 
    │   │   └───packr
    │   ├───golang
    │   │   └───protobuf
    │   │       ├───jsonpb
    │   │       ├───proto
    │   │       └───ptypes
    │   │           ├───any
    │   │           ├───duration
    │   │           └───timestamp
    │   ├───hyperledger
    │   │   ├───fabric-chaincode-go
    │   │   │   ├───pkg
    │   │   │   │   ├───attrmgr
    │   │   │   │   └───cid
    │   │   │   └───shim
    │   │   │       └───internal
    │   │   ├───fabric-contract-api-go
    │   │   │   ├───contractapi
    │   │   │   │   └───utils
    │   │   │   ├───internal
    │   │   │   │   ├───types
    │   │   │   │   └───utils
    │   │   │   ├───metadata
    │   │   │   └───serializer
    │   │   └───fabric-protos-go
    │   │       ├───common
    │   │       ├───ledger
    │   │       │   ├───queryresult
    │   │       │   └───rwset
    │   │       ├───msp
    │   │       └───peer
    │   ├───joho
    │   │   └───godotenv
    │   ├───josharian
    │   │   └───intern
    │   ├───mailru
    │   │   └───easyjson
    │   │       ├───buffer
    │   │       ├───jlexer
    │   │       └───jwriter
    │   ├───rogpeppe
    │   │   └───go-internal
    │   │       └───modfile
    │   └───xeipuuv
    │       ├───gojsonpointer
    │       ├───gojsonreference
    │       └───gojsonschema
    ├───golang.org
    │   └───x
    │       ├───mod
    │       │   ├───internal
    │       │   │   └───lazyregexp
    │       │   ├───modfile
    │       │   ├───module
    │       │   └───semver
    │       ├───net
    │       │   ├───http
    │       │   │   └───httpguts
    │       │   ├───http2
    │       │   │   └───hpack
    │       │   ├───idna
    │       │   ├───internal
    │       │   │   └───timeseries
    │       │   └───trace
    │       ├───sys
    │       │   └───unix
    │       └───text
    │           ├───secure
    │           │   └───bidirule
    │           ├───transform
    │           └───unicode
    │               ├───bidi
    │               └───norm
    ├───google.golang.org
    │   ├───genproto
    │   │   └───googleapis
    │   │       └───rpc
    │   │           └───status
    │   ├───grpc
    │   │   ├───attributes
    │   │   ├───backoff
    │   │   ├───balancer
    │   │   │   ├───base
    │   │   │   ├───grpclb
    │   │   │   │   └───state
    │   │   │   └───roundrobin
    │   │   ├───binarylog
    │   │   │   └───grpc_binarylog_v1
    │   │   ├───channelz
    │   │   ├───codes
    │   │   ├───connectivity
    │   │   ├───credentials
    │   │   │   └───insecure
    │   │   ├───encoding
    │   │   │   └───proto
    │   │   ├───grpclog
    │   │   ├───internal
    │   │   │   ├───backoff
    │   │   │   ├───balancer
    │   │   │   │   └───gracefulswitch
    │   │   │   ├───balancerload
    │   │   │   ├───binarylog
    │   │   │   ├───buffer
    │   │   │   ├───channelz
    │   │   │   ├───credentials
    │   │   │   ├───envconfig
    │   │   │   ├───grpclog
    │   │   │   ├───grpcrand
    │   │   │   ├───grpcsync
    │   │   │   ├───grpcutil
    │   │   │   ├───idle
    │   │   │   ├───metadata
    │   │   │   ├───pretty
    │   │   │   ├───resolver
    │   │   │   │   ├───dns
    │   │   │   │   ├───passthrough
    │   │   │   │   └───unix
    │   │   │   ├───serviceconfig
    │   │   │   ├───status
    │   │   │   ├───syscall
    │   │   │   └───transport
    │   │   │       └───networktype
    │   │   ├───keepalive
    │   │   ├───metadata
    │   │   ├───peer
    │   │   ├───resolver
    │   │   ├───serviceconfig
    │   │   ├───stats
    │   │   ├───status
    │   │   └───tap
    │   └───protobuf
    │       ├───encoding
    │       │   ├───protojson
    │       │   ├───prototext
    │       │   └───protowire
    │       ├───internal
    │       │   ├───descfmt
    │       │   ├───descopts
    │       │   ├───detrand
    │       │   ├───encoding
    │       │   │   ├───defval
    │       │   │   ├───json
    │       │   │   ├───messageset
    │       │   │   ├───tag
    │       │   │   └───text
    │       │   ├───errors
    │       │   ├───filedesc
    │       │   ├───filetype
    │       │   ├───flags
    │       │   ├───genid
    │       │   ├───impl
    │       │   ├───order
    │       │   ├───pragma
    │       │   ├───set
    │       │   ├───strs
    │       │   └───version
    │       ├───proto
    │       ├───reflect
    │       │   ├───protodesc
    │       │   ├───protoreflect
    │       │   └───protoregistry
    │       ├───runtime
    │       │   ├───protoiface
    │       │   └───protoimpl
    │       └───types
    │           ├───descriptorpb
    │           └───known
    │               ├───anypb
    │               ├───durationpb
    │               ├───emptypb
    │               └───timestamppb
    └───gopkg.in
        └───yaml.v3
```
