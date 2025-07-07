use elasticsearch::{Elasticsearch, SearchParts};
use serde_json::{json, Value};
use std::collections::HashMap;

pub async fn search_products(query: HashMap<String, String>) -> Result<Value, Box<dyn std::error::Error>> {
    let client = Elasticsearch::default();

    let search_query = json!({
        "query": {
            "multi_match": {
                "query": query.get("q").unwrap_or(&"".to_string()),
                "fields": ["name", "description", "category", "brand"]
            }
        }
    });

    let response = client
        .search(SearchParts::Index(&["products"]))
        .body(search_query)
        .send()
        .await?;

    let result = response.json::<Value>().await?;
    Ok(result)
}

pub async fn suggest_terms(query: HashMap<String, String>) -> Result<Value, Box<dyn std::error::Error>> {
    let client = Elasticsearch::default();

    let suggest_query = json!({
        "suggest": {
            "text": query.get("q").unwrap_or(&"".to_string()),
            "term-suggest": {
                "term": {
                    "field": "name"
                }
            }
        }
    });

    let response = client
        .search(SearchParts::Index(&["products"]))
        .body(suggest_query)
        .send()
        .await?;

    let result = response.json::<Value>().await?;
    Ok(result)
}
