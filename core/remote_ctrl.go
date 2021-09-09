
//line core/remote_ctrl.rl:1
package core

import (
    "fmt"
)

var (
    status = 0
    tempe = 26
)


//line core/remote_ctrl.go:16
var _remote_ctrl_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
}

var _remote_ctrl_key_offsets []byte = []byte{
	0, 0, 2, 4, 6, 8, 10, 12, 
	14, 16, 17, 19, 23, 25, 
}

var _remote_ctrl_trans_keys []byte = []byte{
	69, 101, 67, 99, 78, 110, 67, 99, 
	82, 114, 85, 117, 82, 114, 78, 110, 
	32, 79, 111, 70, 78, 102, 110, 70, 
	102, 10, 32, 68, 73, 84, 100, 105, 
	116, 
}

var _remote_ctrl_single_lengths []byte = []byte{
	0, 2, 2, 2, 2, 2, 2, 2, 
	2, 1, 2, 4, 2, 8, 
}

var _remote_ctrl_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 
}

var _remote_ctrl_index_offsets []byte = []byte{
	0, 0, 3, 6, 9, 12, 15, 18, 
	21, 24, 26, 29, 34, 37, 
}

var _remote_ctrl_indicies []byte = []byte{
	0, 0, 1, 2, 2, 1, 3, 3, 
	1, 4, 4, 1, 5, 5, 1, 6, 
	6, 1, 7, 7, 1, 8, 8, 1, 
	9, 1, 10, 10, 1, 11, 12, 11, 
	12, 1, 13, 13, 1, 14, 15, 16, 
	17, 18, 16, 17, 18, 1, 
}

var _remote_ctrl_trans_targs []byte = []byte{
	2, 0, 13, 4, 5, 13, 7, 8, 
	9, 10, 11, 12, 13, 13, 13, 13, 
	1, 3, 6, 
}

var _remote_ctrl_trans_actions []byte = []byte{
	0, 0, 9, 0, 0, 7, 0, 0, 
	0, 0, 0, 0, 11, 13, 15, 5, 
	0, 0, 0, 
}

var _remote_ctrl_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 
}

var _remote_ctrl_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 3, 
}

const remote_ctrl_start int = 13
const remote_ctrl_first_final int = 13
const remote_ctrl_error int = 0

const remote_ctrl_en_main int = 13


//line core/remote_ctrl.rl:26


func Command(data string){
    cs, p, pe, eof := 0, 0, len(data), len(data)
    _ = eof
    var ts,te,act int
    _ = act
    _ = ts
    _ = te
    
//line core/remote_ctrl.go:99
	{
	cs = remote_ctrl_start
	ts = 0
	te = 0
	act = 0
	}

//line core/remote_ctrl.go:107
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_remote_ctrl_from_state_actions[cs])
	_nacts = uint(_remote_ctrl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _remote_ctrl_actions[_acts - 1] {
		case 1:
//line NONE:1
ts = p

//line core/remote_ctrl.go:130
		}
	}

	_keys = int(_remote_ctrl_key_offsets[cs])
	_trans = int(_remote_ctrl_index_offsets[cs])

	_klen = int(_remote_ctrl_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _remote_ctrl_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _remote_ctrl_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_remote_ctrl_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _remote_ctrl_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _remote_ctrl_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_remote_ctrl_indicies[_trans])
	cs = int(_remote_ctrl_trans_targs[_trans])

	if _remote_ctrl_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_remote_ctrl_trans_actions[_trans])
	_nacts = uint(_remote_ctrl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _remote_ctrl_actions[_acts-1] {
		case 2:
//line core/remote_ctrl.rl:37
te = p+1

		case 3:
//line core/temperature.rl:5
te = p+1
{
        if status == 1{
            tempe += 1
        }
    }
		case 4:
//line core/temperature.rl:11
te = p+1
{
        if status == 1{
            tempe -= 1
        }
    }
		case 5:
//line core/remote_ctrl.rl:17
te = p+1
{
        status = 1
    }
		case 6:
//line core/remote_ctrl.rl:21
te = p+1
{
        status = 0
    }
		case 7:
//line core/remote_ctrl.rl:42
te = p+1
{return;}
//line core/remote_ctrl.go:235
		}
	}

_again:
	_acts = int(_remote_ctrl_to_state_actions[cs])
	_nacts = uint(_remote_ctrl_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _remote_ctrl_actions[_acts-1] {
		case 0:
//line NONE:1
ts = 0

//line core/remote_ctrl.go:249
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line core/remote_ctrl.rl:47

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