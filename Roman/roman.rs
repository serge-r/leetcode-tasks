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

pub fn int_to_roman(num: i32) -> String {
    const LEN: usize = 13;

    let mut result: String = String::from("");
    let mut res;
    let mut num_copy = num;
    let nums: [i32; LEN] = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    let lit = HashMap::from([
        (1, "I"),
        (4, "IV"),
        (5, "V"),
        (9, "IX"),
        (10, "X"),
        (40, "XL"),
        (50, "L"),
        (90, "XC"),
        (100, "C"),
        (400, "CD"),
        (500, "D"),
        (900, "CM"),
        (1000, "M"),
    ]);

    for j in 0..LEN {
        if num_copy > 0 && num_copy < 4000 {
            res = num_copy / nums[j];
            if res > 3 {
                result.push_str(lit[&nums[j]]);
                result.push_str(lit[&nums[j - 1]]);
            } else {
                let mut i = res;
                while i > 0 {
                    result = result + lit[&nums[j]];
                    i -= 1;
                }
            }
            num_copy -= nums[j] * res;
        } else {
            break;
        }
    }

    result
}
