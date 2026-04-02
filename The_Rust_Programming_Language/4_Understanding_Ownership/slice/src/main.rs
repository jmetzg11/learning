// slice is a kind of reference, does not have ownership


fn main() {
    let s = String::from("dog cat fish");
    let word = first_word(&s);
    println!("first word: {word}");

    let x = "bat rabbit crow"; // is also a reference; string literal
    let word = first_word(x);
    println!("first word: {word}");
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    &s[..]
}


// &s[3..len] == &s[3..]
// &s[0..9] == &[..9]
