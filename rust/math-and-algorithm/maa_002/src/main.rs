use proconio::input;

fn main() {
    input! {
      a:i32,
      b:i32,
      c:i32
    }

    println!("{}", logic(a, b, c));
}

fn logic(a: i32, b: i32, c: i32) -> i32 {
    return a + b + c;
}

#[test]
fn test_logic() {
    assert_eq!(logic(1, 2, 3), 6);
}
