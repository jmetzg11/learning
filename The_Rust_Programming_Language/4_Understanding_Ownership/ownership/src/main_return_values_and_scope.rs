fn main() {
    let s1 = gives_ownership();
    println!("{s1}");
    let s2 = String::from("hello");
    let s3 = takes_and_gives_back(s2); // s2 moves to function and then moves back to s3
    // s2 is dropped
}

fn gives_ownership() -> String {
    let some_string = String::from("yours");
    some_string
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string
}
