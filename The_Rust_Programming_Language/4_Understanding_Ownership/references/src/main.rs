fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);

    println!("The length of '{s1}' is {len}");

    let mut x1 = String::from("hello");
    // let x2 = &x1 would be a problem because two reference to the same mut value
    // Not allowed to prevent race conditions
    change(&mut x1);
    println!("x1 was changed to {x1}");

    let mut y = String::from("x");

}

fn calculate_length(s: &String) -> usize { // s is reference to a string
    s.len()
} // reference goes out of scope but value is not dropped because here we never had ownership

fn change(x: &mut String) { // no other reference to a mutable value allowed, only one
    x.push_str(", dog face")
}


// also not allowed, user of unmutable references don't expect the value to suddenly change
let mut s = String::from("hello");
let r1 = &s;
let r2 = &s; // multiple immutable references are allowed
let r3 = &mut s;

// this works because r1 and r2 leave scope in the println!
let mut s = String::from("hello");
let r1 = &s; // no problem
let r2 = &s; // no problem
println!("{r1} and {r2}");
// Variables r1 and r2 will not be used after this point.
let r3 = &mut s; // no problem
println!("{r3}");

// Dangling References
fn main() {
    let reference_to_nothing = dangle();
}

fn dangle() -> &String {
    let s = String::from("hello");

    &s // returning a reference to a value
} // s goes out out of scope, the refence in main does not have a value

// this works because ownership is returned
fn no_dangle() -> String {
    let s = String::from("hello");
    s
}
