mod constants;
mod types;

use std::time::Duration;

use binance_spot_connector_rust::{
    http::Credentials,
    hyper::BinanceHttpClient,
    trade::{self, new_order::NewOrder},
};
use constants::NUM_OF_DAI_TO_TRADE;
use dotenv::dotenv;
use rust_decimal::Decimal;
use rust_decimal_macros::dec;

use crate::{
    constants::{MARKET_SYMBOL, RECV_WINDOW},
    types::{OpenTradesApiResponse, Strategy},
};

#[tokio::main]
async fn main() {
    // init
    env_logger::Builder::from_default_env()
        .filter(None, log::LevelFilter::Debug)
        .init();
    dotenv().ok();

    let strategies = [
        Strategy::new(dec!(1.001), dec!(0.999)),
        Strategy::new(dec!(1.002), dec!(0.9984)),
        Strategy::new(dec!(1.005), dec!(0.9975)),
    ];

    let api_key = std::env::var("BINANCE_API_KEY").expect("api key required");
    let api_secret = std::env::var("BINANCE_API_SECRET").expect("binance api secret required");
    let credentials = Credentials::from_hmac(api_key, api_secret);
    let client = BinanceHttpClient::default().credentials(credentials);

    loop {
        log::debug!("Starting another round of checking...");
        let open_trades: Vec<OpenTradesApiResponse> = {
            let request = trade::open_orders()
                .symbol(MARKET_SYMBOL)
                .recv_window(RECV_WINDOW);
            let data = client
                .send(request)
                .await
                .expect("Unable to fetch the open orders.")
                .into_body_str()
                .await
                .expect("Unable to parse into a body string");
            log::debug!("opened orders reponse: {}", data);
            serde_json::from_str(&data).expect("Unable to parse into OpenTradesApiRenspose")
        };

        for strategy in &strategies {
            if find_trade(open_trades.clone(), "SELL", strategy.sell_usdt_dai_price).is_none() {
                let order = generate_order(trade::order::Side::Sell, strategy.sell_usdt_dai_price);
                _ = client
                    .send(order)
                    .await
                    .unwrap_or_else(|_| {
                        panic!(
                            "Unable to send the sell order (sell, {})",
                            strategy.sell_usdt_dai_price
                        )
                    })
                    .into_body_str()
                    .await;
                log::info!("Sending order to sell at {}", strategy.sell_usdt_dai_price);
            }
            if find_trade(open_trades.clone(), "BUY", strategy.buy_usdt_dai_price).is_none() {
                let order = generate_order(trade::order::Side::Buy, strategy.buy_usdt_dai_price);
                _ = client
                    .send(order)
                    .await
                    .unwrap_or_else(|_| {
                        panic!(
                            "Unable to send the sell order (buy, {})",
                            strategy.buy_usdt_dai_price
                        )
                    })
                    .into_body_str()
                    .await;
                log::info!("Sending order to buy at {}", strategy.buy_usdt_dai_price);
            }
        }

        tokio::time::sleep(Duration::from_secs(20)).await;
    }
}

fn find_trade(
    open_trades: Vec<OpenTradesApiResponse>,
    side: &str,
    price: Decimal,
) -> Option<OpenTradesApiResponse> {
    open_trades
        .into_iter()
        .find(|s| s.side == side && s.price == price)
}

fn generate_order(side: trade::order::Side, price: Decimal) -> NewOrder {
    trade::new_order(MARKET_SYMBOL, side, "LIMIT")
        .price(price)
        .quantity(NUM_OF_DAI_TO_TRADE)
        .time_in_force(trade::order::TimeInForce::Gtc)
        .recv_window(RECV_WINDOW)
}
