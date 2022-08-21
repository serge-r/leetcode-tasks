use std::collections::HashMap;

pub fn roman_to_int(str: &str) -> u32 {
    // Roman symbols
    let lit = HashMap::from([
        ('I', 1),
        ('V', 5),
        ('X', 10),
        ('L', 50),
        ('C', 100),
        ('D', 500),
        ('M', 1000),
    ]);

    // calculated vars
    let (mut counter, mut sum) = (0, 0);
    let mut prev_char = ' ';
    let mut prev_val;

    let string_len = str.len();

    if string_len <= 0 || string_len > 15 {
        return 0; // Not sure what I need to return if error
    }

    for ch in str.chars() {
        match lit.get(&ch) {
            Some(cur_val) => {
                if !lit.contains_key(&prev_char) {
                    prev_char = ch;
                    sum += cur_val;
                    continue;
                } else {
                    prev_val = lit[&prev_char];
                }

                if cur_val > &prev_val {
                    sum -= prev_val;
                    sum += cur_val - prev_val;
                    counter = 0;
                }

                if cur_val < &prev_val {
                    sum += cur_val;
                    counter = 0;
                }

                if cur_val == &prev_val {
                    match ch {
                        'V' | 'L' | 'D' => return 0,
                        _ => {
                            if counter < 2 {
                                sum += cur_val;
                                counter += 1;
                            } else {
                                return 0;
                            }
                        }
                    }
                }
                prev_char = ch;
            }
            None => return 0,
        }
    }
    sum
}
