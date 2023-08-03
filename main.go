package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseDir string `yaml:"baseDir"`
	Items   map[string]struct {
		Dicts []string `yaml:"dicts"`
		Path  string   `yaml:"path"`
		Alias []string `yaml:"alias"`
	} `yaml:",inline"`
}

func parseConfig(configfile string) (Config, error) {
	// 读取配置文件
	content, err := os.ReadFile(configfile)
	if err != nil {
		fmt.Println("读取配置文件出错：", err)
		return Config{}, err
	}

	// 解析YAML
	var config Config
	errrrrr := yaml.Unmarshal(content, &config)
	if errrrrr != nil {
		// panic(err)
		return Config{}, err
	}

	return config, nil
}
func main() {
	parseConfig("configfile")
	fmt.Println("请输入数据，以两个回车符为结束标志：")

	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// 遇到两个回车符结束输入
			break
		}

		lines = append(lines, line)
	}

	fmt.Println("您输入的数据为：")
	fmt.Printf("长度为: %d\n", len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
}
