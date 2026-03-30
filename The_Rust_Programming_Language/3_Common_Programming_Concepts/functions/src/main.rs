fn main() {
    another_function();
    function_with_parameters(3, 'a');
    let y = plus_five(3);
    println!("The value of the expression function: {y}")
}

fn another_function() {
    println!("Use snake case for functions and variables");
}

// must declare the type of the parameter
fn function_with_parameters(x: i32, unit_label: char) {
    println!("The parameter/arguments are {x} and {unit_label}");
}

fn plus_five(x: i32) -> i32 {
    x + 5
}
