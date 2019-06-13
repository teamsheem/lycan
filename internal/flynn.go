package internal

import (
	"os/exec"
	"regexp"
)

type Flynn struct {}

type CmdFlynn struct {
	Cmd string
	Params []string
	Output []byte
}

func (f *Flynn) CreateCmd(cmd string, params ...string) CmdFlynn {
	c := CmdFlynn{
		Cmd: cmd,
		Params: params,
	}
	return c
}


func (f *Flynn) RunCmd(c *CmdFlynn) (out []byte, err error) {
	out, err = exec.Command(c.Cmd, c.Params...).Output()
	if err != nil {
		c.Output = out
	}
	return
}

func (f *Flynn) AddCluster(tlsPin, clusterName, controllerDomain, controllerKey string) ([]byte, error){
	c := f.CreateCmd("flynn",
		"cluster",
		"add", "-p", tlsPin, clusterName, controllerDomain, controllerKey,
	)
	return f.RunCmd(&c)
}

func (f *Flynn) CreateApp(appName string) ([]byte, error) {
	c := f.CreateCmd("flynn", "create", appName)
	return f.RunCmd(&c)
}

func (f *Flynn) CreateMysqlDB(appName string) ([]byte, error) {
	c := f.CreateCmd("flynn", "-a", appName, "resource", "add", "mysql")
	return f.RunCmd(&c)
}

func (f *Flynn) RunMigration(appName string) ([]byte, error) {
	c := f.CreateCmd("flynn", "-a", appName, "run", "php", "artisan", "migrate", "--force")
	return f.RunCmd(&c)
}

func (f *Flynn) ExtractEnvs(appName string) (map[string]string, error) {
	rex := regexp.MustCompile("(\\w+)=(\\w+)")
	c := f.CreateCmd("flynn", "-a", appName, "env")
	r, err := f.RunCmd(&c)
	if err != nil {
		return nil, err
	}

	data := rex.FindAllStringSubmatch(string(r), -1)

	res := make(map[string]string)
	for _, kv := range data {
		k := kv[1]
		v := kv[2]
		res[k] = v
	}

	return res, nil
}

func (f *Flynn) SetEnv
