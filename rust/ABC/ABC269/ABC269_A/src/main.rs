use proconio::input;

fn main() {
    input! {
        a: i32,
        b: i32,
        c: i32,
        d: i32
    }

    println!("{}", logic(a, b, c, d))
}

fn logic(a: i32, b: i32, c: i32, d: i32) -> String {
    return format!("{}\nTakahashi", (a + b) * (c - d));
}

#[test]
fn test_logic() {
    assert_eq!(logic(1, 2, 5, 3), "6\nTakahashi");
    assert_eq!(logic(10, -20, 30, -40), "-700\nTakahashi")
}
