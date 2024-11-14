// 對讀取設定的行為進行封裝
package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
	/*
		type Viper struct{...}
		configName        string
		configFile        string
	*/
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	/*
			func (v *Viper) SetConfigName(in string) {
			if in != "" {
				v.configName = in
				v.configFile = ""
			}
		}
	*/
	vp.AddConfigPath("configs/")
	/*
			func (v *Viper) AddConfigPath(in string) {
			if in != "" {
				absin := absPathify(in)
				jww.INFO.Println("adding", absin, "to paths to search")
				if !stringInSlice(absin, v.configPaths) {
					v.configPaths = append(v.configPaths, absin)
				}
			}
		}
	*/
	vp.AddConfigPath("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
