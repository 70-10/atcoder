use proconio::input;

fn main() {
    input! {
        x: i32,
    }

    println!("{}", logic(x))
}

fn logic(num: i32) -> i32 {
    return 5 + num;
}


#[test]
fn test_logic() {
    assert_eq!(logic(5), 5);
}