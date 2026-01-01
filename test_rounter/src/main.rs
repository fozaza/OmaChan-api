use anyhow::anyhow;
use cmd_lib::{run_cmd, run_fun};
use regex::{self, Regex};
use serde_json::json;

fn main() -> anyhow::Result<()> {
    println!("Test rust");
    let base_url = "localhost:3000";
    run_cmd!(curl $base_url)?;

    let payload = json!({
        "Email":    "test@email.com",
        "Password": "12345678"
    })
    .to_string();

    println!("payload {}", payload);
    let data = run_fun!(
        curl -S -X POST
            -H "Content-Type: application/json"
            -d $payload
            localhost:3000/login
    )?;
    // yay i know this is not good code but it work
    let vec_data: Vec<_> = data.split(",").collect();
    let mut string_token: String = vec_data[1].split("token").collect();
    string_token = string_token.split(r#":""#).collect();
    string_token = string_token.split(r#""#).collect();
    string_token = string_token.trim_end_matches("\"}").to_string();
    string_token.remove(0);
    string_token.remove(0);
    println!("token {}", string_token);

    let auth = format!("Authorization: Bearer {}", string_token);
    run_cmd!(
        curl -S
            -H "Content-Type: application/json"
            -H $auth
            localhost:3000/admin
    )?;
    Ok(())
}
