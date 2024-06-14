# bursa

<div align="center">
    <img src="./assets/bursa-logo-with-text-horizontal.png" alt="Bursa Logo" width="640">
</div>

Programmatic Cardano Wallet


## Generating a key pair

```bash
# Create a directory to store files in
mkdir ~/tmp/
cd ~/tmp/

# Create a mnemonic
bursa mnemonic create   # prints mnemonic to stdout
MNEMONIC="$(bursa mnemonic create)" # store in $MNEMONIC

# Create payment key pair from mnemonic
bursa key payment --verification-key-file payment.vkey --signing-key-file payment.skey --from-mnemonic $MNEMONIC
# Writes out payment.vkey and payment.skey

# Create staking key pair from mnemonic
bursa key staking --verification-key-file staking.vkey --signing-key-file staking.skey --from-mnemonic $MNEMONIC
```