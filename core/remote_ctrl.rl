package core

import (
    "fmt"
)

var (
    status = 0
    tempe = 26
)

%%{
    machine remote_ctrl;

    include temperature "temperature.rl";

    action on{
        status = 1
    }

    action off{
        status = 0
    }

    write data;
}%%

func Command(data string){
    cs, p, pe, eof := 0, 0, len(data), len(data)
    _ = eof
    var ts,te,act int
    _ = act
    _ = ts
    _ = te
    %%{
        main := |*
            ' ';
            'incr'i => incr;
            'dec'i => dec;
            'turn on'i => on;
            'turn off'i => off;
            '\n' => {return;};
        *|;

        write init;
        write exec;
    }%%
}

func PrintCurrentState(){
    statusStr := "关机"
    if status == 1{
        statusStr = "开机"
    }
    fmt.Printf("当前状态：%s\n当前温度：%d\n",statusStr,tempe)
    fmt.Println("=========================")
    return;
}