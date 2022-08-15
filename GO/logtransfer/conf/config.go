package conf

type LogTransConf struct {
	KafkaInfo `ini:"kafka"`
	ESInfo    `ini:"ES"`
	EtcdInfo  `ini:"etcd"`
}

type KafkaInfo struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESInfo struct {
	Address  string `ini:"address"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

type EtcdInfo struct {
	Address    string `ini:"address"`
	TimeoutSec int    `ini:"timeout"`
	Key        string `ini:"key"`
	// /logagent/log-collect/:[{"path":"./my.log","topic":"testkey"},{"path":"./db.log","topic":"testkey"},{"path":"./serv.log","topic":"testkey"}]
}
