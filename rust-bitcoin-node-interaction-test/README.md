# Setup docker

Btw, <b>it will take long time to setup</b> (it took me 1.5 days on Linux and 7 days on Windows), just be pationate and treat yourself with a good night's sleep. You deserve it! After all, you have to download the whole blockchain...

```
# Generate credentials
# ---
python gen_bitcoin_rpc_auth.py bitcoinduser bitcoinduser

# Setup docker
# Replace {GENERATED_CREDENTIALS} with generated credentials from python
# Example "-rpcauth" can look like: -rpcauth='bitcoinduser:dfc30a59b9a9d09499e6c5cb130c2065$940a9c7bf5882656d74a94c5feebac3da912b8b66ccd1b6840e63eca99cc9841'
# (on Windows replace `pwd` with `pwd -W`)
# ---
docker run -it --rm -p 8333:8333 -p 8332:8332 -v $(pwd)/bitcoin-data:/home/bitcoin/.bitcoin ruimarinho/bitcoin-core:24-alpine -printtoconsole -rpcallowip=172.17.0.0/16 -rpcbind=0.0.0.0 -rpcauth='{GENERATED_CREDENTIALS}'
```

# Test that everything works:

- Run the following command: `curl --data-binary '{"jsonrpc":"1.0","id":"1","method":"getnetworkinfo","params":[]}' http://bitcoinduser:bitcoinduser@localhost:8332/`

- Example working answear should look like this:

```
{
  "result": {
    "version": 240001,
    "subversion": "/Satoshi:24.0.1/",
    "protocolversion": 70016,
    "localservices": "0000000000000409",
    "localservicesnames": ["NETWORK", "WITNESS", "NETWORK_LIMITED"],
    "localrelay": true,
    "timeoffset": 0,
    "networkactive": true,
    "connections": 3,
    "connections_in": 0,
    "connections_out": 3,
    "networks": [
      { "name": "ipv4", "limited": false, "reachable": true, "proxy": "", "proxy_randomize_credentials": false },
      { "name": "ipv6", "limited": false, "reachable": true, "proxy": "", "proxy_randomize_credentials": false },
      { "name": "onion", "limited": true, "reachable": false, "proxy": "", "proxy_randomize_credentials": false },
      { "name": "i2p", "limited": true, "reachable": false, "proxy": "", "proxy_randomize_credentials": false },
      { "name": "cjdns", "limited": true, "reachable": false, "proxy": "", "proxy_randomize_credentials": false }
    ],
    "relayfee": 0.00001,
    "incrementalfee": 0.00001,
    "localaddresses": [],
    "warnings": ""
  },
  "error": null,
  "id": "1"
}
```
