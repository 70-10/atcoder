use proconio::input;

fn main() {
    input! {
        a: i32,
        b: i32
    }

    println!("{}", logic(a, b))
}

fn logic(a: i32, b: i32) -> i32 {
    let sum = a + b;

    if sum > 12 {
        return sum - 12;
    }

    return sum;
}

#[test]
fn test_logic() {
    assert_eq!(logic(1, 2), 3);
    assert_eq!(logic(5, 9), 2);
    assert_eq!(logic(12, 12), 12);
}
