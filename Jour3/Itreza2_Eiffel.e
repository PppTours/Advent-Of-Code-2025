class
    APPLICATION

inherit
    ARGUMENTS_32

create
    make

feature {NONE}

    make

        do
            create input_file.make_open_read ("input.txt")
            final_result := 0
            
            from 
                input_file.read_line
            until
                input_file.exhausted
            loop
                last_index := 0
                
                from
                    j := 0
                until
                    j = 12
                loop
                    last_digit := 0
                    
                    from
                        i := last_index + 1
                    until
                        i > input_file.last_string.count - (11 - j)
                    loop
                        if (input_file.last_string.code (i) - 48).as_integer_32 > last_digit then
                            last_index := i
                            last_digit := (input_file.last_string.code (i) - 48).as_integer_32
                        end
                        i := i + 1
                    end

                    from --this loop is basically just digit * 10^(11-j) ...
                        i := 0
                    until
                        i = 11 - j
                    loop
                        last_digit := last_digit * 10
                        i := i + 1
                    end --this is so stupid, but I'm too lazy to do otherwise
                    
                    final_result := final_result + last_digit
                    j := j + 1
                end

                input_file.read_line    
            end
            input_file.close
            
            print(final_result)
        end

feature

    input_file: PLAIN_TEXT_FILE
    final_result: INTEGER_64
    last_index: INTEGER_32
    last_digit: INTEGER_64
    i: INTEGER_32
    j: INTEGER_32

end