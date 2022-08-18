package define

import "flag"

//***************************************************
//author tim.huang
//2022/8/18
//
//
//***************************************************

//***********************    type    ****************************

//***********************    type_end    ****************************

// ***********************    var    ****************************

var CallDepth = flag.Int("log_depth", 3, "log depth")

var Modules = flag.String("modules", "Chat", "startup modules")

var Address = flag.String("address", "0.0.0.0", "server address")

// 端口如果是0，启动时会随机一个端口
var Port = flag.Int("port", 8081, "server port")

//***********************    var_end    ****************************

//***********************    interface    ****************************

//***********************    interface_end    ****************************

//***********************    struct    ****************************

//***********************    struct_end    ****************************

func init() {
}
