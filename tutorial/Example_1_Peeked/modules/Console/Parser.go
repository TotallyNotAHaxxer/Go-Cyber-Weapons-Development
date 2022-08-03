package Peeked

import Outbound_Functions "main/modules/Functions"

var dt Outbound_Functions.Console_Data

var Func_map = map[string]interface{}{
	"host": dt.Set_Aux_Host,
	"verb": dt.Set_Aux_Verbosity,
}

var Console_Map = map[string]interface{}{
	"settings": dt.Settings,
	"clear":    dt.Clear,
}

var Script_Map = map[string]interface{}{
	"scan": dt.Thread_Call,
}
