use std::str::FromStr;

use rust_decimal::Decimal;
use serde::de::Error as DeError;
use serde::{Deserialize, Deserializer, Serialize};

pub struct Strategy {
    pub sell_usdt_dai_price: Decimal,
    pub buy_usdt_dai_price: Decimal,
}
impl Strategy {
    pub fn new(sell_usdt_dai_price: Decimal, buy_usdt_dai_price: Decimal) -> Strategy {
        Self {
            sell_usdt_dai_price,
            buy_usdt_dai_price,
        }
    }
}

// todo: fucking sucks, check API reference and fix!!!!

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct OpenTradesApiResponse {
    // pub symbol: String,
    // #[serde(rename = "orderId")]
    // pub order_id: i64,
    // #[serde(rename = "orderListId")]
    // pub order_list_id: i64,
    // #[serde(rename = "clientOrderId")]
    // pub client_order_id: String,
    #[serde(deserialize_with = "string_to_decimal")]
    pub price: Decimal,
    // #[serde(rename = "origQty", deserialize_with = "string_to_decimal")]
    // pub orig_qty: Decimal,
    // #[serde(rename = "executedQty", deserialize_with = "string_to_decimal")]
    // pub executed_qty: Decimal,
    // #[serde(rename = "cummulativeQuoteQty", deserialize_with = "string_to_decimal")]
    // pub cummulative_quote_qty: Decimal,
    // pub status: String,
    // #[serde(rename = "timeInForce")]
    // pub time_in_force: String,
    // #[serde(rename = "type")]
    // pub order_type: String,
    pub side: String,
    // #[serde(rename = "stopPrice", deserialize_with = "string_to_decimal")]
    // pub stop_price: Decimal,
    // #[serde(rename = "icebergQty", deserialize_with = "string_to_decimal")]
    // pub iceberg_qty: Decimal,
    // pub time: i64,
    // #[serde(rename = "updateTime")]
    // pub update_time: i64,
    // #[serde(rename = "isWorking")]
    // pub is_working: bool,
    // #[serde(rename = "workingTime")]
    // pub working_time: i64,
    // #[serde(rename = "origQuoteOrderQty", deserialize_with = "string_to_decimal")]
    // pub orig_quote_order_qty: Decimal,
    // #[serde(rename = "selfTradePreventionMode")]
    // pub self_trade_prevention_mode: String,
}

fn string_to_decimal<'de, D>(deserializer: D) -> Result<Decimal, D::Error>
where
    D: Deserializer<'de>,
{
    let s: &str = Deserialize::deserialize(deserializer)?;
    Decimal::from_str(s).map_err(DeError::custom)
}
