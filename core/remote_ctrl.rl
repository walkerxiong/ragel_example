package core

import (
    "fmt"
)

type Mode int

const (
    modHot Mode = iota
    modCold 
    modWind
)

var (
    status = 0
    tempe = 26
    mod = modWind
)

%%{
    machine remote_ctrl;

    include temperature "temperature.rl";
    include mode "mode.rl";

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
            'mod'i => switchMode;
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
    fmt.Printf("当前状态：%s\n当前温度：%d\n当前模式: %s\n",statusStr,tempe,mod)
    fmt.Println("=========================")
    return;
}


func (m Mode)String()string{
    switch m {
        case modHot:
            return "热风模式"
        case modCold:
            return "冷风模式"
        case modWind:
            return "吹风模式"
    }
    return "模式异常"
}