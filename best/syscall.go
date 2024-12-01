package best

import "github.com/zlllgp/dream/config"

func Syscall() {
	//Good 使用全局变量
	config.GetEnv()
	//Bad os.Getenv()
}
