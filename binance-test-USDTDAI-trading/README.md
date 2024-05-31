This code **actually executes trades**, be aware of running it!

The original idea was: "Since DAI and USDT are both stable coins, however the price fluctuates, could I execute trades in a way that I can get like 0.2% per trade?".

Well, yes, I could, even after fees! The problem is that there are like 3 of these trade opportunities per 2 weeks, which leads to +-0.5% gain per month. This amounts to +-6% per year. Not only is crypto too unstable for this to be a viable long term stategy, but I could get `bank time deposits` at the same rate and 1000% more security.

I wouldn't use Rust for communication with Binance again. The official library is kinda trash and [the crate (that seems official)](https://crates.io/crates/binance_spot_connector_rust) is NOT owned by Binance!

I wrote this in 30 minutes, so please don't judge xd

Before running it, copy `.env.example` to `.env` and the required variables.

Oh, I hate crypto btw ... except Monero :)
