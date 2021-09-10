%%{
    machine mode;

    action switchMode{
        if status == 0{
            return
        }
        switch mod {
            case modHot:
                mod = modCold
            case modCold:
                mod = modWind
            case modWind:
                mod = modHot
        }
    }
}%%