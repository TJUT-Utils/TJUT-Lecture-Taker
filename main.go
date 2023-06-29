package main

import (
	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

var config Config

func init() {
	viper.SetConfigName(`lecture.config`)
	viper.SetConfigType(`toml`)
	viper.AddConfigPath(`.`)

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println(`read config failed: `, err)
			return
		}
		//Try to write default config

		viper.SetDefault(`startTime`, time.Now())
		viper.SetDefault(`duration`, 500*time.Millisecond)
		viper.SetDefault(`firstPublicIDs`, []string{})
		viper.SetDefault(`firstPEIDs`, []string{})
		viper.SetDefault(`secondPublicIDs`, []string{})
		viper.SetDefault(`secondPEIDs`, []string{})
		viper.SetDefault(`thirdPublicIDs`, []string{})
		viper.SetDefault(`thirdPEIDs`, []string{})
		viper.SetDefault(`cookies`, ``)

		log.Println(`config file not found, creating a new one...`)
		viper.SetConfigFile(`lecture.config.toml`)
		err = viper.WriteConfig()
		if err != nil {
			log.Println(`failed to write a new config: `, err)
		}
		log.Println(`new config file created, please edit it and restart the program.`)
		os.Exit(0)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(`failed to read config, unidentified config file: `, err)
	}
}

func main() {
	operate()

	s := gocron.NewScheduler(time.UTC)
	_, err := s.StartAt(config.Time).Every(config.Duration).Do(operate)
	if err != nil {
		log.Println(err)
	}

	s.StartBlocking()
}

func operate() {
	client := newClient()

	var WG sync.WaitGroup
	for index := range config.FirstPublicIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleQxgxk`).
				SetQueryParam(`jxbid`, config.FirstPublicIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `1`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.FirstPublicIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.FirstPublicIDs[i], resp.Message)
			}
		}(index)
	}

	for index := range config.FirstPEIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleTykxk`).
				SetQueryParam(`jxbid`, config.FirstPEIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `1`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.FirstPEIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.FirstPEIDs[i], resp.Message)
			}
		}(index)
	}

	for index := range config.SecondPublicIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleQxgxk`).
				SetQueryParam(`jxbid`, config.SecondPublicIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `2`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.SecondPublicIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.SecondPublicIDs[i], resp.Message)
			}
		}(index)
	}

	for index := range config.SecondPEIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleTykxk`).
				SetQueryParam(`jxbid`, config.SecondPEIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `2`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.SecondPEIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.SecondPEIDs[i], resp.Message)
			}
		}(index)
	}

	for index := range config.ThirdPublicIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleQxgxk`).
				SetQueryParam(`jxbid`, config.ThirdPublicIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `3`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.ThirdPublicIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.ThirdPublicIDs[i], resp.Message)
			}
		}(index)
	}

	for index := range config.ThirdPEIDs {
		WG.Add(1)
		go func(i int) {
			defer WG.Done()
			rawResp, err := client.R().
				SetQueryParam(`method`, `handleTykxk`).
				SetQueryParam(`jxbid`, config.ThirdPEIDs[i]).
				SetQueryParam(`glJxbid`, ``).
				SetQueryParam(`xkzy`, `3`).
				Get(`http://xk.tjut.edu.cn/xsxk/xkOper.xk`)
			if err != nil {
				log.Println(err)
				return
			}

			var resp OperateResp
			err = jsoniter.Unmarshal(rawResp.Body(), &resp)
			if err != nil {
				log.Println(err)
				return
			}

			if resp.Success {
				log.Printf(`选课成功: %s`, config.ThirdPEIDs[i])
			} else {
				log.Printf(`选课失败: %s, %s`, config.ThirdPEIDs[i], resp.Message)
			}
		}(index)
	}

	WG.Wait()
}

func newClient() *resty.Client {
	client := resty.New()
	client.SetHeader("Cookie", config.Cookies)
	client.SetHeader("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)

	return client
}

type Config struct {
	FirstPublicIDs  []string      `toml:"first_ids"`
	SecondPublicIDs []string      `toml:"second_ids"`
	ThirdPublicIDs  []string      `toml:"third_ids"`
	FirstPEIDs      []string      `toml:"first_peids"`
	SecondPEIDs     []string      `toml:"second_peids"`
	ThirdPEIDs      []string      `toml:"third_peids"`
	Cookies         string        `toml:"cookies"`
	Time            time.Time     `toml:"time"`
	Duration        time.Duration `toml:"duration"`
}

type OperateResp struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
