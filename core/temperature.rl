
%%{
    machine  temperature;

    action incr{
        if status == 1{
            tempe += 1
        }
    }

    action dec{
        if status == 1{
            tempe -= 1
        }
    }

}%%

