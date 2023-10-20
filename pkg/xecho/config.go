package xecho

//
//import "github.com/Nav1Cr0ss/s-user-storage/pkg/configurator"
//
//// Config is struct to configure HTTP server
//type Config struct {
//	Host    string `default:""`
//	Port    string `default:"8085"`
//	Debug   bool   `default:"false"`
//	EnvName string `default:"" split_words:"true"`
//}
//
//type ConfigHttpServer struct {
//	Config
//	AuthServer string `required:"true" split_words:"true"`
//}
//
//func GetConfig(c *configurator.Configurator) *Config {
//	return c.New("api-http", &Config{}).(*Config)
//}
//
//func GetConfigHttpServer(c *configurator.Configurator) *ConfigHttpServer {
//	return c.New("api-http", &ConfigHttpServer{}).(*ConfigHttpServer)
//}
