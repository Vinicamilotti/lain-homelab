package cli

import (
	"fmt"
	"os"
	"os/exec"

	confighandler "gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/configHandler"
	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/frontend"
	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/models"
)

type AppFunc func(config models.ConfigFile)

type Cli struct {
	Config  models.ConfigFile
	AppOpts map[int]AppFunc
}

func NewCli(config models.ConfigFile) *Cli {
	opts := map[int]AppFunc{}

	return &Cli{
		Config:  config,
		AppOpts: opts,
	}
}

func (c *Cli) RunApp() {

	c.AppOpts[0] = func(config models.ConfigFile) {
		c.CreateConfigFileFromTea(frontend.StartConfigUI(config))
	}

	c.AppOpts[1] = func(config models.ConfigFile) {
		for k, v := range config.AutoStart {
			if !v.Tracked {
				continue
			}
			fmt.Printf("%s: %s\n", k, v.Path)
		}
	}
	c.AppOpts[2] = func(config models.ConfigFile) {
		c.RestartTrackedServices()
	}
	c.AppOpts[3] = func(config models.ConfigFile) {
		c.StopTrackedServices()
	}
	c.AppOpts[4] = func(config models.ConfigFile) {
		os.Exit(0)
	}

	for {
		opt := frontend.StartMainMenu()
		c.AppOpts[opt.Choice](c.Config)
	}

}

func (c *Cli) RestartTrackedServices() {
	for service, config := range c.Config.AutoStart {
		if !config.Tracked {
			continue
		}
		cmd := exec.Command("docker", "compose", "restart")

		// 2. Set the "working directory" for this command
		// This is the equivalent of running 'cd' first.
		cmd.Dir = config.Path

		// 3. Connect the output so you can see what's happening
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Navigating to %s and restarting...\n", config.Path)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error restarting %s: %s\n", service, err.Error())
			return
		}

		fmt.Printf("Service %s restarted successfully.\n", service)

	}
}

func (c *Cli) StopTrackedServices() {
	for service, config := range c.Config.AutoStart {
		if !config.Tracked {
			continue
		}
		cmd := exec.Command("docker", "compose", "down")

		// 2. Set the "working directory" for this command
		// This is the equivalent of running 'cd' first.
		cmd.Dir = config.Path

		// 3. Connect the output so you can see what's happening
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Navigating to %s and restarting...\n", config.Path)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error stopping %s: %s\n", service, err.Error())
			return
		}

		fmt.Printf("Service %s restarted successfully.\n", service)
	}
}

func (c *Cli) CreateConfigFileFromTea(m frontend.ConfigServicesMenu) {
	if !m.Save {
		return
	}
	var config models.ConfigFile
	config.AutoStart = make(map[string]models.ServiceConfig)

	for i, service := range m.Selection {

		_, ok := m.Selected[i]

		c := m.Config.AutoStart[service]
		config.AutoStart[service] = models.ServiceConfig{
			Path:    c.Path,
			Tracked: ok,
		}

	}
	confighandler.SaveConfig(config)
}
