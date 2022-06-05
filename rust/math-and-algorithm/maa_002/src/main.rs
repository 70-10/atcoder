use proconio::input;

fn main() {
    input! {
        x: [usize;3]
    }
    
    println!("{}", x.iter().sum::<usize>());
}
