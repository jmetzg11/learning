fn main() {
    let x = 7;
    if big_number(x) {
        println!("A big number: {x}");
    } else {
        println!("A small number: {x}");
    }

    let condition = true;
    let number = if condition {5} else {6};
    println!("The value of the number with the let if thing is: {number}")


    // loop {
    //     println!("I would run forever!")
    // }
}

fn big_number(x: i32) -> bool {
    if x > 5 {
        true
    } else {
        false
    }
}

//loop laboel
fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");
}

// while loop with conditional
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1;
    }

    println!("LIFTOFF!!!");
}

// for loop
fn main() {
    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {element}");
    }
}
fn main() {
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
