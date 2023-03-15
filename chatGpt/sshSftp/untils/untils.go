package untils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Ssh  ssh  `yaml:"ssh"`
	Path path `yaml:"path"`
	//RemoteFilePath remoteFilePath `yaml:"remoteFilePath"`
	//LocalFilePath  localFilePath  `yaml:"localFilePath"`
}

type path struct {
	Filepath1 filepath1 `yaml:"filepath1"`
	Filepath2 filepath2 `yaml:"filepath2"`
}

type ssh struct {
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Ip     string `yaml:"ip"`
	Port   string `yaml:"port"`
}

type filepath1 struct {
	RemoteDir1 string `yaml:"remoteDir1"`
	LocalDir1  string `yaml:"localDir1"`
}

type filepath2 struct {
	RemoteDir2 string `yaml:"remoteDir2"`
	LocalDir2  string `yaml:"localDir2"`
}

// ConfigData 定义一个全局变量
var ConfigData *config

// LoadConfig 读取配置文件
//func LoadConfig(fileName string) error {
func LoadConfig() error {
	//config := new(config)
	//
	//binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//filePath := filepath.Join(binPath, fileName)
	//data, err := ioutil.ReadFile(filePath)
	//
	////反序列化yaml格式内容到结构体中
	//err = yaml.Unmarshal([]byte(data), &config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 将解析后的配置文件数据存入全局变量中
	//ConfigData = config
	//
	//return nil
	config := new(config)
	// 读取yaml格式的配置文件内容,使用 os.ReadFile
	//data, err := ioutil.ReadFile("/Users/wangyi/go/src/chatGpt/config/config.yaml")  //带缓冲的读取文件,适用于读取大文件
	data, err := os.ReadFile("/Users/wangyi/go/src/chatGpt/config/config.yaml") //一次将文件读完,使用于小文件
	if err != nil {
		return err
	}
	// 反序列化 yaml 格式内容到结构体中
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	// 将解析后的配置文件数据存入全局变量中
	ConfigData = config
	return nil

}