fn main() {
    let s = String::from("hello");

    takes_ownership(s); // s moves into this function, no longer valide here

    let x = 5;

    makes_copy(x); // i32 implements the Copy train

    // x still valid here
}

fn takes_ownership(some_string: String) { // some_string comes into scope
    println!("{some_string}");
} // some_string leaves scope and is dropped

fn makes_copy(come_integer: i32) {
    println!("{some_integer")
} // copy is dropped here
