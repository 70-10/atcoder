use proconio::input;

fn main() {
    input! {
        s: String
    }

    print!("{}", logic(&s))
}

fn logic(s: &str) -> String {
    match s {
        "red" => "SSS".to_string(),
        "blue" => "FFF".to_string(),
        "green" => "MMM".to_string(),
        _ => "Unknown".to_string(),
    }
}

#[test]
fn test_logic() {
    assert_eq!(logic("red"), "SSS");
    assert_eq!(logic("blue"), "FFF");
    assert_eq!(logic("green"), "MMM");
    assert_eq!(logic("atcoder"), "Unknown");
}
