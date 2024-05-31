use rust_decimal::Decimal;
use rust_decimal_macros::dec;

pub const MARKET_SYMBOL: &str = "USDTDAI";

pub const NUM_OF_DAI_TO_TRADE: Decimal = dec!(105);

pub const RECV_WINDOW: u64 = 20_000;
